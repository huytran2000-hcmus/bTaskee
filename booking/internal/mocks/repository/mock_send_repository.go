// Code generated by mockery v2.42.2. DO NOT EDIT.

package repository

import (
	context "context"

	model "github.com/huytran2000-hcmus/bTaskee/booking/internal/model"
	mock "github.com/stretchr/testify/mock"
)

// MockSendRepository is an autogenerated mock type for the SendRepository type
type MockSendRepository struct {
	mock.Mock
}

type MockSendRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSendRepository) EXPECT() *MockSendRepository_Expecter {
	return &MockSendRepository_Expecter{mock: &_m.Mock}
}

// SendTask provides a mock function with given fields: ctx, req
func (_m *MockSendRepository) SendTask(ctx context.Context, req model.SendTaskRequest) (model.SendTaskResponse, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for SendTask")
	}

	var r0 model.SendTaskResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.SendTaskRequest) (model.SendTaskResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.SendTaskRequest) model.SendTaskResponse); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(model.SendTaskResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.SendTaskRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSendRepository_SendTask_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendTask'
type MockSendRepository_SendTask_Call struct {
	*mock.Call
}

// SendTask is a helper method to define mock.On call
//   - ctx context.Context
//   - req model.SendTaskRequest
func (_e *MockSendRepository_Expecter) SendTask(ctx interface{}, req interface{}) *MockSendRepository_SendTask_Call {
	return &MockSendRepository_SendTask_Call{Call: _e.mock.On("SendTask", ctx, req)}
}

func (_c *MockSendRepository_SendTask_Call) Run(run func(ctx context.Context, req model.SendTaskRequest)) *MockSendRepository_SendTask_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.SendTaskRequest))
	})
	return _c
}

func (_c *MockSendRepository_SendTask_Call) Return(_a0 model.SendTaskResponse, _a1 error) *MockSendRepository_SendTask_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSendRepository_SendTask_Call) RunAndReturn(run func(context.Context, model.SendTaskRequest) (model.SendTaskResponse, error)) *MockSendRepository_SendTask_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockSendRepository creates a new instance of MockSendRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSendRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSendRepository {
	mock := &MockSendRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}