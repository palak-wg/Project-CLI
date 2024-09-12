package repositories

import (
	"database/sql"
	"doctor-patient-cli/interfaces"
	"doctor-patient-cli/models"
	"fmt"
)

// NotificationRepositoryImpl implements the NotificationRepository interface
type NotificationRepositoryImpl struct {
	db *sql.DB
}

// NewNotificationRepository creates a new instance of NotificationRepositoryImpl
func NewNotificationRepository(db *sql.DB) interfaces.NotificationRepository {
	return &NotificationRepositoryImpl{db: db}
}

// GetNotificationsByUserID retrieves notifications for a specific user
func (repo *NotificationRepositoryImpl) GetNotificationsByUserID(userID string) ([]models.Notification, error) {
	query := "SELECT user_id, content, timestamp FROM notifications WHERE user_id = ?"
	rows, err := repo.db.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("error querying notifications by user ID: %v", err)
	}
	defer rows.Close()

	var notifications []models.Notification
	for rows.Next() {
		var n models.Notification
		_ = rows.Scan(&n.UserID, &n.Content, &n.Timestamp)
		notifications = append(notifications, n)
	}
	return notifications, nil
}

// GetAllNotifications retrieves all notifications
func (repo *NotificationRepositoryImpl) GetAllNotifications() ([]models.Notification, error) {
	query := "SELECT user_id, content, timestamp FROM notifications"
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying all notifications: %v", err)
	}
	defer rows.Close()

	var notifications []models.Notification
	for rows.Next() {
		var n models.Notification
		_ = rows.Scan(&n.UserID, &n.Content, &n.Timestamp)
		notifications = append(notifications, n)
	}
	return notifications, nil
}

//// MarkNotificationAsRead marks a notification as read (if you choose to implement it)
//func (repo *NotificationRepositoryImpl) MarkNotificationAsRead(notificationID int) error {
//	query := "UPDATE notifications SET is_read = TRUE WHERE notification_id = ?"
//	_, err := repo.db.Exec(query, notificationID)
//	if err != nil {
//		return fmt.Errorf("error marking notification as read: %v", err)
//	}
//	return nil
//}
