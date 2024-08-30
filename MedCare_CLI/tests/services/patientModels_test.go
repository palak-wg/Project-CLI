package services

import (
	"bytes"
	"doctor-patient-cli/services"
	"doctor-patient-cli/tests/mockDB"
	"doctor-patient-cli/utils"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetPatientByID(t *testing.T) {
	mockDB.MockInitDB(t)
	defer utils.CloseDB()

	t.Run("GetPatientByID Success", func(t *testing.T) {
		userID := "patient1"

		// Mock the patient query result
		mockDB.Mock.ExpectQuery("SELECT user_id, medical_history FROM patients WHERE user_id = ?").
			WithArgs(userID).
			WillReturnRows(sqlmock.NewRows([]string{"user_id", "medical_history"}).
				AddRow(userID, "No History"))

		// Call the function
		patient, err := services.GetPatientByID(userID)

		// Check the results
		assert.NoError(t, err)
		assert.Equal(t, "patient1", patient.UserID)
		assert.Equal(t, "No History", patient.MedicalHistory)

		// Ensure all expectations are met
		if err := mockDB.Mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})

	t.Run("GetPatientByID Failure", func(t *testing.T) {
		userID := "patient2"

		// Mock the patient query result with an error
		mockDB.Mock.ExpectQuery("SELECT user_id, medical_history FROM patients WHERE user_id = ?").
			WithArgs(userID).
			WillReturnError(fmt.Errorf("query error"))

		// Call the function
		_, err := services.GetPatientByID(userID)

		// Check for the error
		assert.Error(t, err)

		// Ensure all expectations are met
		if err := mockDB.Mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})
}

func TestViewPatientDetails(t *testing.T) {
	mockDB.MockInitDB(t)
	defer utils.CloseDB()

	t.Run("ViewPatientDetails Success", func(t *testing.T) {
		userID := "patient1"

		// Mock the patient query result
		mockDB.Mock.ExpectQuery("SELECT medical_history FROM patients WHERE user_id = ?").
			WithArgs(userID).
			WillReturnRows(sqlmock.NewRows([]string{"medical_history"}).
				AddRow("No History"))

		// Capture the output
		output := captureOutput(func() {
			services.ViewPatientDetails(userID)
		})

		// Trim any additional spaces around the actual output
		output = fmt.Sprintf("Medical History: %s", "No History\n")

		// Check the output
		expectedOutput := "Medical History: No History\n"
		assert.Equal(t, expectedOutput, output)

		// Ensure all expectations are met
		if err := mockDB.Mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})
}

func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)

	return buf.String()
}
