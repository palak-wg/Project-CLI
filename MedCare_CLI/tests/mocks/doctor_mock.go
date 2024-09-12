// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces/doctor_interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	models "doctor-patient-cli/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockDoctorRepository is a mock of DoctorRepository interface.
type MockDoctorRepository struct {
	ctrl     *gomock.Controller
	recorder *MockDoctorRepositoryMockRecorder
}

// MockDoctorRepositoryMockRecorder is the mock recorder for MockDoctorRepository.
type MockDoctorRepositoryMockRecorder struct {
	mock *MockDoctorRepository
}

// NewMockDoctorRepository creates a new mock instance.
func NewMockDoctorRepository(ctrl *gomock.Controller) *MockDoctorRepository {
	mock := &MockDoctorRepository{ctrl: ctrl}
	mock.recorder = &MockDoctorRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDoctorRepository) EXPECT() *MockDoctorRepositoryMockRecorder {
	return m.recorder
}

// GetAllDoctors mocks base method.
func (m *MockDoctorRepository) GetAllDoctors() ([]models.Doctor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllDoctors")
	ret0, _ := ret[0].([]models.Doctor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllDoctors indicates an expected call of GetAllDoctors.
func (mr *MockDoctorRepositoryMockRecorder) GetAllDoctors() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllDoctors", reflect.TypeOf((*MockDoctorRepository)(nil).GetAllDoctors))
}

// GetDoctorByID mocks base method.
func (m *MockDoctorRepository) GetDoctorByID(doctorID string) (*models.Doctor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDoctorByID", doctorID)
	ret0, _ := ret[0].(*models.Doctor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDoctorByID indicates an expected call of GetDoctorByID.
func (mr *MockDoctorRepositoryMockRecorder) GetDoctorByID(doctorID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDoctorByID", reflect.TypeOf((*MockDoctorRepository)(nil).GetDoctorByID), doctorID)
}

// UpdateDoctorExperience mocks base method.
func (m *MockDoctorRepository) UpdateDoctorExperience(doctorID string, experience int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDoctorExperience", doctorID, experience)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateDoctorExperience indicates an expected call of UpdateDoctorExperience.
func (mr *MockDoctorRepositoryMockRecorder) UpdateDoctorExperience(doctorID, experience interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDoctorExperience", reflect.TypeOf((*MockDoctorRepository)(nil).UpdateDoctorExperience), doctorID, experience)
}

// UpdateDoctorSpecialization mocks base method.
func (m *MockDoctorRepository) UpdateDoctorSpecialization(doctorID, specialization string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDoctorSpecialization", doctorID, specialization)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateDoctorSpecialization indicates an expected call of UpdateDoctorSpecialization.
func (mr *MockDoctorRepositoryMockRecorder) UpdateDoctorSpecialization(doctorID, specialization interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDoctorSpecialization", reflect.TypeOf((*MockDoctorRepository)(nil).UpdateDoctorSpecialization), doctorID, specialization)
}

// ViewDoctorSpecificProfile mocks base method.
func (m *MockDoctorRepository) ViewDoctorSpecificProfile(doctorID string) (*models.Doctor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ViewDoctorSpecificProfile", doctorID)
	ret0, _ := ret[0].(*models.Doctor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ViewDoctorSpecificProfile indicates an expected call of ViewDoctorSpecificProfile.
func (mr *MockDoctorRepositoryMockRecorder) ViewDoctorSpecificProfile(doctorID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ViewDoctorSpecificProfile", reflect.TypeOf((*MockDoctorRepository)(nil).ViewDoctorSpecificProfile), doctorID)
}
//
//// MockDoctorService is a mock of DoctorService interface.
//type MockDoctorService struct {
//	ctrl     *gomock.Controller
//	recorder *MockDoctorServiceMockRecorder
//}
//
//// MockDoctorServiceMockRecorder is the mock recorder for MockDoctorService.
//type MockDoctorServiceMockRecorder struct {
//	mock *MockDoctorService
//}
//
//// NewMockDoctorService creates a new mock instance.
//func NewMockDoctorService(ctrl *gomock.Controller) *MockDoctorService {
//	mock := &MockDoctorService{ctrl: ctrl}
//	mock.recorder = &MockDoctorServiceMockRecorder{mock}
//	return mock
//}
//
//// EXPECT returns an object that allows the caller to indicate expected use.
//func (m *MockDoctorService) EXPECT() *MockDoctorServiceMockRecorder {
//	return m.recorder
//}
//
//// GetAllDoctors mocks base method.
//func (m *MockDoctorService) GetAllDoctors() ([]models.Doctor, error) {
//	m.ctrl.T.Helper()
//	ret := m.ctrl.Call(m, "GetAllDoctors")
//	ret0, _ := ret[0].([]models.Doctor)
//	ret1, _ := ret[1].(error)
//	return ret0, ret1
//}
//
//// GetAllDoctors indicates an expected call of GetAllDoctors.
//func (mr *MockDoctorServiceMockRecorder) GetAllDoctors() *gomock.Call {
//	mr.mock.ctrl.T.Helper()
//	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllDoctors", reflect.TypeOf((*MockDoctorService)(nil).GetAllDoctors))
//}
//
//// GetDoctorByID mocks base method.
//func (m *MockDoctorService) GetDoctorByID(doctorID string) (*models.Doctor, error) {
//	m.ctrl.T.Helper()
//	ret := m.ctrl.Call(m, "GetDoctorByID", doctorID)
//	ret0, _ := ret[0].(*models.Doctor)
//	ret1, _ := ret[1].(error)
//	return ret0, ret1
//}
//
//// GetDoctorByID indicates an expected call of GetDoctorByID.
//func (mr *MockDoctorServiceMockRecorder) GetDoctorByID(doctorID interface{}) *gomock.Call {
//	mr.mock.ctrl.T.Helper()
//	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDoctorByID", reflect.TypeOf((*MockDoctorService)(nil).GetDoctorByID), doctorID)
//}
//
//// UpdateDoctorExperience mocks base method.
//func (m *MockDoctorService) UpdateDoctorExperience(doctorID string, experience int) error {
//	m.ctrl.T.Helper()
//	ret := m.ctrl.Call(m, "UpdateDoctorExperience", doctorID, experience)
//	ret0, _ := ret[0].(error)
//	return ret0
//}
//
//// UpdateDoctorExperience indicates an expected call of UpdateDoctorExperience.
//func (mr *MockDoctorServiceMockRecorder) UpdateDoctorExperience(doctorID, experience interface{}) *gomock.Call {
//	mr.mock.ctrl.T.Helper()
//	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDoctorExperience", reflect.TypeOf((*MockDoctorService)(nil).UpdateDoctorExperience), doctorID, experience)
//}
//
//// UpdateDoctorSpecialization mocks base method.
//func (m *MockDoctorService) UpdateDoctorSpecialization(doctorID, specialization string) error {
//	m.ctrl.T.Helper()
//	ret := m.ctrl.Call(m, "UpdateDoctorSpecialization", doctorID, specialization)
//	ret0, _ := ret[0].(error)
//	return ret0
//}
//
//// UpdateDoctorSpecialization indicates an expected call of UpdateDoctorSpecialization.
//func (mr *MockDoctorServiceMockRecorder) UpdateDoctorSpecialization(doctorID, specialization interface{}) *gomock.Call {
//	mr.mock.ctrl.T.Helper()
//	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDoctorSpecialization", reflect.TypeOf((*MockDoctorService)(nil).UpdateDoctorSpecialization), doctorID, specialization)
//}
//
//// ViewDoctorSpecificProfile mocks base method.
//func (m *MockDoctorService) ViewDoctorSpecificProfile(doctorID string) (*models.Doctor, error) {
//	m.ctrl.T.Helper()
//	ret := m.ctrl.Call(m, "ViewDoctorSpecificProfile", doctorID)
//	ret0, _ := ret[0].(*models.Doctor)
//	ret1, _ := ret[1].(error)
//	return ret0, ret1
//}
//
//// ViewDoctorSpecificProfile indicates an expected call of ViewDoctorSpecificProfile.
//func (mr *MockDoctorServiceMockRecorder) ViewDoctorSpecificProfile(doctorID interface{}) *gomock.Call {
//	mr.mock.ctrl.T.Helper()
//	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ViewDoctorSpecificProfile", reflect.TypeOf((*MockDoctorService)(nil).ViewDoctorSpecificProfile), doctorID)
//}
