package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
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

func TestNotificationHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockNotificationService(ctrl)
	handler := handlers.NewNotificationHandler(mockService)

	router := mux.NewRouter()
	router.HandleFunc("/notifications/{user_id}", handler.GetNotifications)

	t.Run("Success - Get notifications", func(t *testing.T) {
		notifications := []models.Notification{
			{UserID: "user1", Content: "Test Notification"},
		}

		// Set up the mock service expectation
		mockService.EXPECT().GetNotificationsByUserID("user1").Return(notifications, nil).Times(1)

		req := httptest.NewRequest("GET", "/notifications/user1", nil)
		req = req.WithContext(context.WithValue(req.Context(), middlewares.UserIdKey, "user1"))
		req = req.WithContext(context.WithValue(req.Context(), middlewares.RoleKey, "patient"))

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		var response models.APIResponse
		err := json.NewDecoder(rr.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, response.Status)
	})

	t.Run("Error - Access denied", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/notifications/user1", nil)
		req = req.WithContext(context.WithValue(req.Context(), middlewares.UserIdKey, "user2")) // Different user ID
		req = req.WithContext(context.WithValue(req.Context(), middlewares.RoleKey, "patient"))

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusForbidden, rr.Code)

		var response models.APIResponse
		err := json.NewDecoder(rr.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusForbidden, response.Status)
		assert.Equal(t, "Access denied", response.Data)
	})

	t.Run("Error - User not found", func(t *testing.T) {
		// Set up the expectation for the mock service
		mockService.EXPECT().GetNotificationsByUserID("user1").Return(nil, errors.New("user not found")).Times(1)

		req := httptest.NewRequest("GET", "/notifications/user1", nil)
		req = req.WithContext(context.WithValue(req.Context(), middlewares.UserIdKey, "user1"))
		req = req.WithContext(context.WithValue(req.Context(), middlewares.RoleKey, "patient"))

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)

		var response models.APIResponse
		err := json.NewDecoder(rr.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, response.Status)
		assert.Equal(t, http.StatusText(http.StatusNotFound), response.Data)
	})
}
