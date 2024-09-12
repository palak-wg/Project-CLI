// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces/user_interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	models "doctor-patient-cli/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUserRepository) CreateUser(user *models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserRepositoryMockRecorder) CreateUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserRepository)(nil).CreateUser), user)
}

// GetUserByID mocks base method.
func (m *MockUserRepository) GetUserByID(userID string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", userID)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockUserRepositoryMockRecorder) GetUserByID(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockUserRepository)(nil).GetUserByID), userID)
}

// UpdateAge mocks base method.
func (m *MockUserRepository) UpdateAge(userID, age string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAge", userID, age)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAge indicates an expected call of UpdateAge.
func (mr *MockUserRepositoryMockRecorder) UpdateAge(userID, age interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAge", reflect.TypeOf((*MockUserRepository)(nil).UpdateAge), userID, age)
}

// UpdateEmail mocks base method.
func (m *MockUserRepository) UpdateEmail(userID, email string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEmail", userID, email)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateEmail indicates an expected call of UpdateEmail.
func (mr *MockUserRepositoryMockRecorder) UpdateEmail(userID, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEmail", reflect.TypeOf((*MockUserRepository)(nil).UpdateEmail), userID, email)
}

// UpdateGender mocks base method.
func (m *MockUserRepository) UpdateGender(userID, gender string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateGender", userID, gender)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateGender indicates an expected call of UpdateGender.
func (mr *MockUserRepositoryMockRecorder) UpdateGender(userID, gender interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateGender", reflect.TypeOf((*MockUserRepository)(nil).UpdateGender), userID, gender)
}

// UpdateName mocks base method.
func (m *MockUserRepository) UpdateName(userID, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateName", userID, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateName indicates an expected call of UpdateName.
func (mr *MockUserRepositoryMockRecorder) UpdateName(userID, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateName", reflect.TypeOf((*MockUserRepository)(nil).UpdateName), userID, name)
}

// UpdatePassword mocks base method.
func (m *MockUserRepository) UpdatePassword(userID, newPassword string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePassword", userID, newPassword)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePassword indicates an expected call of UpdatePassword.
func (mr *MockUserRepositoryMockRecorder) UpdatePassword(userID, newPassword interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePassword", reflect.TypeOf((*MockUserRepository)(nil).UpdatePassword), userID, newPassword)
}

// UpdatePhoneNumber mocks base method.
func (m *MockUserRepository) UpdatePhoneNumber(userID, phoneNumber string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePhoneNumber", userID, phoneNumber)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePhoneNumber indicates an expected call of UpdatePhoneNumber.
func (mr *MockUserRepositoryMockRecorder) UpdatePhoneNumber(userID, phoneNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePhoneNumber", reflect.TypeOf((*MockUserRepository)(nil).UpdatePhoneNumber), userID, phoneNumber)
}

// MockUserService is a mock of UserService interface.
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService.
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

//// NewMockUserService creates a new mock instance.
//func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
//	mock := &MockUserService{ctrl: ctrl}
//	mock.recorder = &MockUserServiceMockRecorder{mock}
//	return mock
//}
//
//// EXPECT returns an object that allows the caller to indicate expected use.
//func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
//	return m.recorder
//}
//
//// CreateUser mocks base method.
//func (m *MockUserService) CreateUser(user *models.User) error {
//	m.ctrl.T.Helper()
//	ret := m.ctrl.Call(m, "CreateUser", user)
//	ret0, _ := ret[0].(error)
//	return ret0
//}
//
//// CreateUser indicates an expected call of CreateUser.
//func (mr *MockUserServiceMockRecorder) CreateUser(user interface{}) *gomock.Call {
//	mr.mock.ctrl.T.Helper()
//	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserService)(nil).CreateUser), user)
//}
//
//// GetUserByID mocks base method.
//func (m *MockUserService) GetUserByID(userID string) (*models.User, error) {
//	m.ctrl.T.Helper()
//	ret := m.ctrl.Call(m, "GetUserByID", userID)
//	ret0, _ := ret[0].(*models.User)
//	ret1, _ := ret[1].(error)
//	return ret0, ret1
//}
//
//// GetUserByID indicates an expected call of GetUserByID.
//func (mr *MockUserServiceMockRecorder) GetUserByID(userID interface{}) *gomock.Call {
//	mr.mock.ctrl.T.Helper()
//	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockUserService)(nil).GetUserByID), userID)
//}
//
//// UpdateAge mocks base method.
//func (m *MockUserService) UpdateAge(userID, age string) error {
//	m.ctrl.T.Helper()
//	ret := m.ctrl.Call(m, "UpdateAge", userID, age)
//	ret0, _ := ret[0].(error)
//	return ret0
//}
//
//// UpdateAge indicates an expected call of UpdateAge.
//func (mr *MockUserServiceMockRecorder) UpdateAge(userID, age interface{}) *gomock.Call {
//	mr.mock.ctrl.T.Helper()
//	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAge", reflect.TypeOf((*MockUserService)(nil).UpdateAge), userID, age)
//}
//
//// UpdateEmail mocks base method.
//func (m *MockUserService) UpdateEmail(userID, email string) error {
//	m.ctrl.T.Helper()
//	ret := m.ctrl.Call(m, "UpdateEmail", userID, email)
//	ret0, _ := ret[0].(error)
//	return ret0
//}
//
//// UpdateEmail indicates an expected call of UpdateEmail.
//func (mr *MockUserServiceMockRecorder) UpdateEmail(userID, email interface{}) *gomock.Call {
//	mr.mock.ctrl.T.Helper()
//	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEmail", reflect.TypeOf((*MockUserService)(nil).UpdateEmail), userID, email)
//}
//
//// UpdateGender mocks base method.
//func (m *MockUserService) UpdateGender(userID, gender string) error {
//	m.ctrl.T.Helper()
//	ret := m.ctrl.Call(m, "UpdateGender", userID, gender)
//	ret0, _ := ret[0].(error)
//	return ret0
//}
//
//// UpdateGender indicates an expected call of UpdateGender.
//func (mr *MockUserServiceMockRecorder) UpdateGender(userID, gender interface{}) *gomock.Call {
//	mr.mock.ctrl.T.Helper()
//	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateGender", reflect.TypeOf((*MockUserService)(nil).UpdateGender), userID, gender)
//}
//
//// UpdateName mocks base method.
//func (m *MockUserService) UpdateName(userID, name string) error {
//	m.ctrl.T.Helper()
//	ret := m.ctrl.Call(m, "UpdateName", userID, name)
//	ret0, _ := ret[0].(error)
//	return ret0
//}
//
//// UpdateName indicates an expected call of UpdateName.
//func (mr *MockUserServiceMockRecorder) UpdateName(userID, name interface{}) *gomock.Call {
//	mr.mock.ctrl.T.Helper()
//	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateName", reflect.TypeOf((*MockUserService)(nil).UpdateName), userID, name)
//}
//
//// UpdatePassword mocks base method.
//func (m *MockUserService) UpdatePassword(userID, newPassword string) error {
//	m.ctrl.T.Helper()
//	ret := m.ctrl.Call(m, "UpdatePassword", userID, newPassword)
//	ret0, _ := ret[0].(error)
//	return ret0
//}
//
//// UpdatePassword indicates an expected call of UpdatePassword.
//func (mr *MockUserServiceMockRecorder) UpdatePassword(userID, newPassword interface{}) *gomock.Call {
//	mr.mock.ctrl.T.Helper()
//	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePassword", reflect.TypeOf((*MockUserService)(nil).UpdatePassword), userID, newPassword)
//}
//
//// UpdatePhoneNumber mocks base method.
//func (m *MockUserService) UpdatePhoneNumber(userID, phoneNumber string) error {
//	m.ctrl.T.Helper()
//	ret := m.ctrl.Call(m, "UpdatePhoneNumber", userID, phoneNumber)
//	ret0, _ := ret[0].(error)
//	return ret0
//}
//
//// UpdatePhoneNumber indicates an expected call of UpdatePhoneNumber.
//func (mr *MockUserServiceMockRecorder) UpdatePhoneNumber(userID, phoneNumber interface{}) *gomock.Call {
//	mr.mock.ctrl.T.Helper()
//	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePhoneNumber", reflect.TypeOf((*MockUserService)(nil).UpdatePhoneNumber), userID, phoneNumber)
//}
