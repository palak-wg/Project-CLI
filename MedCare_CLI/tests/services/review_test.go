package services

import (
	"doctor-patient-cli/services"
	"doctor-patient-cli/tests/mockDB"
	"doctor-patient-cli/utils"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestAddReview(t *testing.T) {
	mockDB.MockInitDB(t)
	defer utils.CloseDB()

	t.Run("AddReview Success", func(t *testing.T) {
		// Set up mock expectation for the insert query
		mockDB.Mock.ExpectExec("INSERT INTO reviews").
			WithArgs("patient1", "doctor1", "Great doctor!", 5).
			WillReturnResult(sqlmock.NewResult(1, 1))

		// Call the AddReview function
		err := services.AddReview("patient1", "doctor1", "Great doctor!", 5)

		// Check if there was no error
		assert.NoError(t, err)

		// Ensure all expectations are met
		if err := mockDB.Mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})
}

func TestGetAllReviews(t *testing.T) {
	mockDB.MockInitDB(t)
	defer utils.CloseDB()

	t.Run("GetAllReviews Success", func(t *testing.T) {
		// Set up mock rows to return
		rows := sqlmock.NewRows([]string{"patient_id", "doctor_id", "content", "rating"}).
			AddRow("patient1", "doctor1", "Great doctor!", 5).
			AddRow("patient2", "doctor2", "Not bad", 4)

		// Expect the query and set up the rows to return
		mockDB.Mock.ExpectQuery("SELECT patient_id, doctor_id, content, rating FROM reviews").
			WillReturnRows(rows)

		// Call the GetAllReviews function
		reviews, err := services.GetAllReviews()

		// Check if there was no error
		assert.NoError(t, err)
		assert.Len(t, reviews, 2)
		assert.Equal(t, "patient1", reviews[0].PatientID)
		assert.Equal(t, "Great doctor!", reviews[0].Content)

		// Ensure all expectations are met
		if err := mockDB.Mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})

	t.Run("GetAllReviews Query Error", func(t *testing.T) {
		// Expect the query and simulate an error
		mockDB.Mock.ExpectQuery("SELECT patient_id, doctor_id, content, rating FROM reviews").
			WillReturnError(fmt.Errorf("query error"))

		// Call the GetAllReviews function
		_, err := services.GetAllReviews()

		// Check if an error was returned
		assert.Error(t, err)

		// Ensure all expectations are met
		if err := mockDB.Mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})
}
