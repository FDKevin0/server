// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "github.com/bangumi/server/internal/model"
)

// SubjectService is an autogenerated mock type for the SubjectService type
type SubjectService struct {
	mock.Mock
}

type SubjectService_Expecter struct {
	mock *mock.Mock
}

func (_m *SubjectService) EXPECT() *SubjectService_Expecter {
	return &SubjectService_Expecter{mock: &_m.Mock}
}

// Get provides a mock function with given fields: ctx, id
func (_m *SubjectService) Get(ctx context.Context, id model.SubjectID) (model.Subject, error) {
	ret := _m.Called(ctx, id)

	var r0 model.Subject
	if rf, ok := ret.Get(0).(func(context.Context, model.SubjectID) model.Subject); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(model.Subject)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.SubjectID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubjectService_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type SubjectService_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//  - ctx context.Context
//  - id model.SubjectID
func (_e *SubjectService_Expecter) Get(ctx interface{}, id interface{}) *SubjectService_Get_Call {
	return &SubjectService_Get_Call{Call: _e.mock.On("Get", ctx, id)}
}

func (_c *SubjectService_Get_Call) Run(run func(ctx context.Context, id model.SubjectID)) *SubjectService_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.SubjectID))
	})
	return _c
}

func (_c *SubjectService_Get_Call) Return(_a0 model.Subject, _a1 error) *SubjectService_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetActors provides a mock function with given fields: ctx, subjectID, characterIDs
func (_m *SubjectService) GetActors(ctx context.Context, subjectID model.SubjectID, characterIDs ...model.CharacterID) (map[model.CharacterID][]model.Person, error) {
	_va := make([]interface{}, len(characterIDs))
	for _i := range characterIDs {
		_va[_i] = characterIDs[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, subjectID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 map[model.CharacterID][]model.Person
	if rf, ok := ret.Get(0).(func(context.Context, model.SubjectID, ...model.CharacterID) map[model.CharacterID][]model.Person); ok {
		r0 = rf(ctx, subjectID, characterIDs...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[model.CharacterID][]model.Person)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.SubjectID, ...model.CharacterID) error); ok {
		r1 = rf(ctx, subjectID, characterIDs...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubjectService_GetActors_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetActors'
type SubjectService_GetActors_Call struct {
	*mock.Call
}

// GetActors is a helper method to define mock.On call
//  - ctx context.Context
//  - subjectID model.SubjectID
//  - characterIDs ...model.CharacterID
func (_e *SubjectService_Expecter) GetActors(ctx interface{}, subjectID interface{}, characterIDs ...interface{}) *SubjectService_GetActors_Call {
	return &SubjectService_GetActors_Call{Call: _e.mock.On("GetActors",
		append([]interface{}{ctx, subjectID}, characterIDs...)...)}
}

func (_c *SubjectService_GetActors_Call) Run(run func(ctx context.Context, subjectID model.SubjectID, characterIDs ...model.CharacterID)) *SubjectService_GetActors_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]model.CharacterID, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(model.CharacterID)
			}
		}
		run(args[0].(context.Context), args[1].(model.SubjectID), variadicArgs...)
	})
	return _c
}

func (_c *SubjectService_GetActors_Call) Return(_a0 map[model.CharacterID][]model.Person, _a1 error) *SubjectService_GetActors_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetCharacterRelated provides a mock function with given fields: ctx, characterID
func (_m *SubjectService) GetCharacterRelated(ctx context.Context, characterID model.CharacterID) ([]model.SubjectCharacterRelation, error) {
	ret := _m.Called(ctx, characterID)

	var r0 []model.SubjectCharacterRelation
	if rf, ok := ret.Get(0).(func(context.Context, model.CharacterID) []model.SubjectCharacterRelation); ok {
		r0 = rf(ctx, characterID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.SubjectCharacterRelation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.CharacterID) error); ok {
		r1 = rf(ctx, characterID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubjectService_GetCharacterRelated_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCharacterRelated'
type SubjectService_GetCharacterRelated_Call struct {
	*mock.Call
}

// GetCharacterRelated is a helper method to define mock.On call
//  - ctx context.Context
//  - characterID model.CharacterID
func (_e *SubjectService_Expecter) GetCharacterRelated(ctx interface{}, characterID interface{}) *SubjectService_GetCharacterRelated_Call {
	return &SubjectService_GetCharacterRelated_Call{Call: _e.mock.On("GetCharacterRelated", ctx, characterID)}
}

func (_c *SubjectService_GetCharacterRelated_Call) Run(run func(ctx context.Context, characterID model.CharacterID)) *SubjectService_GetCharacterRelated_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.CharacterID))
	})
	return _c
}

func (_c *SubjectService_GetCharacterRelated_Call) Return(_a0 []model.SubjectCharacterRelation, _a1 error) *SubjectService_GetCharacterRelated_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetPersonRelated provides a mock function with given fields: ctx, personID
func (_m *SubjectService) GetPersonRelated(ctx context.Context, personID model.PersonID) ([]model.SubjectPersonRelation, error) {
	ret := _m.Called(ctx, personID)

	var r0 []model.SubjectPersonRelation
	if rf, ok := ret.Get(0).(func(context.Context, model.PersonID) []model.SubjectPersonRelation); ok {
		r0 = rf(ctx, personID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.SubjectPersonRelation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.PersonID) error); ok {
		r1 = rf(ctx, personID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubjectService_GetPersonRelated_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPersonRelated'
type SubjectService_GetPersonRelated_Call struct {
	*mock.Call
}

// GetPersonRelated is a helper method to define mock.On call
//  - ctx context.Context
//  - personID model.PersonID
func (_e *SubjectService_Expecter) GetPersonRelated(ctx interface{}, personID interface{}) *SubjectService_GetPersonRelated_Call {
	return &SubjectService_GetPersonRelated_Call{Call: _e.mock.On("GetPersonRelated", ctx, personID)}
}

func (_c *SubjectService_GetPersonRelated_Call) Run(run func(ctx context.Context, personID model.PersonID)) *SubjectService_GetPersonRelated_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.PersonID))
	})
	return _c
}

func (_c *SubjectService_GetPersonRelated_Call) Return(_a0 []model.SubjectPersonRelation, _a1 error) *SubjectService_GetPersonRelated_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetSubjectRelated provides a mock function with given fields: ctx, subjectID
func (_m *SubjectService) GetSubjectRelated(ctx context.Context, subjectID model.SubjectID) ([]model.SubjectInternalRelation, error) {
	ret := _m.Called(ctx, subjectID)

	var r0 []model.SubjectInternalRelation
	if rf, ok := ret.Get(0).(func(context.Context, model.SubjectID) []model.SubjectInternalRelation); ok {
		r0 = rf(ctx, subjectID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.SubjectInternalRelation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.SubjectID) error); ok {
		r1 = rf(ctx, subjectID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubjectService_GetSubjectRelated_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSubjectRelated'
type SubjectService_GetSubjectRelated_Call struct {
	*mock.Call
}

// GetSubjectRelated is a helper method to define mock.On call
//  - ctx context.Context
//  - subjectID model.SubjectID
func (_e *SubjectService_Expecter) GetSubjectRelated(ctx interface{}, subjectID interface{}) *SubjectService_GetSubjectRelated_Call {
	return &SubjectService_GetSubjectRelated_Call{Call: _e.mock.On("GetSubjectRelated", ctx, subjectID)}
}

func (_c *SubjectService_GetSubjectRelated_Call) Run(run func(ctx context.Context, subjectID model.SubjectID)) *SubjectService_GetSubjectRelated_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.SubjectID))
	})
	return _c
}

func (_c *SubjectService_GetSubjectRelated_Call) Return(_a0 []model.SubjectInternalRelation, _a1 error) *SubjectService_GetSubjectRelated_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewSubjectService interface {
	mock.TestingT
	Cleanup(func())
}

// NewSubjectService creates a new instance of SubjectService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSubjectService(t mockConstructorTestingTNewSubjectService) *SubjectService {
	mock := &SubjectService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
