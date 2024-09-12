package repositories_test

import (
	"doctor-patient-cli/models"
	"doctor-patient-cli/repositories"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetNotificationsByUserID_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock db: %v", err)
	}
	defer db.Close()

	repo := repositories.NewNotificationRepository(db)

	mockRows := sqlmock.NewRows([]string{"user_id", "content", "timestamp"}).
		AddRow("123", "Test Notification", "2024-09-01")

	mock.ExpectQuery("SELECT user_id, content, timestamp FROM notifications WHERE user_id = ?").
		WithArgs("123").
		WillReturnRows(mockRows)

	notifications, err := repo.GetNotificationsByUserID("123")

	expected := []models.Notification{
		{UserID: "123", Content: "Test Notification", Timestamp: []uint8("2024-09-01")},
	}

	assert.NoError(t, err)
	assert.Equal(t, expected, notifications)
}

func TestGetNotificationsByUserID_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repositories.NewNotificationRepository(db)

	mock.ExpectQuery("SELECT user_id, content, timestamp FROM notifications WHERE user_id = ?").
		WithArgs("123").
		WillReturnError(errors.New("database error"))

	result, err := repo.GetNotificationsByUserID("123")

	// Assert that an error occurred
	assert.Nil(t, result)
	assert.EqualError(t, err, "error querying notifications by user ID: database error")

	// Ensure all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetAllNotifications_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock db: %v", err)
	}
	defer db.Close()

	repo := repositories.NewNotificationRepository(db)

	mockRows := sqlmock.NewRows([]string{"user_id", "content", "timestamp"}).
		AddRow("123", "Test Notification 1", "2024-09-01").
		AddRow("124", "Test Notification 2", "2024-09-02")

	mock.ExpectQuery("SELECT user_id, content, timestamp FROM notifications").
		WillReturnRows(mockRows)

	notifications, err := repo.GetAllNotifications()

	expected := []models.Notification{
		{UserID: "123", Content: "Test Notification 1", Timestamp: []uint8("2024-09-01")},
		{UserID: "124", Content: "Test Notification 2", Timestamp: []uint8("2024-09-02")},
	}

	assert.NoError(t, err)
	assert.Equal(t, expected, notifications)
}

func TestGetAllNotifications_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repositories.NewNotificationRepository(db)
	mock.ExpectQuery("SELECT user_id, content, timestamp FROM notifications").
		WillReturnError(errors.New("database error"))

	notifications, err := repo.GetAllNotifications()
	assert.Nil(t, notifications)
	assert.ErrorContains(t, err, "database error")
	assert.NoError(t, mock.ExpectationsWereMet())

}
