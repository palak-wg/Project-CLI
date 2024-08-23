package controllers

import (
	"doctor-patient-cli/models"
	"fmt"
)

func DoctorMenu(user models.User) {
	doctor, err := models.GetDoctorByID(user.UserID)
	if err != nil {
		fmt.Println("Error fetching doctor details:", err)
		return
	}

	if !user.IsApproved {
		fmt.Println("Your account has not been approved by admin yet.")
		return
	}

	for {
		fmt.Println("\n=======Doctor Functionality=======")
		fmt.Println("1. View Profile")
		fmt.Println("2. Check Notifications")
		fmt.Println("3. Respond to Patient Request")
		fmt.Println("4. Suggest Prescription")
		fmt.Println("5. Approve Appointment")
		fmt.Println("6. Update Experience")
		fmt.Println("7. Update Specialization")
		fmt.Println("8. View All Appointments")
		fmt.Println("9. Logout")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			if err != nil {
				fmt.Println("Error fetching profile:", err)
				continue
			}
			fmt.Printf("Profile: %v\n", doctor)

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
				fmt.Println("Response sent to patient.")
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
			fmt.Println("Enter new experience in years:")
			var experience int
			fmt.Scanln(&experience)

			err := models.UpdateDoctorExperience(user.UserID, experience)
			if err != nil {
				fmt.Println("Error updating experience:", err)
			} else {
				fmt.Println("Experience updated.")
			}

		case 7:
			fmt.Println("Enter new specialization:")
			var specialization string
			fmt.Scanln(&specialization)

			err := models.UpdateDoctorSpecialization(user.UserID, specialization)
			if err != nil {
				fmt.Println("Error updating specialization:", err)
			} else {
				fmt.Println("Specialization updated.")
			}

		case 8:
			appointments, err := models.GetAppointmentsByDoctorID(user.UserID)
			if err != nil {
				fmt.Println("Error fetching appointments:", err)
				continue
			}
			for _, appointment := range appointments {
				fmt.Printf("Appointment: %v\n", appointment)
			}

		case 9:
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
