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

// TestApproveDoctorSignup_Success tests the full success path
func TestApproveDoctorSignup_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAdminRepo := mocks.NewMockAdminRepository(ctrl)
	mockUserRepo := mocks.NewMockUserRepository(ctrl)

	service := services.NewAdminService(mockAdminRepo, mockUserRepo)

	user := models.User{
		UserID: "1",
		Email:  "test@example.com",
	}

	// Mock repository responses
	mockAdminRepo.EXPECT().ApproveDoctorSignup("1").Return(nil)
	mockUserRepo.EXPECT().GetUserByID("1").Return(&user, nil)
	mockAdminRepo.EXPECT().CreateNotificationForUser("1", gomock.Any()).Return(nil)

	// Act
	err := service.ApproveDoctorSignup("1")

	// Assert
	assert.NoError(t, err)
}

// TestApproveDoctorSignup_ApproveFail tests failure during doctor approval
func TestApproveDoctorSignup_ApproveFail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAdminRepo := mocks.NewMockAdminRepository(ctrl)
	mockUserRepo := mocks.NewMockUserRepository(ctrl)

	service := services.NewAdminService(mockAdminRepo, mockUserRepo)

	// Mock the failure in approving the doctor signup
	mockAdminRepo.EXPECT().ApproveDoctorSignup("1").Return(errors.New("db error"))

	// Act
	err := service.ApproveDoctorSignup("1")

	// Assert
	assert.Error(t, err)
	assert.Equal(t, "db error", err.Error())
}

// TestApproveDoctorSignup_UserFetchFail tests failure in fetching user details
func TestApproveDoctorSignup_UserFetchFail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAdminRepo := mocks.NewMockAdminRepository(ctrl)
	mockUserRepo := mocks.NewMockUserRepository(ctrl)

	service := services.NewAdminService(mockAdminRepo, mockUserRepo)

	// Mock success in doctor signup approval
	mockAdminRepo.EXPECT().ApproveDoctorSignup("1").Return(nil)

	// Mock failure in fetching user details
	mockUserRepo.EXPECT().GetUserByID("1").Return(nil, errors.New("user not found"))

	// Act
	err := service.ApproveDoctorSignup("1")

	// Assert
	assert.Error(t, err)
	assert.Equal(t, "user not found", err.Error())
}

// TestApproveDoctorSignup_NotificationFail tests failure in creating notification
func TestApproveDoctorSignup_NotificationFail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAdminRepo := mocks.NewMockAdminRepository(ctrl)
	mockUserRepo := mocks.NewMockUserRepository(ctrl)

	service := services.NewAdminService(mockAdminRepo, mockUserRepo)

	user := models.User{
		UserID: "1",
		Email:  "test@example.com",
	}

	// Mock success in doctor signup approval and user fetch
	mockAdminRepo.EXPECT().ApproveDoctorSignup("1").Return(nil)
	mockUserRepo.EXPECT().GetUserByID("1").Return(&user, nil)

	// Mock failure in creating notification
	mockAdminRepo.EXPECT().CreateNotificationForUser("1", gomock.Any()).Return(errors.New("notification error"))

	// Act
	err := service.ApproveDoctorSignup("1")

	// Assert
	assert.Error(t, err)
	assert.Equal(t, "notification error", err.Error())
}

func TestGetPendingDoctorRequests_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAdminRepo := mocks.NewMockAdminRepository(ctrl)
	service := services.NewAdminService(mockAdminRepo, nil)

	doctors := []models.Doctor{
		{User: models.User{UserID: "1", Name: "Doctor A"}},
		{User: models.User{UserID: "2", Name: "Doctor B"}},
	}

	// Mock the repository response
	mockAdminRepo.EXPECT().PendingDoctorSignupRequest().Return(doctors, nil)

	// Act
	result, err := service.GetPendingDoctorRequests()

	// Assert
	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "Doctor A", result[0].Name)
	assert.Equal(t, "Doctor B", result[1].Name)
}

// TestGetPendingDoctorRequests_Failure tests failure in fetching pending doctor requests
func TestGetPendingDoctorRequests_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAdminRepo := mocks.NewMockAdminRepository(ctrl)
	service := services.NewAdminService(mockAdminRepo, nil)

	// Mock the failure case
	mockAdminRepo.EXPECT().PendingDoctorSignupRequest().Return(nil, errors.New("db error"))

	// Act
	result, err := service.GetPendingDoctorRequests()

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "db error", err.Error())
}

func TestGetAllUsers_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAdminRepo := mocks.NewMockAdminRepository(ctrl)
	service := services.NewAdminService(mockAdminRepo, nil)

	users := []models.User{
		{UserID: "1", Name: "User A"},
		{UserID: "2", Name: "User B"},
	}

	// Mock the repository response
	mockAdminRepo.EXPECT().GetAllUsers().Return(users, nil)

	// Act
	result, err := service.GetAllUsers()

	// Assert
	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "User A", result[0].Name)
	assert.Equal(t, "User B", result[1].Name)
}

// TestGetAllUsers_Failure tests failure in fetching all users
func TestGetAllUsers_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAdminRepo := mocks.NewMockAdminRepository(ctrl)
	service := services.NewAdminService(mockAdminRepo, nil)

	// Mock the failure case
	mockAdminRepo.EXPECT().GetAllUsers().Return(nil, errors.New("db error"))

	// Act
	result, err := service.GetAllUsers()

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "db error", err.Error())
}
