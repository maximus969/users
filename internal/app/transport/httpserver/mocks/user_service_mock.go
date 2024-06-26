// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/maximus969/users-app/internal/app/domain"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// UserService is an autogenerated mock type for the UserService type
type UserService struct {
	mock.Mock
}

type UserService_Expecter struct {
	mock *mock.Mock
}

func (_m *UserService) EXPECT() *UserService_Expecter {
	return &UserService_Expecter{mock: &_m.Mock}
}

// CreateUser provides a mock function with given fields: ctx, user
func (_m *UserService) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	ret := _m.Called(ctx, user)

	if len(ret) == 0 {
		panic("no return value specified for CreateUser")
	}

	var r0 domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.User) (domain.User, error)); ok {
		return rf(ctx, user)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.User) domain.User); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserService_CreateUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateUser'
type UserService_CreateUser_Call struct {
	*mock.Call
}

// CreateUser is a helper method to define mock.On call
//   - ctx context.Context
//   - user domain.User
func (_e *UserService_Expecter) CreateUser(ctx interface{}, user interface{}) *UserService_CreateUser_Call {
	return &UserService_CreateUser_Call{Call: _e.mock.On("CreateUser", ctx, user)}
}

func (_c *UserService_CreateUser_Call) Run(run func(ctx context.Context, user domain.User)) *UserService_CreateUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.User))
	})
	return _c
}

func (_c *UserService_CreateUser_Call) Return(_a0 domain.User, _a1 error) *UserService_CreateUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserService_CreateUser_Call) RunAndReturn(run func(context.Context, domain.User) (domain.User, error)) *UserService_CreateUser_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteUser provides a mock function with given fields: ctx, id
func (_m *UserService) DeleteUser(ctx context.Context, id uuid.UUID) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserService_DeleteUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteUser'
type UserService_DeleteUser_Call struct {
	*mock.Call
}

// DeleteUser is a helper method to define mock.On call
//   - ctx context.Context
//   - id uuid.UUID
func (_e *UserService_Expecter) DeleteUser(ctx interface{}, id interface{}) *UserService_DeleteUser_Call {
	return &UserService_DeleteUser_Call{Call: _e.mock.On("DeleteUser", ctx, id)}
}

func (_c *UserService_DeleteUser_Call) Run(run func(ctx context.Context, id uuid.UUID)) *UserService_DeleteUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *UserService_DeleteUser_Call) Return(_a0 error) *UserService_DeleteUser_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *UserService_DeleteUser_Call) RunAndReturn(run func(context.Context, uuid.UUID) error) *UserService_DeleteUser_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserById provides a mock function with given fields: ctx, id
func (_m *UserService) GetUserById(ctx context.Context, id uuid.UUID) (domain.User, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetUserById")
	}

	var r0 domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (domain.User, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) domain.User); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserService_GetUserById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserById'
type UserService_GetUserById_Call struct {
	*mock.Call
}

// GetUserById is a helper method to define mock.On call
//   - ctx context.Context
//   - id uuid.UUID
func (_e *UserService_Expecter) GetUserById(ctx interface{}, id interface{}) *UserService_GetUserById_Call {
	return &UserService_GetUserById_Call{Call: _e.mock.On("GetUserById", ctx, id)}
}

func (_c *UserService_GetUserById_Call) Run(run func(ctx context.Context, id uuid.UUID)) *UserService_GetUserById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *UserService_GetUserById_Call) Return(_a0 domain.User, _a1 error) *UserService_GetUserById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserService_GetUserById_Call) RunAndReturn(run func(context.Context, uuid.UUID) (domain.User, error)) *UserService_GetUserById_Call {
	_c.Call.Return(run)
	return _c
}

// GetUsers provides a mock function with given fields: ctx
func (_m *UserService) GetUsers(ctx context.Context) ([]domain.User, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetUsers")
	}

	var r0 []domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]domain.User, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []domain.User); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserService_GetUsers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUsers'
type UserService_GetUsers_Call struct {
	*mock.Call
}

// GetUsers is a helper method to define mock.On call
//   - ctx context.Context
func (_e *UserService_Expecter) GetUsers(ctx interface{}) *UserService_GetUsers_Call {
	return &UserService_GetUsers_Call{Call: _e.mock.On("GetUsers", ctx)}
}

func (_c *UserService_GetUsers_Call) Run(run func(ctx context.Context)) *UserService_GetUsers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *UserService_GetUsers_Call) Return(_a0 []domain.User, _a1 error) *UserService_GetUsers_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserService_GetUsers_Call) RunAndReturn(run func(context.Context) ([]domain.User, error)) *UserService_GetUsers_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateUser provides a mock function with given fields: ctx, user
func (_m *UserService) UpdateUser(ctx context.Context, user domain.User) (domain.User, error) {
	ret := _m.Called(ctx, user)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUser")
	}

	var r0 domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.User) (domain.User, error)); ok {
		return rf(ctx, user)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.User) domain.User); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserService_UpdateUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateUser'
type UserService_UpdateUser_Call struct {
	*mock.Call
}

// UpdateUser is a helper method to define mock.On call
//   - ctx context.Context
//   - user domain.User
func (_e *UserService_Expecter) UpdateUser(ctx interface{}, user interface{}) *UserService_UpdateUser_Call {
	return &UserService_UpdateUser_Call{Call: _e.mock.On("UpdateUser", ctx, user)}
}

func (_c *UserService_UpdateUser_Call) Run(run func(ctx context.Context, user domain.User)) *UserService_UpdateUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.User))
	})
	return _c
}

func (_c *UserService_UpdateUser_Call) Return(_a0 domain.User, _a1 error) *UserService_UpdateUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserService_UpdateUser_Call) RunAndReturn(run func(context.Context, domain.User) (domain.User, error)) *UserService_UpdateUser_Call {
	_c.Call.Return(run)
	return _c
}

// NewUserService creates a new instance of UserService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserService(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserService {
	mock := &UserService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
