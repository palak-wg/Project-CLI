package services

import (
	"doctor-patient-cli/interfaces"
	"doctor-patient-cli/models"
	"github.com/fatih/color"
)

type NotificationService struct {
	repo interfaces.NotificationRepository
}

func NewNotificationService(repo interfaces.NotificationRepository) *NotificationService {
	return &NotificationService{repo: repo}
}

// GetNotificationsByUserID retrieves notifications for a specific user
func (service *NotificationService) GetNotificationsByUserID(userID string) ([]models.Notification, error) {
	notifications, err := service.repo.GetNotificationsByUserID(userID)
	if err != nil {
		color.Red("Error retrieving notifications for user ID %s: %v", userID, err)
		return nil, err
	}
	return notifications, nil
}

// GetAllNotifications retrieves all notifications
func (service *NotificationService) GetAllNotifications() ([]models.Notification, error) {
	notifications, err := service.repo.GetAllNotifications()
	if err != nil {
		color.Red("Error retrieving all notifications: %v", err)
		return nil, err
	}
	return notifications, nil
}

//// MarkNotificationAsRead marks a notification as read (if you choose to implement it)
//func (svc *NotificationServiceImpl) MarkNotificationAsRead(notificationID int) error {
//	return svc.repo.MarkNotificationAsRead(notificationID)
//}
