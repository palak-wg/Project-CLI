package services_test

import (
	"doctor-patient-cli/models"
	"doctor-patient-cli/services"
	"doctor-patient-cli/tests/mocks"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestPatientService_GetPatientByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockPatientRepository(ctrl)
	service := services.NewPatientService(mockRepo)

	testCases := []struct {
		name           string
		patientID      string
		mockReturnData *models.Patient
		mockReturnErr  error
		expectedResult *models.Patient
		expectedErr    error
	}{
		{
			name:           "Success case - patient retrieved",
			patientID:      "123",
			mockReturnData: &models.Patient{User: models.User{UserID: "123", Name: "John Doe", Age: 30, Gender: "Male", Email: "john.doe@example.com", PhoneNumber: "555-5555"}},
			mockReturnErr:  nil,
			expectedResult: &models.Patient{User: models.User{UserID: "123", Name: "John Doe", Age: 30, Gender: "Male", Email: "john.doe@example.com", PhoneNumber: "555-5555"}},
			expectedErr:    nil,
		},
		{
			name:           "Failure case - patient not found",
			patientID:      "123",
			mockReturnData: nil,
			mockReturnErr:  errors.New("patient not found"),
			expectedResult: nil,
			expectedErr:    errors.New("patient not found"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo.EXPECT().GetPatientByID(tc.patientID).Return(tc.mockReturnData, tc.mockReturnErr)

			patient, err := service.GetPatientByID(tc.patientID)

			assert.Equal(t, tc.expectedResult, patient)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}

func TestPatientService_UpdatePatientDetails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockPatientRepository(ctrl)
	service := services.NewPatientService(mockRepo)

	patient := &models.Patient{
		User: models.User{
			UserID:      "123",
			Name:        "John Doe",
			Age:         30,
			Gender:      "Male",
			Email:       "john.doe@example.com",
			PhoneNumber: "555-5555",
		},

		MedicalHistory: "No History",
	}

	testCases := []struct {
		name          string
		patient       *models.Patient
		mockReturnErr error
		expectedErr   error
	}{
		{
			name:          "Success case - patient details updated",
			patient:       patient,
			mockReturnErr: nil,
			expectedErr:   nil,
		},
		{
			name:          "Failure case - error updating patient details",
			patient:       patient,
			mockReturnErr: errors.New("update error"),
			expectedErr:   errors.New("update error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo.EXPECT().UpdatePatientDetails(tc.patient).Return(tc.mockReturnErr)

			err := service.UpdatePatientDetails(tc.patient)

			assert.Equal(t, tc.expectedErr, err)
		})
	}
}

//package services
//
//import (
//	"bytes"
//	"doctor-patient-cli/services"
//	"doctor-patient-cli/tests/mockDB"
//	"doctor-patient-cli/utils"
//	"fmt"
//	"io"
//	"os"
//	"testing"
//
//	"github.com/DATA-DOG/go-sqlmock"
//	"github.com/stretchr/testify/assert"
//)
//
//func TestGetPatientByID(t *testing.T) {
//	mockDB.MockInitDB(t)
//	defer utils.CloseDB()
//
//	t.Run("GetPatientByID Success", func(t *testing.T) {
//		userID := "patient1"
//
//		// Mock the patient query result
//		mockDB.Mock.ExpectQuery("SELECT user_id, medical_history FROM patients WHERE user_id = ?").
//			WithArgs(userID).
//			WillReturnRows(sqlmock.NewRows([]string{"user_id", "medical_history"}).
//				AddRow(userID, "No History"))
//
//		// Call the function
//		patient, err := services.GetPatientByID(userID)
//
//		// Check the results
//		assert.NoError(t, err)
//		assert.Equal(t, "patient1", patient.UserID)
//		assert.Equal(t, "No History", patient.MedicalHistory)
//
//		// Ensure all expectations are met
//		if err := mockDB.Mock.ExpectationsWereMet(); err != nil {
//			t.Errorf("there were unfulfilled expectations: %v", err)
//		}
//	})
//
//	t.Run("GetPatientByID Failure", func(t *testing.T) {
//		userID := "patient2"
//
//		// Mock the patient query result with an error
//		mockDB.Mock.ExpectQuery("SELECT user_id, medical_history FROM patients WHERE user_id = ?").
//			WithArgs(userID).
//			WillReturnError(fmt.Errorf("query error"))
//
//		// Call the function
//		_, err := services.GetPatientByID(userID)
//
//		// Check for the error
//		assert.Error(t, err)
//
//		// Ensure all expectations are met
//		if err := mockDB.Mock.ExpectationsWereMet(); err != nil {
//			t.Errorf("there were unfulfilled expectations: %v", err)
//		}
//	})
//}
//
//func TestViewPatientDetails(t *testing.T) {
//	mockDB.MockInitDB(t)
//	defer utils.CloseDB()
//
//	t.Run("ViewPatientDetails Success", func(t *testing.T) {
//		userID := "patient1"
//
//		// Mock the patient query result
//		mockDB.Mock.ExpectQuery("SELECT medical_history FROM patients WHERE user_id = ?").
//			WithArgs(userID).
//			WillReturnRows(sqlmock.NewRows([]string{"medical_history"}).
//				AddRow("No History"))
//
//		// Capture the output
//		output := captureOutput(func() {
//			services.ViewPatientDetails(userID)
//		})
//
//		// Trim any additional spaces around the actual output
//		output = fmt.Sprintf("Medical History: %s", "No History\n")
//
//		// Check the output
//		expectedOutput := "Medical History: No History\n"
//		assert.Equal(t, expectedOutput, output)
//
//		// Ensure all expectations are met
//		if err := mockDB.Mock.ExpectationsWereMet(); err != nil {
//			t.Errorf("there were unfulfilled expectations: %v", err)
//		}
//	})
//}
//
//func captureOutput(f func()) string {
//	old := os.Stdout
//	r, w, _ := os.Pipe()
//	os.Stdout = w
//
//	f()
//
//	w.Close()
//	os.Stdout = old
//
//	var buf bytes.Buffer
//	io.Copy(&buf, r)
//
//	return buf.String()
//}
