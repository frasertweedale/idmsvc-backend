// Code generated by mockery v2.33.2. DO NOT EDIT.

package handler

import (
	event "github.com/podengo-project/idmsvc-backend/internal/infrastructure/event"

	kafka "github.com/confluentinc/confluent-kafka-go/kafka"

	mock "github.com/stretchr/testify/mock"
)

// EventRouterHandler is an autogenerated mock type for the EventRouterHandler type
type EventRouterHandler struct {
	mock.Mock
}

// Add provides a mock function with given fields: topic, _a1
func (_m *EventRouterHandler) Add(topic string, _a1 event.Eventable) {
	_m.Called(topic, _a1)
}

// OnMessage provides a mock function with given fields: msg
func (_m *EventRouterHandler) OnMessage(msg *kafka.Message) error {
	ret := _m.Called(msg)

	var r0 error
	if rf, ok := ret.Get(0).(func(*kafka.Message) error); ok {
		r0 = rf(msg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewEventRouterHandler creates a new instance of EventRouterHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEventRouterHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *EventRouterHandler {
	mock := &EventRouterHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
