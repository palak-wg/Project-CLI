package services

import (
	"doctor-patient-cli/tests/mockDB"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
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
func TestApproveDoctorSignup(t *testing.T) {
	// Initialize sqlmock
	mockDB.MockInitDB(t)
	defer utils.CloseDB()

	// Test cases
	tests := []struct {
		name        string
		userID      string
		mockSetup   func()
		expectedErr error
	}{
		{
			name:   "Success case",
			userID: "doctor123",
			mockSetup: func() {
				// Mock the Update query
				mockDB.Mock.ExpectExec("UPDATE users SET is_approved = \\? WHERE user_id = \\?").
					WithArgs(true, "doctor123").
					WillReturnResult(sqlmock.NewResult(1, 1))

				// Mock the Insert into doctors table
				mockDB.Mock.ExpectExec("INSERT INTO doctors").
					WithArgs("doctor123", "xxx", 0, 2).
					WillReturnResult(sqlmock.NewResult(1, 1))

				// Mock the Insert into notifications
				mockDB.Mock.ExpectExec("INSERT INTO notifications").
					WithArgs("doctor123", "Your signup request has been approved by the admin.").
					WillReturnResult(sqlmock.NewResult(1, 1))

				// Mock the GetUserByID query
				rows := sqlmock.NewRows([]string{"user_id", "password", "username", "age", "gender", "email", "phone_number", "user_type", "is_approved"}).
					AddRow("doctor123", "hashedpassword", "Dr. Who", 45, "Male", "doctor@example.com", "1234567890", "doctor", true)
				mockDB.Mock.ExpectQuery("SELECT user_id, password, username, age, gender, email, phone_number, user_type, is_approved FROM users WHERE user_id = \\?").
					WithArgs("doctor123").
					WillReturnRows(rows)

				// Mock the email sending - this is a separate function and is often not mocked directly.
			},
			expectedErr: nil,
		},
		{
			name:   "Failure case - error approving doctor signup",
			userID: "doctor123",
			mockSetup: func() {
				mockDB.Mock.ExpectExec("UPDATE users SET is_approved = \\? WHERE user_id = \\?").
					WithArgs(true, "doctor123").
					WillReturnError(fmt.Errorf("database error"))
			},
			expectedErr: fmt.Errorf("database error"),
		},
		{
			name:   "Failure case - error creating notification",
			userID: "doctor123",
			mockSetup: func() {
				mockDB.Mock.ExpectExec("UPDATE users SET is_approved = \\? WHERE user_id = \\?").
					WithArgs(true, "doctor123").
					WillReturnResult(sqlmock.NewResult(1, 1))

				mockDB.Mock.ExpectExec("INSERT INTO doctors").
					WithArgs("doctor123", "xxx", 0, 2).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mockDB.Mock.ExpectExec("INSERT INTO notifications").
					WithArgs("doctor123", "Your signup request has been approved by the admin.").
					WillReturnError(fmt.Errorf("notification error"))
			},
			expectedErr: fmt.Errorf("notification error"),
		},
		{
			name:   "Failure case - error fetching doctor",
			userID: "doctor123",
			mockSetup: func() {
				mockDB.Mock.ExpectExec("UPDATE users SET is_approved = \\? WHERE user_id = \\?").
					WithArgs(true, "doctor123").
					WillReturnResult(sqlmock.NewResult(1, 1))

				mockDB.Mock.ExpectExec("INSERT INTO doctors").
					WithArgs("doctor123", "xxx", 0, 2).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mockDB.Mock.ExpectExec("INSERT INTO notifications").
					WithArgs("doctor123", "Your signup request has been approved by the admin.").
					WillReturnResult(sqlmock.NewResult(1, 1))

				// Simulate an error when fetching the doctor
				mockDB.Mock.ExpectQuery("SELECT user_id, password, username, age, gender, email, phone_number, user_type, is_approved FROM users WHERE user_id = \\?").
					WithArgs("doctor123").
					WillReturnError(fmt.Errorf("fetch error"))
			},
			expectedErr: fmt.Errorf("fetch error"),
		},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			err := services.ApproveDoctorSignup(tt.userID)
			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}

			// Ensure all expectations were met
			if err := mockDB.Mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %v", err)
			}
		})
	}
}
