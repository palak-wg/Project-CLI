package services

import (
	"doctor-patient-cli/models"
	"doctor-patient-cli/utils"
)

func GetNotificationsByUserID(userID string) ([]models.Notification, error) {
	db := utils.GetDB()
	rows, err := db.Query("SELECT user_id, content, timestamp FROM notifications WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifications []models.Notification
	for rows.Next() {
		var notification models.Notification
		_ = rows.Scan(&notification.UserID, &notification.Content, &notification.Timestamp)
		notifications = append(notifications, notification)
	}
	return notifications, nil
}

func GetAllNotifications() ([]models.Notification, error) {
	db := utils.GetDB()
	rows, err := db.Query("SELECT user_id, content, timestamp FROM notifications")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifications []models.Notification
	for rows.Next() {
		var notification models.Notification
		_ = rows.Scan(&notification.UserID, &notification.Content, &notification.Timestamp)
		notifications = append(notifications, notification)
	}
	return notifications, nil
}
