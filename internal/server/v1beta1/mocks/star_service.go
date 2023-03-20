// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	context "context"

	asset "github.com/odpf/compass/core/asset"

	mock "github.com/stretchr/testify/mock"

	namespace "github.com/odpf/compass/core/namespace"

	star "github.com/odpf/compass/core/star"

	user "github.com/odpf/compass/core/user"
)

// StarService is an autogenerated mock type for the StarService type
type StarService struct {
	mock.Mock
}

type StarService_Expecter struct {
	mock *mock.Mock
}

func (_m *StarService) EXPECT() *StarService_Expecter {
	return &StarService_Expecter{mock: &_m.Mock}
}

// GetStargazers provides a mock function with given fields: _a0, _a1, _a2
func (_m *StarService) GetStargazers(_a0 context.Context, _a1 star.Filter, _a2 string) ([]user.User, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 []user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, star.Filter, string) ([]user.User, error)); ok {
		return rf(_a0, _a1, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, star.Filter, string) []user.User); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]user.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, star.Filter, string) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StarService_GetStargazers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetStargazers'
type StarService_GetStargazers_Call struct {
	*mock.Call
}

// GetStargazers is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 star.Filter
//   - _a2 string
func (_e *StarService_Expecter) GetStargazers(_a0 interface{}, _a1 interface{}, _a2 interface{}) *StarService_GetStargazers_Call {
	return &StarService_GetStargazers_Call{Call: _e.mock.On("GetStargazers", _a0, _a1, _a2)}
}

func (_c *StarService_GetStargazers_Call) Run(run func(_a0 context.Context, _a1 star.Filter, _a2 string)) *StarService_GetStargazers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(star.Filter), args[2].(string))
	})
	return _c
}

func (_c *StarService_GetStargazers_Call) Return(_a0 []user.User, _a1 error) *StarService_GetStargazers_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *StarService_GetStargazers_Call) RunAndReturn(run func(context.Context, star.Filter, string) ([]user.User, error)) *StarService_GetStargazers_Call {
	_c.Call.Return(run)
	return _c
}

// GetStarredAssetByUserID provides a mock function with given fields: _a0, _a1, _a2
func (_m *StarService) GetStarredAssetByUserID(_a0 context.Context, _a1 string, _a2 string) (asset.Asset, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 asset.Asset
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (asset.Asset, error)); ok {
		return rf(_a0, _a1, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) asset.Asset); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(asset.Asset)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StarService_GetStarredAssetByUserID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetStarredAssetByUserID'
type StarService_GetStarredAssetByUserID_Call struct {
	*mock.Call
}

// GetStarredAssetByUserID is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
//   - _a2 string
func (_e *StarService_Expecter) GetStarredAssetByUserID(_a0 interface{}, _a1 interface{}, _a2 interface{}) *StarService_GetStarredAssetByUserID_Call {
	return &StarService_GetStarredAssetByUserID_Call{Call: _e.mock.On("GetStarredAssetByUserID", _a0, _a1, _a2)}
}

func (_c *StarService_GetStarredAssetByUserID_Call) Run(run func(_a0 context.Context, _a1 string, _a2 string)) *StarService_GetStarredAssetByUserID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *StarService_GetStarredAssetByUserID_Call) Return(_a0 asset.Asset, _a1 error) *StarService_GetStarredAssetByUserID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *StarService_GetStarredAssetByUserID_Call) RunAndReturn(run func(context.Context, string, string) (asset.Asset, error)) *StarService_GetStarredAssetByUserID_Call {
	_c.Call.Return(run)
	return _c
}

// GetStarredAssetsByUserID provides a mock function with given fields: _a0, _a1, _a2
func (_m *StarService) GetStarredAssetsByUserID(_a0 context.Context, _a1 star.Filter, _a2 string) ([]asset.Asset, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 []asset.Asset
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, star.Filter, string) ([]asset.Asset, error)); ok {
		return rf(_a0, _a1, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, star.Filter, string) []asset.Asset); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]asset.Asset)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, star.Filter, string) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StarService_GetStarredAssetsByUserID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetStarredAssetsByUserID'
type StarService_GetStarredAssetsByUserID_Call struct {
	*mock.Call
}

// GetStarredAssetsByUserID is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 star.Filter
//   - _a2 string
func (_e *StarService_Expecter) GetStarredAssetsByUserID(_a0 interface{}, _a1 interface{}, _a2 interface{}) *StarService_GetStarredAssetsByUserID_Call {
	return &StarService_GetStarredAssetsByUserID_Call{Call: _e.mock.On("GetStarredAssetsByUserID", _a0, _a1, _a2)}
}

func (_c *StarService_GetStarredAssetsByUserID_Call) Run(run func(_a0 context.Context, _a1 star.Filter, _a2 string)) *StarService_GetStarredAssetsByUserID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(star.Filter), args[2].(string))
	})
	return _c
}

func (_c *StarService_GetStarredAssetsByUserID_Call) Return(_a0 []asset.Asset, _a1 error) *StarService_GetStarredAssetsByUserID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *StarService_GetStarredAssetsByUserID_Call) RunAndReturn(run func(context.Context, star.Filter, string) ([]asset.Asset, error)) *StarService_GetStarredAssetsByUserID_Call {
	_c.Call.Return(run)
	return _c
}

// Stars provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *StarService) Stars(_a0 context.Context, _a1 *namespace.Namespace, _a2 string, _a3 string) (string, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *namespace.Namespace, string, string) (string, error)); ok {
		return rf(_a0, _a1, _a2, _a3)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *namespace.Namespace, string, string) string); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *namespace.Namespace, string, string) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StarService_Stars_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Stars'
type StarService_Stars_Call struct {
	*mock.Call
}

// Stars is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *namespace.Namespace
//   - _a2 string
//   - _a3 string
func (_e *StarService_Expecter) Stars(_a0 interface{}, _a1 interface{}, _a2 interface{}, _a3 interface{}) *StarService_Stars_Call {
	return &StarService_Stars_Call{Call: _e.mock.On("Stars", _a0, _a1, _a2, _a3)}
}

func (_c *StarService_Stars_Call) Run(run func(_a0 context.Context, _a1 *namespace.Namespace, _a2 string, _a3 string)) *StarService_Stars_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*namespace.Namespace), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *StarService_Stars_Call) Return(_a0 string, _a1 error) *StarService_Stars_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *StarService_Stars_Call) RunAndReturn(run func(context.Context, *namespace.Namespace, string, string) (string, error)) *StarService_Stars_Call {
	_c.Call.Return(run)
	return _c
}

// Unstars provides a mock function with given fields: _a0, _a1, _a2
func (_m *StarService) Unstars(_a0 context.Context, _a1 string, _a2 string) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StarService_Unstars_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Unstars'
type StarService_Unstars_Call struct {
	*mock.Call
}

// Unstars is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
//   - _a2 string
func (_e *StarService_Expecter) Unstars(_a0 interface{}, _a1 interface{}, _a2 interface{}) *StarService_Unstars_Call {
	return &StarService_Unstars_Call{Call: _e.mock.On("Unstars", _a0, _a1, _a2)}
}

func (_c *StarService_Unstars_Call) Run(run func(_a0 context.Context, _a1 string, _a2 string)) *StarService_Unstars_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *StarService_Unstars_Call) Return(_a0 error) *StarService_Unstars_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *StarService_Unstars_Call) RunAndReturn(run func(context.Context, string, string) error) *StarService_Unstars_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewStarService interface {
	mock.TestingT
	Cleanup(func())
}

// NewStarService creates a new instance of StarService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStarService(t mockConstructorTestingTNewStarService) *StarService {
	mock := &StarService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
