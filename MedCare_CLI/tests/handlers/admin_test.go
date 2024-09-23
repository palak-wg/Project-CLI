package handlers

import (
	"bytes"
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

func TestGetUsers(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockAdminService(ctrl)
	handler := handlers.NewAdminHandler(mockService)

	t.Run("success fetching users", func(t *testing.T) {
		mockUsers := []models.User{
			{
				UserID:      "1",
				Name:        "John Doe",
				Age:         30,
				Gender:      "Male",
				Email:       "john@example.com",
				PhoneNumber: "1234567890",
			},
		}

		// Mocking the service method
		mockService.EXPECT().GetAllUsers().Return(mockUsers, nil)

		req, err := http.NewRequest("GET", "/users", nil)
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		handler.GetUsers(rr, req)

		// Check the response code
		assert.Equal(t, http.StatusOK, rr.Code)

		// Create the expected response
		expectedResponse := models.APIResponse{
			Status: http.StatusOK,
			Data: []models.APIResponseUser{
				{
					UserID:      "1",
					Name:        "John Doe",
					Age:         30,
					Gender:      "Male",
					Email:       "john@example.com",
					PhoneNumber: "1234567890",
				},
			},
		}

		// Decode the actual response
		var response models.APIResponse
		err = json.NewDecoder(rr.Body).Decode(&response)
		assert.NoError(t, err)

		// Assert the response matches the expected response
		assert.Equal(t, expectedResponse.Status, response.Status)
	})

	t.Run("error fetching users", func(t *testing.T) {
		mockService.EXPECT().GetAllUsers().Return(nil, errors.New("some error"))

		req, err := http.NewRequest("GET", "/users", nil)
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		handler.GetUsers(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)

		expectedResponse := models.APIResponse{
			Status: http.StatusInternalServerError,
			Data:   "Internal Server Error fetching user profiles",
		}

		var response models.APIResponse
		err = json.NewDecoder(rr.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, expectedResponse, response)
	})
}

func TestGetPendingDoctors(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockAdminService(ctrl)
	handler := handlers.NewAdminHandler(mockService)

	t.Run("success fetching pending doctors", func(t *testing.T) {
		mockPendingDoctors := []models.Doctor{
			{
				User: models.User{
					UserID: "1",
					Name:   "Dr. Smith",
				},
			},
		}

		mockService.EXPECT().GetPendingDoctorRequests().Return(mockPendingDoctors, nil)

		req, err := http.NewRequest("GET", "/doctors/approval", nil)
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		handler.GetPendingDoctors(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		expectedResponse := models.APIResponse{
			Status: http.StatusOK,
			Data: []models.APIResponsePendingSignup{
				{
					ID:   "1",
					Name: "Dr. Smith",
				},
			},
		}

		var response models.APIResponse
		err = json.NewDecoder(rr.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, expectedResponse.Status, response.Status)
	})

	t.Run("error fetching pending doctors", func(t *testing.T) {
		mockService.EXPECT().GetPendingDoctorRequests().Return(nil, errors.New("some error"))

		req, err := http.NewRequest("GET", "/doctors/approval", nil)
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		handler.GetPendingDoctors(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)

		expectedResponse := models.APIResponse{
			Status: http.StatusInternalServerError,
			Data:   "Error fetching pending requests",
		}

		var response models.APIResponse
		err = json.NewDecoder(rr.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, expectedResponse, response)
	})
}

func TestApprovePendingDoctors(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockAdminService(ctrl)
	handler := handlers.NewAdminHandler(mockService)

	t.Run("success approving doctor signup", func(t *testing.T) {
		mockService.EXPECT().ApproveDoctorSignup("1").Return(nil)

		doctor := models.APIResponsePendingSignup{
			ID: "1",
		}
		body, err := json.Marshal(doctor)
		assert.NoError(t, err)

		req, err := http.NewRequest("POST", "/approve-doctor", bytes.NewReader(body))
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		handler.ApprovePendingDoctors(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		expectedResponse := models.APIResponse{
			Status: http.StatusOK,
			Data:   "Approved doctor 1 signup",
		}

		var response models.APIResponse
		err = json.NewDecoder(rr.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, expectedResponse, response)
	})

	t.Run("error decoding request body", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/approve-doctor", bytes.NewReader([]byte("{")))
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		handler.ApprovePendingDoctors(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)

		expectedResponse := models.APIResponse{
			Status: http.StatusInternalServerError,
			Data:   "Error decoding body",
		}

		var response models.APIResponse
		err = json.NewDecoder(rr.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, expectedResponse, response)
	})

	t.Run("error approving doctor", func(t *testing.T) {
		mockService.EXPECT().ApproveDoctorSignup("1").Return(errors.New("some error"))

		doctor := models.APIResponsePendingSignup{
			ID: "1",
		}
		body, err := json.Marshal(doctor)
		assert.NoError(t, err)

		req, err := http.NewRequest("POST", "/approve-doctor", bytes.NewReader(body))
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		handler.ApprovePendingDoctors(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)

		expectedResponse := models.APIResponse{
			Status: http.StatusInternalServerError,
			Data:   "Error approving doctor profile for signup",
		}

		var response models.APIResponse
		err = json.NewDecoder(rr.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, expectedResponse, response)
	})
}
