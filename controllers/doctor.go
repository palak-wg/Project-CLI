package controllers

import (
	"doctor-patient-cli/models"
	"doctor-patient-cli/utils"
	"fmt"
	"github.com/fatih/color"
)

func DoctorMenu(user models.User) {
	_, err := models.GetDoctorByID(user.UserID)
	if err != nil {
		fmt.Println("Error fetching doctor details:", err)
		return
	}

	if !user.IsApproved {
		fmt.Println("Your account has not been approved by admin yet.")
		return
	}

	for {
		fmt.Println("\n===========================================")
		fmt.Println("\t\tDoctor Functionality")
		fmt.Println("===========================================")
		fmt.Println("1. View Profile")
		fmt.Println("2. Check Notifications")
		fmt.Println("3. Respond to Patient Message Request")
		fmt.Println("4. Suggest Prescription")
		fmt.Println("5. Approve Appointment")
		fmt.Println("6. Update Profile")
		fmt.Println("7. View All Appointments")
		fmt.Println("8. Check unread messages")
		fmt.Println("9. Logout")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			if err != nil {
				fmt.Println("Error fetching profile:", err)
				continue
			}
			models.ViewProfile(user)

		case 2:
			notifications, err := models.GetNotificationsByUserID(user.UserID)
			if err != nil {
				fmt.Println("Error fetching notifications:", err)
				continue
			}
			for _, notification := range notifications {
				fmt.Printf("Notification: %s, Timestamp: %s\n", notification.Content, notification.Timestamp)
			}

		case 3:
			fmt.Println("Enter Patient User ID to respond:")
			var patientID string
			fmt.Scanln(&patientID)

			fmt.Println("Enter your response:")
			var response string
			fmt.Scanln(&response)

			err := models.RespondToPatientRequest(user.UserID, patientID, response)
			if err != nil {
				fmt.Println("Error responding to patient:", err)
			} else {
				color.Green("Response sent to patient.")
			}

		case 4:
			fmt.Println("Enter Patient User ID to suggest prescription:")
			var patientID string
			fmt.Scanln(&patientID)

			fmt.Println("Enter prescription details:")
			var prescription string
			fmt.Scanln(&prescription)

			err := models.SuggestPrescription(user.UserID, patientID, prescription)
			if err != nil {
				fmt.Println("Error suggesting prescription:", err)
			} else {
				fmt.Println("Prescription sent to patient.")
			}

		case 5:
			fmt.Println("Enter Appointment ID to approve:")
			var appointmentID string
			fmt.Scanln(&appointmentID)

			err := models.ApproveAppointment(appointmentID)
			if err != nil {
				fmt.Println("Error approving appointment:", err)
			} else {
				fmt.Println("Appointment approved.")
			}

		case 6:

			fmt.Println("\nUpdate your profile:")
			fmt.Println("1. Update First Name")
			fmt.Println("2. Update Age")
			fmt.Println("3. Update Gender")
			fmt.Println("4. Update Email")
			fmt.Println("5. Update Phone Number")
			fmt.Println("6. Update Password")
			fmt.Println("7. Update Experience")
			fmt.Println("8. Update Specialization")
			fmt.Print("Enter your choice: ")

			var updateChoice int
			fmt.Scanln(&updateChoice)

			switch updateChoice {
			case 1:
				fmt.Print("Enter new firstname: ")
				var newFirstname string
				fmt.Scanln(&newFirstname)
				err := models.UpdateUsername(user.UserID, newFirstname)
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
			case 7:
				fmt.Println("Enter new experience in years:")
				var experience int
				fmt.Scanln(&experience)

				err := models.UpdateDoctorExperience(user.UserID, experience)
				if err != nil {
					fmt.Println("Error updating experience:", err)
				} else {
					fmt.Println("Experience updated.")
				}

			case 8:
				fmt.Println("Enter new specialization:")
				var specialization string
				fmt.Scanln(&specialization)

				err := models.UpdateDoctorSpecialization(user.UserID, specialization)
				if err != nil {
					fmt.Println("Error updating specialization:", err)
				} else {
					fmt.Println("Specialization updated.")
				}

			default:
				fmt.Println("Invalid choice. Please try again.")
			}

		case 7:
			appointments, err := models.GetAppointmentsByDoctorID(user.UserID)
			if err != nil {
				fmt.Println("Error fetching appointments:", err)
				continue
			}
			fmt.Println("\n============ APPOINTMENTS ===============")
			for _, appointment := range appointments {
				fmt.Printf("AppointmentID: %s, PatientID: %s, Time: %s, Status: %v\n",
					appointment.AppointmentID, appointment.PatientID, appointment.DateTime, appointment.IsApproved)
			}

		case 8:
			fmt.Println("\nCheck messages:")
			fmt.Println("1. All unread messages")
			fmt.Println("2. Specific patient")
			fmt.Print("Enter your choice: ")

			var choice int
			fmt.Scanln(&choice)

			switch choice {
			case 1:
				messages, err := models.GetUnreadMessage(user.UserID)
				if err != nil {
					fmt.Println("Error fetching messages:", err)
				}
				for _, message := range messages {
					fmt.Printf("From: %s, Message: %s, Timestamp: %s\n", message.Sender, message.Content, message.Timestamp)
				}
			case 2:
				fmt.Println("Enter patient ID: ")
				var ID string
				fmt.Scanln(&ID)
				messages, err := models.GetUnreadMessagesByUserID(ID, user.UserID)
				if err != nil {
					fmt.Println("Error fetching messages:", err)
				}
				for _, message := range messages {
					fmt.Printf("Message: %s, Timestamp: %s\n", message.Content, message.Timestamp)
				}
			default:
				color.Red("Invalid Choice. Try again.")
			}

		case 9:
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
