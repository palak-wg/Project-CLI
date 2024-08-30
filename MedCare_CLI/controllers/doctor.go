package controllers

import (
	"doctor-patient-cli/models"
	"doctor-patient-cli/services"
	"doctor-patient-cli/utils"
	"fmt"
	"github.com/fatih/color"
)

func DoctorMenu(user models.User) {
	_, err := services.GetDoctorByID(user.UserID)
	if err != nil {
		color.Red("🚨 Error fetching doctor details: %v", err)
		return
	}

	if !user.IsApproved {
		color.Yellow("⚠️ Your account has not been approved by admin yet.")
		return
	}

	for {
		color.Cyan("\n===========================================")
		color.Cyan("\tDoctor Functionality")
		color.Cyan("===========================================")
		color.Magenta("1. View Profile")
		color.Magenta("2. Check Notifications")
		color.Magenta("3. Respond to Patient Message Request")
		color.Magenta("4. Suggest Prescription")
		color.Magenta("5. Approve Appointment")
		color.Magenta("6. Update Profile")
		color.Magenta("7. View All Appointments")
		color.Magenta("8. Check Unread Messages")
		color.Magenta("9. Logout")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			if err != nil {
				color.Red("🚨 Error fetching profile: %v", err)
				continue
			}
			color.Cyan("\n================== PROFILE ==================")
			services.ViewProfile(user)

		case 2:
			notifications, err := services.GetNotificationsByUserID(user.UserID)
			if err != nil {
				color.Red("🚨 Error fetching notifications: %v", err)
				continue
			}
			color.Cyan("\n============ NOTIFICATIONS ===============")
			for _, notification := range notifications {
				fmt.Printf("Notification: %s, Timestamp: %s\n", notification.Content, notification.Timestamp)
			}

		case 3:
			color.Cyan("\nResponding to patient request:")
			color.Magenta("Enter Patient User ID to respond:")
			var patientID string
			fmt.Scanln(&patientID)

			color.Magenta("Enter your response:")
			var response string
			fmt.Scanln(&response)

			err := services.RespondToPatientRequest(user.UserID, patientID, response)
			if err != nil {
				color.Red("🚨 Error responding to patient: %v", err)
			} else {
				color.Green("✅ Response sent to patient.")
			}

		case 4:
			color.Magenta("Enter Patient User ID to suggest prescription:")
			var patientID string
			fmt.Scanln(&patientID)

			color.Magenta("Enter prescription details:")
			var prescription string
			fmt.Scanln(&prescription)

			err := services.SuggestPrescription(user.UserID, patientID, prescription)
			if err != nil {
				color.Red("🚨 Error suggesting prescription: %v", err)
			} else {
				color.Green("✅ Prescription sent to patient.")
			}

		case 5:
			color.Magenta("Enter Appointment ID to approve:")
			var appointmentID string
			fmt.Scanln(&appointmentID)

			err := services.ApproveAppointment(appointmentID)
			if err != nil {
				color.Red("🚨 Error approving appointment: %v", err)
			} else {
				color.Green("✅ Appointment approved.")
			}

		case 6:
			color.Cyan("\nUpdate your profile:")
			color.Magenta("1. Update First Name")
			color.Magenta("2. Update Age")
			color.Magenta("3. Update Gender")
			color.Magenta("4. Update Email")
			color.Magenta("5. Update Phone Number")
			color.Magenta("6. Update Password")
			color.Magenta("7. Update Experience")
			color.Magenta("8. Update Specialization")
			fmt.Print("Enter your choice: ")

			var updateChoice int
			fmt.Scanln(&updateChoice)

			switch updateChoice {
			case 1:
				color.Magenta("Enter new first name: ")
				var newFirstname string
				fmt.Scanln(&newFirstname)
				err := services.UpdateUsername(user.UserID, newFirstname)
				if err != nil {
					color.Red("🚨 Error updating username: %v", err)
				} else {
					color.Green("✅ Username updated.")
				}
			case 2:
				color.Magenta("Enter new age: ")
				var newAge int
				fmt.Scanln(&newAge)
				err := services.UpdateAge(user.UserID, newAge)
				if err != nil {
					color.Red("🚨 Error updating age: %v", err)
				} else {
					color.Green("✅ Age updated.")
				}
			case 3:
				color.Magenta("Enter new gender: ")
				var newGender string
				fmt.Scanln(&newGender)
				err := services.UpdateGender(user.UserID, newGender)
				if err != nil {
					color.Red("🚨 Error updating gender: %v", err)
				} else {
					color.Green("✅ Gender updated.")
				}
			case 4:
				color.Magenta("Enter new email: ")
				var newEmail string
				fmt.Scanln(&newEmail)
				err := services.UpdateEmail(user.UserID, newEmail)
				if err != nil {
					color.Red("🚨 Error updating email: %v", err)
				} else {
					color.Green("✅ Email updated.")
				}
			case 5:
				color.Magenta("Enter new phone number: ")
				var newPhoneNumber string
				fmt.Scanln(&newPhoneNumber)
				err := services.UpdatePhoneNumber(user.UserID, newPhoneNumber)
				if err != nil {
					color.Red("🚨 Error updating phone number: %v", err)
				} else {
					color.Green("✅ Phone number updated.")
				}
			case 6:
				color.Magenta("Enter new password: ")
				var newPassword string
				fmt.Scanln(&newPassword)
				err := services.UpdatePassword(user.UserID, utils.HashPassword(newPassword))
				if err != nil {
					color.Red("🚨 Error updating password: %v", err)
				} else {
					color.Green("✅ Password updated.")
				}
			case 7:
				color.Magenta("Enter new experience in years:")
				var experience int
				fmt.Scanln(&experience)

				err := services.UpdateDoctorExperience(user.UserID, experience)
				if err != nil {
					color.Red("🚨 Error updating experience: %v", err)
				} else {
					color.Green("✅ Experience updated.")
				}

			case 8:
				color.Magenta("Enter new specialization:")
				var specialization string
				fmt.Scanln(&specialization)

				err := services.UpdateDoctorSpecialization(user.UserID, specialization)
				if err != nil {
					color.Red("🚨 Error updating specialization: %v", err)
				} else {
					color.Green("✅ Specialization updated.")
				}

			default:
				color.Red("🚨 Invalid choice. Please try again.")
			}

		case 7:
			appointments, err := services.GetAppointmentsByDoctorID(user.UserID)
			if err != nil {
				color.Red("🚨 Error fetching appointments: %v", err)
				continue
			}
			color.Cyan("\n============ APPOINTMENTS ===============")
			for _, appointment := range appointments {
				fmt.Printf("AppointmentID: %d, PatientID: %s, Time: %s, Status: %v\n",
					appointment.AppointmentID, appointment.PatientID, appointment.DateTime, appointment.IsApproved)
			}

		case 8:
			color.Cyan("\nCheck messages:")
			color.Magenta("1. All unread messages")
			color.Magenta("2. Specific patient")
			fmt.Print("Enter your choice: ")

			var choice int
			fmt.Scanln(&choice)

			switch choice {
			case 1:
				messages, err := services.GetUnreadMessage(user.UserID)
				if err != nil {
					color.Red("🚨 Error fetching messages: %v", err)
				}
				for _, message := range messages {
					fmt.Printf("From: %s, Message: %s, Timestamp: %s\n", message.Sender, message.Content, message.Timestamp)
				}
			case 2:
				color.Magenta("Enter patient ID: ")
				var ID string
				fmt.Scanln(&ID)
				messages, err := services.GetUnreadMessagesByUserID(ID, user.UserID)
				if err != nil {
					color.Red("🚨 Error fetching messages: %v", err)
				}
				for _, message := range messages {
					fmt.Printf("Message: %s, Timestamp: %s\n", message.Content, message.Timestamp)
				}
			default:
				color.Red("🚨 Invalid choice. Try again.")
			}

		case 9:
			color.Green("✅ Logging out. Goodbye!")
			return

		default:
			color.Red("🚨 Invalid choice. Please try again.")
		}
	}
}
