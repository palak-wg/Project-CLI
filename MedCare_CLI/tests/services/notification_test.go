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

func TestGetNotificationsByUserID(t *testing.T) {
	// Initialize the mock database
	mockDB.MockInitDB(t)
	defer utils.CloseDB()

	t.Run("GetNotificationsByUserID Success", func(t *testing.T) {
		// Set up mock rows to return
		rows := sqlmock.NewRows([]string{"user_id", "content", "timestamp"}).
			AddRow("user1", "Notification 1", "2023-01-01 10:00:00").
			AddRow("user1", "Notification 2", "2023-01-02 11:00:00")

		// Expect the query and return the mock rows
		mockDB.Mock.ExpectQuery("SELECT user_id, content, timestamp FROM notifications WHERE user_id = ?").
			WithArgs("user1").
			WillReturnRows(rows)

		// Call the GetNotificationsByUserID function
		notifications, err := services.GetNotificationsByUserID("user1")

		// Check for no errors
		assert.NoError(t, err)
		assert.Len(t, notifications, 2)
		assert.Equal(t, "Notification 1", notifications[0].Content)

		// Ensure all expectations are met
		if err := mockDB.Mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})

	t.Run("GetNotificationsByUserID Query Error", func(t *testing.T) {
		// Expect the query and simulate an error
		mockDB.Mock.ExpectQuery("SELECT user_id, content, timestamp FROM notifications WHERE user_id = ?").
			WithArgs("user1").
			WillReturnError(fmt.Errorf("query error"))

		// Call the GetNotificationsByUserID function
		_, err := services.GetNotificationsByUserID("user1")

		// Check that an error was returned
		assert.Error(t, err)

		// Ensure all expectations are met
		if err := mockDB.Mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})
}

func TestGetAllNotifications(t *testing.T) {
	// Initialize the mock database
	mockDB.MockInitDB(t)
	defer utils.CloseDB()

	t.Run("GetAllNotifications Success", func(t *testing.T) {
		// Set up mock rows to return
		rows := sqlmock.NewRows([]string{"user_id", "content", "timestamp"}).
			AddRow("user1", "Notification 1", "2023-01-01 10:00:00").
			AddRow("user2", "Notification 2", "2023-01-02 11:00:00")

		// Expect the query and return the mock rows
		mockDB.Mock.ExpectQuery("SELECT user_id, content, timestamp FROM notifications").
			WillReturnRows(rows)

		// Call the GetAllNotifications function
		notifications, err := services.GetAllNotifications()

		// Check for no errors
		assert.NoError(t, err)
		assert.Len(t, notifications, 2)
		assert.Equal(t, "Notification 1", notifications[0].Content)

		// Ensure all expectations are met
		if err := mockDB.Mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})

	t.Run("GetAllNotifications Query Error", func(t *testing.T) {
		// Expect the query and simulate an error
		mockDB.Mock.ExpectQuery("SELECT user_id, content, timestamp FROM notifications").
			WillReturnError(fmt.Errorf("query error"))

		// Call the GetAllNotifications function
		_, err := services.GetAllNotifications()

		// Check that an error was returned
		assert.Error(t, err)

		// Ensure all expectations are met
		if err := mockDB.Mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})
}
