package controllers

import (
	"doctor-patient-cli/models"
	"doctor-patient-cli/utils"
	"fmt"
)

func PatientMenu(user models.User) {
	for {
		fmt.Println("\n========Patient Functionality=========")
		fmt.Println("1. View Profile")
		fmt.Println("2. Check Notifications")
		fmt.Println("3. View All Doctors")
		fmt.Println("4. Send Message to Doctor")
		fmt.Println("5. Add Review")
		fmt.Println("6. Update Profile")
		fmt.Println("7. Logout")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			patient, err := models.GetPatientByID(user.UserID)
			if err != nil {
				fmt.Println("Error fetching profile:", err)
				continue
			}
			fmt.Printf("Profile: %v\n", patient)

		case 2:
			notifications, err := models.GetNotificationsByUserID(user.UserID)
			fmt.Printf("Notifications: %v\n", notifications)
			if err != nil {
				fmt.Println("Error fetching notifications:", err)
				continue
			}
			for _, notification := range notifications {
				fmt.Printf("Notification: %s, Timestamp: %s\n", notification.Content, notification.Timestamp)
			}

		case 3:
			doctors, err := models.GetAllDoctors()
			if err != nil {
				fmt.Println("Error fetching doctors:", err)
				continue
			}
			for _, doctor := range doctors {
				fmt.Printf("Doctor ID: %s, Specialization: %s, Experience: %d, Rating: %.2f\n", doctor.UserID, doctor.Specialization, doctor.Experience, doctor.Rating)
			}

		case 4:
			fmt.Println("Enter Doctor User ID to send a message:")
			var doctorID string
			fmt.Scanln(&doctorID)

			fmt.Println("Enter your message:")
			var message string
			fmt.Scanln(&message)

			err := models.SendMessageToDoctor(user.UserID, doctorID, message)
			if err != nil {
				fmt.Println("Error sending message:", err)
			} else {
				fmt.Println("Message sent to doctor.")
			}

		case 5:
			fmt.Println("Enter Doctor User ID to add a review:")
			var doctorID string
			fmt.Scanln(&doctorID)

			fmt.Println("Enter your review:")
			var review string
			fmt.Scanln(&review)

			fmt.Println("Enter your rating (1-5):")
			var rating int
			fmt.Scanln(&rating)

			err := models.AddReview(user.UserID, doctorID, review, rating)
			if err != nil {
				fmt.Println("Error adding review:", err)
			} else {
				fmt.Println("Review added.")
			}

		case 6:
			fmt.Println("\nUpdate your profile:")
			fmt.Println("1. Update Username")
			fmt.Println("2. Update Age")
			fmt.Println("3. Update Gender")
			fmt.Println("4. Update Email")
			fmt.Println("5. Update Phone Number")
			fmt.Println("6. Update Password")

			var updateChoice int
			fmt.Scanln(&updateChoice)

			switch updateChoice {
			case 1:
				fmt.Print("Enter new username: ")
				var newUsername string
				fmt.Scanln(&newUsername)
				err := models.UpdateUsername(user.UserID, newUsername)
				if err != nil {
					fmt.Println("Error updating username:", err)
				} else {
					fmt.Println("Username updated.")
				}
			case 2:
				fmt.Print("Enter new age: ")
				var newAge int
				fmt.Scanln(&newAge)
				err := models.UpdateAge(user.UserID, newAge)
				if err != nil {
					fmt.Println("Error updating age:", err)
				} else {
					fmt.Println("Age updated.")
				}
			case 3:
				fmt.Print("Enter new gender: ")
				var newGender string
				fmt.Scanln(&newGender)
				err := models.UpdateGender(user.UserID, newGender)
				if err != nil {
					fmt.Println("Error updating gender:", err)
				} else {
					fmt.Println("Gender updated.")
				}
			case 4:
				fmt.Print("Enter new email: ")
				var newEmail string
				fmt.Scanln(&newEmail)
				err := models.UpdateEmail(user.UserID, newEmail)
				if err != nil {
					fmt.Println("Error updating email:", err)
				} else {
					fmt.Println("Email updated.")
				}
			case 5:
				fmt.Print("Enter new phone number: ")
				var newPhoneNumber string
				fmt.Scanln(&newPhoneNumber)
				err := models.UpdatePhoneNumber(user.UserID, newPhoneNumber)
				if err != nil {
					fmt.Println("Error updating phone number:", err)
				} else {
					fmt.Println("Phone number updated.")
				}
			case 6:
				fmt.Print("Enter new password: ")
				var newPassword string
				fmt.Scanln(&newPassword)
				err := models.UpdatePassword(user.UserID, utils.HashPassword(newPassword))
				if err != nil {
					fmt.Println("Error updating password:", err)
				} else {
					fmt.Println("Password updated.")
				}
			default:
				fmt.Println("Invalid choice. Please try again.")
			}

		case 7:
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
