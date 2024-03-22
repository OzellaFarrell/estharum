// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"
	big "math/big"

	client "github.com/synapsecns/sanguine/ethergo/client"

	mock "github.com/stretchr/testify/mock"
)

// ClientFetcher is an autogenerated mock type for the ClientFetcher type
type ClientFetcher struct {
	mock.Mock
}

// GetClient provides a mock function with given fields: ctx, chainID
func (_m *ClientFetcher) GetClient(ctx context.Context, chainID *big.Int) (client.EVM, error) {
	ret := _m.Called(ctx, chainID)

	var r0 client.EVM
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int) client.EVM); ok {
		r0 = rf(ctx, chainID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(client.EVM)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *big.Int) error); ok {
		r1 = rf(ctx, chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewClientFetcher interface {
	mock.TestingT
	Cleanup(func())
}

// NewClientFetcher creates a new instance of ClientFetcher. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewClientFetcher(t mockConstructorTestingTNewClientFetcher) *ClientFetcher {
	mock := &ClientFetcher{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}