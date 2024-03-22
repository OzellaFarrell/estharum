// autogenerated file

package inbox

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// IInboxCaller ...
type IInboxCaller interface {
	// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
	//
	// Solidity: function agentManager() view returns(address)
	AgentManager(opts *bind.CallOpts) (common.Address, error)
	// Destination is a free data retrieval call binding the contract method 0xb269681d.
	//
	// Solidity: function destination() view returns(address)
	Destination(opts *bind.CallOpts) (common.Address, error)
	// GetGuardReport is a free data retrieval call binding the contract method 0xc495912b.
	//
	// Solidity: function getGuardReport(uint256 index) view returns(bytes statementPayload, bytes reportSignature)
	GetGuardReport(opts *bind.CallOpts, index *big.Int) (struct {
		StatementPayload []byte
		ReportSignature  []byte
	}, error)
	// GetReportsAmount is a free data retrieval call binding the contract method 0x756ed01d.
	//
	// Solidity: function getReportsAmount() view returns(uint256)
	GetReportsAmount(opts *bind.CallOpts) (*big.Int, error)
	// GetStoredSignature is a free data retrieval call binding the contract method 0xddeffa66.
	//
	// Solidity: function getStoredSignature(uint256 index) view returns(bytes)
	GetStoredSignature(opts *bind.CallOpts, index *big.Int) ([]byte, error)
	// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
	//
	// Solidity: function localDomain() view returns(uint32)
	LocalDomain(opts *bind.CallOpts) (uint32, error)
	// Origin is a free data retrieval call binding the contract method 0x938b5f32.
	//
	// Solidity: function origin() view returns(address)
	Origin(opts *bind.CallOpts) (common.Address, error)
	// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
	//
	// Solidity: function owner() view returns(address)
	Owner(opts *bind.CallOpts) (common.Address, error)
	// Summit is a free data retrieval call binding the contract method 0x9fbcb9cb.
	//
	// Solidity: function summit() view returns(address)
	Summit(opts *bind.CallOpts) (common.Address, error)
	// Version is a free data retrieval call binding the contract method 0x54fd4d50.
	//
	// Solidity: function version() view returns(string versionString)
	Version(opts *bind.CallOpts) (string, error)
}
