// Code generated by mockery v2.42.2. DO NOT EDIT.

package processor

import (
	scan "marcuson/scriptman/internal/script/internal/scan"

	mock "github.com/stretchr/testify/mock"
)

// MockProcessor is an autogenerated mock type for the Processor type
type MockProcessor struct {
	mock.Mock
}

type MockProcessor_Expecter struct {
	mock *mock.Mock
}

func (_m *MockProcessor) EXPECT() *MockProcessor_Expecter {
	return &MockProcessor_Expecter{mock: &_m.Mock}
}

// IsProcessCompletedEarly provides a mock function with given fields:
func (_m *MockProcessor) IsProcessCompletedEarly() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsProcessCompletedEarly")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockProcessor_IsProcessCompletedEarly_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsProcessCompletedEarly'
type MockProcessor_IsProcessCompletedEarly_Call struct {
	*mock.Call
}

// IsProcessCompletedEarly is a helper method to define mock.On call
func (_e *MockProcessor_Expecter) IsProcessCompletedEarly() *MockProcessor_IsProcessCompletedEarly_Call {
	return &MockProcessor_IsProcessCompletedEarly_Call{Call: _e.mock.On("IsProcessCompletedEarly")}
}

func (_c *MockProcessor_IsProcessCompletedEarly_Call) Run(run func()) *MockProcessor_IsProcessCompletedEarly_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockProcessor_IsProcessCompletedEarly_Call) Return(_a0 bool) *MockProcessor_IsProcessCompletedEarly_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockProcessor_IsProcessCompletedEarly_Call) RunAndReturn(run func() bool) *MockProcessor_IsProcessCompletedEarly_Call {
	_c.Call.Return(run)
	return _c
}

// ProcessEnd provides a mock function with given fields:
func (_m *MockProcessor) ProcessEnd() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ProcessEnd")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockProcessor_ProcessEnd_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ProcessEnd'
type MockProcessor_ProcessEnd_Call struct {
	*mock.Call
}

// ProcessEnd is a helper method to define mock.On call
func (_e *MockProcessor_Expecter) ProcessEnd() *MockProcessor_ProcessEnd_Call {
	return &MockProcessor_ProcessEnd_Call{Call: _e.mock.On("ProcessEnd")}
}

func (_c *MockProcessor_ProcessEnd_Call) Run(run func()) *MockProcessor_ProcessEnd_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockProcessor_ProcessEnd_Call) Return(_a0 error) *MockProcessor_ProcessEnd_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockProcessor_ProcessEnd_Call) RunAndReturn(run func() error) *MockProcessor_ProcessEnd_Call {
	_c.Call.Return(run)
	return _c
}

// ProcessLine provides a mock function with given fields: line
func (_m *MockProcessor) ProcessLine(line *scan.LineScript) error {
	ret := _m.Called(line)

	if len(ret) == 0 {
		panic("no return value specified for ProcessLine")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*scan.LineScript) error); ok {
		r0 = rf(line)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockProcessor_ProcessLine_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ProcessLine'
type MockProcessor_ProcessLine_Call struct {
	*mock.Call
}

// ProcessLine is a helper method to define mock.On call
//   - line *scan.LineScript
func (_e *MockProcessor_Expecter) ProcessLine(line interface{}) *MockProcessor_ProcessLine_Call {
	return &MockProcessor_ProcessLine_Call{Call: _e.mock.On("ProcessLine", line)}
}

func (_c *MockProcessor_ProcessLine_Call) Run(run func(line *scan.LineScript)) *MockProcessor_ProcessLine_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*scan.LineScript))
	})
	return _c
}

func (_c *MockProcessor_ProcessLine_Call) Return(_a0 error) *MockProcessor_ProcessLine_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockProcessor_ProcessLine_Call) RunAndReturn(run func(*scan.LineScript) error) *MockProcessor_ProcessLine_Call {
	_c.Call.Return(run)
	return _c
}

// ProcessStart provides a mock function with given fields:
func (_m *MockProcessor) ProcessStart() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ProcessStart")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockProcessor_ProcessStart_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ProcessStart'
type MockProcessor_ProcessStart_Call struct {
	*mock.Call
}

// ProcessStart is a helper method to define mock.On call
func (_e *MockProcessor_Expecter) ProcessStart() *MockProcessor_ProcessStart_Call {
	return &MockProcessor_ProcessStart_Call{Call: _e.mock.On("ProcessStart")}
}

func (_c *MockProcessor_ProcessStart_Call) Run(run func()) *MockProcessor_ProcessStart_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockProcessor_ProcessStart_Call) Return(_a0 error) *MockProcessor_ProcessStart_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockProcessor_ProcessStart_Call) RunAndReturn(run func() error) *MockProcessor_ProcessStart_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockProcessor creates a new instance of MockProcessor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockProcessor(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockProcessor {
	mock := &MockProcessor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
