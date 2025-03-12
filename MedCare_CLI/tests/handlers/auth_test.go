package handlers

import (
	"bytes"
	"doctor-patient-cli/handlers"
	"doctor-patient-cli/models"
	"doctor-patient-cli/tests/mocks"
	"doctor-patient-cli/utils"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserHandler_Signup(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockUserService(ctrl)
	handler := handlers.NewUserHandler(mockService)

	t.Run("Success - User signup", func(t *testing.T) {
		user := models.User{UserID: "user1", Password: "password"} // Use plaintext password for test

		// Set up expectations for the mocked methods
		mockService.EXPECT().GetUserByID(user.UserID).Return(nil, nil)
		mockService.EXPECT().CreateUser(gomock.Any()).Do(func(u *models.User) {
			// Ensure the password is hashed
			assert.NotEqual(t, "password", u.Password)
		}).Return(nil)

		body, _ := json.Marshal(user)
		req := httptest.NewRequest("POST", "/signup", bytes.NewBuffer(body))
		rr := httptest.NewRecorder()

		handler.Signup(rr, req)

		assert.Equal(t, http.StatusCreated, rr.Code)
		var response models.APIResponse
		json.NewDecoder(rr.Body).Decode(&response)
		assert.Equal(t, http.StatusCreated, response.Status)
	})

	t.Run("Error - Username already exists", func(t *testing.T) {
		user := models.User{UserID: "user1", Password: "password"}

		mockService.EXPECT().GetUserByID(user.UserID).Return(&user, nil)

		body, _ := json.Marshal(user)
		req := httptest.NewRequest("POST", "/signup", bytes.NewBuffer(body))
		rr := httptest.NewRecorder()

		handler.Signup(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		var response models.APIResponse
		json.NewDecoder(rr.Body).Decode(&response)
		assert.Equal(t, http.StatusBadRequest, response.Status)
		assert.Equal(t, "This username already exists.", response.Data)
	})

	t.Run("Error - Invalid input", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/signup", nil)
		rr := httptest.NewRecorder()

		handler.Signup(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		var response models.APIResponse
		json.NewDecoder(rr.Body).Decode(&response)
		assert.Equal(t, http.StatusBadRequest, response.Status)
	})
}

func TestUserHandler_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockUserService(ctrl)
	handler := handlers.NewUserHandler(mockService)

	t.Run("Success - User login", func(t *testing.T) {
		user := models.User{UserID: "user1", Password: utils.HashPassword("password")}
		client := models.User{UserID: "user1", Password: "password"}

		mockService.EXPECT().GetUserByID(client.UserID).Return(&user, nil)

		body, _ := json.Marshal(client)
		req := httptest.NewRequest("POST", "/login", bytes.NewBuffer(body))
		rr := httptest.NewRecorder()

		handler.Login(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		var response struct {
			Code  int    `json:"status_code"`
			Token string `json:"token"`
		}
		json.NewDecoder(rr.Body).Decode(&response)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("Error - Wrong password", func(t *testing.T) {
		user := models.User{UserID: "user1", Password: utils.HashPassword("correctPassword")}
		client := models.User{UserID: "user1", Password: "wrongPassword"}

		mockService.EXPECT().GetUserByID(client.UserID).Return(&user, nil)

		body, _ := json.Marshal(client)
		req := httptest.NewRequest("POST", "/login", bytes.NewBuffer(body))
		rr := httptest.NewRecorder()

		handler.Login(rr, req)

		assert.Equal(t, http.StatusUnauthorized, rr.Code)
		var response models.APIResponse
		json.NewDecoder(rr.Body).Decode(&response)
		assert.Equal(t, http.StatusUnauthorized, response.Status)
	})

	t.Run("Error - Invalid input", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/login", nil)
		rr := httptest.NewRecorder()

		handler.Login(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		var response models.APIResponse
		json.NewDecoder(rr.Body).Decode(&response)
		assert.Equal(t, http.StatusBadRequest, response.Status)
	})
}
