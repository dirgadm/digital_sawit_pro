// Code generated by MockGen. DO NOT EDIT.
// Source: repository/interfaces.go

// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRepositoryInterface is a mock of RepositoryInterface interface.
type MockRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryInterfaceMockRecorder
}

// MockRepositoryInterfaceMockRecorder is the mock recorder for MockRepositoryInterface.
type MockRepositoryInterfaceMockRecorder struct {
	mock *MockRepositoryInterface
}

// NewMockRepositoryInterface creates a new mock instance.
func NewMockRepositoryInterface(ctrl *gomock.Controller) *MockRepositoryInterface {
	mock := &MockRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepositoryInterface) EXPECT() *MockRepositoryInterfaceMockRecorder {
	return m.recorder
}

// GetMyProfile mocks base method.
func (m *MockRepositoryInterface) GetMyProfile(ctx context.Context, userID int, anotherAttr ...string) (string, string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, userID}
	for _, a := range anotherAttr {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMyProfile", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetMyProfile indicates an expected call of GetMyProfile.
func (mr *MockRepositoryInterfaceMockRecorder) GetMyProfile(ctx, userID interface{}, anotherAttr ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, userID}, anotherAttr...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMyProfile", reflect.TypeOf((*MockRepositoryInterface)(nil).GetMyProfile), varargs...)
}

// GetTestById mocks base method.
func (m *MockRepositoryInterface) GetTestById(ctx context.Context, input GetTestByIdInput) (GetTestByIdOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTestById", ctx, input)
	ret0, _ := ret[0].(GetTestByIdOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTestById indicates an expected call of GetTestById.
func (mr *MockRepositoryInterfaceMockRecorder) GetTestById(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTestById", reflect.TypeOf((*MockRepositoryInterface)(nil).GetTestById), ctx, input)
}

// LoginUser mocks base method.
func (m *MockRepositoryInterface) LoginUser(ctx context.Context, phoneNumber, passwordHash string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginUser", ctx, phoneNumber, passwordHash)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoginUser indicates an expected call of LoginUser.
func (mr *MockRepositoryInterfaceMockRecorder) LoginUser(ctx, phoneNumber, passwordHash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginUser", reflect.TypeOf((*MockRepositoryInterface)(nil).LoginUser), ctx, phoneNumber, passwordHash)
}

// RegisterUser mocks base method.
func (m *MockRepositoryInterface) RegisterUser(ctx context.Context, phoneNumber, fullName, passwordHash string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterUser", ctx, phoneNumber, fullName, passwordHash)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterUser indicates an expected call of RegisterUser.
func (mr *MockRepositoryInterfaceMockRecorder) RegisterUser(ctx, phoneNumber, fullName, passwordHash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterUser", reflect.TypeOf((*MockRepositoryInterface)(nil).RegisterUser), ctx, phoneNumber, fullName, passwordHash)
}

// UpdateMyProfile mocks base method.
func (m *MockRepositoryInterface) UpdateMyProfile(ctx context.Context, userID int, phoneNumber, fullName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMyProfile", ctx, userID, phoneNumber, fullName)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMyProfile indicates an expected call of UpdateMyProfile.
func (mr *MockRepositoryInterfaceMockRecorder) UpdateMyProfile(ctx, userID, phoneNumber, fullName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMyProfile", reflect.TypeOf((*MockRepositoryInterface)(nil).UpdateMyProfile), ctx, userID, phoneNumber, fullName)
}
