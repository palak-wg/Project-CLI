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

func TestCreateAppointment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockAppointmentService(ctrl)
	handler := handlers.NewAppointmentHandler(mockService)

	// Set up a router for testing
	r := mux.NewRouter()
	r.HandleFunc("/appointments", handler.CreateAppointment).Methods("POST")

	t.Run("Success - Valid patient creates appointment", func(t *testing.T) {
		// Mock the appointment data
		appointment := models.Appointment{
			PatientID: "123",
			DoctorID:  "456",
		}

		// Set expectation for SendAppointmentRequest
		mockService.EXPECT().SendAppointmentRequest("123", "456").Return(nil)

		// Create a new HTTP request with valid data
		body, _ := json.Marshal(appointment)
		req := httptest.NewRequest("POST", "/appointments", bytes.NewBuffer(body))
		req = req.WithContext(context.WithValue(req.Context(), middlewares.RoleKey, "patient"))
		req = req.WithContext(context.WithValue(req.Context(), middlewares.UserIdKey, "123"))

		// Create a ResponseRecorder to record the response
		rr := httptest.NewRecorder()

		// Serve the request
		r.ServeHTTP(rr, req)

		// Check if status code is 200 OK
		assert.Equal(t, http.StatusOK, rr.Code)

		// Parse the response body
		var response models.APIResponse
		err := json.NewDecoder(rr.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, response.Status)
	})

	t.Run("Error - Decoding request body", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/appointments", nil) // No body
		req = req.WithContext(context.WithValue(req.Context(), middlewares.RoleKey, "patient"))
		req = req.WithContext(context.WithValue(req.Context(), middlewares.UserIdKey, "123"))

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)

		var response models.APIResponse
		err := json.NewDecoder(rr.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, response.Status)
	})

	t.Run("Error - Invalid PatientID", func(t *testing.T) {
		appointment := models.Appointment{
			PatientID: "456", // Different from user ID
			DoctorID:  "456",
		}

		body, _ := json.Marshal(appointment)
		req := httptest.NewRequest("POST", "/appointments", bytes.NewBuffer(body))
		req = req.WithContext(context.WithValue(req.Context(), middlewares.RoleKey, "patient"))
		req = req.WithContext(context.WithValue(req.Context(), middlewares.UserIdKey, "123"))

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)

		var response models.APIResponse
		err := json.NewDecoder(rr.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, response.Status)
	})

	t.Run("Error - Sending appointment fails", func(t *testing.T) {
		appointment := models.Appointment{
			PatientID: "123",
			DoctorID:  "456",
		}

		// Set expectation for SendAppointmentRequest to return an error
		mockService.EXPECT().SendAppointmentRequest("123", "456").Return(errors.New("service error"))

		body, _ := json.Marshal(appointment)
		req := httptest.NewRequest("POST", "/appointments", bytes.NewBuffer(body))
		req = req.WithContext(context.WithValue(req.Context(), middlewares.RoleKey, "patient"))
		req = req.WithContext(context.WithValue(req.Context(), middlewares.UserIdKey, "123"))

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)

		var response models.APIResponse
		err := json.NewDecoder(rr.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, response.Status)
	})
}

func TestUpdateAppointment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockAppointmentService(ctrl)
	handler := handlers.NewAppointmentHandler(mockService)

	r := mux.NewRouter()
	r.HandleFunc("/appointments", handler.UpdateAppointment).Methods("PUT")

	t.Run("Success - Valid doctor approves appointment", func(t *testing.T) {
		appointment := models.Appointment{AppointmentID: 123}
		body, _ := json.Marshal(appointment)

		req := httptest.NewRequest("PUT", "/appointments", bytes.NewBuffer(body))
		req.Header.Set("Authorization", "Bearer valid_token")
		req = req.WithContext(context.WithValue(req.Context(), middlewares.RoleKey, "doctor"))

		// Set up the expectation before the request is executed
		mockService.EXPECT().ApproveAppointment(123).Return(nil)

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		var response models.APIResponse
		err := json.NewDecoder(rr.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, response.Status)
		assert.Equal(t, "Appointment 123 approved successfully", response.Data)
	})

	t.Run("Error - Patient role tries to approve appointment", func(t *testing.T) {
		appointment := models.Appointment{AppointmentID: 123}
		body, _ := json.Marshal(appointment)

		req := httptest.NewRequest("PUT", "/appointments", bytes.NewBuffer(body))
		req.Header.Set("Authorization", "Bearer valid_token")
		req = req.WithContext(context.WithValue(req.Context(), middlewares.RoleKey, "patient"))

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)

		var response models.APIResponse
		err := json.NewDecoder(rr.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, response.Status)
	})

	t.Run("Error - Invalid request body", func(t *testing.T) {
		req := httptest.NewRequest("PUT", "/appointments", bytes.NewBuffer([]byte("invalid json")))
		req.Header.Set("Authorization", "Bearer valid_token")
		req = req.WithContext(context.WithValue(req.Context(), middlewares.RoleKey, "doctor"))

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)

		var response models.APIResponse
		err := json.NewDecoder(rr.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, response.Status)
	})

	t.Run("Error - Approving appointment fails", func(t *testing.T) {
		appointment := models.Appointment{AppointmentID: 123}
		body, _ := json.Marshal(appointment)

		req := httptest.NewRequest("PUT", "/appointments", bytes.NewBuffer(body))
		req.Header.Set("Authorization", "Bearer valid_token")
		req = req.WithContext(context.WithValue(req.Context(), middlewares.RoleKey, "doctor"))

		mockService.EXPECT().ApproveAppointment(123).Return(errors.New("some error"))

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)

		var response models.APIResponse
		err := json.NewDecoder(rr.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, response.Status)
		assert.Equal(t, "Error approving appointment", response.Data)
	})
}

func TestGetAppointments(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockAppointmentService(ctrl)
	handler := handlers.NewAppointmentHandler(mockService)

	r := mux.NewRouter()
	r.HandleFunc("/appointments", handler.GetAppointments).Methods("GET")

	t.Run("Success - Doctor fetches appointments", func(t *testing.T) {
		appointments := []models.Appointment{
			{AppointmentID: 1, PatientID: "patient1", DoctorID: "doctor1"},
			{AppointmentID: 2, PatientID: "patient2", DoctorID: "doctor1"},
		}

		mockService.EXPECT().GetAppointmentsByDoctorID("doctor1").Return(appointments, nil)

		req := httptest.NewRequest("GET", "/appointments", nil)
		req = req.WithContext(context.WithValue(req.Context(), middlewares.RoleKey, "doctor"))
		req = req.WithContext(context.WithValue(req.Context(), middlewares.UserIdKey, "doctor1"))

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		var response models.APIResponse
		err := json.NewDecoder(rr.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, response.Status)
	})

	//t.Run("Success - Admin fetches appointments", func(t *testing.T) {
	//	appointments := []models.Appointment{
	//		{AppointmentID: 1, PatientID: "patient1", DoctorID: "doctor1"},
	//	}
	//
	//	// Mock the service call
	//	mockService.EXPECT().GetAppointmentsByDoctorID("doctor1").Return(appointments, nil).Times(1)
	//
	//	// Create a new request
	//	req := httptest.NewRequest("GET", "/appointments", nil)
	//	req = req.WithContext(context.WithValue(req.Context(), middlewares.RoleKey, "admin"))
	//	req = req.WithContext(context.WithValue(req.Context(), middlewares.UserIdKey, "doctor1")) // ID used for fetching
	//
	//	// Create a response recorder
	//	rr := httptest.NewRecorder()
	//	r.ServeHTTP(rr, req)
	//
	//	// Assert the status code
	//	assert.Equal(t, http.StatusOK, rr.Code)
	//
	//	// Decode the response
	//	var response models.APIResponse
	//	err := json.NewDecoder(rr.Body).Decode(&response)
	//	assert.NoError(t, err)
	//	assert.Equal(t, http.StatusOK, response.Status)
	//
	//	// Assert the data
	//	assert.Equal(t, appointments, response.Data)
	//})

	t.Run("Error - Patient attempts to fetch appointments", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/appointments", nil)
		req = req.WithContext(context.WithValue(req.Context(), middlewares.RoleKey, "patient"))
		req = req.WithContext(context.WithValue(req.Context(), middlewares.UserIdKey, "patient1"))

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusForbidden, rr.Code)

		var response models.APIResponse
		err := json.NewDecoder(rr.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusForbidden, response.Status)
		assert.Equal(t, "Access denied", response.Data)
	})

	t.Run("Error - Fetch appointments fails", func(t *testing.T) {
		mockService.EXPECT().GetAppointmentsByDoctorID("doctor1").Return(nil, errors.New("service error"))

		req := httptest.NewRequest("GET", "/appointments", nil)
		req = req.WithContext(context.WithValue(req.Context(), middlewares.RoleKey, "doctor"))
		req = req.WithContext(context.WithValue(req.Context(), middlewares.UserIdKey, "doctor1"))

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)

		var response models.APIResponse
		err := json.NewDecoder(rr.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, response.Status)
		assert.Equal(t, "Error fetching appointments", response.Data)
	})
}
