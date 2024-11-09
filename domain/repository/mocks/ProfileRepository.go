// Code generated by mockery v2.46.3. DO NOT EDIT.

package repo_mocks

import (
	context "context"

	entity "github.com/kbiits/dealls-take-home-test/domain/entity"
	mock "github.com/stretchr/testify/mock"
)

// ProfileRepository is an autogenerated mock type for the ProfileRepository type
type ProfileRepository struct {
	mock.Mock
}

// AddProfile provides a mock function with given fields: ctx, profile
func (_m *ProfileRepository) AddProfile(ctx context.Context, profile entity.Profile) (entity.Profile, error) {
	ret := _m.Called(ctx, profile)

	if len(ret) == 0 {
		panic("no return value specified for AddProfile")
	}

	var r0 entity.Profile
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.Profile) (entity.Profile, error)); ok {
		return rf(ctx, profile)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entity.Profile) entity.Profile); ok {
		r0 = rf(ctx, profile)
	} else {
		r0 = ret.Get(0).(entity.Profile)
	}

	if rf, ok := ret.Get(1).(func(context.Context, entity.Profile) error); ok {
		r1 = rf(ctx, profile)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProfileByID provides a mock function with given fields: ctx, userID
func (_m *ProfileRepository) GetProfileByID(ctx context.Context, userID string) (entity.Profile, error) {
	ret := _m.Called(ctx, userID)

	if len(ret) == 0 {
		panic("no return value specified for GetProfileByID")
	}

	var r0 entity.Profile
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (entity.Profile, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) entity.Profile); ok {
		r0 = rf(ctx, userID)
	} else {
		r0 = ret.Get(0).(entity.Profile)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProfileByUserID provides a mock function with given fields: ctx, userID
func (_m *ProfileRepository) GetProfileByUserID(ctx context.Context, userID string) (entity.Profile, error) {
	ret := _m.Called(ctx, userID)

	if len(ret) == 0 {
		panic("no return value specified for GetProfileByUserID")
	}

	var r0 entity.Profile
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (entity.Profile, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) entity.Profile); ok {
		r0 = rf(ctx, userID)
	} else {
		r0 = ret.Get(0).(entity.Profile)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRandomProfilesInSameDistrict provides a mock function with given fields: ctx, loggedInUserID, districtID, limit
func (_m *ProfileRepository) GetRandomProfilesInSameDistrict(ctx context.Context, loggedInUserID string, districtID string, limit int) ([]entity.Profile, error) {
	ret := _m.Called(ctx, loggedInUserID, districtID, limit)

	if len(ret) == 0 {
		panic("no return value specified for GetRandomProfilesInSameDistrict")
	}

	var r0 []entity.Profile
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int) ([]entity.Profile, error)); ok {
		return rf(ctx, loggedInUserID, districtID, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int) []entity.Profile); ok {
		r0 = rf(ctx, loggedInUserID, districtID, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Profile)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, int) error); ok {
		r1 = rf(ctx, loggedInUserID, districtID, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateProfileByUserID provides a mock function with given fields: ctx, profile
func (_m *ProfileRepository) UpdateProfileByUserID(ctx context.Context, profile entity.Profile) (entity.Profile, error) {
	ret := _m.Called(ctx, profile)

	if len(ret) == 0 {
		panic("no return value specified for UpdateProfileByUserID")
	}

	var r0 entity.Profile
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.Profile) (entity.Profile, error)); ok {
		return rf(ctx, profile)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entity.Profile) entity.Profile); ok {
		r0 = rf(ctx, profile)
	} else {
		r0 = ret.Get(0).(entity.Profile)
	}

	if rf, ok := ret.Get(1).(func(context.Context, entity.Profile) error); ok {
		r1 = rf(ctx, profile)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewProfileRepository creates a new instance of ProfileRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProfileRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *ProfileRepository {
	mock := &ProfileRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}