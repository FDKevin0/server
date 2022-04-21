// Code generated by mockery v2.10.6. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/bangumi/server/domain"
	mock "github.com/stretchr/testify/mock"

	model "github.com/bangumi/server/model"
)

// IndexRepo is an autogenerated mock type for the IndexRepo type
type IndexRepo struct {
	mock.Mock
}

type IndexRepo_Expecter struct {
	mock *mock.Mock
}

func (_m *IndexRepo) EXPECT() *IndexRepo_Expecter {
	return &IndexRepo_Expecter{mock: &_m.Mock}
}

// CountSubjects provides a mock function with given fields: ctx, id, subjectType
func (_m *IndexRepo) CountSubjects(ctx context.Context, id uint32, subjectType uint8) (int64, error) {
	ret := _m.Called(ctx, id, subjectType)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, uint32, uint8) int64); ok {
		r0 = rf(ctx, id, subjectType)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint32, uint8) error); ok {
		r1 = rf(ctx, id, subjectType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IndexRepo_CountSubjects_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CountSubjects'
type IndexRepo_CountSubjects_Call struct {
	*mock.Call
}

// CountSubjects is a helper method to define mock.On call
//  - ctx context.Context
//  - id uint32
//  - subjectType uint8
func (_e *IndexRepo_Expecter) CountSubjects(ctx interface{}, id interface{}, subjectType interface{}) *IndexRepo_CountSubjects_Call {
	return &IndexRepo_CountSubjects_Call{Call: _e.mock.On("CountSubjects", ctx, id, subjectType)}
}

func (_c *IndexRepo_CountSubjects_Call) Run(run func(ctx context.Context, id uint32, subjectType uint8)) *IndexRepo_CountSubjects_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32), args[2].(uint8))
	})
	return _c
}

func (_c *IndexRepo_CountSubjects_Call) Return(_a0 int64, _a1 error) *IndexRepo_CountSubjects_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Get provides a mock function with given fields: ctx, id
func (_m *IndexRepo) Get(ctx context.Context, id uint32) (model.Index, error) {
	ret := _m.Called(ctx, id)

	var r0 model.Index
	if rf, ok := ret.Get(0).(func(context.Context, uint32) model.Index); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(model.Index)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint32) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IndexRepo_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type IndexRepo_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//  - ctx context.Context
//  - id uint32
func (_e *IndexRepo_Expecter) Get(ctx interface{}, id interface{}) *IndexRepo_Get_Call {
	return &IndexRepo_Get_Call{Call: _e.mock.On("Get", ctx, id)}
}

func (_c *IndexRepo_Get_Call) Run(run func(ctx context.Context, id uint32)) *IndexRepo_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32))
	})
	return _c
}

func (_c *IndexRepo_Get_Call) Return(_a0 model.Index, _a1 error) *IndexRepo_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// ListSubjects provides a mock function with given fields: ctx, id, subjectType, limit, offset
func (_m *IndexRepo) ListSubjects(ctx context.Context, id uint32, subjectType uint8, limit int, offset int) ([]domain.IndexSubject, error) {
	ret := _m.Called(ctx, id, subjectType, limit, offset)

	var r0 []domain.IndexSubject
	if rf, ok := ret.Get(0).(func(context.Context, uint32, uint8, int, int) []domain.IndexSubject); ok {
		r0 = rf(ctx, id, subjectType, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.IndexSubject)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint32, uint8, int, int) error); ok {
		r1 = rf(ctx, id, subjectType, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IndexRepo_ListSubjects_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListSubjects'
type IndexRepo_ListSubjects_Call struct {
	*mock.Call
}

// ListSubjects is a helper method to define mock.On call
//  - ctx context.Context
//  - id uint32
//  - subjectType uint8
//  - limit int
//  - offset int
func (_e *IndexRepo_Expecter) ListSubjects(ctx interface{}, id interface{}, subjectType interface{}, limit interface{}, offset interface{}) *IndexRepo_ListSubjects_Call {
	return &IndexRepo_ListSubjects_Call{Call: _e.mock.On("ListSubjects", ctx, id, subjectType, limit, offset)}
}

func (_c *IndexRepo_ListSubjects_Call) Run(run func(ctx context.Context, id uint32, subjectType uint8, limit int, offset int)) *IndexRepo_ListSubjects_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32), args[2].(uint8), args[3].(int), args[4].(int))
	})
	return _c
}

func (_c *IndexRepo_ListSubjects_Call) Return(_a0 []domain.IndexSubject, _a1 error) *IndexRepo_ListSubjects_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}
