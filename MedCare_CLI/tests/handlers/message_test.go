package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"doctor-patient-cli/handlers"
	"doctor-patient-cli/middlewares"
	"doctor-patient-cli/models"
	"doctor-patient-cli/tests/mocks"
)

func TestMessageHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockMessageService(ctrl)
	handler := handlers.NewMessageHandler(mockService)

	// Setup a context with user ID
	userID := "user1"
	ctx := context.WithValue(context.Background(), middlewares.UserIdKey, userID)

	t.Run("Success - Get unread messages", func(t *testing.T) {
		messages := []models.Message{
			{Sender: "user1", Receiver: "user2", Content: "Hello!"},
		}

		mockService.EXPECT().GetUnreadMessages(userID).Return(messages, nil).Times(1)

		req := httptest.NewRequest("GET", "/messages", nil).WithContext(ctx)
		rr := httptest.NewRecorder()
		handler.GetMessages(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		var response models.APIResponse
		err := json.NewDecoder(rr.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, response.Status)
	})

	t.Run("Error - Fetching messages fails", func(t *testing.T) {
		mockService.EXPECT().GetUnreadMessages(userID).Return(nil, assert.AnError).Times(1)

		req := httptest.NewRequest("GET", "/messages", nil).WithContext(ctx)
		rr := httptest.NewRecorder()
		handler.GetMessages(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)

		var response models.APIResponse
		err := json.NewDecoder(rr.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, response.Status)
	})

	t.Run("Success - Add message", func(t *testing.T) {
		message := models.Message{
			Sender:   userID,
			Receiver: "user2",
			Content:  "Hello!",
		}

		mockService.EXPECT().SendMessage(message.Sender, message.Receiver, message.Content).Return(nil).Times(1)

		body, _ := json.Marshal(message)
		req := httptest.NewRequest("POST", "/messages", bytes.NewBuffer(body)).WithContext(ctx)
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()
		handler.AddMessage(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		var response models.APIResponse
		err := json.NewDecoder(rr.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, response.Status)
		assert.Equal(t, "Message sent successfully", response.Data)
	})

	t.Run("Error - Bad request when adding message", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/messages", nil).WithContext(ctx)
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()
		handler.AddMessage(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)

		var response models.APIResponse
		err := json.NewDecoder(rr.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, response.Status)
	})

	t.Run("Error - Cannot send message", func(t *testing.T) {
		message := models.Message{
			Sender:   userID,
			Receiver: "user2",
			Content:  "Hello!",
		}

		mockService.EXPECT().SendMessage(message.Sender, message.Receiver, message.Content).Return(assert.AnError).Times(1)

		body, _ := json.Marshal(message)
		req := httptest.NewRequest("POST", "/messages", bytes.NewBuffer(body)).WithContext(ctx)
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()
		handler.AddMessage(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)

		var response models.APIResponse
		err := json.NewDecoder(rr.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, response.Status)
	})
}
