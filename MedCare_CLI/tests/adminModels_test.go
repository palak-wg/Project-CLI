package tests

import (
	"doctor-patient-cli/models"
	"doctor-patient-cli/utils"
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

// MockSendEmail is a helper function to mock the SendEmail functionality.
func MockSendEmail(to, subject, body string) {
	fmt.Printf("Mock email sent to: %s, subject: %s, body: %s\n", to, subject, body)
}

func TestApproveDoctorSignup(t *testing.T) {
	MockInitDB(t)
	defer utils.CloseDB()

	t.Run("ApproveDoctorSignup Success", func(t *testing.T) {
		userID := "doctor1"

		// Expect the update query to approve the doctor
		mock.ExpectExec(regexp.QuoteMeta("UPDATE users SET is_approved = ? WHERE user_id = ?")).
			WithArgs(true, userID).
			WillReturnResult(sqlmock.NewResult(1, 1))

		// Expect the insert query to add the doctor to the doctors table
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO doctors (user_id, specialization, experience, rating) VALUES (?, ?, ?, ?)")).
			WithArgs(userID, "xxx", 0, 2).
			WillReturnResult(sqlmock.NewResult(1, 1))

		// Expect the insert query to add a notification
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO notifications (user_id, content, timestamp) VALUES (?, ?, ?)")).
			WithArgs(userID, "Your signup request has been approved by the admin.", sqlmock.AnyArg()).
			WillReturnResult(sqlmock.NewResult(1, 1))

		// Expect a call to GetUserByID to retrieve the doctor's email
		mock.ExpectQuery(regexp.QuoteMeta("SELECT user_id, email FROM users WHERE user_id = ?")).
			WithArgs(userID).
			WillReturnRows(sqlmock.NewRows([]string{"user_id", "email"}).AddRow(userID, "doctor1@example.com"))

		// Call the ApproveDoctorSignup function
		err := models.ApproveDoctorSignup(userID)
		assert.NoError(t, err)

		// Ensure all expectations are met
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})

	// Additional failure test cases can be added in a similar manner.
}
