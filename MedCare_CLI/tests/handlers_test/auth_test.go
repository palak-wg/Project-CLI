package handlers_test

import (
	"bytes"
	"doctor-patient-cli/handlers"
	"doctor-patient-cli/models"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// Mock the UserService interface for testing
type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) GetUserByID(userID string) (*models.User, error) {
	args := m.Called(userID)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserService) CreateUser(user *models.User) error {
	return m.Called(user).Error(0)
}

func TestSignup(t *testing.T) {
	// Initialize the mock service
	mockService := new(MockUserService)
	handler := &handlers.UserHandler{}

	// Create test cases
	tests := []struct {
		name           string
		inputBody      interface{}
		mockGetUser    *models.User
		mockGetUserErr error
		mockCreateUser error
		expectedCode   int
		expectedBody   string
	}{
		{
			name:         "Invalid request body",
			inputBody:    "invalid json",
			expectedCode: http.StatusBadRequest,
			expectedBody: "Invalid request payload",
		},
		{
			name:           "User already exists",
			inputBody:      models.User{UserID: "existing_user", Password: "password123"},
			mockGetUser:    &models.User{UserID: "existing_user"},
			mockGetUserErr: nil,
			expectedCode:   http.StatusBadRequest,
			expectedBody:   "This username already exists.",
		},
		{
			name:           "Successful signup",
			inputBody:      models.User{UserID: "new_user", Password: "password123"},
			mockGetUser:    nil,
			mockGetUserErr: errors.New("user not found"),
			mockCreateUser: nil,
			expectedCode:   http.StatusCreated,
			expectedBody:   "User created successfully",
		},
		{
			name:           "Internal server error during user creation",
			inputBody:      models.User{UserID: "new_user", Password: "password123"},
			mockGetUser:    nil,
			mockGetUserErr: errors.New("user not found"),
			mockCreateUser: errors.New("error creating user"),
			expectedCode:   http.StatusInternalServerError,
			expectedBody:   "Could not create user",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Prepare the request body
			body, _ := json.Marshal(tc.inputBody)
			req, err := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(body))
			require.NoError(t, err)

			// Mock service method responses
			mockService.On("GetUserByID", tc.inputBody.(models.User).UserID).Return(tc.mockGetUser, tc.mockGetUserErr)
			if tc.mockGetUser == nil {
				mockService.On("CreateUser", mock.AnythingOfType("*models.User")).Return(tc.mockCreateUser)
			}

			// Create a response recorder
			rr := httptest.NewRecorder()

			// Call the handler
			handler.Signup(rr, req)

			// Assert the status code and response body
			require.Equal(t, tc.expectedCode, rr.Code)
			require.Contains(t, rr.Body.String(), tc.expectedBody)
		})
	}
}

//
//func TestLogin(t *testing.T) {
//	// Initialize the mock service
//	mockService := new(MockUserService)
//	handler := &handlers.UserHandler{}
//
//	// Hash a password for testing
//	hashedPassword := utils.HashPassword("password123")
//
//	// Create test cases
//	tests := []struct {
//		name           string
//		inputBody      interface{}
//		mockGetUser    *models.User
//		mockGetUserErr error
//		mockToken      string
//		mockTokenErr   error
//		expectedCode   int
//		expectedBody   string
//	}{
//		{
//			name:         "Invalid request body",
//			inputBody:    "invalid json",
//			expectedCode: http.StatusBadRequest,
//			expectedBody: "Invalid request payload",
//		},
//		{
//			name:           "User not found",
//			inputBody:      models.User{UserID: "non_existing_user", Password: "password123"},
//			mockGetUser:    nil,
//			mockGetUserErr: errors.New("user not found"),
//			expectedCode:   http.StatusUnauthorized,
//			expectedBody:   "Unauthorized",
//		},
//		{
//			name:         "Invalid password",
//			inputBody:    models.User{UserID: "existing_user", Password: "wrongpassword"},
//			mockGetUser:  &models.User{UserID: "existing_user", Password: hashedPassword},
//			expectedCode: http.StatusUnauthorized,
//			expectedBody: "Unauthorized",
//		},
//		{
//			name:         "Successful login",
//			inputBody:    models.User{UserID: "existing_user", Password: "password123"},
//			mockGetUser:  &models.User{UserID: "existing_user", Password: hashedPassword},
//			mockToken:    "mocktoken123",
//			expectedCode: http.StatusOK,
//			expectedBody: "mocktoken123",
//		},
//	}
//
//	for _, tc := range tests {
//		t.Run(tc.name, func(t *testing.T) {
//			// Prepare the request body
//			body, _ := json.Marshal(tc.inputBody)
//			req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
//			require.NoError(t, err)
//
//			// Mock service method responses
//			mockService.On("GetUserByID", tc.inputBody.(models.User).UserID).Return(tc.mockGetUser, tc.mockGetUserErr)
//			if tc.mockGetUser != nil && utils.CheckPasswordHash(tc.inputBody.(models.User).Password, tc.mockGetUser.Password) {
//				_, _ = tokens.GenerateToken(models.User{})
//			}
//
//			// Create a response recorder
//			rr := httptest.NewRecorder()
//
//			// Call the handler
//			handler.Login(rr, req)
//
//			// Assert the status code and response body
//			require.Equal(t, tc.expectedCode, rr.Code)
//			require.Contains(t, rr.Body.String(), tc.expectedBody)
//		})
//	}
//}
