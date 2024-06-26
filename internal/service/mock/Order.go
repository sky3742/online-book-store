// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"
	model "online-book-store/internal/model"

	mock "github.com/stretchr/testify/mock"
)

// Order is an autogenerated mock type for the Order type
type Order struct {
	mock.Mock
}

// CreateOrder provides a mock function with given fields: ctx, userID, order
func (_m *Order) CreateOrder(ctx context.Context, userID string, order *model.Order) (*model.Order, *model.AppError) {
	ret := _m.Called(ctx, userID, order)

	if len(ret) == 0 {
		panic("no return value specified for CreateOrder")
	}

	var r0 *model.Order
	var r1 *model.AppError
	if rf, ok := ret.Get(0).(func(context.Context, string, *model.Order) (*model.Order, *model.AppError)); ok {
		return rf(ctx, userID, order)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *model.Order) *model.Order); ok {
		r0 = rf(ctx, userID, order)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Order)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *model.Order) *model.AppError); ok {
		r1 = rf(ctx, userID, order)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetOrder provides a mock function with given fields: ctx, userID, orderNo
func (_m *Order) GetOrder(ctx context.Context, userID string, orderNo string) (*model.Order, *model.AppError) {
	ret := _m.Called(ctx, userID, orderNo)

	if len(ret) == 0 {
		panic("no return value specified for GetOrder")
	}

	var r0 *model.Order
	var r1 *model.AppError
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*model.Order, *model.AppError)); ok {
		return rf(ctx, userID, orderNo)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *model.Order); ok {
		r0 = rf(ctx, userID, orderNo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Order)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) *model.AppError); ok {
		r1 = rf(ctx, userID, orderNo)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetOrders provides a mock function with given fields: ctx, userID
func (_m *Order) GetOrders(ctx context.Context, userID string) ([]model.Order, *model.AppError) {
	ret := _m.Called(ctx, userID)

	if len(ret) == 0 {
		panic("no return value specified for GetOrders")
	}

	var r0 []model.Order
	var r1 *model.AppError
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]model.Order, *model.AppError)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []model.Order); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Order)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) *model.AppError); ok {
		r1 = rf(ctx, userID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// NewOrder creates a new instance of Order. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOrder(t interface {
	mock.TestingT
	Cleanup(func())
}) *Order {
	mock := &Order{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
