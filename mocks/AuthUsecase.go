// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	request "github.com/syahidfrd/go-boilerplate/transport/request"
)

// AuthUsecase is an autogenerated mock type for the AuthUsecase type
type AuthUsecase struct {
	mock.Mock
}

// SignIn provides a mock function with given fields: ctx, _a1
func (_m *AuthUsecase) SignIn(ctx context.Context, _a1 *request.SignInReq) (string, error) {
	ret := _m.Called(ctx, _a1)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, *request.SignInReq) string); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *request.SignInReq) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SignUp provides a mock function with given fields: ctx, _a1
func (_m *AuthUsecase) SignUp(ctx context.Context, _a1 *request.SignUpReq) error {
	ret := _m.Called(ctx, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *request.SignUpReq) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewAuthUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewAuthUsecase creates a new instance of AuthUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAuthUsecase(t mockConstructorTestingTNewAuthUsecase) *AuthUsecase {
	mock := &AuthUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
