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

func TestGetDoctorByIDService_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockDoctorRepository(ctrl)
	service := services.NewDoctorService(mockRepo)

	user := models.User{
		UserID: "doctor1",
	}
	expectedDoctor := &models.Doctor{
		User:           user,
		Specialization: "Cardiology",
		Experience:     10,
		Rating:         4.5,
	}

	// Success case
	mockRepo.EXPECT().GetDoctorByID("doctor1").Return(expectedDoctor, nil)

	doctor, err := service.GetDoctorByID("doctor1")
	assert.NoError(t, err)
	assert.NotNil(t, doctor)
	assert.Equal(t, "doctor1", doctor.User.UserID)
}

// Test case when doctor ID is not found
func TestGetDoctorByIDService_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockDoctorRepository(ctrl)
	service := services.NewDoctorService(mockRepo)

	// Failure case
	mockRepo.EXPECT().GetDoctorByID("doctor1").Return(nil, errors.New("doctor not found"))

	doctor, err := service.GetDoctorByID("doctor1")
	assert.Error(t, err)
	assert.Nil(t, doctor)
	assert.Equal(t, "doctor not found", err.Error())
}

func TestGetAllDoctorsService_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockDoctorRepository(ctrl)
	service := services.NewDoctorService(mockRepo)

	expectedDoctors := []models.Doctor{
		{User: models.User{UserID: "doctor1"}, Specialization: "Cardiology", Experience: 10, Rating: 4.5},
		{User: models.User{UserID: "doctor2"}, Specialization: "Dermatology", Experience: 5, Rating: 4.0},
	}

	// Success case
	mockRepo.EXPECT().GetAllDoctors().Return(expectedDoctors, nil)

	doctors, err := service.GetAllDoctors()
	assert.NoError(t, err)
	assert.Len(t, doctors, 2)
	assert.Equal(t, "doctor1", doctors[0].User.UserID)
	assert.Equal(t, "doctor2", doctors[1].User.UserID)
}

// Test case for empty list of doctors
func TestGetAllDoctorsService_Empty(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockDoctorRepository(ctrl)
	service := services.NewDoctorService(mockRepo)

	// Empty case
	mockRepo.EXPECT().GetAllDoctors().Return([]models.Doctor{}, nil)

	doctors, err := service.GetAllDoctors()
	assert.NoError(t, err)
	assert.Len(t, doctors, 0)
}

// Test case for failure in retrieving all doctors
func TestGetAllDoctorsService_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockDoctorRepository(ctrl)
	service := services.NewDoctorService(mockRepo)

	// Failure case
	mockRepo.EXPECT().GetAllDoctors().Return(nil, errors.New("failed to retrieve doctors"))

	doctors, err := service.GetAllDoctors()
	assert.Error(t, err)
	assert.Nil(t, doctors)
	assert.Equal(t, "failed to retrieve doctors", err.Error())
}

func TestUpdateDoctorExperienceService_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockDoctorRepository(ctrl)
	service := services.NewDoctorService(mockRepo)

	// Success case
	mockRepo.EXPECT().UpdateDoctorExperience("doctor1", 15).Return(nil)

	err := service.UpdateDoctorExperience("doctor1", 15)
	assert.NoError(t, err)
}

// Test case for failure in updating doctor experience
func TestUpdateDoctorExperienceService_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockDoctorRepository(ctrl)
	service := services.NewDoctorService(mockRepo)

	// Failure case
	mockRepo.EXPECT().UpdateDoctorExperience("doctor1", 15).Return(errors.New("failed to update experience"))

	err := service.UpdateDoctorExperience("doctor1", 15)
	assert.Error(t, err)
	assert.Equal(t, "failed to update experience", err.Error())
}

func TestUpdateDoctorSpecializationService_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockDoctorRepository(ctrl)
	service := services.NewDoctorService(mockRepo)

	// Success case
	mockRepo.EXPECT().UpdateDoctorSpecialization("doctor1", "Neurology").Return(nil)

	err := service.UpdateDoctorSpecialization("doctor1", "Neurology")
	assert.NoError(t, err)
}

// Test case for failure in updating doctor specialization
func TestUpdateDoctorSpecializationService_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockDoctorRepository(ctrl)
	service := services.NewDoctorService(mockRepo)

	// Failure case
	mockRepo.EXPECT().UpdateDoctorSpecialization("doctor1", "Neurology").Return(errors.New("failed to update specialization"))

	err := service.UpdateDoctorSpecialization("doctor1", "Neurology")
	assert.Error(t, err)
	assert.Equal(t, "failed to update specialization", err.Error())
}

func TestViewDoctorSpecificProfileService_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockDoctorRepository(ctrl)
	service := services.NewDoctorService(mockRepo)

	expectedDoctor := &models.Doctor{
		User:           models.User{UserID: "doctor1"},
		Specialization: "Cardiology",
		Experience:     10,
		Rating:         4.5,
	}

	// Success case
	mockRepo.EXPECT().ViewDoctorSpecificProfile("doctor1").Return(expectedDoctor, nil)

	doctor, err := service.ViewDoctorSpecificProfile("doctor1")
	assert.NoError(t, err)
	assert.NotNil(t, doctor)
	assert.Equal(t, "doctor1", doctor.User.UserID)
}

// Test case for failure in viewing specific profile
func TestViewDoctorSpecificProfileService_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockDoctorRepository(ctrl)
	service := services.NewDoctorService(mockRepo)

	// Failure case
	mockRepo.EXPECT().ViewDoctorSpecificProfile("doctor1").Return(nil, errors.New("doctor not found"))

	doctor, err := service.ViewDoctorSpecificProfile("doctor1")
	assert.Error(t, err)
	assert.Nil(t, doctor)
	assert.Equal(t, "doctor not found", err.Error())
}
