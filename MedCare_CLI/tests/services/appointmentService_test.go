package services_test

import (
	"doctor-patient-cli/models"
	"doctor-patient-cli/services"
	"doctor-patient-cli/tests/mocks"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSendAppointmentRequestService_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockAppointmentRepository(ctrl)
	service := services.NewAppointmentService(mockRepo)

	// Success case
	mockRepo.EXPECT().SendAppointmentRequest("patient1", "doctor1").Return(nil)

	err := service.SendAppointmentRequest("patient1", "doctor1")
	assert.NoError(t, err)
}

// Test case for failure in sending an appointment request
func TestSendAppointmentRequestService_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockAppointmentRepository(ctrl)
	service := services.NewAppointmentService(mockRepo)

	// Failure case
	mockRepo.EXPECT().SendAppointmentRequest("patient1", "doctor1").Return(errors.New("failed to send request"))

	err := service.SendAppointmentRequest("patient1", "doctor1")
	assert.Error(t, err)
	assert.Equal(t, "failed to send request", err.Error())
}

func TestGetAppointmentsByDoctorIDService_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockAppointmentRepository(ctrl)
	service := services.NewAppointmentService(mockRepo)

	expectedAppointments := []models.Appointment{
		{AppointmentID: 1, DoctorID: "doctor1", PatientID: "patient1", IsApproved: true},
	}

	// Success case
	mockRepo.EXPECT().GetAppointmentsByDoctorID("doctor1").Return(expectedAppointments, nil)

	appointments, err := service.GetAppointmentsByDoctorID("doctor1")
	assert.NoError(t, err)
	assert.Len(t, appointments, 1)
	assert.Equal(t, "doctor1", appointments[0].DoctorID)
}

// Test case for failure in getting appointments by doctor ID
func TestGetAppointmentsByDoctorIDService_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockAppointmentRepository(ctrl)
	service := services.NewAppointmentService(mockRepo)

	// Failure case
	mockRepo.EXPECT().GetAppointmentsByDoctorID("doctor1").Return(nil, errors.New("failed to retrieve appointments"))

	appointments, err := service.GetAppointmentsByDoctorID("doctor1")
	assert.Error(t, err)
	assert.Nil(t, appointments)
	assert.Equal(t, "failed to retrieve appointments", err.Error())
}

func TestApproveAppointmentService_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockAppointmentRepository(ctrl)
	service := services.NewAppointmentService(mockRepo)

	// Success case
	mockRepo.EXPECT().ApproveAppointment(1).Return(nil)

	err := service.ApproveAppointment(1)
	assert.NoError(t, err)
}

// Test case for failure in approving an appointment
func TestApproveAppointmentService_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockAppointmentRepository(ctrl)
	service := services.NewAppointmentService(mockRepo)

	// Failure case
	mockRepo.EXPECT().ApproveAppointment(1).Return(errors.New("failed to approve appointment"))

	err := service.ApproveAppointment(1)
	assert.Error(t, err)
	assert.Equal(t, "failed to approve appointment", err.Error())
}

func TestGetPendingAppointmentsByDoctorIDService_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockAppointmentRepository(ctrl)
	service := services.NewAppointmentService(mockRepo)

	expectedAppointments := []models.Appointment{
		{AppointmentID: 1, PatientID: "patient1"},
	}

	// Success case
	mockRepo.EXPECT().GetPendingAppointmentsByDoctorID("doctor1").Return(expectedAppointments, nil)

	appointments, err := service.GetPendingAppointmentsByDoctorID("doctor1")
	assert.NoError(t, err)
	assert.Len(t, appointments, 1)
	assert.Equal(t, "patient1", appointments[0].PatientID)
}

// Test case for failure in getting pending appointments
func TestGetPendingAppointmentsByDoctorIDService_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockAppointmentRepository(ctrl)
	service := services.NewAppointmentService(mockRepo)

	// Failure case
	mockRepo.EXPECT().GetPendingAppointmentsByDoctorID("doctor1").Return(nil, errors.New("failed to retrieve pending appointments"))

	appointments, err := service.GetPendingAppointmentsByDoctorID("doctor1")
	assert.Error(t, err)
	assert.Nil(t, appointments)
	assert.Equal(t, "failed to retrieve pending appointments", err.Error())
}
