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

	//t.Run("Success - Fetch user profile", func(t *testing.T) {
	//	userID := "user1"
	//	expectedUser := &models.User{
	//		UserID:      userID,
	//		Name:        "John Doe",
	//		Age:         30,
	//		Gender:      "Male",
	//		Email:       "johndoe@example.com",
	//		PhoneNumber: "1234567890",
	//	}
	//
	//	// Set up the mock expectation
	//	mockService.EXPECT().GetUserByID(userID).Return(expectedUser, nil).Times(1)
	//
	//	req := httptest.NewRequest("GET", "/users/"+userID, nil)
	//	req = req.WithContext(context.WithValue(req.Context(), middlewares.RoleKey, "doctor"))
	//	req = req.WithContext(context.WithValue(req.Context(), middlewares.UserIdKey, userID))
	//
	//	rr := httptest.NewRecorder()
	//	r.ServeHTTP(rr, req)
	//
	//	// Check for the correct response code
	//	if status := rr.Code; status != http.StatusOK {
	//		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	//	}
	//
	//	var response models.APIResponse
	//	err := json.NewDecoder(rr.Body).Decode(&response)
	//	if err != nil {
	//		t.Fatalf("Failed to decode response: %v", err)
	//	}
	//
	//	// Check if the response status is OK
	//	if response.Status != http.StatusOK {
	//		t.Errorf("Expected status %v, got %v", http.StatusOK, response.Status)
	//	}
	//
	//	// Ensure the response.Data is of the expected type
	//	userData, ok := response.Data.(*models.User)
	//	if !ok {
	//		t.Errorf("Expected response.Data to be of type *models.User")
	//	}
	//
	//	// Assert the user data
	//	if userData.UserID != expectedUser.UserID || userData.Name != expectedUser.Name ||
	//		userData.Age != expectedUser.Age || userData.Gender != expectedUser.Gender ||
	//		userData.Email != expectedUser.Email || userData.PhoneNumber != expectedUser.PhoneNumber {
	//		t.Error("Response data does not match expected user data")
	//	}
	//})

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
