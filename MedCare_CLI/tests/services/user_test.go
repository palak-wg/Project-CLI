package services

import (
	"doctor-patient-cli/models"
	"doctor-patient-cli/services"
	"doctor-patient-cli/tests/mockDB"
	"doctor-patient-cli/utils"
	"fmt"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	mockDB.MockInitDB(t)
	defer utils.CloseDB()

	t.Run("CreateUser Success", func(t *testing.T) {
		user := models.User{
			UserID:      "user1",
			Password:    "password123",
			Username:    "John Doe",
			Age:         30,
			Gender:      "Male",
			Email:       "john.doe@example.com",
			PhoneNumber: "1234567890",
			UserType:    "patient",
		}

		mockDB.Mock.ExpectExec("INSERT INTO users").WithArgs(user.UserID, user.Password, user.Username, user.Age, user.Gender, user.Email, user.PhoneNumber, user.UserType, 0).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mockDB.Mock.ExpectExec("INSERT INTO patients").WithArgs(user.UserID, "No History").
			WillReturnResult(sqlmock.NewResult(1, 1))
		mockDB.Mock.ExpectExec("INSERT INTO notifications").WithArgs(user.UserID, "welcome user1 to the application.").
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := services.CreateUser(user)
		assert.NoError(t, err)

		if err := mockDB.Mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})

	t.Run("CreateUser Failure", func(t *testing.T) {
		user := models.User{
			UserID:      "user2",
			Password:    "password123",
			Username:    "Jane Doe",
			Age:         28,
			Gender:      "Female",
			Email:       "jane.doe@example.com",
			PhoneNumber: "0987654321",
			UserType:    "doctor",
		}

		mockDB.Mock.ExpectExec("INSERT INTO users").WithArgs(user.UserID, user.Password, user.Username, user.Age, user.Gender, user.Email, user.PhoneNumber, user.UserType, 0).
			WillReturnError(fmt.Errorf("insert error"))

		err := services.CreateUser(user)
		assert.Error(t, err)

		if err := mockDB.Mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})
}

func TestGetUserByID(t *testing.T) {
	mockDB.MockInitDB(t)
	defer utils.CloseDB()

	t.Run("GetUserByID Success", func(t *testing.T) {
		userID := "user1"
		mockDB.Mock.ExpectQuery("SELECT user_id, password, username, age, gender, email, phone_number, user_type, is_approved FROM users WHERE user_id = ?").
			WithArgs(userID).
			WillReturnRows(sqlmock.NewRows([]string{"user_id", "password", "username", "age", "gender", "email", "phone_number", "user_type", "is_approved"}).
				AddRow(userID, "password123", "John Doe", 30, "Male", "john.doe@example.com", "1234567890", "patient", true))

		user, err := services.GetUserByID(userID)
		assert.NoError(t, err)
		assert.Equal(t, "John Doe", user.Username)

		if err := mockDB.Mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})

	t.Run("GetUserByID Failure", func(t *testing.T) {
		userID := "nonexistent"
		mockDB.Mock.ExpectQuery("SELECT user_id, password, username, age, gender, email, phone_number, user_type, is_approved FROM users WHERE user_id = ?").
			WithArgs(userID).
			WillReturnError(fmt.Errorf("query error"))

		_, err := services.GetUserByID(userID)
		assert.Error(t, err)

		if err := mockDB.Mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})
}

func TestGetAllUserIDs(t *testing.T) {
	mockDB.MockInitDB(t)
	defer utils.CloseDB()

	t.Run("GetAllUserIDs Success", func(t *testing.T) {
		mockDB.Mock.ExpectQuery("SELECT user_id FROM users").
			WillReturnRows(sqlmock.NewRows([]string{"user_id"}).AddRow("user1").AddRow("user2"))

		userIDs, err := services.GetAllUserIDs()
		assert.NoError(t, err)
		assert.Equal(t, []string{"user1", "user2"}, userIDs)

		if err := mockDB.Mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})

	t.Run("GetAllUserIDs Failure", func(t *testing.T) {
		mockDB.Mock.ExpectQuery("SELECT user_id FROM users").
			WillReturnError(fmt.Errorf("query error"))

		_, err := services.GetAllUserIDs()
		assert.Error(t, err)

		if err := mockDB.Mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})
}

func TestUpdateUsername(t *testing.T) {
	mockDB.MockInitDB(t)
	defer utils.CloseDB()

	t.Run("UpdateUsername Success", func(t *testing.T) {
		userID := "user1"
		newUsername := "NewUsername"

		// Expect the exact SQL query pattern
		mockDB.Mock.ExpectExec(`^UPDATE users SET username = \? WHERE user_id = \?$`).
			WithArgs(newUsername, userID).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := services.UpdateUsername(userID, newUsername)
		assert.NoError(t, err)

		if err := mockDB.Mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})

	t.Run("UpdateUsername Failure", func(t *testing.T) {
		userID := "user1"
		newUsername := "NewUsername"

		// Expect the exact SQL query pattern
		mockDB.Mock.ExpectExec(`^UPDATE users SET username = \? WHERE user_id = \?$`).
			WithArgs(newUsername, userID).
			WillReturnError(fmt.Errorf("update error"))

		err := services.UpdateUsername(userID, newUsername)
		assert.Error(t, err)

		if err := mockDB.Mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})
}

func TestUpdateAge(t *testing.T) {
	mockDB.MockInitDB(t)
	defer utils.CloseDB()

	t.Run("UpdateAge Success", func(t *testing.T) {
		userID := "user1"
		newAge := 35

		// Ensure the SQL pattern matches exactly
		mockDB.Mock.ExpectExec(`^UPDATE users SET age = \? WHERE user_id = \?$`).
			WithArgs(newAge, userID).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := services.UpdateAge(userID, newAge)
		assert.NoError(t, err)

		if err := mockDB.Mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})

	t.Run("UpdateAge Failure", func(t *testing.T) {
		userID := "user1"
		newAge := 35

		// Ensure the SQL pattern matches exactly
		mockDB.Mock.ExpectExec(`^UPDATE users SET age = \? WHERE user_id = \?$`).
			WithArgs(newAge, userID).
			WillReturnError(fmt.Errorf("update error"))

		err := services.UpdateAge(userID, newAge)
		assert.Error(t, err)

		if err := mockDB.Mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})
}

func TestUpdateGender(t *testing.T) {
	mockDB.MockInitDB(t)
	defer utils.CloseDB()

	t.Run("Success", func(t *testing.T) {
		userID := "user1"
		gender := "Male"

		mockDB.Mock.ExpectExec(`^UPDATE users SET gender = \? WHERE user_id = \?$`).
			WithArgs(gender, userID).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := services.UpdateGender(userID, gender)
		assert.NoError(t, err)

		if err := mockDB.Mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})

	t.Run("Failure", func(t *testing.T) {
		userID := "user1"
		gender := "Male"

		mockDB.Mock.ExpectExec(`^UPDATE users SET gender = \? WHERE user_id = \?$`).
			WithArgs(gender, userID).
			WillReturnError(fmt.Errorf("update error"))

		err := services.UpdateGender(userID, gender)
		assert.Error(t, err)

		if err := mockDB.Mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})
}

func TestUpdateEmail(t *testing.T) {
	mockDB.MockInitDB(t)
	defer utils.CloseDB()

	t.Run("Success", func(t *testing.T) {
		userID := "user1"
		email := "newemail@example.com"

		mockDB.Mock.ExpectExec(`^UPDATE users SET email = \? WHERE user_id = \?$`).
			WithArgs(email, userID).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := services.UpdateEmail(userID, email)
		assert.NoError(t, err)

		if err := mockDB.Mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})

	t.Run("Failure", func(t *testing.T) {
		userID := "user1"
		email := "newemail@example.com"

		mockDB.Mock.ExpectExec(`^UPDATE users SET email = \? WHERE user_id = \?$`).
			WithArgs(email, userID).
			WillReturnError(fmt.Errorf("update error"))

		err := services.UpdateEmail(userID, email)
		assert.Error(t, err)

		if err := mockDB.Mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})
}

func TestUpdatePhoneNumber(t *testing.T) {
	mockDB.MockInitDB(t)
	defer utils.CloseDB()

	t.Run("Success", func(t *testing.T) {
		userID := "user1"
		phoneNumber := "1234567890"

		mockDB.Mock.ExpectExec(`^UPDATE users SET phone_number = \? WHERE user_id = \?$`).
			WithArgs(phoneNumber, userID).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := services.UpdatePhoneNumber(userID, phoneNumber)
		assert.NoError(t, err)

		if err := mockDB.Mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})

	t.Run("Failure", func(t *testing.T) {
		userID := "user1"
		phoneNumber := "1234567890"

		mockDB.Mock.ExpectExec(`^UPDATE users SET phone_number = \? WHERE user_id = \?$`).
			WithArgs(phoneNumber, userID).
			WillReturnError(fmt.Errorf("update error"))

		err := services.UpdatePhoneNumber(userID, phoneNumber)
		assert.Error(t, err)

		if err := mockDB.Mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})
}

func TestUpdatePassword(t *testing.T) {
	mockDB.MockInitDB(t)
	defer utils.CloseDB()

	t.Run("Success", func(t *testing.T) {
		userID := "user1"
		password := "newpassword"

		mockDB.Mock.ExpectExec(`^UPDATE users SET password = \? WHERE user_id = \?$`).
			WithArgs(password, userID).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := services.UpdatePassword(userID, password)
		assert.NoError(t, err)

		if err := mockDB.Mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})

	t.Run("Failure", func(t *testing.T) {
		userID := "user1"
		password := "newpassword"

		mockDB.Mock.ExpectExec(`^UPDATE users SET password = \? WHERE user_id = \?$`).
			WithArgs(password, userID).
			WillReturnError(fmt.Errorf("update error"))

		err := services.UpdatePassword(userID, password)
		assert.Error(t, err)

		if err := mockDB.Mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})
}

func TestViewProfile_Success(t *testing.T) {
	// Mocking the database
	mockDB.MockInitDB(t)
	defer utils.CloseDB()

	t.Run("ViewProfile RedirectToPatient", func(t *testing.T) {
		// Define test data
		user := models.User{
			UserID:      "user1",
			Username:    "JohnDoe",
			Age:         30,
			Gender:      "Male",
			Email:       "john.doe@example.com",
			PhoneNumber: "123-456-7890",
			UserType:    "doctor",
			IsApproved:  true,
		}

		// Mocking the database query
		mockDB.Mock.ExpectQuery("SELECT user_id, username, age, gender,email, phone_number, user_type FROM users WHERE user_id = ?").
			WithArgs(user.UserID).
			WillReturnRows(sqlmock.NewRows([]string{"user_id", "username", "age", "gender", "email", "phone_number", "user_type"}).
				AddRow(user.UserID, user.Username, user.Age, user.Gender, user.Email, user.PhoneNumber, user.UserType))

		// Redirecting stdout to capture output
		output := captureOutput(func() {
			services.ViewProfile(user)
		})

		// Assertions on the captured output
		expectedOutput := fmt.Sprintf("User ID: %v\nFirst Name: %v\nAge: %v\nGender: %v\nEmail: %v\nPhoneNumber: %v\n", user.UserID, user.Username, user.Age, user.Gender, user.Email, user.PhoneNumber)
		if !contains(output, expectedOutput) {
			t.Errorf("expected output to contain: %v, but got: %v", expectedOutput, output)
		}

		// Check that all expectations were met
		if err := mockDB.Mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})

	t.Run("ViewProfile RedirectToPatient", func(t *testing.T) {
		// Define test data
		user := models.User{
			UserID:      "user2",
			Username:    "John",
			Age:         30,
			Gender:      "Male",
			Email:       "john23@example.com",
			PhoneNumber: "123-456-7890",
			UserType:    "patient",
			IsApproved:  false,
		}

		// Mocking the database query
		mockDB.Mock.ExpectQuery("SELECT user_id, username, age, gender,email, phone_number, user_type FROM users WHERE user_id = ?").
			WithArgs(user.UserID).
			WillReturnRows(sqlmock.NewRows([]string{"user_id", "username", "age", "gender", "email", "phone_number", "user_type"}).
				AddRow(user.UserID, user.Username, user.Age, user.Gender, user.Email, user.PhoneNumber, user.UserType))

		// Redirecting stdout to capture output
		output := captureOutput(func() {
			services.ViewProfile(user)
		})

		// Assertions on the captured output
		expectedOutput := fmt.Sprintf("User ID: %v\nFirst Name: %v\nAge: %v\nGender: %v\nEmail: %v\nPhoneNumber: %v\n", user.UserID, user.Username, user.Age, user.Gender, user.Email, user.PhoneNumber)
		if !contains(output, expectedOutput) {
			t.Errorf("expected output to contain: %v, but got: %v", expectedOutput, output)
		}

		// Check that all expectations were met
		if err := mockDB.Mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})
}

func contains(s, substr string) bool {
	return strings.Contains(s, substr)
}
