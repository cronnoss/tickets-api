// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/cronnoss/tickets-api/internal/app/domain"

	mock "github.com/stretchr/testify/mock"
)

// EventService is an autogenerated mock type for the EventService type
type EventService struct {
	mock.Mock
}

type EventService_Expecter struct {
	mock *mock.Mock
}

func (_m *EventService) EXPECT() *EventService_Expecter {
	return &EventService_Expecter{mock: &_m.Mock}
}

// CreateEvent provides a mock function with given fields: ctx, event
func (_m *EventService) CreateEvent(ctx context.Context, event domain.Event) (domain.Event, error) {
	ret := _m.Called(ctx, event)

	if len(ret) == 0 {
		panic("no return value specified for CreateEvent")
	}

	var r0 domain.Event
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Event) (domain.Event, error)); ok {
		return rf(ctx, event)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.Event) domain.Event); ok {
		r0 = rf(ctx, event)
	} else {
		r0 = ret.Get(0).(domain.Event)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.Event) error); ok {
		r1 = rf(ctx, event)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EventService_CreateEvent_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateEvent'
type EventService_CreateEvent_Call struct {
	*mock.Call
}

// CreateEvent is a helper method to define mock.On call
//   - ctx context.Context
//   - event domain.Event
func (_e *EventService_Expecter) CreateEvent(ctx interface{}, event interface{}) *EventService_CreateEvent_Call {
	return &EventService_CreateEvent_Call{Call: _e.mock.On("CreateEvent", ctx, event)}
}

func (_c *EventService_CreateEvent_Call) Run(run func(ctx context.Context, event domain.Event)) *EventService_CreateEvent_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.Event))
	})
	return _c
}

func (_c *EventService_CreateEvent_Call) Return(_a0 domain.Event, _a1 error) *EventService_CreateEvent_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *EventService_CreateEvent_Call) RunAndReturn(run func(context.Context, domain.Event) (domain.Event, error)) *EventService_CreateEvent_Call {
	_c.Call.Return(run)
	return _c
}

// GetEvents provides a mock function with given fields: ctx
func (_m *EventService) GetEvents(ctx context.Context) ([]domain.Event, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetEvents")
	}

	var r0 []domain.Event
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]domain.Event, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []domain.Event); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Event)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EventService_GetEvents_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetEvents'
type EventService_GetEvents_Call struct {
	*mock.Call
}

// GetEvents is a helper method to define mock.On call
//   - ctx context.Context
func (_e *EventService_Expecter) GetEvents(ctx interface{}) *EventService_GetEvents_Call {
	return &EventService_GetEvents_Call{Call: _e.mock.On("GetEvents", ctx)}
}

func (_c *EventService_GetEvents_Call) Run(run func(ctx context.Context)) *EventService_GetEvents_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *EventService_GetEvents_Call) Return(_a0 []domain.Event, _a1 error) *EventService_GetEvents_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *EventService_GetEvents_Call) RunAndReturn(run func(context.Context) ([]domain.Event, error)) *EventService_GetEvents_Call {
	_c.Call.Return(run)
	return _c
}

// NewEventService creates a new instance of EventService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEventService(t interface {
	mock.TestingT
	Cleanup(func())
}) *EventService {
	mock := &EventService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
