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

// Test SendMessage Success
func TestMessageService_SendMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockMessageRepository(ctrl)
	service := services.NewMessageService(mockRepo)

	// Successful case
	mockRepo.EXPECT().
		SendMessage("1", "2", "Hello Doctor").
		Return(nil)

	err := service.SendMessage("1", "2", "Hello Doctor")
	assert.NoError(t, err)
}

// Test SendMessage Failure
func TestMessageService_SendMessage_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockMessageRepository(ctrl)
	service := services.NewMessageService(mockRepo)

	// Failure case
	mockRepo.EXPECT().
		SendMessage("1", "2", "Hello Doctor").
		Return(errors.New("failed to send message"))

	err := service.SendMessage("1", "2", "Hello Doctor")
	assert.Error(t, err)
	assert.Equal(t, "failed to send message", err.Error())
}

// Test GetUnreadMessages Success
func TestMessageService_GetUnreadMessages(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockMessageRepository(ctrl)
	service := services.NewMessageService(mockRepo)

	expectedMessages := []models.Message{
		{Sender: "1", Content: "Hello", Timestamp: []uint8("2024-09-09")},
	}

	// Successful case
	mockRepo.EXPECT().
		GetUnreadMessages("2").
		Return(expectedMessages, nil)

	messages, err := service.GetUnreadMessages("2")
	assert.NoError(t, err)
	assert.Equal(t, expectedMessages, messages)
}

// Test GetUnreadMessages Failure
func TestMessageService_GetUnreadMessages_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockMessageRepository(ctrl)
	service := services.NewMessageService(mockRepo)

	// Failure case
	mockRepo.EXPECT().
		GetUnreadMessages("2").
		Return(nil, errors.New("failed to retrieve messages"))

	messages, err := service.GetUnreadMessages("2")
	assert.Error(t, err)
	assert.Nil(t, messages)
	assert.Equal(t, "failed to retrieve messages", err.Error())
}

// Test RespondToPatient Success
func TestMessageService_RespondToPatient(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockMessageRepository(ctrl)
	service := services.NewMessageService(mockRepo)

	// Successful case
	mockRepo.EXPECT().
		RespondToPatient("2", "1", "Here's your prescription").
		Return(nil)

	err := service.RespondToPatient("2", "1", "Here's your prescription")
	assert.NoError(t, err)
}

// Test RespondToPatient Failure
func TestMessageService_RespondToPatient_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockMessageRepository(ctrl)
	service := services.NewMessageService(mockRepo)

	// Failure case
	mockRepo.EXPECT().
		RespondToPatient("2", "1", "Here's your prescription").
		Return(errors.New("failed to respond"))

	err := service.RespondToPatient("2", "1", "Here's your prescription")
	assert.Error(t, err)
	assert.Equal(t, "failed to respond", err.Error())
}

// Test GetUnreadMessagesById Success
func TestMessageService_GetUnreadMessagesById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockMessageRepository(ctrl)
	service := services.NewMessageService(mockRepo)

	expectedMessages := []models.Message{
		{Sender: "1", Content: "Hi", Timestamp: []uint8("2024-09-08")},
	}

	// Successful case
	mockRepo.EXPECT().
		GetUnreadMessagesById("1", "2").
		Return(expectedMessages, nil)

	messages, err := service.GetUnreadMessagesById("1", "2")
	assert.NoError(t, err)
	assert.Equal(t, expectedMessages, messages)
}

// Test GetUnreadMessagesById Failure
func TestMessageService_GetUnreadMessagesById_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockMessageRepository(ctrl)
	service := services.NewMessageService(mockRepo)

	// Failure case
	mockRepo.EXPECT().
		GetUnreadMessagesById("1", "2").
		Return(nil, errors.New("failed to retrieve messages"))

	messages, err := service.GetUnreadMessagesById("1", "2")
	assert.Error(t, err)
	assert.Nil(t, messages)
	assert.Equal(t, "failed to retrieve messages", err.Error())
}
