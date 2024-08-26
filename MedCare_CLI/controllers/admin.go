package controllers

import (
	"doctor-patient-cli/models"
	"fmt"
)

func AdminMenu() {
	for {
		fmt.Println("\n===========================================")
		fmt.Println("\t\tAdmin Functionality")
		fmt.Println("===========================================")
		fmt.Println("1. Check Notifications")
		fmt.Println("2. Approve Doctor Signup")
		fmt.Println("3. Get specific user profile")
		fmt.Println("4. Get All User IDs")
		fmt.Println("5. View All Reviews")
		fmt.Println("6. View All Notifications")
		fmt.Println("7. Delete User")
		fmt.Println("8. Logout")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			notifications, err := models.GetNotificationsByUserID("admin")
			if err != nil {
				fmt.Println("Error fetching notifications:", err)
				continue
			}
			for _, notification := range notifications {
				fmt.Printf("Notification: %s, Timestamp: %s\n", notification.Content, notification.Timestamp)
			}

		case 2:
			models.PendingDoctorSignupRequest()
			fmt.Print("Enter Doctor UserID to approve: ")
			var userID string
			fmt.Scanln(&userID)
			err := models.ApproveDoctorSignup(userID)
			if err != nil {
				fmt.Println("Error approving doctor signup:", err)
			} else {
				fmt.Println("Doctor signup approved and notification sent.")
			}

		case 3:
			user := models.User{}
			fmt.Print("Enter userID: ")
			fmt.Scanln(&user.UserID)
			models.ViewProfile(user)

		case 4:
			userIDs, err := models.GetAllUserIDs()
			if err != nil {
				fmt.Println("Error fetching user IDs:", err)
				continue
			}
			fmt.Println("User IDs:", userIDs)

		case 5:
			reviews, err := models.GetAllReviews()
			if err != nil {
				fmt.Println("Error fetching reviews:", err)
				continue
			}
			for _, review := range reviews {
				fmt.Printf("Review: %s, Rating: %d, Patient: %s, Doctor: %s\n", review.Content,
					review.Rating, review.PatientID, review.DoctorID)
			}

		case 6:
			notifications, err := models.GetAllNotifications()
			if err != nil {
				fmt.Println("Error fetching notifications:", err)
				continue
			}
			for _, notification := range notifications {
				fmt.Printf("Notification: %s, Timestamp: %s\n", notification.Content, notification.Timestamp)
			}

		case 7:
			var userID string
			fmt.Print("Enter the User ID to delete: ")
			fmt.Scanln(&userID)

			err := models.DeleteUser(userID)
			if err != nil {
				fmt.Printf("Failed to delete user: %v\n", err)
			} else {
				fmt.Println("User deleted successfully.")
			}

		case 8:
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
