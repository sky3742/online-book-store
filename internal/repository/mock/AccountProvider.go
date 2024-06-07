// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"
	model "online-book-store/internal/model"

	mock "github.com/stretchr/testify/mock"
)

// AccountProvider is an autogenerated mock type for the AccountProvider type
type AccountProvider struct {
	mock.Mock
}

// CreateAccount provides a mock function with given fields: ctx, account
func (_m *AccountProvider) CreateAccount(ctx context.Context, account *model.Account) (*model.Account, error) {
	ret := _m.Called(ctx, account)

	if len(ret) == 0 {
		panic("no return value specified for CreateAccount")
	}

	var r0 *model.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Account) (*model.Account, error)); ok {
		return rf(ctx, account)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.Account) *model.Account); ok {
		r0 = rf(ctx, account)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Account)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.Account) error); ok {
		r1 = rf(ctx, account)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAccount provides a mock function with given fields: ctx, account
func (_m *AccountProvider) GetAccount(ctx context.Context, account *model.Account) (*model.Account, error) {
	ret := _m.Called(ctx, account)

	if len(ret) == 0 {
		panic("no return value specified for GetAccount")
	}

	var r0 *model.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Account) (*model.Account, error)); ok {
		return rf(ctx, account)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.Account) *model.Account); ok {
		r0 = rf(ctx, account)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Account)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.Account) error); ok {
		r1 = rf(ctx, account)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAccountProvider creates a new instance of AccountProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAccountProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *AccountProvider {
	mock := &AccountProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
