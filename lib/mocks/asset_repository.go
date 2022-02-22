// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	asset "github.com/odpf/columbus/asset"

	mock "github.com/stretchr/testify/mock"
)

// AssetRepository is an autogenerated mock type for the Repository type
type AssetRepository struct {
	mock.Mock
}

// Delete provides a mock function with given fields: ctx, id
func (_m *AssetRepository) Delete(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: _a0, _a1
func (_m *AssetRepository) Get(_a0 context.Context, _a1 asset.Config) ([]asset.Asset, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []asset.Asset
	if rf, ok := ret.Get(0).(func(context.Context, asset.Config) []asset.Asset); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]asset.Asset)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, asset.Config) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *AssetRepository) GetByID(ctx context.Context, id string) (asset.Asset, error) {
	ret := _m.Called(ctx, id)

	var r0 asset.Asset
	if rf, ok := ret.Get(0).(func(context.Context, string) asset.Asset); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(asset.Asset)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByVersion provides a mock function with given fields: ctx, id, version
func (_m *AssetRepository) GetByVersion(ctx context.Context, id string, version string) (asset.Asset, error) {
	ret := _m.Called(ctx, id, version)

	var r0 asset.Asset
	if rf, ok := ret.Get(0).(func(context.Context, string, string) asset.Asset); ok {
		r0 = rf(ctx, id, version)
	} else {
		r0 = ret.Get(0).(asset.Asset)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, id, version)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCount provides a mock function with given fields: _a0, _a1
func (_m *AssetRepository) GetCount(_a0 context.Context, _a1 asset.Config) (int, error) {
	ret := _m.Called(_a0, _a1)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, asset.Config) int); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, asset.Config) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPrevVersions provides a mock function with given fields: ctx, cfg, id
func (_m *AssetRepository) GetPrevVersions(ctx context.Context, cfg asset.Config, id string) ([]asset.AssetVersion, error) {
	ret := _m.Called(ctx, cfg, id)

	var r0 []asset.AssetVersion
	if rf, ok := ret.Get(0).(func(context.Context, asset.Config, string) []asset.AssetVersion); ok {
		r0 = rf(ctx, cfg, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]asset.AssetVersion)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, asset.Config, string) error); ok {
		r1 = rf(ctx, cfg, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Upsert provides a mock function with given fields: ctx, userID, ast
func (_m *AssetRepository) Upsert(ctx context.Context, userID string, ast *asset.Asset) (string, error) {
	ret := _m.Called(ctx, userID, ast)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string, *asset.Asset) string); ok {
		r0 = rf(ctx, userID, ast)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, *asset.Asset) error); ok {
		r1 = rf(ctx, userID, ast)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
