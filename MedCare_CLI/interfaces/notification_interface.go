package interfaces

import "doctor-patient-cli/models"

// NotificationRepository defines the methods related to notifications
type NotificationRepository interface {
	GetNotificationsByUserID(userID string) ([]models.Notification, error)
	GetAllNotifications() ([]models.Notification, error)
	//MarkNotificationAsRead(notificationID int) error
}

// NotificationService defines the methods for notification-related operations
type NotificationService interface {
	GetNotificationsByUserID(userID string) ([]models.Notification, error)
	GetAllNotifications() ([]models.Notification, error)
	//MarkNotificationAsRead(notificationID int) error
}
