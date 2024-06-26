// Code generated by mockery v2.42.2. DO NOT EDIT.

package rewriter

import (
	scan "marcuson/scriptman/internal/script/internal/scan"

	mock "github.com/stretchr/testify/mock"
)

// MockRewriter is an autogenerated mock type for the Rewriter type
type MockRewriter struct {
	mock.Mock
}

type MockRewriter_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRewriter) EXPECT() *MockRewriter_Expecter {
	return &MockRewriter_Expecter{mock: &_m.Mock}
}

// RewriteAfterAll provides a mock function with given fields:
func (_m *MockRewriter) RewriteAfterAll() (string, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for RewriteAfterAll")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func() (string, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRewriter_RewriteAfterAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RewriteAfterAll'
type MockRewriter_RewriteAfterAll_Call struct {
	*mock.Call
}

// RewriteAfterAll is a helper method to define mock.On call
func (_e *MockRewriter_Expecter) RewriteAfterAll() *MockRewriter_RewriteAfterAll_Call {
	return &MockRewriter_RewriteAfterAll_Call{Call: _e.mock.On("RewriteAfterAll")}
}

func (_c *MockRewriter_RewriteAfterAll_Call) Run(run func()) *MockRewriter_RewriteAfterAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockRewriter_RewriteAfterAll_Call) Return(_a0 string, _a1 error) *MockRewriter_RewriteAfterAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRewriter_RewriteAfterAll_Call) RunAndReturn(run func() (string, error)) *MockRewriter_RewriteAfterAll_Call {
	_c.Call.Return(run)
	return _c
}

// RewriteBeforeAll provides a mock function with given fields:
func (_m *MockRewriter) RewriteBeforeAll() (string, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for RewriteBeforeAll")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func() (string, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRewriter_RewriteBeforeAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RewriteBeforeAll'
type MockRewriter_RewriteBeforeAll_Call struct {
	*mock.Call
}

// RewriteBeforeAll is a helper method to define mock.On call
func (_e *MockRewriter_Expecter) RewriteBeforeAll() *MockRewriter_RewriteBeforeAll_Call {
	return &MockRewriter_RewriteBeforeAll_Call{Call: _e.mock.On("RewriteBeforeAll")}
}

func (_c *MockRewriter_RewriteBeforeAll_Call) Run(run func()) *MockRewriter_RewriteBeforeAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockRewriter_RewriteBeforeAll_Call) Return(_a0 string, _a1 error) *MockRewriter_RewriteBeforeAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRewriter_RewriteBeforeAll_Call) RunAndReturn(run func() (string, error)) *MockRewriter_RewriteBeforeAll_Call {
	_c.Call.Return(run)
	return _c
}

// RewriteLine provides a mock function with given fields: line
func (_m *MockRewriter) RewriteLine(line *scan.LineScript) (string, error) {
	ret := _m.Called(line)

	if len(ret) == 0 {
		panic("no return value specified for RewriteLine")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(*scan.LineScript) (string, error)); ok {
		return rf(line)
	}
	if rf, ok := ret.Get(0).(func(*scan.LineScript) string); ok {
		r0 = rf(line)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(*scan.LineScript) error); ok {
		r1 = rf(line)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRewriter_RewriteLine_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RewriteLine'
type MockRewriter_RewriteLine_Call struct {
	*mock.Call
}

// RewriteLine is a helper method to define mock.On call
//   - line *scan.LineScript
func (_e *MockRewriter_Expecter) RewriteLine(line interface{}) *MockRewriter_RewriteLine_Call {
	return &MockRewriter_RewriteLine_Call{Call: _e.mock.On("RewriteLine", line)}
}

func (_c *MockRewriter_RewriteLine_Call) Run(run func(line *scan.LineScript)) *MockRewriter_RewriteLine_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*scan.LineScript))
	})
	return _c
}

func (_c *MockRewriter_RewriteLine_Call) Return(_a0 string, _a1 error) *MockRewriter_RewriteLine_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRewriter_RewriteLine_Call) RunAndReturn(run func(*scan.LineScript) (string, error)) *MockRewriter_RewriteLine_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockRewriter creates a new instance of MockRewriter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRewriter(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRewriter {
	mock := &MockRewriter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
