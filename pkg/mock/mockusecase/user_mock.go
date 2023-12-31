// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/usecase/interfaces/user.go
//
// Generated by this command:
//
//	mockgen -source=pkg/usecase/interfaces/user.go -destination=pkg/mock/mockusecase/user_mock.go -package=mockusecase
//
// Package mockusecase is a generated GoMock package.
package mockusecase

import (
	reflect "reflect"

	domain "github.com/karthikkalarikal/ecommerce-project/pkg/domain"
	models "github.com/karthikkalarikal/ecommerce-project/pkg/utils/models"
	gomock "go.uber.org/mock/gomock"
)

// MockUserUseCase is a mock of UserUseCase interface.
type MockUserUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUserUseCaseMockRecorder
}

// MockUserUseCaseMockRecorder is the mock recorder for MockUserUseCase.
type MockUserUseCaseMockRecorder struct {
	mock *MockUserUseCase
}

// NewMockUserUseCase creates a new mock instance.
func NewMockUserUseCase(ctrl *gomock.Controller) *MockUserUseCase {
	mock := &MockUserUseCase{ctrl: ctrl}
	mock.recorder = &MockUserUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserUseCase) EXPECT() *MockUserUseCaseMockRecorder {
	return m.recorder
}

// AddAddress mocks base method.
func (m *MockUserUseCase) AddAddress(arg0 models.Address, arg1 int) ([]models.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAddress", arg0, arg1)
	ret0, _ := ret[0].([]models.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddAddress indicates an expected call of AddAddress.
func (mr *MockUserUseCaseMockRecorder) AddAddress(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAddress", reflect.TypeOf((*MockUserUseCase)(nil).AddAddress), arg0, arg1)
}

// ChangePassword mocks base method.
func (m *MockUserUseCase) ChangePassword(pass models.ChangePassword, userId int) (models.UserSignInResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangePassword", pass, userId)
	ret0, _ := ret[0].(models.UserSignInResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangePassword indicates an expected call of ChangePassword.
func (mr *MockUserUseCaseMockRecorder) ChangePassword(pass, userId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangePassword", reflect.TypeOf((*MockUserUseCase)(nil).ChangePassword), pass, userId)
}

// EditUserDetails mocks base method.
func (m *MockUserUseCase) EditUserDetails(arg0 int, arg1 models.UserDetailsResponse) (models.UserDetailsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditUserDetails", arg0, arg1)
	ret0, _ := ret[0].(models.UserDetailsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EditUserDetails indicates an expected call of EditUserDetails.
func (mr *MockUserUseCaseMockRecorder) EditUserDetails(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditUserDetails", reflect.TypeOf((*MockUserUseCase)(nil).EditUserDetails), arg0, arg1)
}

// FindAddressByUI mocks base method.
func (m *MockUserUseCase) FindAddressByUI(arg0 int) ([]models.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAddressByUI", arg0)
	ret0, _ := ret[0].([]models.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAddressByUI indicates an expected call of FindAddressByUI.
func (mr *MockUserUseCaseMockRecorder) FindAddressByUI(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAddressByUI", reflect.TypeOf((*MockUserUseCase)(nil).FindAddressByUI), arg0)
}

// FindUserById mocks base method.
func (m *MockUserUseCase) FindUserById(arg0 int) (domain.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserById", arg0)
	ret0, _ := ret[0].(domain.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserById indicates an expected call of FindUserById.
func (mr *MockUserUseCaseMockRecorder) FindUserById(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserById", reflect.TypeOf((*MockUserUseCase)(nil).FindUserById), arg0)
}

// LoginHandler mocks base method.
func (m *MockUserUseCase) LoginHandler(user models.UserLogin) (any, error, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginHandler", user)
	ret0, _ := ret[0].(any)
	ret1, _ := ret[1].(error)
	ret2, _ := ret[2].(bool)
	return ret0, ret1, ret2
}

// LoginHandler indicates an expected call of LoginHandler.
func (mr *MockUserUseCaseMockRecorder) LoginHandler(user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginHandler", reflect.TypeOf((*MockUserUseCase)(nil).LoginHandler), user)
}

// SelectAddress mocks base method.
func (m *MockUserUseCase) SelectAddress(arg0 int, arg1 bool) (models.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectAddress", arg0, arg1)
	ret0, _ := ret[0].(models.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectAddress indicates an expected call of SelectAddress.
func (mr *MockUserUseCaseMockRecorder) SelectAddress(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectAddress", reflect.TypeOf((*MockUserUseCase)(nil).SelectAddress), arg0, arg1)
}

// UserSignUp mocks base method.
func (m *MockUserUseCase) UserSignUp(User models.UserDetails) (models.TokenUsers, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserSignUp", User)
	ret0, _ := ret[0].(models.TokenUsers)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserSignUp indicates an expected call of UserSignUp.
func (mr *MockUserUseCaseMockRecorder) UserSignUp(User any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserSignUp", reflect.TypeOf((*MockUserUseCase)(nil).UserSignUp), User)
}
