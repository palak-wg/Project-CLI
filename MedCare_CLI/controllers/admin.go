package controllers

import (
	"doctor-patient-cli/services"
	"fmt"
	"github.com/fatih/color"
)

func AdminMenu() {
	for {
		color.Cyan("\n===========================================")
		color.Cyan("\tAdmin Functionality")
		color.Cyan("===========================================")
		color.Magenta("1. Check Notifications")
		color.Magenta("2. Approve Doctor Signup")
		color.Magenta("3. Get Specific User Profile")
		color.Magenta("4. Get All User IDs")
		color.Magenta("5. View All Reviews")
		color.Magenta("6. View All Notifications")
		color.Magenta("7. Logout")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			color.Blue("📬 Fetching notifications...")
			notifications, err := services.GetNotificationsByUserID("admin")
			if err != nil {
				color.Red("🚨 Error fetching notifications: %v", err)
				continue
			}
			for _, notification := range notifications {
				fmt.Printf("Notification: %s, Timestamp: %s\n", notification.Content, notification.Timestamp)
			}

		case 2:
			color.Blue("🔍 Checking pending doctor signups...")
			services.PendingDoctorSignupRequest()
			fmt.Print("Enter Doctor UserID to approve: ")
			var userID string
			fmt.Scanln(&userID)
			err := services.ApproveDoctorSignup(userID)
			if err != nil {
				color.Red("🚨 Error approving doctor signup: %v", err)
				continue
			}
			color.Green("Doctor signup approved and notification sent.")

		case 3:
			color.Blue("🔍 Fetching user profile...")
			var UserID string
			fmt.Print("Enter userID: ")
			fmt.Scanln(&UserID)
			user, err := services.GetUserByID(UserID)
			if err != nil {
				color.Red("🚨 No such user exists")
				continue
			}
			color.Cyan("\n================== PROFILE ==================")
			services.ViewProfile(user)

		case 4:
			color.Blue("📋 Fetching all user IDs...")
			userIDs, err := services.GetAllUserIDs()
			if err != nil {
				color.Red("🚨 Error fetching user IDs: %v", err)
				continue
			}
			fmt.Println("User IDs:", userIDs)

		case 5:
			color.Blue("📜 Fetching all reviews...")
			reviews, err := services.GetAllReviews()
			if err != nil {
				color.Red("🚨 Error fetching reviews: %v", err)
				continue
			}
			for _, review := range reviews {
				fmt.Printf("Review: %s, Rating: %d, Patient: %s, Doctor: %s\n", review.Content,
					review.Rating, review.PatientID, review.DoctorID)
			}

		case 6:
			color.Blue("📬 Fetching all notifications...")
			notifications, err := services.GetAllNotifications()
			if err != nil {
				color.Red("🚨 Error fetching notifications: %v", err)
				continue
			}
			for _, notification := range notifications {
				fmt.Printf("Notification: %s, Timestamp: %s\n", notification.Content, notification.Timestamp)
			}

		case 7:
			color.Green("👋 Logging out...")
			return

		default:
			color.Red("🚫 Invalid choice. Please try again.")
		}
	}
}
