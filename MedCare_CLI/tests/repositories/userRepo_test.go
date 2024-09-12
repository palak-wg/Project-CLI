package repositories_test

import (
	"database/sql"
	"doctor-patient-cli/models"
	"doctor-patient-cli/repositories"
	"errors"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser_PatientSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repositories.NewUserRepository(db)

	user := &models.User{
		UserID:      "testpatient",
		Password:    "password",
		Name:        "Test Patient",
		Age:         30,
		Gender:      "Male",
		Email:       "testpatient@example.com",
		PhoneNumber: "1234567890",
		UserType:    "patient",
		IsApproved:  true,
	}

	// Expect INSERT query to add the user
	mock.ExpectExec("INSERT INTO users").
		WithArgs(user.UserID, user.Password, user.Name, user.Age, user.Gender, user.Email, user.PhoneNumber, user.UserType, 0).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Expect INSERT query for patient in the patients table
	mock.ExpectExec("INSERT INTO patients").
		WithArgs(user.UserID, "No History").
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Expect notification for patient registration
	mock.ExpectExec("INSERT INTO notifications").
		WithArgs(user.UserID, fmt.Sprintf("welcome %s to the application.", user.UserID)).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.CreateUser(user)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCreateUser_DoctorSignupRequest(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repositories.NewUserRepository(db)

	user := &models.User{
		UserID:      "testdoctor",
		Password:    "password",
		Name:        "Test Doctor",
		Age:         35,
		Gender:      "Female",
		Email:       "testdoctor@example.com",
		PhoneNumber: "9876543210",
		UserType:    "doctor",
		IsApproved:  false,
	}

	// Expect INSERT query to add the doctor user
	mock.ExpectExec("INSERT INTO users").
		WithArgs(user.UserID, user.Password, user.Name, user.Age, user.Gender, user.Email, user.PhoneNumber, user.UserType, 0).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Expect a notification for admin to approve doctor signup
	mock.ExpectExec("INSERT INTO notifications").
		WithArgs("admin", fmt.Sprintf("Please approve %s signup request for doctor role.", user.UserID)).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.CreateUser(user)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCreateUser_ErrorHandling(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repositories.NewUserRepository(db)

	user := &models.User{
		UserID:      "erroruser",
		Password:    "password",
		Name:        "Error User",
		Age:         25,
		Gender:      "Other",
		Email:       "error@example.com",
		PhoneNumber: "1111111111",
		UserType:    "patient",
		IsApproved:  true,
	}

	// Simulate an error in the first INSERT statement
	mock.ExpectExec("INSERT INTO users").
		WithArgs(user.UserID, user.Password, user.Name, user.Age, user.Gender, user.Email, user.PhoneNumber, user.UserType, 0).
		WillReturnError(errors.New("insert failed"))

	err = repo.CreateUser(user)
	assert.Error(t, err)
	assert.Equal(t, "insert failed", err.Error())
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCreateUser_DoctorNotification_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repositories.NewUserRepository(db)

	user := &models.User{
		UserID:      "testdoctor",
		Password:    "password",
		Name:        "Test Doctor",
		Age:         35,
		Gender:      "Female",
		Email:       "testdoctor@example.com",
		PhoneNumber: "9876543210",
		UserType:    "doctor",
		IsApproved:  false,
	}

	// Expect INSERT query to add the doctor user
	mock.ExpectExec("INSERT INTO users").
		WithArgs(user.UserID, user.Password, user.Name, user.Age, user.Gender, user.Email, user.PhoneNumber, user.UserType, 0).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Simulate an error during the notification creation
	mock.ExpectExec("INSERT INTO notifications").
		WithArgs("admin", fmt.Sprintf("Please approve %s signup request for doctor role.", user.UserID)).
		WillReturnError(errors.New("notification failed"))

	err = repo.CreateUser(user)
	assert.Error(t, err)
	assert.Equal(t, "notification failed", err.Error())
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUserByID_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repositories.NewUserRepository(db)

	user := &models.User{
		UserID:      "testuser",
		Password:    "password",
		Name:        "Test User",
		Age:         30,
		Gender:      "Male",
		Email:       "test@example.com",
		PhoneNumber: "1234567890",
		UserType:    "patient",
		IsApproved:  true,
	}

	// Expect SELECT query to get the user
	rows := sqlmock.NewRows([]string{"user_id", "password", "username", "age", "gender", "email", "phone_number", "user_type", "is_approved"}).
		AddRow(user.UserID, user.Password, user.Name, user.Age, user.Gender, user.Email, user.PhoneNumber, user.UserType, user.IsApproved)
	mock.ExpectQuery("SELECT user_id, password, username, age, gender, email, phone_number, user_type, is_approved FROM users WHERE user_id = ?").
		WithArgs(user.UserID).
		WillReturnRows(rows)

	result, err := repo.GetUserByID(user.UserID)
	assert.NoError(t, err)
	assert.Equal(t, user, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUserByID_NotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repositories.NewUserRepository(db)

	// Expect SELECT query to return no rows
	mock.ExpectQuery("SELECT user_id, password, username, age, gender, email, phone_number, user_type, is_approved FROM users WHERE user_id = ?").
		WithArgs("nonexistent").
		WillReturnError(sql.ErrNoRows)

	result, err := repo.GetUserByID("nonexistent")
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Equal(t, "user not found", err.Error())
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateName(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock database connection: %v", err)
	}
	defer db.Close()

	repo := repositories.NewUserRepository(db)

	// Define the parameters for the update
	userID := "testuser"
	newName := "Updated Name"

	// Expect the UPDATE query to be executed
	mock.ExpectExec("^UPDATE users SET Name = \\? WHERE UserID = \\?$").
		WithArgs(newName, userID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Test UpdateName
	err = repo.UpdateName(userID, newName)
	assert.NoError(t, err)
}

func TestUpdateAge(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock database connection: %v", err)
	}
	defer db.Close()

	repo := repositories.NewUserRepository(db)

	// Define the parameters for the update
	userID := "testuser"
	newAge := "31"

	// Expect the UPDATE query to be executed
	mock.ExpectExec("^UPDATE users SET Age = \\? WHERE UserID = \\?$").
		WithArgs(newAge, userID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Test UpdateAge
	err = repo.UpdateAge(userID, newAge)
	assert.NoError(t, err)
}

func TestUpdateEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock database connection: %v", err)
	}
	defer db.Close()

	repo := repositories.NewUserRepository(db)

	// Define the parameters for the update
	userID := "testuser"
	newEmail := "updated@example.com"

	// Expect the UPDATE query to be executed
	mock.ExpectExec("^UPDATE users SET Email = \\? WHERE UserID = \\?$").
		WithArgs(newEmail, userID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Test UpdateEmail
	err = repo.UpdateEmail(userID, newEmail)
	assert.NoError(t, err)
}

func TestUpdatePhoneNumber(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock database connection: %v", err)
	}
	defer db.Close()

	repo := repositories.NewUserRepository(db)

	// Define the parameters for the update
	userID := "testuser"
	newPhoneNumber := "0987654321"

	// Expect the UPDATE query to be executed
	mock.ExpectExec("^UPDATE users SET PhoneNumber = \\? WHERE UserID = \\?$").
		WithArgs(newPhoneNumber, userID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Test UpdatePhoneNumber
	err = repo.UpdatePhoneNumber(userID, newPhoneNumber)
	assert.NoError(t, err)
}

func TestUpdateGender(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock database connection: %v", err)
	}
	defer db.Close()

	repo := repositories.NewUserRepository(db)

	// Define the parameters for the update
	userID := "testuser"
	newGender := "Female"

	// Expect the UPDATE query to be executed
	mock.ExpectExec("^UPDATE users SET Gender = \\? WHERE UserID = \\?$").
		WithArgs(newGender, userID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Test UpdateGender
	err = repo.UpdateGender(userID, newGender)
	assert.NoError(t, err)
}

func TestUpdatePassword(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock database connection: %v", err)
	}
	defer db.Close()

	repo := repositories.NewUserRepository(db)

	// Define the parameters for the update
	userID := "testuser"
	newPassword := "newpassword"

	// Expect the UPDATE query to be executed
	mock.ExpectExec("^UPDATE users SET Password = \\? WHERE UserID = \\?$").
		WithArgs(newPassword, userID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Test UpdatePassword
	err = repo.UpdatePassword(userID, newPassword)
	assert.NoError(t, err)
}
