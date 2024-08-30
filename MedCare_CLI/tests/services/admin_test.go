package services

import (
	"doctor-patient-cli/tests/mockDB"
	"errors"
	"testing"

	"doctor-patient-cli/services"
	"doctor-patient-cli/utils"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/fatih/color"
)

func TestPendingDoctorSignupRequest(t *testing.T) {
	// Mocking the database
	mockDB.MockInitDB(t)
	defer utils.CloseDB()

	t.Run("Success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"user_id"}).
			AddRow("doctor1").
			AddRow("doctor2")

		// Notice the exact match with the SQL query
		mockDB.Mock.ExpectQuery(`SELECT user_id FROM users WHERE user_type ='doctor' AND is_approved=0`).
			WillReturnRows(rows)

		// Capturing the output
		color.NoColor = false
		services.PendingDoctorSignupRequest()

		// Ensure all expectations were met
		if err := mockDB.Mock.ExpectationsWereMet(); err != nil {
			t.Errorf("There were unfulfilled expectations: %s", err)
		}
	})

	t.Run("Query Error", func(t *testing.T) {
		mockDB.Mock.ExpectQuery(`SELECT user_id FROM users WHERE user_type ='doctor' AND is_approved=0`).
			WillReturnError(errors.New("query error"))

		// Capturing the output
		color.NoColor = false
		services.PendingDoctorSignupRequest()

		// Ensure all expectations were met
		if err := mockDB.Mock.ExpectationsWereMet(); err != nil {
			t.Errorf("There were unfulfilled expectations: %s", err)
		}
	})

	t.Run("Scan Error", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"user_id"}).
			AddRow("doctor1").
			RowError(0, errors.New("scan error"))

		mockDB.Mock.ExpectQuery(`SELECT user_id FROM users WHERE user_type ='doctor' AND is_approved=0`).
			WillReturnRows(rows)

		// Capturing the output
		color.NoColor = false
		services.PendingDoctorSignupRequest()

		// Ensure all expectations were met
		if err := mockDB.Mock.ExpectationsWereMet(); err != nil {
			t.Errorf("There were unfulfilled expectations: %s", err)
		}

	})
}
