package handlers_test

import (
	"doctor-patient-cli/handlers"
	"doctor-patient-cli/models"
	"doctor-patient-cli/tests/mocks"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetMessages_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockMessageService(ctrl)
	handler := handlers.NewMessageHandler(mockService)

	mockMessages := []models.Message{
		{Sender: "user1", Receiver: "user123", Content: "Hello"},
	}
	mockService.EXPECT().GetUnreadMessages("user123").Return(mockMessages, nil)

	req, err := http.NewRequest(http.MethodGet, "/messages", nil)
	require.NoError(t, err)
	req.Header.Set("Authorization", "Bearer mockToken")
	rr := httptest.NewRecorder()

	handler.GetMessages(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response models.APIResponse
	err = json.NewDecoder(rr.Body).Decode(&response)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, response.Status)
	assert.Equal(t, mockMessages, response.Data)
}

func TestGetMessages_Unauthorized(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockMessageService(ctrl)
	handler := handlers.NewMessageHandler(mockService)

	req, err := http.NewRequest(http.MethodGet, "/messages", nil)
	require.NoError(t, err)
	req.Header.Set("Authorization", "Bearer mockToken")
	rr := httptest.NewRecorder()

	handler.GetMessages(rr, req)

	assert.Equal(t, http.StatusUnauthorized, rr.Code)

	var response models.APIResponse
	err = json.NewDecoder(rr.Body).Decode(&response)
	require.NoError(t, err)

	assert.Equal(t, http.StatusUnauthorized, response.Status)
}

func TestAddMessage_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockMessageService(ctrl)
	handler := handlers.NewMessageHandler(mockService)

	message := `{"sender": "user123", "receiver": "user456", "content": "Hello"}`

	mockService.EXPECT().SendMessage("user123", "user456", "Hello").Return(nil)

	req, err := http.NewRequest(http.MethodPost, "/messages", strings.NewReader(message))
	require.NoError(t, err)
	req.Header.Set("Authorization", "Bearer mockToken")
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.AddMessage(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response models.APIResponse
	err = json.NewDecoder(rr.Body).Decode(&response)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, response.Status)
	assert.Equal(t, "Message sent successfully", response.Data)
}

func TestAddMessage_BadRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockMessageService(ctrl)
	handler := handlers.NewMessageHandler(mockService)

	message := `{"sender": "user456", "receiver": "user789", "content": "Hello"}`

	req, err := http.NewRequest(http.MethodPost, "/messages", strings.NewReader(message))
	require.NoError(t, err)
	req.Header.Set("Authorization", "Bearer mockToken")
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.AddMessage(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)

	var response models.APIResponse
	err = json.NewDecoder(rr.Body).Decode(&response)
	require.NoError(t, err)

	assert.Equal(t, http.StatusBadRequest, response.Status)
}
