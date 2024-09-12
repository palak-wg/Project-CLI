package controllers

import (
	"bufio"
	"doctor-patient-cli/models"
	"doctor-patient-cli/utils"
	"fmt"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
	"strings"
)

func DoctorMenu(user models.User) {

	_, err := userService.GetUserByID(user.UserID)
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

			doc, _ := doctorService.GetDoctorByID(user.UserID)

			color.Cyan("\n================== PROFILE ==================")
			profileTable := tablewriter.NewWriter(os.Stdout)
			profileTable.SetHeader([]string{"Field", "Value"})

			profileTable.Append([]string{"User ID", user.UserID})
			profileTable.Append([]string{"First Name", user.Name})
			profileTable.Append([]string{"Age", fmt.Sprintf("%v years", user.Age)})
			profileTable.Append([]string{"Gender", user.Gender})
			profileTable.Append([]string{"Email", user.Email})
			profileTable.Append([]string{"Phone Number", user.PhoneNumber})
			profileTable.Append([]string{"Experience", fmt.Sprintf("%v years", doc.Experience)})
			profileTable.Append([]string{"Specialization", doc.Specialization})
			profileTable.Append([]string{"Rating", fmt.Sprintf("%.2f ⭐", doc.Rating)})
			profileTable.Render()

		case 2:
			notifications, err := notificationService.GetNotificationsByUserID(user.UserID)
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

		case 3:
			color.Cyan("\nResponding to patient request:")
			color.Magenta("Enter Patient User ID to respond:")
			var patientID string
			fmt.Scanln(&patientID)

			color.Magenta("Enter your response (end with an empty line):")
			reader := bufio.NewReader(os.Stdin)
			response, _ := reader.ReadString('\n') // Reads until newline character

			// Using services to respond to the patient request
			err := messageService.RespondToPatient(user.UserID, patientID, response)
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
			reader := bufio.NewReader(os.Stdin)
			prescription, _ := reader.ReadString('\n') // Reads until newline character

			// Trim any leading or trailing whitespace (including newlines) from the prescription string
			prescription = strings.TrimSpace(prescription)

			// Call the service to suggest the prescription
			err := messageService.SendMessage(user.UserID, patientID, prescription)
			if err != nil {
				color.Red("🚨 Error suggesting prescription: %v", err)
			} else {
				color.Green("✅ Prescription sent to patient.")
			}

		case 5:
			color.Blue("🔍 Checking pending appointment request...")
			pendingAppointments, err := appointmentService.GetPendingAppointmentsByDoctorID(user.UserID)
			if err != nil {
				color.Red("Error fetching pending requests: %v", err)
				continue
			}
			pendingAppointmentTable := tablewriter.NewWriter(os.Stdout)
			pendingAppointmentTable.SetHeader([]string{"Appointment ID", "Patient ID"})

			for _, pendingAppointment := range pendingAppointments {
				pendingAppointmentTable.Append([]string{strconv.Itoa(pendingAppointment.AppointmentID), pendingAppointment.PatientID})
			}
			pendingAppointmentTable.Render()

			color.Magenta("Enter Appointment ID to approve:")
			var appointmentID int
			fmt.Scanln(&appointmentID)

			err = appointmentService.ApproveAppointment(appointmentID)
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
			case 7:
				color.Magenta("Enter new experience in years:")
				var experience int
				fmt.Scanln(&experience)

				err := doctorService.UpdateDoctorExperience(user.UserID, experience)
				if err != nil {
					color.Red("🚨 Error updating experience: %v", err)
				} else {
					color.Green("✅ Experience updated.")
				}

			case 8:
				color.Magenta("Enter new specialization:")
				var specialization string
				fmt.Scanln(&specialization)

				err := doctorService.UpdateDoctorSpecialization(user.UserID, specialization)
				if err != nil {
					color.Red("🚨 Error updating specialization: %v", err)
				} else {
					color.Green("✅ Specialization updated.")
				}

			default:
				color.Red("🚨 Invalid choice. Please try again.")
			}

		case 7:
			appointments, err := appointmentService.GetAppointmentsByDoctorID(user.UserID)
			if err != nil {
				color.Red("🚨 Error fetching appointments: %v", err)
				continue
			}
			color.Cyan("\n============ APPOINTMENTS ===============")

			// Create a new table writer
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"Appointment ID", "Patient ID", "Time", "Status"})

			// Add rows to the table
			for _, appointment := range appointments {
				table.Append([]string{
					fmt.Sprintf("%d", appointment.AppointmentID),
					appointment.PatientID,
					string(appointment.DateTime),
					fmt.Sprintf("%v", appointment.IsApproved),
				})
			}

			// Render the table
			table.Render()

		case 8:
			color.Cyan("\nCheck messages:")
			color.Magenta("1. All unread messages")
			color.Magenta("2. Specific patient")
			fmt.Print("Enter your choice: ")

			var choice int
			fmt.Scanln(&choice)

			switch choice {
			case 1:
				messages, err := messageService.GetUnreadMessages(user.UserID)
				if err != nil {
					color.Red("🚨 Error fetching messages: %v", err)
					continue
				}

				if len(messages) == 0 {
					color.Yellow("No unread messages.")
				} else {
					color.Cyan("\n=========== UNREAD MESSAGES ============")
					// Create a new table writer
					table := tablewriter.NewWriter(os.Stdout)
					table.SetHeader([]string{"From", "Message", "Timestamp"})

					// Add rows to the table
					for _, message := range messages {
						table.Append([]string{
							message.Sender,
							message.Content,
							string(message.Timestamp),
						})
					}

					// Render the table
					table.Render()
				}

			case 2:
				color.Magenta("Enter patient ID: ")
				var patientID string
				fmt.Scanln(&patientID)

				messages, err := messageService.GetUnreadMessagesById(patientID, user.UserID)
				if err != nil {
					color.Red("🚨 Error fetching messages: %v", err)
					continue
				}

				if len(messages) == 0 {
					color.Yellow("No unread messages from this patient.")
				} else {
					color.Cyan("\n=========== MESSAGES FROM PATIENT ============")
					// Create a new table writer
					table := tablewriter.NewWriter(os.Stdout)
					table.SetHeader([]string{"Message", "Timestamp"})

					// Add rows to the table
					for _, message := range messages {
						table.Append([]string{
							message.Content,
							string(message.Timestamp),
						})
					}

					// Render the table
					table.Render()
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
