package models

import (
	"doctor-patient-cli/utils"
)

func GetNotificationsByUserID(userID string) ([]Notification, error) {
	db := utils.GetDB()
	rows, err := db.Query("SELECT user_id, content, timestamp FROM notifications WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifications []Notification
	for rows.Next() {
		var notification Notification
		err = rows.Scan(&notification.UserID, &notification.Content, &notification.Timestamp)
		if err != nil {
			return nil, err
		}
		notifications = append(notifications, notification)
	}
	return notifications, nil
}

func GetAllNotifications() ([]Notification, error) {
	db := utils.GetDB()
	rows, err := db.Query("SELECT user_id, content, timestamp FROM notifications")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifications []Notification
	for rows.Next() {
		var notification Notification
		err = rows.Scan(&notification.UserID, &notification.Content, &notification.Timestamp)
		if err != nil {
			return nil, err
		}
		notifications = append(notifications, notification)
	}
	return notifications, nil
}
