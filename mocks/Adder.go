// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Adder is an autogenerated mock type for the Adder type
type Adder struct {
	mock.Mock
}

// Add provides a mock function with given fields: x, y
func (_m *Adder) Add(x float64, y float64) float64 {
	ret := _m.Called(x, y)

	if len(ret) == 0 {
		panic("no return value specified for Add")
	}

	var r0 float64
	if rf, ok := ret.Get(0).(func(float64, float64) float64); ok {
		r0 = rf(x, y)
	} else {
		r0 = ret.Get(0).(float64)
	}

	return r0
}

// NewAdder creates a new instance of Adder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAdder(t interface {
	mock.TestingT
	Cleanup(func())
}) *Adder {
	mock := &Adder{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
