package repositories_test

import (
	"database/sql"
	"doctor-patient-cli/models"
	"doctor-patient-cli/repositories"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetPatientByID_Success(t *testing.T) {
	// Create a new mock database
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Initialize the repository with the mock DB
	repo := repositories.NewPatientRepository(db)

	// Define the expected patient
	expectedPatient := &models.Patient{
		User: models.User{
			UserID:      "patient123",
			Name:        "John Doe",
			Age:         30,
			Gender:      "Male",
			Email:       "john.doe@example.com",
			PhoneNumber: "1234567890",
		},
	}

	// Expect the SELECT query to be executed and return the expected patient details
	rows := sqlmock.NewRows([]string{"user_id", "username", "age", "gender", "email", "phone_number"}).
		AddRow(expectedPatient.UserID, expectedPatient.Name, expectedPatient.Age, expectedPatient.Gender, expectedPatient.Email, expectedPatient.PhoneNumber)
	mock.ExpectQuery("SELECT user_id, username, age, gender,email, phone_number FROM users WHERE user_id = ?").
		WithArgs(expectedPatient.UserID).
		WillReturnRows(rows)

	// Call the GetPatientByID function
	result, err := repo.GetPatientByID(expectedPatient.UserID)

	// Assert no errors
	assert.NoError(t, err)
	// Assert that the returned patient matches the expected patient
	assert.Equal(t, expectedPatient, result)

	// Ensure all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetPatientByID_NotFound(t *testing.T) {
	// Create a new mock database
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Initialize the repository with the mock DB
	repo := repositories.NewPatientRepository(db)

	// Expect the SELECT query to return no rows
	mock.ExpectQuery("SELECT user_id, username, age, gender,email, phone_number FROM users WHERE user_id = ?").
		WithArgs("nonexistentID").
		WillReturnError(sql.ErrNoRows)

	// Call the GetPatientByID function with a non-existent patient ID
	result, err := repo.GetPatientByID("nonexistentID")

	// Assert that an error occurred and the result is nil
	assert.Nil(t, result)
	assert.EqualError(t, err, "patient not found")

	// Ensure all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetPatientByID_QueryError(t *testing.T) {
	// Create a new mock database
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Initialize the repository with the mock DB
	repo := repositories.NewPatientRepository(db)

	// Simulate a query error (e.g., database connection issue)
	mock.ExpectQuery("SELECT user_id, username, age, gender,email, phone_number FROM users WHERE user_id = ?").
		WithArgs("patient123").
		WillReturnError(errors.New("database error"))

	// Call the GetPatientByID function
	result, err := repo.GetPatientByID("patient123")

	// Assert that an error occurred
	assert.Nil(t, result)
	assert.EqualError(t, err, "database error")

	// Ensure all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdatePatientDetails_Success(t *testing.T) {
	// Create a new mock database
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Initialize the repository with the mock DB
	repo := repositories.NewPatientRepository(db)

	// Define the patient with updated details
	updatedPatient := &models.Patient{
		User: models.User{
			UserID:      "patient123",
			Name:        "John Doe Updated",
			Age:         31,
			Gender:      "Male",
			Email:       "john.updated@example.com",
			PhoneNumber: "0987654321",
		},
		MedicalHistory: "No History",
	}

	// Expect the UPDATE query to be executed successfully
	mock.ExpectExec("UPDATE patients SET name = \\?, age = \\?, gender = \\?, email = \\?, phone_number = \\?, medical_history = \\? WHERE user_id = \\?").
		WithArgs(updatedPatient.Name, updatedPatient.Age, updatedPatient.Gender, updatedPatient.Email, updatedPatient.PhoneNumber, updatedPatient.MedicalHistory, updatedPatient.UserID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Call the UpdatePatientDetails function
	err = repo.UpdatePatientDetails(updatedPatient)

	// Assert no errors
	assert.NoError(t, err)

	// Ensure all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdatePatientDetails_Failure(t *testing.T) {
	// Create a new mock database
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Initialize the repository with the mock DB
	repo := repositories.NewPatientRepository(db)

	// Define the patient with updated details
	updatedPatient := &models.Patient{
		User: models.User{
			UserID:      "patient123",
			Name:        "John Doe Updated",
			Age:         31,
			Gender:      "Male",
			Email:       "john.updated@example.com",
			PhoneNumber: "0987654321",
		},
		MedicalHistory: "No History",
	}

	// Simulate an error during the UPDATE query
	mock.ExpectExec("UPDATE users SET name = \\?, age = \\?, gender = \\?, email = \\?, phone_number = \\? WHERE user_id = \\?").
		WithArgs(updatedPatient.Name, updatedPatient.Age, updatedPatient.Gender, updatedPatient.Email, updatedPatient.PhoneNumber, updatedPatient.UserID).
		WillReturnError(errors.New("update failed"))

	// Call the UpdatePatientDetails function
	err = repo.UpdatePatientDetails(updatedPatient)

	// Assert that an error occurred
	assert.EqualError(t, err, "update failed")

	// Ensure all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}
