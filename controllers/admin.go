package controllers

import (
	"doctor-patient-cli/models"
	"fmt"
)

func AdminMenu(user models.User) {
	for {
		fmt.Println("\n=======Admin Functionality=======")
		fmt.Println("1. Get All User IDs")
		fmt.Println("2. Approve Doctor Signup")
		fmt.Println("3. View All Reviews")
		fmt.Println("4. View All Notifications")
		fmt.Println("5. Delete User")
		fmt.Println("6. Logout")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			userIDs, err := models.GetAllUserIDs()
			if err != nil {
				fmt.Println("Error fetching user IDs:", err)
				continue
			}
			fmt.Println("User IDs:", userIDs)

		case 2:
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
			reviews, err := models.GetAllReviews()
			if err != nil {
				fmt.Println("Error fetching reviews:", err)
				continue
			}
			for _, review := range reviews {
				fmt.Printf("Review: %s, Rating: %d\n", review.Content, review.Rating)
			}

		case 4:
			notifications, err := models.GetAllNotifications()
			if err != nil {
				fmt.Println("Error fetching notifications:", err)
				continue
			}
			for _, notification := range notifications {
				fmt.Printf("Notification: %s, Timestamp: %s\n", notification.Content, notification.Timestamp)
			}

		case 5:
			deleteUser()

		case 6:
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func deleteUser() {}
