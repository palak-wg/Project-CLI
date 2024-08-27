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
		color.Cyan("\tWelcome To The MedCare")
		color.Cyan("===========================================")

		// Menu Options
		fmt.Println("\nPlease choose an option")
		fmt.Println("1. Login\n2. Signup\n3. Exit")
		fmt.Print("Enter your choice: ")

		// User input
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			user := controllers.Login()
			switch user.UserType {
			case "admin":
				controllers.AdminMenu()
			case "doctor":
				controllers.DoctorMenu(user)
			case "patient":
				controllers.PatientMenu(user)
			default:
				color.Red("Invalid user type")
			}
		case 2:
			controllers.Signup()
		case 3:
			color.Green("Exiting... Goodbye!")
			return
		default:
			color.Red("Invalid choice. Please try again.")
		}
	}
}
