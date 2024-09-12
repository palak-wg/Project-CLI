package services_test

import (
	"doctor-patient-cli/services"
	"testing"

	"doctor-patient-cli/models"
	"doctor-patient-cli/tests/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)
	service := services.NewUserService(mockRepo)

	// Define the user to be created
	user := &models.User{
		UserID:      "testuser",
		Password:    "password",
		Name:        "Test User",
		Age:         30,
		Gender:      "Male",
		Email:       "test@example.com",
		PhoneNumber: "1234567890",
		UserType:    "patient",
		IsApproved:  true,
	}

	// Expect CreateUser to be called
	mockRepo.EXPECT().CreateUser(user).Return(nil)

	// Test CreateUser
	err := service.CreateUser(user)
	assert.NoError(t, err)
}

func TestGetUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)
	service := services.NewUserService(mockRepo)

	// Define the user to be retrieved
	userID := "testuser"
	user := &models.User{
		UserID:      userID,
		Password:    "password",
		Name:        "Test User",
		Age:         30,
		Gender:      "Male",
		Email:       "test@example.com",
		PhoneNumber: "1234567890",
		UserType:    "patient",
		IsApproved:  true,
	}

	// Expect GetUserByID to be called
	mockRepo.EXPECT().GetUserByID(userID).Return(user, nil)

	// Test GetUserByID
	result, err := service.GetUserByID(userID)
	assert.NoError(t, err)
	assert.Equal(t, user, result)
}

func TestUpdateName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)
	service := services.NewUserService(mockRepo)

	// Define the parameters for the update
	userID := "testuser"
	newName := "Updated Name"

	// Expect UpdateName to be called
	mockRepo.EXPECT().UpdateName(userID, newName).Return(nil)

	// Test UpdateName
	err := service.UpdateName(userID, newName)
	assert.NoError(t, err)
}

func TestUpdateAge(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)
	service := services.NewUserService(mockRepo)

	// Define the parameters for the update
	userID := "testuser"
	newAge := "31"

	// Expect UpdateAge to be called
	mockRepo.EXPECT().UpdateAge(userID, newAge).Return(nil)

	// Test UpdateAge
	err := service.UpdateAge(userID, newAge)
	assert.NoError(t, err)
}

func TestUpdateEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)
	service := services.NewUserService(mockRepo)

	// Define the parameters for the update
	userID := "testuser"
	newEmail := "updated@example.com"

	// Expect UpdateEmail to be called
	mockRepo.EXPECT().UpdateEmail(userID, newEmail).Return(nil)

	// Test UpdateEmail
	err := service.UpdateEmail(userID, newEmail)
	assert.NoError(t, err)
}

func TestUpdatePhoneNumber(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)
	service := services.NewUserService(mockRepo)

	// Define the parameters for the update
	userID := "testuser"
	newPhoneNumber := "0987654321"

	// Expect UpdatePhoneNumber to be called
	mockRepo.EXPECT().UpdatePhoneNumber(userID, newPhoneNumber).Return(nil)

	// Test UpdatePhoneNumber
	err := service.UpdatePhoneNumber(userID, newPhoneNumber)
	assert.NoError(t, err)
}

func TestUpdateGender(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)
	service := services.NewUserService(mockRepo)

	// Define the parameters for the update
	userID := "testuser"
	newGender := "Female"

	// Expect UpdateGender to be called
	mockRepo.EXPECT().UpdateGender(userID, newGender).Return(nil)

	// Test UpdateGender
	err := service.UpdateGender(userID, newGender)
	assert.NoError(t, err)
}

func TestUpdatePassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)
	service := services.NewUserService(mockRepo)

	// Define the parameters for the update
	userID := "testuser"
	newPassword := "newpassword"

	// Expect UpdatePassword to be called
	mockRepo.EXPECT().UpdatePassword(userID, newPassword).Return(nil)

	// Test UpdatePassword
	err := service.UpdatePassword(userID, newPassword)
	assert.NoError(t, err)
}
