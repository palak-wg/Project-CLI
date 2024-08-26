package tests

import (
	"database/sql"
	"doctor-patient-cli/models"
	"doctor-patient-cli/utils"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetPatientByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	utils.InitDB()

	tests := []struct {
		name      string
		userID    string
		wantErr   bool
		mockSetup func()
	}{
		{
			name:    "Valid Patient",
			userID:  "patient-123",
			wantErr: false,
			mockSetup: func() {
				mock.ExpectQuery("SELECT user_id, username, age, gender,email, phone_number FROM users").
					WithArgs("patient-123").
					WillReturnRows(sqlmock.NewRows([]string{"user_id", "username", "age", "gender", "email", "phone_number"}).
						AddRow("patient-123", "John Doe", 30, "Male", "john@example.com", "1234567890"))

				mock.ExpectQuery("SELECT user_id, medical_history FROM patients").
					WithArgs("patient-123").
					WillReturnRows(sqlmock.NewRows([]string{"user_id", "medical_history"}).
						AddRow("patient-123", "No significant history"))
			},
		},
		{
			name:    "Patient Not Found",
			userID:  "patient-999",
			wantErr: true,
			mockSetup: func() {
				mock.ExpectQuery("SELECT user_id, username, age, gender,email, phone_number FROM users").
					WithArgs("patient-999").
					WillReturnError(sql.ErrNoRows)
			},
		},
		{
			name:    "Database Error",
			userID:  "patient-123",
			wantErr: true,
			mockSetup: func() {
				mock.ExpectQuery("SELECT user_id, username, age, gender,email, phone_number FROM users").
					WithArgs("patient-123").
					WillReturnError(errors.New("some database error"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			_, err := models.GetPatientByID(tt.userID)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestSendMessageToDoctor(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	utils.InitDB()

	patientID := "patient-123"
	doctorID := "doctor-456"
	message := "Hello, I need help."

	mock.ExpectExec("INSERT INTO messages").
		WithArgs(sqlmock.AnyArg(), patientID, doctorID, message).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec("INSERT INTO notifications").
		WithArgs(doctorID, sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = models.SendMessageToDoctor(patientID, doctorID, message)
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
