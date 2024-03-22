package anvil

import (
	"context"
	"fmt"
	"github.com/Flaque/filet"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/google/uuid"
	"github.com/ipfs/go-log"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/dockerutil"
	"github.com/synapsecns/sanguine/core/mapmutex"
	"github.com/synapsecns/sanguine/core/processlog"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/base"
	"github.com/synapsecns/sanguine/ethergo/chain"
	"github.com/synapsecns/sanguine/ethergo/chain/client"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	"github.com/teivah/onecontext"
	"math"
	"math/big"
	"os"
	"strings"
	"sync"
	"testing"
)

const gasLimit = 10000000

// Backend contains the anvil test backend.
type Backend struct {
	*base.Backend
	// fundingMux is used to lock the wallets while funding
	// since FundAccount is expected to add to existing balance
	fundingMux mapmutex.StringerMapMutex
	// store stores the accounts
	store *base.InMemoryKeyStore
	// chainConfig is the chain config
	chainConfig *params.ChainConfig
	// impersonationMux is used to lock the impersonation
	impersonationMux sync.Mutex
}

// BackendName is the name of the anvil backend.
const BackendName = "anvil"

// BackendName returns the name of the backend.
func (f *Backend) BackendName() string {
	return BackendName
}

// NewAnvilBackend creates a test anvil backend.
// nolint: cyclop
func NewAnvilBackend(ctx context.Context, t *testing.T, args *OptionBuilder) *Backend {
	t.Helper()

	pool, err := dockertest.NewPool("")
	assert.Nil(t, err)

	pool.MaxWait = args.maxWait
	if err != nil {
		assert.Nil(t, err)
	}

	commandArgs, err := args.Build()
	assert.Nil(t, err)

	runOptions := &dockertest.RunOptions{
		Repository: "ghcr.io/foundry-rs/foundry",
		// Note: https://github.com/foundry-rs/foundry/commit/6e041f9751efa6b75420689b862df05b0934022b introduces a breaking change with regards to
		// eth_BsendTransaction. The commit changes the way tx fields are detected. This will be fixed (on the anvil or ethergo sides) in a future version.
		Tag: "nightly-7398b65e831f2339d1d0a0bb05ade799e4f9d01e",
		Cmd: []string{strings.Join(append([]string{"anvil"}, commandArgs...), " ")},
		Labels: map[string]string{
			"test-id": uuid.New().String(),
		},
		ExposedPorts: []string{"8545"},
	}

	resource, err := pool.RunWithOptions(runOptions, func(config *docker.HostConfig) {
		config.AutoRemove = args.autoremove
		if args.restartPolicy != nil {
			config.RestartPolicy = *args.restartPolicy
		}
	})
	assert.Nil(t, err)

	logInfoChan := make(chan processlog.LogMetadata)
	go func() {
		defer close(logInfoChan)
		err = dockerutil.TailContainerLogs(dockerutil.WithContext(ctx), dockerutil.WithResource(resource), dockerutil.WithPool(pool), dockerutil.WithProcessLogOptions(args.processOptions...), dockerutil.WithFollow(true), dockerutil.WithCallback(func(ctx context.Context, metadata processlog.LogMetadata) {
			select {
			case <-ctx.Done():
				return
			case logInfoChan <- metadata:
			}
		}))

		if ctx.Err() != nil {
			logger.Warn(err)
		}
	}()

	// Docker will hard kill the container in expiryseconds seconds (this is a test env).
	// containers should be removed on their own, but this is a safety net.
	// to prevent old containers from piling up, we set a timeout to remove the container.
	assert.Nil(t, resource.Expire(args.expirySeconds))

	address := fmt.Sprintf("%s:%s", "http://localhost", resource.GetPort("8545/tcp"))

	var chainID *big.Int
	if err := pool.Retry(func() error {
		rpcClient, err := ethclient.DialContext(ctx, address)
		if err != nil {
			return fmt.Errorf("failed to connect")
		}
		chainID, err = rpcClient.ChainID(ctx)
		if err != nil {
			return fmt.Errorf("failed to get chain id: %w", err)
		}
		return nil
	}); err != nil {
		assert.Nil(t, err)
	}

	chainConfig := args.GetHardfork().ToChainConfig(chainID)

	chn, err := chain.New(ctx, &client.Config{
		RPCUrl:  []string{address},
		ChainID: int(chainConfig.ChainID.Int64()),
	})
	assert.Nilf(t, err, "failed to create chain for chain id %s: %v", chainID, err)

	chn.SetChainConfig(chainConfig)

	select {
	case <-ctx.Done():
		t.Errorf("context canceled before anvil node started")
	case logInfo := <-logInfoChan:
		logger.Warnf("started anvil node for chain %s as container %s. Logs will be stored at %s", chainID, strings.TrimPrefix(resource.Container.Name, "/"), logInfo.LogDir())
	}

	baseBackend, err := base.NewBaseBackend(ctx, t, chn)
	assert.Nil(t, err)

	backend := Backend{
		Backend:     baseBackend,
		fundingMux:  mapmutex.NewStringerMapMutex(),
		store:       base.NewInMemoryKeyStore(),
		chainConfig: chainConfig,
	}

	err = backend.storeWallets(args)
	assert.Nilf(t, err, "failed to store wallets on chain id %s: %v", chainID, err)

	go func() {
		<-ctx.Done()
		_ = pool.Purge(resource)
	}()

	return &backend
}

var logger = log.Logger("anvil-docker")

// storeWallets stores preseeded wallets w/ balances.
func (f *Backend) storeWallets(args *OptionBuilder) error {
	derivationPath := args.GetDerivationPath()
	derivIter := accounts.DefaultIterator(derivationPath)
	maxAccounts := args.GetAccounts()
	for i := 0; i < int(maxAccounts); i++ {
		account := derivIter()

		wall, err := wallet.FromSeedPhrase(args.GetMnemonic(), account)
		if err != nil {
			return fmt.Errorf("could not get seed phrase: %w", err)
		}

		f.store.Store(walletToKey(f.Backend.T(), wall))
	}
	return nil
}

// TODO(trajan0x): add a test for this.
func walletToKey(tb testing.TB, wall wallet.Wallet) *keystore.Key {
	tb.Helper()

	kstr := keystore.NewKeyStore(filet.TmpDir(tb, ""), base.VeryLightScryptN, base.VeryLightScryptP)
	password := gofakeit.Password(true, true, true, false, false, 10)

	acct, err := kstr.ImportECDSA(wall.PrivateKey(), password)
	assert.Nil(tb, err)

	data, err := os.ReadFile(acct.URL.Path)
	assert.Nil(tb, err)

	key, err := keystore.DecryptKey(data, password)
	assert.Nil(tb, err)
	return key
}

// ChainConfig gets the chain config.
func (f *Backend) ChainConfig() *params.ChainConfig {
	return f.chainConfig
}

// Signer gets the signer for the chain.
func (f *Backend) Signer() types.Signer {
	latestBlock, err := f.BlockNumber(f.Context())
	assert.Nil(f.T(), err)

	return types.MakeSigner(f.ChainConfig(), new(big.Int).SetUint64(latestBlock))
}

// FundAccount funds an account with the given amount.
func (f *Backend) FundAccount(ctx context.Context, address common.Address, amount big.Int) {
	ctx, cancel := onecontext.Merge(ctx, f.Context())
	defer cancel()

	anvilClient, err := Dial(ctx, f.RPCAddress())
	assert.Nilf(f.T(), err, "failed to dial anvil client on chain %d: %v", f.GetChainID(), err)

	unlocker := f.fundingMux.Lock(address)
	defer unlocker.Unlock()

	prevBalance, err := f.Backend.BalanceAt(ctx, address, nil)
	assert.Nil(f.T(), err)

	newBal := new(big.Int).Add(prevBalance, &amount)

	// TODO(trajan0x): this can be fixed by looping through new accounts and sending eth when over the limit
	// probably not worth it yet.
	if !newBal.IsUint64() {
		warnUint64Once.Do(func() {
			logger.Warn("new balance overflows uint64, which is not allowed by the rpc api, using max_uint64 instead. Future warnings will be suppressed.")
		})
		newBal = new(big.Int).SetUint64(math.MaxUint64)
	}

	// TODO: this may cause issues when newBal overflows uint64
	err = anvilClient.SetBalance(ctx, address, newBal.Uint64())
	assert.Nil(f.T(), err)
}

// WaitForConfirmation checks confirmation if the transaction is signed.
// nolint: cyclop
func (f *Backend) WaitForConfirmation(ctx context.Context, tx *types.Transaction) {
	assert.NotNil(f.T(), tx, "tx is nil")
	v, r, s := tx.RawSignatureValues()
	isUnsigned := isZero(v) && isZero(r) && isZero(s)
	if isUnsigned {
		warnUnsignedOnce.Do(func() {
			logger.Warn("WaitForConfirmation called on unsigned (likely impersonated) transaction, this does nothing")
		})
		return
	}

	f.Backend.WaitForConfirmation(ctx, tx)
}

func isZero(val *big.Int) bool {
	return val.Cmp(big.NewInt(0)) == 0
}

// GetFundedAccount gets a funded account.
func (f *Backend) GetFundedAccount(ctx context.Context, requestBalance *big.Int) *keystore.Key {
	key := f.MockAccount()

	f.store.Store(key)

	f.FundAccount(ctx, key.Address, *requestBalance)

	return key
}

// GetTxContext gets the tx context for the given address.
// TODO: dedupe w/ geth.
func (f *Backend) GetTxContext(ctx context.Context, address *common.Address) (res backends.AuthType) {
	ctx, cancel := onecontext.Merge(ctx, f.Context())
	defer cancel()

	var acct *keystore.Key
	// TODO handle storing accounts to conform to get tx context
	if address != nil {
		acct = f.store.GetAccount(*address)
		if acct == nil {
			f.T().Errorf("could not get account %s", address.String())
			return res
		}
	} else {
		acct = f.GetFundedAccount(ctx, new(big.Int).SetUint64(math.MaxUint64))
		f.store.Store(acct)
	}

	auth, err := f.NewKeyedTransactorFromKey(acct.PrivateKey)
	assert.Nilf(f.T(), err, "could not get transactor for chain %d: %v", f.GetChainID(), err)

	auth.GasPrice, err = f.SuggestGasPrice(ctx)
	assert.Nilf(f.T(), err, "could not get gas price for chain %d: %v", f.GetChainID(), err)

	auth.GasLimit = gasLimit

	return backends.AuthType{
		TransactOpts: auth,
		PrivateKey:   acct.PrivateKey,
	}
}

// ImpersonateAccount impersonates an account.
//
// Note *any* other call made to the backend will impersonate while this is being called
// in a future version, we'll wrap something like omnirpc to prevent other transactions submission calls from taking place
// in the meantime, this may cause race conditions.
//
// We also print a warning message to the console as an added precaution.
func (f *Backend) ImpersonateAccount(ctx context.Context, address common.Address, transact func(opts *bind.TransactOpts) *types.Transaction) error {
	f.impersonationMux.Lock()
	defer f.impersonationMux.Unlock()

	f.warnImpersonation()

	anvilClient, err := Dial(ctx, f.RPCAddress())
	assert.Nilf(f.T(), err, "could not dial anvil client rpc at %s for chain %d: %v", f.RPCAddress(), f.GetChainID(), err)

	err = anvilClient.ImpersonateAccount(ctx, address)
	assert.Nilf(f.T(), err, "could not impersonate account %s for chain %d: %v", address.String(), f.GetChainID(), err)

	defer func() {
		err = anvilClient.StopImpersonatingAccount(ctx, address)
		assert.Nilf(f.T(), err, "could not stop impersonating account %s for chain %d: %v", address.String(), f.GetChainID(), err)
	}()

	tx := transact(&bind.TransactOpts{
		Context:  ctx,
		From:     address,
		Signer:   ImpersonatedSigner,
		GasLimit: gasLimit,
		NoSend:   true,
	})

	err = anvilClient.SendUnsignedTransaction(ctx, address, tx)
	assert.Nilf(f.T(), err, "could not send unsigned transaction for chain %d: %v from %s", f.GetChainID(), err, address.String())

	return nil
}

func (f *Backend) warnImpersonation() {
	warnImpersonationOnce.Do(func() {
		f.T().Logf(`
				Using Account Impersonation.
				WARNING: This cannot be called concurrently with other impersonation calls.
				Please make sure your callers are concurrency safe against account impersonation.
				`)
	})
}

// ImpersonatedSigner is a signer that does nothing for use in account impersonation w/ contracts.
func ImpersonatedSigner(address common.Address, transaction *types.Transaction) (*types.Transaction, error) {
	return transaction, nil
}

// warnImpersonationOnce is used to log the impersonation warning message once.
// this is a global variable to prevent the message from being logged multiple times.
// normally, global variables are strongly discouraged, but we make an exception here
// considering how unexpected behavior can be if impersonate account is not used correctly.
var warnImpersonationOnce sync.Once

// warnUnsignedOnce warns if a tx is unsigned and thus not confirmable.
var warnUnsignedOnce sync.Once

var warnUint64Once sync.Once
var _ backends.SimulatedTestBackend = &Backend{}
