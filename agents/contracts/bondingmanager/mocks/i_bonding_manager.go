// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	big "math/big"

	bind "github.com/ethereum/go-ethereum/accounts/abi/bind"
	bondingmanager "github.com/synapsecns/sanguine/agents/contracts/bondingmanager"

	common "github.com/ethereum/go-ethereum/common"

	event "github.com/ethereum/go-ethereum/event"

	mock "github.com/stretchr/testify/mock"

	types "github.com/ethereum/go-ethereum/core/types"
)

// IBondingManager is an autogenerated mock type for the IBondingManager type
type IBondingManager struct {
	mock.Mock
}

// AddAgent provides a mock function with given fields: opts, domain, agent, proof
func (_m *IBondingManager) AddAgent(opts *bind.TransactOpts, domain uint32, agent common.Address, proof [][32]byte) (*types.Transaction, error) {
	ret := _m.Called(opts, domain, agent, proof)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, uint32, common.Address, [][32]byte) *types.Transaction); ok {
		r0 = rf(opts, domain, agent, proof)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, uint32, common.Address, [][32]byte) error); ok {
		r1 = rf(opts, domain, agent, proof)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Address provides a mock function with given fields:
func (_m *IBondingManager) Address() common.Address {
	ret := _m.Called()

	var r0 common.Address
	if rf, ok := ret.Get(0).(func() common.Address); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Address)
		}
	}

	return r0
}

// AgentLeaf provides a mock function with given fields: opts, agent
func (_m *IBondingManager) AgentLeaf(opts *bind.CallOpts, agent common.Address) ([32]byte, error) {
	ret := _m.Called(opts, agent)

	var r0 [32]byte
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, common.Address) [32]byte); ok {
		r0 = rf(opts, agent)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([32]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts, common.Address) error); ok {
		r1 = rf(opts, agent)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AgentRoot provides a mock function with given fields: opts
func (_m *IBondingManager) AgentRoot(opts *bind.CallOpts) ([32]byte, error) {
	ret := _m.Called(opts)

	var r0 [32]byte
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) [32]byte); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([32]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AgentStatus provides a mock function with given fields: opts, agent
func (_m *IBondingManager) AgentStatus(opts *bind.CallOpts, agent common.Address) (bondingmanager.AgentStatus, error) {
	ret := _m.Called(opts, agent)

	var r0 bondingmanager.AgentStatus
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, common.Address) bondingmanager.AgentStatus); ok {
		r0 = rf(opts, agent)
	} else {
		r0 = ret.Get(0).(bondingmanager.AgentStatus)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts, common.Address) error); ok {
		r1 = rf(opts, agent)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AllLeafs provides a mock function with given fields: opts
func (_m *IBondingManager) AllLeafs(opts *bind.CallOpts) ([][32]byte, error) {
	ret := _m.Called(opts)

	var r0 [][32]byte
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) [][32]byte); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([][32]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CompleteSlashing provides a mock function with given fields: opts, domain, agent, proof
func (_m *IBondingManager) CompleteSlashing(opts *bind.TransactOpts, domain uint32, agent common.Address, proof [][32]byte) (*types.Transaction, error) {
	ret := _m.Called(opts, domain, agent, proof)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, uint32, common.Address, [][32]byte) *types.Transaction); ok {
		r0 = rf(opts, domain, agent, proof)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, uint32, common.Address, [][32]byte) error); ok {
		r1 = rf(opts, domain, agent, proof)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CompleteUnstaking provides a mock function with given fields: opts, domain, agent, proof
func (_m *IBondingManager) CompleteUnstaking(opts *bind.TransactOpts, domain uint32, agent common.Address, proof [][32]byte) (*types.Transaction, error) {
	ret := _m.Called(opts, domain, agent, proof)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, uint32, common.Address, [][32]byte) *types.Transaction); ok {
		r0 = rf(opts, domain, agent, proof)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, uint32, common.Address, [][32]byte) error); ok {
		r1 = rf(opts, domain, agent, proof)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Destination provides a mock function with given fields: opts
func (_m *IBondingManager) Destination(opts *bind.CallOpts) (common.Address, error) {
	ret := _m.Called(opts)

	var r0 common.Address
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) common.Address); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Address)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DisputeStatus provides a mock function with given fields: opts, agent
func (_m *IBondingManager) DisputeStatus(opts *bind.CallOpts, agent common.Address) (struct {
	Flag        uint8
	Rival       common.Address
	FraudProver common.Address
	DisputePtr  *big.Int
}, error) {
	ret := _m.Called(opts, agent)

	var r0 struct {
		Flag        uint8
		Rival       common.Address
		FraudProver common.Address
		DisputePtr  *big.Int
	}
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, common.Address) struct {
		Flag        uint8
		Rival       common.Address
		FraudProver common.Address
		DisputePtr  *big.Int
	}); ok {
		r0 = rf(opts, agent)
	} else {
		r0 = ret.Get(0).(struct {
			Flag        uint8
			Rival       common.Address
			FraudProver common.Address
			DisputePtr  *big.Int
		})
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts, common.Address) error); ok {
		r1 = rf(opts, agent)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterDisputeOpened provides a mock function with given fields: opts
func (_m *IBondingManager) FilterDisputeOpened(opts *bind.FilterOpts) (*bondingmanager.BondingManagerDisputeOpenedIterator, error) {
	ret := _m.Called(opts)

	var r0 *bondingmanager.BondingManagerDisputeOpenedIterator
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts) *bondingmanager.BondingManagerDisputeOpenedIterator); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bondingmanager.BondingManagerDisputeOpenedIterator)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.FilterOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterDisputeResolved provides a mock function with given fields: opts
func (_m *IBondingManager) FilterDisputeResolved(opts *bind.FilterOpts) (*bondingmanager.BondingManagerDisputeResolvedIterator, error) {
	ret := _m.Called(opts)

	var r0 *bondingmanager.BondingManagerDisputeResolvedIterator
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts) *bondingmanager.BondingManagerDisputeResolvedIterator); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bondingmanager.BondingManagerDisputeResolvedIterator)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.FilterOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterInitialized provides a mock function with given fields: opts
func (_m *IBondingManager) FilterInitialized(opts *bind.FilterOpts) (*bondingmanager.BondingManagerInitializedIterator, error) {
	ret := _m.Called(opts)

	var r0 *bondingmanager.BondingManagerInitializedIterator
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts) *bondingmanager.BondingManagerInitializedIterator); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bondingmanager.BondingManagerInitializedIterator)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.FilterOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterOwnershipTransferred provides a mock function with given fields: opts, previousOwner, newOwner
func (_m *IBondingManager) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*bondingmanager.BondingManagerOwnershipTransferredIterator, error) {
	ret := _m.Called(opts, previousOwner, newOwner)

	var r0 *bondingmanager.BondingManagerOwnershipTransferredIterator
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts, []common.Address, []common.Address) *bondingmanager.BondingManagerOwnershipTransferredIterator); ok {
		r0 = rf(opts, previousOwner, newOwner)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bondingmanager.BondingManagerOwnershipTransferredIterator)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.FilterOpts, []common.Address, []common.Address) error); ok {
		r1 = rf(opts, previousOwner, newOwner)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterRootUpdated provides a mock function with given fields: opts
func (_m *IBondingManager) FilterRootUpdated(opts *bind.FilterOpts) (*bondingmanager.BondingManagerRootUpdatedIterator, error) {
	ret := _m.Called(opts)

	var r0 *bondingmanager.BondingManagerRootUpdatedIterator
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts) *bondingmanager.BondingManagerRootUpdatedIterator); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bondingmanager.BondingManagerRootUpdatedIterator)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.FilterOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterStatusUpdated provides a mock function with given fields: opts, domain, agent
func (_m *IBondingManager) FilterStatusUpdated(opts *bind.FilterOpts, domain []uint32, agent []common.Address) (*bondingmanager.BondingManagerStatusUpdatedIterator, error) {
	ret := _m.Called(opts, domain, agent)

	var r0 *bondingmanager.BondingManagerStatusUpdatedIterator
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts, []uint32, []common.Address) *bondingmanager.BondingManagerStatusUpdatedIterator); ok {
		r0 = rf(opts, domain, agent)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bondingmanager.BondingManagerStatusUpdatedIterator)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.FilterOpts, []uint32, []common.Address) error); ok {
		r1 = rf(opts, domain, agent)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetActiveAgents provides a mock function with given fields: opts, domain
func (_m *IBondingManager) GetActiveAgents(opts *bind.CallOpts, domain uint32) ([]common.Address, error) {
	ret := _m.Called(opts, domain)

	var r0 []common.Address
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, uint32) []common.Address); ok {
		r0 = rf(opts, domain)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]common.Address)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts, uint32) error); ok {
		r1 = rf(opts, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAgent provides a mock function with given fields: opts, index
func (_m *IBondingManager) GetAgent(opts *bind.CallOpts, index *big.Int) (struct {
	Agent  common.Address
	Status bondingmanager.AgentStatus
}, error) {
	ret := _m.Called(opts, index)

	var r0 struct {
		Agent  common.Address
		Status bondingmanager.AgentStatus
	}
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, *big.Int) struct {
		Agent  common.Address
		Status bondingmanager.AgentStatus
	}); ok {
		r0 = rf(opts, index)
	} else {
		r0 = ret.Get(0).(struct {
			Agent  common.Address
			Status bondingmanager.AgentStatus
		})
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts, *big.Int) error); ok {
		r1 = rf(opts, index)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDispute provides a mock function with given fields: opts, index
func (_m *IBondingManager) GetDispute(opts *bind.CallOpts, index *big.Int) (struct {
	Guard           common.Address
	Notary          common.Address
	SlashedAgent    common.Address
	FraudProver     common.Address
	ReportPayload   []byte
	ReportSignature []byte
}, error) {
	ret := _m.Called(opts, index)

	var r0 struct {
		Guard           common.Address
		Notary          common.Address
		SlashedAgent    common.Address
		FraudProver     common.Address
		ReportPayload   []byte
		ReportSignature []byte
	}
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, *big.Int) struct {
		Guard           common.Address
		Notary          common.Address
		SlashedAgent    common.Address
		FraudProver     common.Address
		ReportPayload   []byte
		ReportSignature []byte
	}); ok {
		r0 = rf(opts, index)
	} else {
		r0 = ret.Get(0).(struct {
			Guard           common.Address
			Notary          common.Address
			SlashedAgent    common.Address
			FraudProver     common.Address
			ReportPayload   []byte
			ReportSignature []byte
		})
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts, *big.Int) error); ok {
		r1 = rf(opts, index)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDisputesAmount provides a mock function with given fields: opts
func (_m *IBondingManager) GetDisputesAmount(opts *bind.CallOpts) (*big.Int, error) {
	ret := _m.Called(opts)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) *big.Int); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLeafs provides a mock function with given fields: opts, indexFrom, amount
func (_m *IBondingManager) GetLeafs(opts *bind.CallOpts, indexFrom *big.Int, amount *big.Int) ([][32]byte, error) {
	ret := _m.Called(opts, indexFrom, amount)

	var r0 [][32]byte
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, *big.Int, *big.Int) [][32]byte); ok {
		r0 = rf(opts, indexFrom, amount)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([][32]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts, *big.Int, *big.Int) error); ok {
		r1 = rf(opts, indexFrom, amount)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProof provides a mock function with given fields: opts, agent
func (_m *IBondingManager) GetProof(opts *bind.CallOpts, agent common.Address) ([][32]byte, error) {
	ret := _m.Called(opts, agent)

	var r0 [][32]byte
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, common.Address) [][32]byte); ok {
		r0 = rf(opts, agent)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([][32]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts, common.Address) error); ok {
		r1 = rf(opts, agent)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Inbox provides a mock function with given fields: opts
func (_m *IBondingManager) Inbox(opts *bind.CallOpts) (common.Address, error) {
	ret := _m.Called(opts)

	var r0 common.Address
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) common.Address); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Address)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Initialize provides a mock function with given fields: opts, origin_, destination_, inbox_, summit_
func (_m *IBondingManager) Initialize(opts *bind.TransactOpts, origin_ common.Address, destination_ common.Address, inbox_ common.Address, summit_ common.Address) (*types.Transaction, error) {
	ret := _m.Called(opts, origin_, destination_, inbox_, summit_)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, common.Address, common.Address, common.Address, common.Address) *types.Transaction); ok {
		r0 = rf(opts, origin_, destination_, inbox_, summit_)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, common.Address, common.Address, common.Address, common.Address) error); ok {
		r1 = rf(opts, origin_, destination_, inbox_, summit_)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InitiateUnstaking provides a mock function with given fields: opts, domain, agent, proof
func (_m *IBondingManager) InitiateUnstaking(opts *bind.TransactOpts, domain uint32, agent common.Address, proof [][32]byte) (*types.Transaction, error) {
	ret := _m.Called(opts, domain, agent, proof)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, uint32, common.Address, [][32]byte) *types.Transaction); ok {
		r0 = rf(opts, domain, agent, proof)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, uint32, common.Address, [][32]byte) error); ok {
		r1 = rf(opts, domain, agent, proof)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LeafsAmount provides a mock function with given fields: opts
func (_m *IBondingManager) LeafsAmount(opts *bind.CallOpts) (*big.Int, error) {
	ret := _m.Called(opts)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) *big.Int); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LocalDomain provides a mock function with given fields: opts
func (_m *IBondingManager) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	ret := _m.Called(opts)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) uint32); ok {
		r0 = rf(opts)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Multicall provides a mock function with given fields: opts, calls
func (_m *IBondingManager) Multicall(opts *bind.TransactOpts, calls []bondingmanager.MultiCallableCall) (*types.Transaction, error) {
	ret := _m.Called(opts, calls)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, []bondingmanager.MultiCallableCall) *types.Transaction); ok {
		r0 = rf(opts, calls)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, []bondingmanager.MultiCallableCall) error); ok {
		r1 = rf(opts, calls)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OpenDispute provides a mock function with given fields: opts, guardIndex, notaryIndex
func (_m *IBondingManager) OpenDispute(opts *bind.TransactOpts, guardIndex uint32, notaryIndex uint32) (*types.Transaction, error) {
	ret := _m.Called(opts, guardIndex, notaryIndex)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, uint32, uint32) *types.Transaction); ok {
		r0 = rf(opts, guardIndex, notaryIndex)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, uint32, uint32) error); ok {
		r1 = rf(opts, guardIndex, notaryIndex)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Origin provides a mock function with given fields: opts
func (_m *IBondingManager) Origin(opts *bind.CallOpts) (common.Address, error) {
	ret := _m.Called(opts)

	var r0 common.Address
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) common.Address); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Address)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Owner provides a mock function with given fields: opts
func (_m *IBondingManager) Owner(opts *bind.CallOpts) (common.Address, error) {
	ret := _m.Called(opts)

	var r0 common.Address
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) common.Address); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Address)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseDisputeOpened provides a mock function with given fields: log
func (_m *IBondingManager) ParseDisputeOpened(log types.Log) (*bondingmanager.BondingManagerDisputeOpened, error) {
	ret := _m.Called(log)

	var r0 *bondingmanager.BondingManagerDisputeOpened
	if rf, ok := ret.Get(0).(func(types.Log) *bondingmanager.BondingManagerDisputeOpened); ok {
		r0 = rf(log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bondingmanager.BondingManagerDisputeOpened)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseDisputeResolved provides a mock function with given fields: log
func (_m *IBondingManager) ParseDisputeResolved(log types.Log) (*bondingmanager.BondingManagerDisputeResolved, error) {
	ret := _m.Called(log)

	var r0 *bondingmanager.BondingManagerDisputeResolved
	if rf, ok := ret.Get(0).(func(types.Log) *bondingmanager.BondingManagerDisputeResolved); ok {
		r0 = rf(log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bondingmanager.BondingManagerDisputeResolved)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseInitialized provides a mock function with given fields: log
func (_m *IBondingManager) ParseInitialized(log types.Log) (*bondingmanager.BondingManagerInitialized, error) {
	ret := _m.Called(log)

	var r0 *bondingmanager.BondingManagerInitialized
	if rf, ok := ret.Get(0).(func(types.Log) *bondingmanager.BondingManagerInitialized); ok {
		r0 = rf(log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bondingmanager.BondingManagerInitialized)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseOwnershipTransferred provides a mock function with given fields: log
func (_m *IBondingManager) ParseOwnershipTransferred(log types.Log) (*bondingmanager.BondingManagerOwnershipTransferred, error) {
	ret := _m.Called(log)

	var r0 *bondingmanager.BondingManagerOwnershipTransferred
	if rf, ok := ret.Get(0).(func(types.Log) *bondingmanager.BondingManagerOwnershipTransferred); ok {
		r0 = rf(log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bondingmanager.BondingManagerOwnershipTransferred)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseRootUpdated provides a mock function with given fields: log
func (_m *IBondingManager) ParseRootUpdated(log types.Log) (*bondingmanager.BondingManagerRootUpdated, error) {
	ret := _m.Called(log)

	var r0 *bondingmanager.BondingManagerRootUpdated
	if rf, ok := ret.Get(0).(func(types.Log) *bondingmanager.BondingManagerRootUpdated); ok {
		r0 = rf(log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bondingmanager.BondingManagerRootUpdated)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseStatusUpdated provides a mock function with given fields: log
func (_m *IBondingManager) ParseStatusUpdated(log types.Log) (*bondingmanager.BondingManagerStatusUpdated, error) {
	ret := _m.Called(log)

	var r0 *bondingmanager.BondingManagerStatusUpdated
	if rf, ok := ret.Get(0).(func(types.Log) *bondingmanager.BondingManagerStatusUpdated); ok {
		r0 = rf(log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bondingmanager.BondingManagerStatusUpdated)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoteSlashAgent provides a mock function with given fields: opts, msgOrigin, proofMaturity, domain, agent, prover
func (_m *IBondingManager) RemoteSlashAgent(opts *bind.TransactOpts, msgOrigin uint32, proofMaturity *big.Int, domain uint32, agent common.Address, prover common.Address) (*types.Transaction, error) {
	ret := _m.Called(opts, msgOrigin, proofMaturity, domain, agent, prover)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, uint32, *big.Int, uint32, common.Address, common.Address) *types.Transaction); ok {
		r0 = rf(opts, msgOrigin, proofMaturity, domain, agent, prover)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, uint32, *big.Int, uint32, common.Address, common.Address) error); ok {
		r1 = rf(opts, msgOrigin, proofMaturity, domain, agent, prover)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RenounceOwnership provides a mock function with given fields: opts
func (_m *IBondingManager) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	ret := _m.Called(opts)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts) *types.Transaction); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.TransactOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResolveStuckDispute provides a mock function with given fields: opts, domain, slashedAgent
func (_m *IBondingManager) ResolveStuckDispute(opts *bind.TransactOpts, domain uint32, slashedAgent common.Address) (*types.Transaction, error) {
	ret := _m.Called(opts, domain, slashedAgent)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, uint32, common.Address) *types.Transaction); ok {
		r0 = rf(opts, domain, slashedAgent)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, uint32, common.Address) error); ok {
		r1 = rf(opts, domain, slashedAgent)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SlashAgent provides a mock function with given fields: opts, domain, agent, prover
func (_m *IBondingManager) SlashAgent(opts *bind.TransactOpts, domain uint32, agent common.Address, prover common.Address) (*types.Transaction, error) {
	ret := _m.Called(opts, domain, agent, prover)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, uint32, common.Address, common.Address) *types.Transaction); ok {
		r0 = rf(opts, domain, agent, prover)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, uint32, common.Address, common.Address) error); ok {
		r1 = rf(opts, domain, agent, prover)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Summit provides a mock function with given fields: opts
func (_m *IBondingManager) Summit(opts *bind.CallOpts) (common.Address, error) {
	ret := _m.Called(opts)

	var r0 common.Address
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) common.Address); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Address)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TransferOwnership provides a mock function with given fields: opts, newOwner
func (_m *IBondingManager) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	ret := _m.Called(opts, newOwner)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, common.Address) *types.Transaction); ok {
		r0 = rf(opts, newOwner)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, common.Address) error); ok {
		r1 = rf(opts, newOwner)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Version provides a mock function with given fields: opts
func (_m *IBondingManager) Version(opts *bind.CallOpts) (string, error) {
	ret := _m.Called(opts)

	var r0 string
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) string); ok {
		r0 = rf(opts)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchDisputeOpened provides a mock function with given fields: opts, sink
func (_m *IBondingManager) WatchDisputeOpened(opts *bind.WatchOpts, sink chan<- *bondingmanager.BondingManagerDisputeOpened) (event.Subscription, error) {
	ret := _m.Called(opts, sink)

	var r0 event.Subscription
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *bondingmanager.BondingManagerDisputeOpened) event.Subscription); ok {
		r0 = rf(opts, sink)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(event.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.WatchOpts, chan<- *bondingmanager.BondingManagerDisputeOpened) error); ok {
		r1 = rf(opts, sink)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchDisputeResolved provides a mock function with given fields: opts, sink
func (_m *IBondingManager) WatchDisputeResolved(opts *bind.WatchOpts, sink chan<- *bondingmanager.BondingManagerDisputeResolved) (event.Subscription, error) {
	ret := _m.Called(opts, sink)

	var r0 event.Subscription
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *bondingmanager.BondingManagerDisputeResolved) event.Subscription); ok {
		r0 = rf(opts, sink)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(event.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.WatchOpts, chan<- *bondingmanager.BondingManagerDisputeResolved) error); ok {
		r1 = rf(opts, sink)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchInitialized provides a mock function with given fields: opts, sink
func (_m *IBondingManager) WatchInitialized(opts *bind.WatchOpts, sink chan<- *bondingmanager.BondingManagerInitialized) (event.Subscription, error) {
	ret := _m.Called(opts, sink)

	var r0 event.Subscription
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *bondingmanager.BondingManagerInitialized) event.Subscription); ok {
		r0 = rf(opts, sink)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(event.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.WatchOpts, chan<- *bondingmanager.BondingManagerInitialized) error); ok {
		r1 = rf(opts, sink)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchOwnershipTransferred provides a mock function with given fields: opts, sink, previousOwner, newOwner
func (_m *IBondingManager) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *bondingmanager.BondingManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {
	ret := _m.Called(opts, sink, previousOwner, newOwner)

	var r0 event.Subscription
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *bondingmanager.BondingManagerOwnershipTransferred, []common.Address, []common.Address) event.Subscription); ok {
		r0 = rf(opts, sink, previousOwner, newOwner)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(event.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.WatchOpts, chan<- *bondingmanager.BondingManagerOwnershipTransferred, []common.Address, []common.Address) error); ok {
		r1 = rf(opts, sink, previousOwner, newOwner)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchRootUpdated provides a mock function with given fields: opts, sink
func (_m *IBondingManager) WatchRootUpdated(opts *bind.WatchOpts, sink chan<- *bondingmanager.BondingManagerRootUpdated) (event.Subscription, error) {
	ret := _m.Called(opts, sink)

	var r0 event.Subscription
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *bondingmanager.BondingManagerRootUpdated) event.Subscription); ok {
		r0 = rf(opts, sink)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(event.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.WatchOpts, chan<- *bondingmanager.BondingManagerRootUpdated) error); ok {
		r1 = rf(opts, sink)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchStatusUpdated provides a mock function with given fields: opts, sink, domain, agent
func (_m *IBondingManager) WatchStatusUpdated(opts *bind.WatchOpts, sink chan<- *bondingmanager.BondingManagerStatusUpdated, domain []uint32, agent []common.Address) (event.Subscription, error) {
	ret := _m.Called(opts, sink, domain, agent)

	var r0 event.Subscription
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *bondingmanager.BondingManagerStatusUpdated, []uint32, []common.Address) event.Subscription); ok {
		r0 = rf(opts, sink, domain, agent)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(event.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.WatchOpts, chan<- *bondingmanager.BondingManagerStatusUpdated, []uint32, []common.Address) error); ok {
		r1 = rf(opts, sink, domain, agent)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WithdrawTips provides a mock function with given fields: opts, recipient, origin_, amount
func (_m *IBondingManager) WithdrawTips(opts *bind.TransactOpts, recipient common.Address, origin_ uint32, amount *big.Int) (*types.Transaction, error) {
	ret := _m.Called(opts, recipient, origin_, amount)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, common.Address, uint32, *big.Int) *types.Transaction); ok {
		r0 = rf(opts, recipient, origin_, amount)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, common.Address, uint32, *big.Int) error); ok {
		r1 = rf(opts, recipient, origin_, amount)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewIBondingManager interface {
	mock.TestingT
	Cleanup(func())
}

// NewIBondingManager creates a new instance of IBondingManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIBondingManager(t mockConstructorTestingTNewIBondingManager) *IBondingManager {
	mock := &IBondingManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
