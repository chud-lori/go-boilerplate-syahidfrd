// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	entity "github.com/syahidfrd/go-boilerplate/entity"
)

// TodoRepository is an autogenerated mock type for the TodoRepository type
type TodoRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, todo
func (_m *TodoRepository) Create(ctx context.Context, todo *entity.Todo) error {
	ret := _m.Called(ctx, todo)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Todo) error); ok {
		r0 = rf(ctx, todo)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, id
func (_m *TodoRepository) Delete(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Fetch provides a mock function with given fields: ctx
func (_m *TodoRepository) Fetch(ctx context.Context) ([]entity.Todo, error) {
	ret := _m.Called(ctx)

	var r0 []entity.Todo
	if rf, ok := ret.Get(0).(func(context.Context) []entity.Todo); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Todo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *TodoRepository) GetByID(ctx context.Context, id int64) (entity.Todo, error) {
	ret := _m.Called(ctx, id)

	var r0 entity.Todo
	if rf, ok := ret.Get(0).(func(context.Context, int64) entity.Todo); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(entity.Todo)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, todo
func (_m *TodoRepository) Update(ctx context.Context, todo *entity.Todo) error {
	ret := _m.Called(ctx, todo)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Todo) error); ok {
		r0 = rf(ctx, todo)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewTodoRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewTodoRepository creates a new instance of TodoRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTodoRepository(t mockConstructorTestingTNewTodoRepository) *TodoRepository {
	mock := &TodoRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
