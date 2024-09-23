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
	"github.com/stretchr/testify/assert"
)

func TestUserHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockUserService(ctrl)
	handler := handlers.NewUserHandler(mockService)

	t.Run("GetUser - Success", func(t *testing.T) {
		user := models.User{
			UserID:      "123",
			Name:        "John Doe",
			Age:         30,
			Gender:      "Male",
			Email:       "john@example.com",
			PhoneNumber: "1234567890",
		}
		mockService.EXPECT().GetUserByID("123").Return(&user, nil)

		req := httptest.NewRequest("GET", "/user/123", nil)
		req.Header.Set("Authorization", "Bearer "+generateToken("123", "patient"))
		rr := httptest.NewRecorder()

		handler.GetUser(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		var response models.APIResponse
		err := json.NewDecoder(rr.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, response.Status)
		assert.Equal(t, user.UserID, response.Data.(*models.APIResponseUser).UserID)
	})

	t.Run("GetUser - User Not Found", func(t *testing.T) {
		mockService.EXPECT().GetUserByID("123").Return(nil, errors.New("user not found"))

		req := httptest.NewRequest("GET", "/user/123", nil)
		req.Header.Set("Authorization", "Bearer "+generateToken("123", "patient"))
		rr := httptest.NewRecorder()

		handler.GetUser(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)

		var response models.APIResponse
		err := json.NewDecoder(rr.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, response.Status)
	})

	t.Run("GetUser - Access Denied", func(t *testing.T) {
		user := models.User{
			UserID:      "456",
			Name:        "Jane Doe",
			Age:         28,
			Gender:      "Female",
			Email:       "jane@example.com",
			PhoneNumber: "0987654321",
		}
		mockService.EXPECT().GetUserByID("456").Return(&user, nil)

		req := httptest.NewRequest("GET", "/user/456", nil)
		req.Header.Set("Authorization", "Bearer "+generateToken("123", "patient")) // Token ID does not match user ID
		rr := httptest.NewRecorder()

		handler.GetUser(rr, req)

		assert.Equal(t, http.StatusForbidden, rr.Code)

		var response models.APIResponse
		err := json.NewDecoder(rr.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusForbidden, response.Status)
	})

	t.Run("GetUser - Internal Server Error on Claim Extraction", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/user/123", nil)
		req.Header.Set("Authorization", "Bearer invalid_token")
		rr := httptest.NewRecorder()

		handler.GetUser(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)

		var response models.APIResponse
		err := json.NewDecoder(rr.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, response.Status)
	})
}
