package mocks

import (
	models "doctor-patient-cli/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAdminRepository is a mock of AdminRepository interface.
type MockAdminRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAdminRepositoryMockRecorder
}

// MockAdminRepositoryMockRecorder is the mock recorder for MockAdminRepository.
type MockAdminRepositoryMockRecorder struct {
	mock *MockAdminRepository
}

// NewMockAdminRepository creates a new mock instance.
func NewMockAdminRepository(ctrl *gomock.Controller) *MockAdminRepository {
	mock := &MockAdminRepository{ctrl: ctrl}
	mock.recorder = &MockAdminRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAdminRepository) EXPECT() *MockAdminRepositoryMockRecorder {
	return m.recorder
}

// ApproveDoctorSignup mocks base method.
func (m *MockAdminRepository) ApproveDoctorSignup(userID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApproveDoctorSignup", userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApproveDoctorSignup indicates an expected call of ApproveDoctorSignup.
func (mr *MockAdminRepositoryMockRecorder) ApproveDoctorSignup(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApproveDoctorSignup", reflect.TypeOf((*MockAdminRepository)(nil).ApproveDoctorSignup), userID)
}

// CreateNotificationForUser mocks base method.
func (m *MockAdminRepository) CreateNotificationForUser(userID, content string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNotificationForUser", userID, content)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateNotificationForUser indicates an expected call of CreateNotificationForUser.
func (mr *MockAdminRepositoryMockRecorder) CreateNotificationForUser(userID, content interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNotificationForUser", reflect.TypeOf((*MockAdminRepository)(nil).CreateNotificationForUser), userID, content)
}

// GetAllUsers mocks base method.
func (m *MockAdminRepository) GetAllUsers() ([]models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUsers")
	ret0, _ := ret[0].([]models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUsers indicates an expected call of GetAllUsers.
func (mr *MockAdminRepositoryMockRecorder) GetAllUsers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUsers", reflect.TypeOf((*MockAdminRepository)(nil).GetAllUsers))
}

// PendingDoctorSignupRequest mocks base method.
func (m *MockAdminRepository) PendingDoctorSignupRequest() ([]models.Doctor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PendingDoctorSignupRequest")
	ret0, _ := ret[0].([]models.Doctor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PendingDoctorSignupRequest indicates an expected call of PendingDoctorSignupRequest.
func (mr *MockAdminRepositoryMockRecorder) PendingDoctorSignupRequest() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PendingDoctorSignupRequest", reflect.TypeOf((*MockAdminRepository)(nil).PendingDoctorSignupRequest))
}

// MockAdminService is a mock of AdminService interface.
type MockAdminService struct {
	ctrl     *gomock.Controller
	recorder *MockAdminServiceMockRecorder
}

// MockAdminServiceMockRecorder is the mock recorder for MockAdminService.
type MockAdminServiceMockRecorder struct {
	mock *MockAdminService
}

//// NewMockAdminService creates a new mock instance.
//func NewMockAdminService(ctrl *gomock.Controller) *MockAdminService {
//	mock := &MockAdminService{ctrl: ctrl}
//	mock.recorder = &MockAdminServiceMockRecorder{mock}
//	return mock
//}
//
//// EXPECT returns an object that allows the caller to indicate expected use.
//func (m *MockAdminService) EXPECT() *MockAdminServiceMockRecorder {
//	return m.recorder
//}
//
//// ApproveDoctorSignup mocks base method.
//func (m *MockAdminService) ApproveDoctorSignup(userID string) error {
//	m.ctrl.T.Helper()
//	ret := m.ctrl.Call(m, "ApproveDoctorSignup", userID)
//	ret0, _ := ret[0].(error)
//	return ret0
//}
//
//// ApproveDoctorSignup indicates an expected call of ApproveDoctorSignup.
//func (mr *MockAdminServiceMockRecorder) ApproveDoctorSignup(userID interface{}) *gomock.Call {
//	mr.mock.ctrl.T.Helper()
//	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApproveDoctorSignup", reflect.TypeOf((*MockAdminService)(nil).ApproveDoctorSignup), userID)
//}
//
//// GetAllUsers mocks base method.
//func (m *MockAdminService) GetAllUsers() ([]models.User, error) {
//	m.ctrl.T.Helper()
//	ret := m.ctrl.Call(m, "GetAllUsers")
//	ret0, _ := ret[0].([]models.User)
//	ret1, _ := ret[1].(error)
//	return ret0, ret1
//}
//
//// GetAllUsers indicates an expected call of GetAllUsers.
//func (mr *MockAdminServiceMockRecorder) GetAllUsers() *gomock.Call {
//	mr.mock.ctrl.T.Helper()
//	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUsers", reflect.TypeOf((*MockAdminService)(nil).GetAllUsers))
//}
//
//// PendingDoctorSignupRequest mocks base method.
//func (m *MockAdminService) PendingDoctorSignupRequest() ([]models.Doctor, error) {
//	m.ctrl.T.Helper()
//	ret := m.ctrl.Call(m, "PendingDoctorSignupRequest")
//	ret0, _ := ret[0].([]models.Doctor)
//	ret1, _ := ret[1].(error)
//	return ret0, ret1
//}
//
//// PendingDoctorSignupRequest indicates an expected call of PendingDoctorSignupRequest.
//func (mr *MockAdminServiceMockRecorder) PendingDoctorSignupRequest() *gomock.Call {
//	mr.mock.ctrl.T.Helper()
//	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PendingDoctorSignupRequest", reflect.TypeOf((*MockAdminService)(nil).PendingDoctorSignupRequest))
//}
