package main

import (
	"doctor-patient-cli/controllers"
	"doctor-patient-cli/utils"
	"fmt"
	"github.com/fatih/color"
)

func main() {
	go utils.InitDB()
	defer utils.CloseDB()
	StartApp()
}

// StartApp runs the application logic
func StartApp() {
	for {
		// Title and Welcome Message
		color.Cyan("===========================================")
		color.Cyan("\t🚀 Welcome To The MedCare 🚀")
		color.Cyan("===========================================")

		// Menu Options
		color.Magenta("\nPlease choose an option:")
		fmt.Println("1. Login")
		fmt.Println("2. Signup")
		fmt.Println("3. Exit")
		fmt.Print("\nEnter your choice: ")

		// User input
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			color.Blue("🔑 Logging in...")
			user := controllers.Login()
			switch user.UserType {
			case "admin":
				color.Yellow("👨‍💼 Welcome, Admin!")
				controllers.AdminMenu()
			case "doctor":
				color.Yellow("👨‍⚕️ Welcome, Doctor!")
				controllers.DoctorMenu(user)
			case "patient":
				color.Yellow("🧑‍⚕️ Welcome, Patient!")
				controllers.PatientMenu(user)
			default:
				color.Red("🚨 Invalid user type")
			}
		case 2:
			color.Blue("📝 Signing up...")
			controllers.Signup()
		case 3:
			color.Green("👋 Exiting... Goodbye!")
			return
		default:
			color.Red("🚫 Invalid choice. Please try again.")
		}
	}
}
