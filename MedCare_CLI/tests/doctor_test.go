package tests

import (
	"bytes"
	"doctor-patient-cli/models"
	"doctor-patient-cli/utils"
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestDoctorFunctions(t *testing.T) {
	MockInitDB(t)
	defer utils.CloseDB()

	t.Run("GetDoctorByID Success", func(t *testing.T) {
		userID := "doctor1"
		rows := sqlmock.NewRows([]string{"user_id", "specialization", "experience", "rating"}).
			AddRow(userID, "Cardiologist", 10, 4.5)

		mock.ExpectQuery("SELECT user_id, specialization, experience, rating FROM doctors WHERE user_id = ?").
			WithArgs(userID).
			WillReturnRows(rows)

		doctor, err := models.GetDoctorByID(userID)
		assert.NoError(t, err)
		assert.Equal(t, userID, doctor.UserID)
		assert.Equal(t, "Cardiologist", doctor.Specialization)
		assert.Equal(t, 10, doctor.Experience)
		assert.Equal(t, 4.5, doctor.Rating)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})

	t.Run("GetDoctorByID Failure", func(t *testing.T) {
		userID := "doctor1"
		mock.ExpectQuery("SELECT user_id, specialization, experience, rating FROM doctors WHERE user_id = ?").
			WithArgs(userID).
			WillReturnError(fmt.Errorf("query error"))

		doctor, err := models.GetDoctorByID(userID)
		assert.Error(t, err)
		assert.Equal(t, models.Doctor{}, doctor)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})

	t.Run("GetAllDoctors Success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"user_id", "specialization", "experience", "rating"}).
			AddRow("doctor1", "Cardiologist", 10, 4.5).
			AddRow("doctor2", "Neurologist", 8, 4.0)

		mock.ExpectQuery("SELECT user_id, specialization, experience, rating FROM doctors").
			WillReturnRows(rows)

		doctors, err := models.GetAllDoctors()
		assert.NoError(t, err)
		assert.Len(t, doctors, 2)
		assert.Equal(t, "doctor1", doctors[0].UserID)
		assert.Equal(t, "Cardiologist", doctors[0].Specialization)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})

	t.Run("GetAllDoctors Query Error", func(t *testing.T) {
		mock.ExpectQuery("SELECT user_id, specialization, experience, rating FROM doctors").
			WillReturnError(fmt.Errorf("query error"))

		doctors, err := models.GetAllDoctors()
		assert.Error(t, err)
		assert.Nil(t, doctors)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})

	t.Run("UpdateDoctorExperience Success", func(t *testing.T) {
		userID := "doctor1"
		experience := 12

		// Expect the exact SQL query and arguments
		mock.ExpectExec(regexp.QuoteMeta("UPDATE doctors SET experience = ? WHERE user_id = ?")).
			WithArgs(experience, userID).
			WillReturnResult(sqlmock.NewResult(1, 1))

		// Call the UpdateDoctorExperience function
		err := models.UpdateDoctorExperience(userID, experience)
		assert.NoError(t, err)

		// Ensure all expectations are met
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})

	t.Run("UpdateDoctorExperience Failure", func(t *testing.T) {
		userID := "doctor1"
		experience := 12

		// Expect the exact SQL query with regexp.QuoteMeta
		mock.ExpectExec(regexp.QuoteMeta("UPDATE doctors SET experience = ? WHERE user_id = ?")).
			WithArgs(experience, userID).
			WillReturnError(fmt.Errorf("update error"))

		// Call the UpdateDoctorExperience function
		err := models.UpdateDoctorExperience(userID, experience)

		// Check that an error is returned
		assert.Error(t, err)

		// Ensure all expectations are met
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})

	t.Run("UpdateDoctorSpecialization Success", func(t *testing.T) {
		userID := "doctor1"
		specialization := "Surgeon"

		// Expect the exact SQL query with regexp.QuoteMeta
		mock.ExpectExec(regexp.QuoteMeta("UPDATE doctors SET specialization = ? WHERE user_id = ?")).
			WithArgs(specialization, userID).
			WillReturnResult(sqlmock.NewResult(1, 1))

		// Call the UpdateDoctorSpecialization function
		err := models.UpdateDoctorSpecialization(userID, specialization)
		assert.NoError(t, err)

		// Ensure all expectations are met
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})

	t.Run("UpdateDoctorSpecialization Failure", func(t *testing.T) {
		userID := "doctor1"
		specialization := "Surgeon"

		// Expect the exact SQL query with regexp.QuoteMeta
		mock.ExpectExec(regexp.QuoteMeta("UPDATE doctors SET specialization = ? WHERE user_id = ?")).
			WithArgs(specialization, userID).
			WillReturnError(fmt.Errorf("update error"))

		// Call the UpdateDoctorSpecialization function
		err := models.UpdateDoctorSpecialization(userID, specialization)
		assert.Error(t, err)

		// Ensure all expectations are met
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})

	t.Run("ViewDoctorSpecificProfile Success", func(t *testing.T) {
		userID := "doctor1"
		expectedSpecialization := "Cardiologist"
		expectedExperience := 10
		expectedRating := 4.5

		// Set up mock rows to return
		rows := mock.NewRows([]string{"specialization", "experience", "rating"}).
			AddRow(expectedSpecialization, expectedExperience, expectedRating)

		// Expect the exact SQL query
		mock.ExpectQuery(regexp.QuoteMeta("SELECT specialization, experience, rating FROM doctors WHERE user_id = ?")).
			WithArgs(userID).
			WillReturnRows(rows)

		// Capture the output of fmt.Println
		var buf bytes.Buffer
		stdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		// Call the function
		models.ViewDoctorSpecificProfile(userID)

		// Restore the original stdout
		w.Close()
		os.Stdout = stdout
		_, _ = buf.ReadFrom(r)

		// Adjust the expected output to include the extra spaces
		expectedOutput := fmt.Sprintf("Specialization:  %s\nExperience:  %d\nRating:  %.1f\n", expectedSpecialization, expectedExperience, expectedRating)
		assert.Equal(t, expectedOutput, buf.String())

		// Ensure all expectations are met
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})
}
