package main

import (
	"doctor-patient-cli/controllers"
	"doctor-patient-cli/utils"
	"fmt"
)

func main() {
	utils.InitDB()
	defer utils.CloseDB()

	for {
		fmt.Println("=========Welcome To The Application========")
		fmt.Println("1. Login")
		fmt.Println("2. Signup")
		fmt.Println("3. Exit")
		fmt.Print("Enter your choice:")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			user := controllers.Login()
			switch user.UserType {
			case "admin":
				controllers.AdminMenu(user)
			case "doctor":
				controllers.DoctorMenu(user)
			case "patient":
				controllers.PatientMenu(user)
			default:
				fmt.Println("Invalid user type")
			}
		case 2:
			controllers.Signup()
		case 3:
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
