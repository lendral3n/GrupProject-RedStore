// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

// import (
// 	user "MyEcommerce/features/user"

// 	mock "github.com/stretchr/testify/mock"
// )

// // UserData is an autogenerated mock type for the UserDataInterface type
// type UserData struct {
// 	mock.Mock
// }

// // Delete provides a mock function with given fields: userIdLogin
// func (_m *UserData) Delete(userIdLogin int) error {
// 	ret := _m.Called(userIdLogin)

// 	if len(ret) == 0 {
// 		panic("no return value specified for Delete")
// 	}

// 	var r0 error
// 	if rf, ok := ret.Get(0).(func(int) error); ok {
// 		r0 = rf(userIdLogin)
// 	} else {
// 		r0 = ret.Error(0)
// 	}

// 	return r0
// }

// // Insert provides a mock function with given fields: input
// func (_m *UserData) Insert(input user.Core) error {
// 	ret := _m.Called(input)

// 	if len(ret) == 0 {
// 		panic("no return value specified for Insert")
// 	}

// 	var r0 error
// 	if rf, ok := ret.Get(0).(func(user.Core) error); ok {
// 		r0 = rf(input)
// 	} else {
// 		r0 = ret.Error(0)
// 	}

// 	return r0
// }

// // Login provides a mock function with given fields: email, password
// func (_m *UserData) Login(email string, password string) (*user.Core, error) {
// 	ret := _m.Called(email, password)

// 	if len(ret) == 0 {
// 		panic("no return value specified for Login")
// 	}

// 	var r0 *user.Core
// 	var r1 error
// 	if rf, ok := ret.Get(0).(func(string, string) (*user.Core, error)); ok {
// 		return rf(email, password)
// 	}
// 	if rf, ok := ret.Get(0).(func(string, string) *user.Core); ok {
// 		r0 = rf(email, password)
// 	} else {
// 		if ret.Get(0) != nil {
// 			r0 = ret.Get(0).(*user.Core)
// 		}
// 	}

// 	if rf, ok := ret.Get(1).(func(string, string) error); ok {
// 		r1 = rf(email, password)
// 	} else {
// 		r1 = ret.Error(1)
// 	}

// 	return r0, r1
// }

// // SelectAdminUsers provides a mock function with given fields: page, limit
// func (_m *UserData) SelectAdminUsers(page int, limit int) ([]user.Core, error) {
// 	ret := _m.Called(page, limit)

// 	if len(ret) == 0 {
// 		panic("no return value specified for SelectAdminUsers")
// 	}

// 	var r0 []user.Core
// 	var r1 error
// 	if rf, ok := ret.Get(0).(func(int, int) ([]user.Core, error)); ok {
// 		return rf(page, limit)
// 	}
// 	if rf, ok := ret.Get(0).(func(int, int) []user.Core); ok {
// 		r0 = rf(page, limit)
// 	} else {
// 		if ret.Get(0) != nil {
// 			r0 = ret.Get(0).([]user.Core)
// 		}
// 	}

// 	if rf, ok := ret.Get(1).(func(int, int) error); ok {
// 		r1 = rf(page, limit)
// 	} else {
// 		r1 = ret.Error(1)
// 	}

// 	return r0, r1
// }

// // SelectById provides a mock function with given fields: userIdLogin
// func (_m *UserData) SelectById(userIdLogin int) (*user.Core, error) {
// 	ret := _m.Called(userIdLogin)

// 	if len(ret) == 0 {
// 		panic("no return value specified for SelectById")
// 	}

// 	var r0 *user.Core
// 	var r1 error
// 	if rf, ok := ret.Get(0).(func(int) (*user.Core, error)); ok {
// 		return rf(userIdLogin)
// 	}
// 	if rf, ok := ret.Get(0).(func(int) *user.Core); ok {
// 		r0 = rf(userIdLogin)
// 	} else {
// 		if ret.Get(0) != nil {
// 			r0 = ret.Get(0).(*user.Core)
// 		}
// 	}

// 	if rf, ok := ret.Get(1).(func(int) error); ok {
// 		r1 = rf(userIdLogin)
// 	} else {
// 		r1 = ret.Error(1)
// 	}

// 	return r0, r1
// }

// // Update provides a mock function with given fields: userIdLogin, input
// func (_m *UserData) Update(userIdLogin int, input user.Core) error {
// 	ret := _m.Called(userIdLogin, input)

// 	if len(ret) == 0 {
// 		panic("no return value specified for Update")
// 	}

// 	var r0 error
// 	if rf, ok := ret.Get(0).(func(int, user.Core) error); ok {
// 		r0 = rf(userIdLogin, input)
// 	} else {
// 		r0 = ret.Error(0)
// 	}

// 	return r0
// }

// // NewUserData creates a new instance of UserData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// // The first argument is typically a *testing.T value.
// func NewUserData(t interface {
// 	mock.TestingT
// 	Cleanup(func())
// }) *UserData {
// 	mock := &UserData{}
// 	mock.Mock.Test(t)

// 	t.Cleanup(func() { mock.AssertExpectations(t) })

// 	return mock
// }
