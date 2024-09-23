package handlers

import (
	"bytes"
	"context"
	"doctor-patient-cli/handlers"
	"doctor-patient-cli/middlewares"
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

func TestPatientHandler_UpdateProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockPatientService(ctrl)
	handler := handlers.NewPatientHandler(mockService)

	router := mux.NewRouter()
	router.HandleFunc("/patients/profile", handler.UpdateProfile).Methods("PUT")

	t.Run("Success - Update profile", func(t *testing.T) {
		user := models.Patient{
			User: models.User{
				UserID:      "user1",
				Name:        "John Doe",
				Email:       "john@example.com",
				PhoneNumber: "1234567890",
			},
		}

		mockService.EXPECT().UpdatePatientDetails(&user).Return(nil).Times(1)

		body, _ := json.Marshal(user)
		req := httptest.NewRequest("PUT", "/patients/profile", bytes.NewBuffer(body))
		req = req.WithContext(context.WithValue(req.Context(), middlewares.UserIdKey, "user1"))

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		var response models.APIResponse
		err := json.NewDecoder(rr.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, response.Status)
		assert.Equal(t, "Updated John Doe profile", response.Data)
	})

	t.Run("Error - Bad request", func(t *testing.T) {
		req := httptest.NewRequest("PUT", "/patients/profile", bytes.NewBuffer([]byte("invalid json")))
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)

		var response models.APIResponse
		err := json.NewDecoder(rr.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, response.Status)
		assert.Equal(t, http.StatusText(http.StatusBadRequest), response.Data)
	})

	t.Run("Error - Update failed", func(t *testing.T) {
		user := models.Patient{
			User: models.User{
				UserID:      "user1",
				Name:        "John Doe",
				Email:       "john@example.com",
				PhoneNumber: "1234567890",
			},
		}

		mockService.EXPECT().UpdatePatientDetails(&user).Return(errors.New("update error")).Times(1)

		body, _ := json.Marshal(user)
		req := httptest.NewRequest("PUT", "/patients/profile", bytes.NewBuffer(body))
		req = req.WithContext(context.WithValue(req.Context(), middlewares.UserIdKey, "user1"))

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)

		var response models.APIResponse
		err := json.NewDecoder(rr.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, response.Status)
		assert.Equal(t, http.StatusText(http.StatusBadRequest), response.Data)
	})
}
