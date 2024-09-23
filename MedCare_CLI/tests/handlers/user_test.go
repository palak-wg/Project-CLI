package handlers

import (
	"context"
	"doctor-patient-cli/handlers"
	"doctor-patient-cli/middlewares"
	"doctor-patient-cli/models"
	"doctor-patient-cli/tests/mocks"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockUserService(ctrl)
	handler := handlers.NewUserHandler(mockService)

	r := mux.NewRouter()
	r.HandleFunc("/user/{user_id}", handler.GetUser)

	t.Run("Success - Fetch user profile", func(t *testing.T) {
		userID := "user1"
		expectedUser := &models.User{
			UserID:      userID,
			Name:        "John Doe",
			Age:         30,
			Gender:      "Male",
			Email:       "johndoe@example.com",
			PhoneNumber: "1234567890",
		}

		mockService.EXPECT().GetUserByID(userID).Return(expectedUser, nil).Times(1)

		req := httptest.NewRequest("GET", "/users/"+userID, nil)
		req = req.WithContext(context.WithValue(req.Context(), middlewares.RoleKey, "doctor"))
		req = req.WithContext(context.WithValue(req.Context(), middlewares.UserIdKey, userID))

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		var response models.APIResponse
		err := json.NewDecoder(rr.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, response.Status)

		// Expect response.Data to be of type *models.User
		userData, ok := response.Data.(*models.User)
		assert.True(t, ok, "Expected response.Data to be of type *models.User")

		assert.Equal(t, expectedUser.UserID, userData.UserID)
		assert.Equal(t, expectedUser.Name, userData.Name)
		assert.Equal(t, expectedUser.Age, userData.Age)
		assert.Equal(t, expectedUser.Gender, userData.Gender)
		assert.Equal(t, expectedUser.Email, userData.Email)
		assert.Equal(t, expectedUser.PhoneNumber, userData.PhoneNumber)
	})

	t.Run("Forbidden - Access denied", func(t *testing.T) {
		userID := "user1"

		req := httptest.NewRequest("GET", "/user/"+userID, nil)
		req = req.WithContext(context.WithValue(req.Context(), middlewares.RoleKey, "doctor"))
		req = req.WithContext(context.WithValue(req.Context(), middlewares.UserIdKey, "differentUserID")) // Different token ID

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusForbidden, rr.Code)

		var response models.APIResponse
		err := json.NewDecoder(rr.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusForbidden, response.Status)
		assert.Equal(t, "Access denied", response.Data)
	})

	t.Run("Not Found - User does not exist", func(t *testing.T) {
		userID := "nonexistentUser"

		mockService.EXPECT().GetUserByID(userID).Return(nil, errors.New("user not found")).Times(1)

		req := httptest.NewRequest("GET", "/user/"+userID, nil)
		req = req.WithContext(context.WithValue(req.Context(), middlewares.RoleKey, "doctor"))
		req = req.WithContext(context.WithValue(req.Context(), middlewares.UserIdKey, userID))

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)

		var response models.APIResponse
		err := json.NewDecoder(rr.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, response.Status)
		assert.Equal(t, http.StatusText(http.StatusNotFound), response.Data)
	})
}
