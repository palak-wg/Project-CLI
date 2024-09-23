package handlers_test

import (
	"doctor-patient-cli/handlers"
	"doctor-patient-cli/models"
	"doctor-patient-cli/tests/mocks"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestNotificationHandler_GetNotifications(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockNotificationService(ctrl)
	handler := handlers.NewNotificationHandler(mockService)

	t.Run("Success - Valid token and notifications fetched", func(t *testing.T) {
		mockService.EXPECT().GetNotificationsByUserID("123").Return([]models.Notification{
			{Content: "Test Notification"},
		}, nil)

		req := httptest.NewRequest("GET", "/notifications/123", nil)
		req.Header.Set("Authorization", "Bearer mockToken")

		rr := httptest.NewRecorder()
		req = mux.SetURLVars(req, map[string]string{"user_id": "123"})

		handler.GetNotifications(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		var response models.APIResponse
		err := json.NewDecoder(rr.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, response.Status)
		assert.NotEmpty(t, response.Data)
	})

	t.Run("Error - Invalid token", func(t *testing.T) {

		req := httptest.NewRequest("GET", "/notifications/123", nil)
		req.Header.Set("Authorization", "Bearer invalidToken")

		rr := httptest.NewRecorder()
		req = mux.SetURLVars(req, map[string]string{"user_id": "123"})

		handler.GetNotifications(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)

		var response models.APIResponse
		err := json.NewDecoder(rr.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, response.Status)
		assert.Equal(t, "Error extracting claims", response.Data)
	})

	t.Run("Error - Access denied", func(t *testing.T) {

		req := httptest.NewRequest("GET", "/notifications/456", nil)
		req.Header.Set("Authorization", "Bearer validToken")

		rr := httptest.NewRecorder()
		req = mux.SetURLVars(req, map[string]string{"user_id": "456"})

		handler.GetNotifications(rr, req)

		assert.Equal(t, http.StatusForbidden, rr.Code)

		var response models.APIResponse
		err := json.NewDecoder(rr.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusForbidden, response.Status)
		assert.Equal(t, "Access denied", response.Data)
	})

	t.Run("Error - User not found", func(t *testing.T) {

		mockService.EXPECT().GetNotificationsByUserID("123").Return(nil, errors.New("user not found"))

		req := httptest.NewRequest("GET", "/notifications/123", nil)
		req.Header.Set("Authorization", "Bearer mockToken")

		rr := httptest.NewRecorder()
		req = mux.SetURLVars(req, map[string]string{"user_id": "123"})

		handler.GetNotifications(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)

		var response models.APIResponse
		err := json.NewDecoder(rr.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, response.Status)
		assert.Equal(t, http.StatusText(http.StatusNotFound), response.Data)
	})
}
