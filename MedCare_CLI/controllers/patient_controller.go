package controllers

import (
	"doctor-patient-cli/models"
	"doctor-patient-cli/utils"
	"fmt"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
)

func PatientMenu(user models.User) {
	for {
		color.Cyan("\n===========================================")
		color.Cyan("\tYour Dashboard 🌟")
		color.Cyan("===========================================")
		color.Magenta("1. View Profile 🧑‍⚕️")
		color.Magenta("2. Check Notifications 🔔")
		color.Magenta("3. View All Doctors 🩺")
		color.Magenta("4. Send Message to Doctor 💬")
		color.Magenta("5. Send Appointment Request 📅")
		color.Magenta("6. Add Review ⭐")
		color.Magenta("7. Update Profile ✏️")
		color.Magenta("8. Logout 🚪")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			pat, err := patientService.GetPatientByID(user.UserID)
			if err != nil {
				color.Red("🚨 Error fetching profile: %v", err)
				continue
			}

			if err != nil {
				color.Red("🚨 Error fetching profile: %v", err)
				continue
			}

			color.Cyan("\n================== PROFILE ==================")
			profileTable := tablewriter.NewWriter(os.Stdout)
			profileTable.SetHeader([]string{"Field", "Value"})

			profileTable.Append([]string{"User ID", pat.UserID})
			profileTable.Append([]string{"First Name", pat.Name})
			profileTable.Append([]string{"Age", fmt.Sprintf("%d", pat.Age)})
			profileTable.Append([]string{"Gender", pat.Gender})
			profileTable.Append([]string{"Email", pat.Email})
			profileTable.Append([]string{"Phone Number", pat.PhoneNumber})
			profileTable.Render()

		case 2:
			color.Cyan("\n============== NOTIFICATIONS ================")
			notifications, err := notificationService.GetNotificationsByUserID(user.UserID)
			if err != nil {
				color.Red("🚨 Error fetching notifications: %v", err)
				continue
			}
			if len(notifications) == 0 {
				color.Yellow("No notifications available.")
			} else {
				// Create a table for notifications
				table := tablewriter.NewWriter(os.Stdout)
				table.SetHeader([]string{"Notification", "Timestamp"})
				for _, notification := range notifications {
					table.Append([]string{notification.Content, string(notification.Timestamp)})
				}
				table.Render()
			}

		case 3:
			doctors, err := doctorService.GetAllDoctors()
			if err != nil {
				color.Red("🚨 Error fetching doctors: %v", err)
				continue
			}
			color.Cyan("\n============== DOCTOR(S) ================")
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"Doctor ID", "Specialization", "Experience", "Rating"})
			for _, doctor := range doctors {
				table.Append([]string{doctor.UserID, doctor.Specialization, fmt.Sprintf("%d years", doctor.Experience), fmt.Sprintf("%.2f", doctor.Rating)})
			}
			table.Render()

		case 4:
			color.Magenta("Enter Doctor User ID to send a message:")
			var doctorID string
			fmt.Scanln(&doctorID)

			color.Magenta("Enter your message: ")
			message, _ := reader.ReadString('\n')

			err := messageService.SendMessage(user.UserID, doctorID, message)
			if err != nil {
				color.Red("🚨 Error sending message: %v", err)
			} else {
				color.Green("✅ Message sent to doctor.")
			}

		case 5:
			color.Magenta("Enter Doctor User ID to send appointment request: ")
			var doctorID string
			fmt.Scanln(&doctorID)

			err := appointmentService.SendAppointmentRequest(user.UserID, doctorID)
			if err != nil {
				color.Red("🚨 Error sending appointment request: %v", err)
			} else {
				color.Green("✅ Appointment request sent.")
			}

		case 6:
			color.Magenta("Enter Doctor User ID to add a review: ")
			var doctorID string
			fmt.Scanln(&doctorID)

			color.Magenta("Enter your review: ")
			review, _ := reader.ReadString('\n')

			color.Magenta("Enter your rating (1-5): ")
			var rating int
			fmt.Scanln(&rating)

			err := reviewService.AddReview(&models.Review{
				PatientID: user.UserID,
				DoctorID:  doctorID,
				Content:   review,
				Rating:    rating,
			})

			if err != nil {
				color.Red("🚨 Error adding review: %v", err)
			} else {
				color.Green("✅ Review added.")
			}

		case 7:
			color.Cyan("\nUpdate your profile:")
			color.Magenta("1. Update First Name ✏️")
			color.Magenta("2. Update Age 📅")
			color.Magenta("3. Update Gender 🚻")
			color.Magenta("4. Update Email 📧")
			color.Magenta("5. Update Phone Number 📞")
			color.Magenta("6. Update Password 🔑")
			fmt.Print("Enter your choice: ")

			var updateChoice int
			fmt.Scanln(&updateChoice)

			switch updateChoice {
			case 1:
				color.Magenta("Enter new firstname: ")
				var newFirstname string
				fmt.Scanln(&newFirstname)
				err := userService.UpdateName(user.UserID, newFirstname)
				if err != nil {
					color.Red("🚨 Error updating username: %v", err)
				} else {
					color.Green("✅ Username updated.")
				}
			case 2:
				color.Magenta("Enter new age: ")
				var newAge int
				fmt.Scanln(&newAge)

				age := strconv.Itoa(newAge)
				err := userService.UpdateAge(user.UserID, age)
				if err != nil {
					color.Red("🚨 Error updating age: %v", err)
				} else {
					color.Green("✅ Age updated.")
				}
			case 3:
				color.Magenta("Enter new gender: ")
				var newGender string
				fmt.Scanln(&newGender)
				err := userService.UpdateGender(user.UserID, newGender)
				if err != nil {
					color.Red("🚨 Error updating gender: %v", err)
				} else {
					color.Green("✅ Gender updated.")
				}
			case 4:
				color.Magenta("Enter new email: ")
				var newEmail string
				fmt.Scanln(&newEmail)
				err := userService.UpdateEmail(user.UserID, newEmail)
				if err != nil {
					color.Red("🚨 Error updating email: %v", err)
				} else {
					color.Green("✅ Email updated.")
				}
			case 5:
				color.Magenta("Enter new phone number: ")
				var newPhoneNumber string
				fmt.Scanln(&newPhoneNumber)
				err := userService.UpdatePhoneNumber(user.UserID, newPhoneNumber)
				if err != nil {
					color.Red("🚨 Error updating phone number: %v", err)
				} else {
					color.Green("✅ Phone number updated.")
				}
			case 6:
				color.Magenta("Enter new password: ")
				var newPassword string
				fmt.Scanln(&newPassword)
				err := userService.UpdatePassword(user.UserID, utils.HashPassword(newPassword))
				if err != nil {
					color.Red("🚨 Error updating password: %v", err)
				} else {
					color.Green("✅ Password updated.")
				}
			default:
				color.Red("🚨 Invalid choice. Please try again.")
			}

		case 8:
			color.Green("✅ Logging out. Goodbye!")
			return

		default:
			color.Red("🚨 Invalid choice. Please try again.")
		}
	}
}
