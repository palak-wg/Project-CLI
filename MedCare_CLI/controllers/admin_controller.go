package controllers

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"os"
	"strings"
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
			notifications, err := notificationService.GetNotificationsByUserID("admin")
			if err != nil {
				color.Red("🚨 Error fetching notifications: %v", err)
				continue
			}
			if len(notifications) == 0 {
				color.Yellow("No notifications available.")
			} else {
				color.Cyan("\n============ NOTIFICATIONS ===============")
				notificationTable := tablewriter.NewWriter(os.Stdout)
				notificationTable.SetHeader([]string{"Content", "Timestamp"})

				for _, notification := range notifications {
					notificationTable.Append([]string{notification.Content, string(notification.Timestamp)})
				}

				notificationTable.Render()
			}

		case 2:
			color.Blue("🔍 Checking pending doctor signups...")
			pendingDoctors, err := adminService.GetPendingDoctorRequests()
			if err != nil {
				color.Red("Error fetching pending requests: %v", err)
				continue
			}
			pendingDoctorTable := tablewriter.NewWriter(os.Stdout)
			pendingDoctorTable.SetHeader([]string{"User ID", "Doctor Name"})

			for _, pendingDoc := range pendingDoctors {
				pendingDoctorTable.Append([]string{pendingDoc.UserID, pendingDoc.Name})
			}

			pendingDoctorTable.Render()

			fmt.Print("Enter Doctor UserID to approve: ")
			var userID string
			fmt.Scanln(&userID)
			err = adminService.ApproveDoctorSignup(userID)
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
			user, err := userService.GetUserByID(UserID)
			if err != nil {
				color.Red("🚨 No such user exists")
				continue
			}
			color.Cyan("\n================== USER PROFILE ==================")

			// Print headers for the table
			fmt.Printf("\n%-20s | %-20s\n", "Field", "Details")
			fmt.Printf("%s\n", strings.Repeat("-", 43))

			// Print user details in a table-like format
			fmt.Printf("%-20s | %-20s\n", "User ID", user.UserID)
			fmt.Printf("%-20s | %-20s\n", "Name", user.Name)
			fmt.Printf("%-20s | %-20s\n", "Email", user.Email)
			fmt.Printf("%-20s | %-20d\n", "Age", user.Age)
			fmt.Printf("%-20s | %-20s\n", "Gender", user.Gender)
			fmt.Printf("%-20s | %-20s\n", "Phone Num", user.PhoneNumber)

			color.Cyan("\n===================================================")

		case 4:
			color.Blue("📋 Fetching all user IDs...")
			userIDs, err := adminService.GetAllUsers()
			if err != nil {
				color.Red("🚨 Error fetching user IDs: %v", err)
				continue
			}
			// Check if users list is empty
			if len(userIDs) == 0 {
				color.Yellow("No users found.")
			} else {
				// Print the user IDs in a list format
				color.Cyan("\n================== USER IDs ==================")
				userTable := tablewriter.NewWriter(os.Stdout)
				userTable.SetHeader([]string{"User ID", "Name", "Role"})
				for _, user := range userIDs {
					userTable.Append([]string{user.UserID, user.Name, user.UserType})
				}
				color.Cyan("=============================================")
				userTable.Render()
			}

		case 5:
			color.Blue("📜 Fetching all reviews...")
			reviews, err := reviewService.GetAllReviews()
			if err != nil {
				color.Red("🚨 Error fetching reviews: %v", err)
				continue
			}
			// Check if reviews list is empty
			if len(reviews) == 0 {
				color.Yellow("No reviews found.")
			} else {
				// Display the reviews in a table-like format
				color.Cyan("\n================== REVIEWS ==================")
				fmt.Printf("%-20s %-20s %-50s %-10s\n", "Patient ID", "Doctor ID", "Review", "Rating")
				color.Cyan("------------------------------------------" +
					"-------------------------------------------------------------------------------------")
				for _, review := range reviews {
					fmt.Printf("%-20s %-20s %-50s %-10d\n", review.PatientID, review.DoctorID, review.Content, review.Rating)
				}
				color.Cyan("===========================================" +
					"=====================================================================================")
			}

		case 6:
			color.Blue("📬 Fetching all notifications...")
			notifications, err := notificationService.GetAllNotifications()
			if err != nil {
				color.Red("🚨 Error fetching notifications: %v", err)
				continue
			}

			// Check if notifications list is empty
			if len(notifications) == 0 {
				color.Yellow("No notifications found.")
			} else {
				// Display the notifications in a table-like format
				color.Cyan("\n================== NOTIFICATIONS ==================")
				fmt.Printf("%-100s %-30s\n", "Notification Content", "Timestamp")
				color.Cyan("----------------------------------------------------" +
					"-------------------------------------------------------------------------------")
				for _, notification := range notifications {
					fmt.Printf("%-100s %-30s\n", notification.Content, notification.Timestamp)
				}
				color.Cyan("=======================================================" +
					"===============================================================================")
			}

		case 7:
			color.Green("👋 Logging out...")
			return

		default:
			color.Red("🚫 Invalid choice. Please try again.")
		}
	}
}
