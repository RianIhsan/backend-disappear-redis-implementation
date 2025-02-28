// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	domain "github.com/capstone-kelompok-7/backend-disappear/module/entities"
	mock "github.com/stretchr/testify/mock"
)

// ServiceUserInterface is an autogenerated mock type for the ServiceUserInterface type
type ServiceUserInterface struct {
	mock.Mock
}

// GetAllUsers provides a mock function with given fields:
func (_m *ServiceUserInterface) GetAllUsers() ([]*domain.UserModels, error) {
	ret := _m.Called()

	var r0 []*domain.UserModels
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*domain.UserModels, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*domain.UserModels); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.UserModels)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUsersByEmail provides a mock function with given fields: email
func (_m *ServiceUserInterface) GetUsersByEmail(email string) (*domain.UserModels, error) {
	ret := _m.Called(email)

	var r0 *domain.UserModels
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*domain.UserModels, error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.UserModels); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.UserModels)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUsersById provides a mock function with given fields: userId
func (_m *ServiceUserInterface) GetUsersById(userId uint64) (*domain.UserModels, error) {
	ret := _m.Called(userId)

	var r0 *domain.UserModels
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64) (*domain.UserModels, error)); ok {
		return rf(userId)
	}
	if rf, ok := ret.Get(0).(func(uint64) *domain.UserModels); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.UserModels)
		}
	}

	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewServiceUserInterface creates a new instance of ServiceUserInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewServiceUserInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *ServiceUserInterface {
	mock := &ServiceUserInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
