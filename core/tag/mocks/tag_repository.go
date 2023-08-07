// Code generated by mockery v2.28.1. DO NOT EDIT.

package mocks

import (
	context "context"

	tag "github.com/goto/compass/core/tag"
	mock "github.com/stretchr/testify/mock"
)

// TagRepository is an autogenerated mock type for the TagRepository type
type TagRepository struct {
	mock.Mock
}

type TagRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *TagRepository) EXPECT() *TagRepository_Expecter {
	return &TagRepository_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, _a1
func (_m *TagRepository) Create(ctx context.Context, _a1 *tag.Tag) error {
	ret := _m.Called(ctx, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *tag.Tag) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TagRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type TagRepository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - _a1 *tag.Tag
func (_e *TagRepository_Expecter) Create(ctx interface{}, _a1 interface{}) *TagRepository_Create_Call {
	return &TagRepository_Create_Call{Call: _e.mock.On("Create", ctx, _a1)}
}

func (_c *TagRepository_Create_Call) Run(run func(ctx context.Context, _a1 *tag.Tag)) *TagRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*tag.Tag))
	})
	return _c
}

func (_c *TagRepository_Create_Call) Return(_a0 error) *TagRepository_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TagRepository_Create_Call) RunAndReturn(run func(context.Context, *tag.Tag) error) *TagRepository_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, filter
func (_m *TagRepository) Delete(ctx context.Context, filter tag.Tag) error {
	ret := _m.Called(ctx, filter)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, tag.Tag) error); ok {
		r0 = rf(ctx, filter)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TagRepository_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type TagRepository_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - filter tag.Tag
func (_e *TagRepository_Expecter) Delete(ctx interface{}, filter interface{}) *TagRepository_Delete_Call {
	return &TagRepository_Delete_Call{Call: _e.mock.On("Delete", ctx, filter)}
}

func (_c *TagRepository_Delete_Call) Run(run func(ctx context.Context, filter tag.Tag)) *TagRepository_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(tag.Tag))
	})
	return _c
}

func (_c *TagRepository_Delete_Call) Return(_a0 error) *TagRepository_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TagRepository_Delete_Call) RunAndReturn(run func(context.Context, tag.Tag) error) *TagRepository_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Read provides a mock function with given fields: ctx, filter
func (_m *TagRepository) Read(ctx context.Context, filter tag.Tag) ([]tag.Tag, error) {
	ret := _m.Called(ctx, filter)

	var r0 []tag.Tag
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, tag.Tag) ([]tag.Tag, error)); ok {
		return rf(ctx, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, tag.Tag) []tag.Tag); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]tag.Tag)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, tag.Tag) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TagRepository_Read_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Read'
type TagRepository_Read_Call struct {
	*mock.Call
}

// Read is a helper method to define mock.On call
//   - ctx context.Context
//   - filter tag.Tag
func (_e *TagRepository_Expecter) Read(ctx interface{}, filter interface{}) *TagRepository_Read_Call {
	return &TagRepository_Read_Call{Call: _e.mock.On("Read", ctx, filter)}
}

func (_c *TagRepository_Read_Call) Run(run func(ctx context.Context, filter tag.Tag)) *TagRepository_Read_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(tag.Tag))
	})
	return _c
}

func (_c *TagRepository_Read_Call) Return(_a0 []tag.Tag, _a1 error) *TagRepository_Read_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *TagRepository_Read_Call) RunAndReturn(run func(context.Context, tag.Tag) ([]tag.Tag, error)) *TagRepository_Read_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, _a1
func (_m *TagRepository) Update(ctx context.Context, _a1 *tag.Tag) error {
	ret := _m.Called(ctx, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *tag.Tag) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TagRepository_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type TagRepository_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - _a1 *tag.Tag
func (_e *TagRepository_Expecter) Update(ctx interface{}, _a1 interface{}) *TagRepository_Update_Call {
	return &TagRepository_Update_Call{Call: _e.mock.On("Update", ctx, _a1)}
}

func (_c *TagRepository_Update_Call) Run(run func(ctx context.Context, _a1 *tag.Tag)) *TagRepository_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*tag.Tag))
	})
	return _c
}

func (_c *TagRepository_Update_Call) Return(_a0 error) *TagRepository_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TagRepository_Update_Call) RunAndReturn(run func(context.Context, *tag.Tag) error) *TagRepository_Update_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewTagRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewTagRepository creates a new instance of TagRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTagRepository(t mockConstructorTestingTNewTagRepository) *TagRepository {
	mock := &TagRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
