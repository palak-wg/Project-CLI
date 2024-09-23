package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"

	"doctor-patient-cli/handlers"
	"doctor-patient-cli/models"
	"doctor-patient-cli/tests/mocks"
)

func TestDoctorHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockDoctorService(ctrl)
	handler := handlers.NewDoctorHandler(mockService)
	r := mux.NewRouter()
	r.HandleFunc("/doctors", handler.GetDoctors).Methods("GET")
	r.HandleFunc("/doctors/{doctor_id}", handler.GetDoctor).Methods("GET")

	t.Run("Success - Get all doctors", func(t *testing.T) {
		doctors := []models.Doctor{
			{User: models.User{UserID: "doctor1"}, Specialization: "Cardiology", Experience: 10, Rating: 4.5},
			{User: models.User{UserID: "doctor2"}, Specialization: "Neurology", Experience: 8, Rating: 4.7},
		}

		mockService.EXPECT().GetAllDoctors().Return(doctors, nil).Times(1)

		req := httptest.NewRequest("GET", "/doctors", nil)
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		var response models.APIResponse
		err := json.NewDecoder(rr.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, response.Status)

		// Assert that the Data field is of type []interface{}
		var responseDoctors []models.APIResponseDoctor
		dataBytes, err := json.Marshal(response.Data)
		assert.NoError(t, err)
		err = json.Unmarshal(dataBytes, &responseDoctors)
		assert.NoError(t, err)
		assert.Len(t, responseDoctors, 2)
	})

	t.Run("Error - Get all doctors fails", func(t *testing.T) {
		mockService.EXPECT().GetAllDoctors().Return(nil, assert.AnError).Times(1)

		req := httptest.NewRequest("GET", "/doctors", nil)
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)

		var response models.APIResponse
		err := json.NewDecoder(rr.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, response.Status)
	})

	t.Run("Success - Get doctor by ID", func(t *testing.T) {
		doctor := models.Doctor{
			User:           models.User{UserID: "doctor1"},
			Specialization: "Cardiology",
			Experience:     10,
			Rating:         4.5,
		}

		mockService.EXPECT().GetDoctorByID("doctor1").Return(&doctor, nil).Times(1)

		req := httptest.NewRequest("GET", "/doctors/doctor1", nil)
		req = mux.SetURLVars(req, map[string]string{"doctor_id": "doctor1"})
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		var response models.APIResponse
		err := json.NewDecoder(rr.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, response.Status)

		// Assert that the Data field is of type models.APIResponseDoctor
		var responseDoctor models.APIResponseDoctor
		dataBytes, err := json.Marshal(response.Data)
		assert.NoError(t, err)
		err = json.Unmarshal(dataBytes, &responseDoctor)
		assert.NoError(t, err)
		assert.Equal(t, doctor.UserID, responseDoctor.DoctorID)
	})

	t.Run("Error - Doctor not found", func(t *testing.T) {
		mockService.EXPECT().GetDoctorByID("doctor1").Return(nil, assert.AnError).Times(1)

		req := httptest.NewRequest("GET", "/doctors/doctor1", nil)
		req = mux.SetURLVars(req, map[string]string{"doctor_id": "doctor1"})
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)

		var response models.APIResponse
		err := json.NewDecoder(rr.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, response.Status)
	})
}
