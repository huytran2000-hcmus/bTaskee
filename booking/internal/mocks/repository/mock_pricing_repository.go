// Code generated by mockery v2.42.2. DO NOT EDIT.

package repository

import (
	context "context"

	model "github.com/huytran2000-hcmus/bTaskee/booking/internal/model"
	mock "github.com/stretchr/testify/mock"
)

// MockPricingRepository is an autogenerated mock type for the PricingRepository type
type MockPricingRepository struct {
	mock.Mock
}

type MockPricingRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPricingRepository) EXPECT() *MockPricingRepository_Expecter {
	return &MockPricingRepository_Expecter{mock: &_m.Mock}
}

// GetPricing provides a mock function with given fields: ctx, req
func (_m *MockPricingRepository) GetPricing(ctx context.Context, req model.CalculatePriceRequest) (model.CalculatePriceResponse, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for GetPricing")
	}

	var r0 model.CalculatePriceResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.CalculatePriceRequest) (model.CalculatePriceResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.CalculatePriceRequest) model.CalculatePriceResponse); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(model.CalculatePriceResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.CalculatePriceRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPricingRepository_GetPricing_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPricing'
type MockPricingRepository_GetPricing_Call struct {
	*mock.Call
}

// GetPricing is a helper method to define mock.On call
//   - ctx context.Context
//   - req model.CalculatePriceRequest
func (_e *MockPricingRepository_Expecter) GetPricing(ctx interface{}, req interface{}) *MockPricingRepository_GetPricing_Call {
	return &MockPricingRepository_GetPricing_Call{Call: _e.mock.On("GetPricing", ctx, req)}
}

func (_c *MockPricingRepository_GetPricing_Call) Run(run func(ctx context.Context, req model.CalculatePriceRequest)) *MockPricingRepository_GetPricing_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.CalculatePriceRequest))
	})
	return _c
}

func (_c *MockPricingRepository_GetPricing_Call) Return(_a0 model.CalculatePriceResponse, _a1 error) *MockPricingRepository_GetPricing_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPricingRepository_GetPricing_Call) RunAndReturn(run func(context.Context, model.CalculatePriceRequest) (model.CalculatePriceResponse, error)) *MockPricingRepository_GetPricing_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPricingRepository creates a new instance of MockPricingRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPricingRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPricingRepository {
	mock := &MockPricingRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}