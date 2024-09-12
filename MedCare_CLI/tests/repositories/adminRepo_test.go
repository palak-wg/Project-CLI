package repositories_test

import (
	"database/sql"
	"doctor-patient-cli/repositories"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestApproveDoctorSignup(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := repositories.NewAdminRepository(db)

	t.Run("success case", func(t *testing.T) {
		mock.ExpectExec("UPDATE users SET is_approved").WithArgs(true, "user1").
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectExec("INSERT INTO doctors").WithArgs("user1", "xxx", 0, 2).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := repo.ApproveDoctorSignup("user1")
		assert.NoError(t, err)
		mock.ExpectationsWereMet()
	})

	t.Run("failure case: updating approval status", func(t *testing.T) {
		mock.ExpectExec("UPDATE users SET is_approved").WithArgs(true, "user1").
			WillReturnError(sql.ErrConnDone)

		err := repo.ApproveDoctorSignup("user1")
		assert.Error(t, err)
		mock.ExpectationsWereMet()
	})

	t.Run("failure case: inserting doctor record", func(t *testing.T) {
		mock.ExpectExec("UPDATE users SET is_approved").WithArgs(true, "user1").
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectExec("INSERT INTO doctors").WithArgs("user1", "xxx", 0, 2).
			WillReturnError(sql.ErrConnDone)

		err := repo.ApproveDoctorSignup("user1")
		assert.Error(t, err)
		mock.ExpectationsWereMet()
	})
}

func TestCreateNotificationForUser(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := repositories.NewAdminRepository(db)

	t.Run("success case", func(t *testing.T) {
		mock.ExpectExec("INSERT INTO notifications").WithArgs("user1", "Notification content").
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := repo.CreateNotificationForUser("user1", "Notification content")
		assert.NoError(t, err)
		mock.ExpectationsWereMet()
	})

	t.Run("failure case", func(t *testing.T) {
		mock.ExpectExec("INSERT INTO notifications").WithArgs("user1", "Notification content").
			WillReturnError(sql.ErrConnDone)

		err := repo.CreateNotificationForUser("user1", "Notification content")
		assert.Error(t, err)
		mock.ExpectationsWereMet()
	})
}

func TestPendingDoctorSignupRequest(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := repositories.NewAdminRepository(db)

	t.Run("success with pending doctor signups", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"user_id", "username"}).
			AddRow("user1", "doctor1").
			AddRow("user2", "doctor2")

		mock.ExpectQuery("SELECT user_id, username FROM users WHERE user_type = `doctor` AND is_approved = 0").
			WillReturnRows(rows)

		doctors, err := repo.PendingDoctorSignupRequest()
		assert.NoError(t, err)
		assert.Len(t, doctors, 2)
		mock.ExpectationsWereMet()
	})

	t.Run("success with no pending doctor signups", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"user_id", "username"})
		mock.ExpectQuery("SELECT user_id, username FROM users WHERE user_type = `doctor` AND is_approved = 0").
			WillReturnRows(rows)

		doctors, err := repo.PendingDoctorSignupRequest()
		assert.NoError(t, err)
		assert.Len(t, doctors, 0)
		mock.ExpectationsWereMet()
	})

	t.Run("failure case", func(t *testing.T) {
		mock.ExpectQuery("SELECT user_id, username FROM users WHERE user_type = `doctor` AND is_approved = 0").
			WillReturnError(sql.ErrConnDone)

		doctors, err := repo.PendingDoctorSignupRequest()
		assert.Error(t, err)
		assert.Nil(t, doctors)
		mock.ExpectationsWereMet()
	})
}

func TestGetAllUsers(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := repositories.NewAdminRepository(db)

	t.Run("success with users", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"user_id", "password", "username", "age", "gender", "email", "phone_number", "user_type", "is_approved"}).
			AddRow("user1", "password1", "User1", 30, "M", "user1@example.com", "1234567890", "doctor", true).
			AddRow("user2", "password2", "User2", 25, "F", "user2@example.com", "0987654321", "patient", false)

		mock.ExpectQuery("^SELECT \\* FROM users$").WillReturnRows(rows)

		users, err := repo.GetAllUsers()
		assert.NoError(t, err)
		assert.Len(t, users, 2)
		mock.ExpectationsWereMet()
	})

	t.Run("success with no users", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"user_id", "password", "username", "age", "gender", "email", "phone_number", "user_type", "is_approved"})
		mock.ExpectQuery("^SELECT \\* FROM users$").WillReturnRows(rows)

		users, err := repo.GetAllUsers()
		assert.NoError(t, err)
		assert.Len(t, users, 0)
		mock.ExpectationsWereMet()
	})

	t.Run("failure case", func(t *testing.T) {
		mock.ExpectQuery("^SELECT \\* FROM users$").WillReturnError(sql.ErrConnDone)

		users, err := repo.GetAllUsers()
		assert.Error(t, err)
		assert.Nil(t, users)
		mock.ExpectationsWereMet()
	})
}
