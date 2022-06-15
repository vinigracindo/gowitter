// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	entities "github.com/vinigracindo/gowitter/internal/entities"

	uuid "github.com/google/uuid"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: name, email, username, password
func (_m *UserRepository) Create(name string, email string, username string, password string) (*entities.User, error) {
	ret := _m.Called(name, email, username, password)

	var r0 *entities.User
	if rf, ok := ret.Get(0).(func(string, string, string, string) *entities.User); ok {
		r0 = rf(name, email, username, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, string) error); ok {
		r1 = rf(name, email, username, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByUUID provides a mock function with given fields: _a0
func (_m *UserRepository) FindByUUID(_a0 uuid.UUID) (*entities.User, error) {
	ret := _m.Called(_a0)

	var r0 *entities.User
	if rf, ok := ret.Get(0).(func(uuid.UUID) *entities.User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: user
func (_m *UserRepository) Update(user *entities.User) error {
	ret := _m.Called(user)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entities.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type NewUserRepositoryT interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserRepository creates a new instance of UserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserRepository(t NewUserRepositoryT) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
