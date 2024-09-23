package handlers_test

import (
	"doctor-patient-cli/handlers"
	"doctor-patient-cli/models"
	"doctor-patient-cli/tests/mocks"
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetDoctors(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockDoctorService(ctrl)
	handler := handlers.NewDoctorHandler(mockService)

	tests := []struct {
		name             string
		mockReturn       []models.APIResponseDoctor
		mockError        error
		expectedCode     int
		expectedResponse interface{}
	}{
		{
			name: "Success - Doctors fetched",
			mockReturn: []models.APIResponseDoctor{
				{DoctorID: "doc1", Specialization: "Cardiology", Experience: 10, Rating: 4.5},
				{DoctorID: "doc2", Specialization: "Neurology", Experience: 8, Rating: 4.7},
			},
			mockError:    nil,
			expectedCode: http.StatusOK,
			expectedResponse: []models.APIResponseDoctor{
				{DoctorID: "doc1", Specialization: "Cardiology", Experience: 10, Rating: 4.5},
				{DoctorID: "doc2", Specialization: "Neurology", Experience: 8, Rating: 4.7},
			},
		},
		{
			name:             "Error - Failed to fetch doctors",
			mockReturn:       nil,
			mockError:        errors.New("database error"),
			expectedCode:     http.StatusInternalServerError,
			expectedResponse: "Error fetching doctors profiles",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Set up the mock behavior
			mockService.EXPECT().GetAllDoctors().Return(tc.mockReturn, tc.mockError)

			// Prepare the request and response recorder
			req, _ := http.NewRequest(http.MethodGet, "/doctors", nil)
			rr := httptest.NewRecorder()

			// Call the handler
			handler.GetDoctors(rr, req)

			// Assert the response
			require.Equal(t, tc.expectedCode, rr.Code)

			// Check the body of the response
			if tc.expectedCode == http.StatusOK {
				var response models.APIResponse
				err := json.NewDecoder(rr.Body).Decode(&response)
				require.NoError(t, err)

				// Assert that the body contains the expected response
				require.Equal(t, tc.expectedResponse, response.Data)
			} else {
				var response models.APIResponse
				err := json.NewDecoder(rr.Body).Decode(&response)
				require.NoError(t, err)

				require.Equal(t, tc.expectedResponse, response.Data)
			}
		})
	}
}

func TestGetDoctor(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockDoctorService(ctrl)
	handler := handlers.NewDoctorHandler(mockService)

	tests := []struct {
		name             string
		doctorID         string
		mockReturn       *models.APIResponseDoctor
		mockError        error
		expectedCode     int
		expectedResponse interface{}
	}{
		{
			name:     "Success - Doctor found",
			doctorID: "doc1",
			mockReturn: &models.APIResponseDoctor{
				DoctorID: "doc1", Specialization: "Cardiology", Experience: 10, Rating: 4.5,
			},
			mockError:    nil,
			expectedCode: http.StatusOK,
			expectedResponse: models.APIResponseDoctor{
				DoctorID:       "doc1",
				Specialization: "Cardiology",
				Experience:     10,
				Rating:         4.5,
			},
		},
		{
			name:             "Error - Doctor not found",
			doctorID:         "doc999",
			mockReturn:       nil,
			mockError:        errors.New("not found"),
			expectedCode:     http.StatusNotFound,
			expectedResponse: http.StatusText(http.StatusNotFound),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Set up the mock behavior
			mockService.EXPECT().GetDoctorByID(tc.doctorID).Return(tc.mockReturn, tc.mockError)

			// Prepare the request and response recorder
			req, _ := http.NewRequest(http.MethodGet, "/doctor/"+tc.doctorID, nil)
			rr := httptest.NewRecorder()

			// Set the path variables for mux
			req = mux.SetURLVars(req, map[string]string{"doctor_id": tc.doctorID})

			// Call the handler
			handler.GetDoctor(rr, req)

			// Assert the response
			require.Equal(t, tc.expectedCode, rr.Code)

			// Check the body of the response
			if tc.expectedCode == http.StatusOK {
				var response models.APIResponse
				err := json.NewDecoder(rr.Body).Decode(&response)
				require.NoError(t, err)

				// Assert that the body contains the expected response
				require.Equal(t, tc.expectedResponse, response.Data)
			} else {
				var response models.APIResponse
				err := json.NewDecoder(rr.Body).Decode(&response)
				require.NoError(t, err)

				require.Equal(t, tc.expectedResponse, response.Data)
			}
		})
	}
}
