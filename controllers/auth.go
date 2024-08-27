package controllers

import (
	"doctor-patient-cli/models"
	"doctor-patient-cli/utils"
	"fmt"
)

func Signup() {
	fmt.Println("\n\n===============SIGNUP================")
	fmt.Println("\n==========Enter Your Details=========")
	user := models.User{}

	for {
		fmt.Print("Enter Role (doctor/patient): ")
		fmt.Scanln(&user.UserType)

		if !(utils.ValidateRole(user.UserType)) {
			fmt.Println("Invalid Role. It must be doctor or patient.")
			continue
		}
		break
	}

	for {
		fmt.Print("Enter UserID: ")
		fmt.Scanln(&user.UserID)
		if !utils.ValidateUserID(user.UserID) {
			fmt.Println("Invalid userID")
			continue
		}
		break
	}

	for {
		fmt.Print("Enter Password: ")
		fmt.Scanln(&user.Password)
		if !utils.ValidatePassword(user.Password) {
			fmt.Println("Password criteria doesn't match")
			continue
		}
		break
	}

	for {
		fmt.Print("Enter First Name: ")
		fmt.Scanln(&user.Username)
		if !utils.ValidateUsername(user.Username) {
			fmt.Println("Invalid username")
			continue
		}
		break
	}

	for {
		fmt.Print("Enter Age: ")
		fmt.Scanln(&user.Age)
		if !utils.ValidateAge(user.Age) {
			fmt.Println("Invalid Age")
			continue
		}
		break
	}

	for {
		fmt.Print("Enter Gender: ")
		fmt.Scanln(&user.Gender)
		if !utils.ValidateGender(user.Gender) {
			fmt.Println("Invalid gender")
			continue
		}
		break
	}

	for {
		fmt.Print("Enter Email: ")
		fmt.Scanln(&user.Email)
		if !utils.ValidateEmail(user.Email) {
			fmt.Println("Invalid email")
			continue
		}
		break
	}

	for {
		fmt.Print("Enter Phone Number (10 digits): ")
		fmt.Scanln(&user.PhoneNumber)
		if !utils.ValidatePhoneNumber(user.PhoneNumber) {
			fmt.Println("Invalid PhoneNumber")
			continue
		}
		break
	}

	user.Password = utils.HashPassword(user.Password)

	err := models.CreateUser(user)
	if err != nil {
		fmt.Println("Error creating user: ", err)
		return
	}

}

func Login() models.User {
	fmt.Println("\n\n===============LOGIN=================")
	fmt.Println("\n==========Enter Your Details=========")
	fmt.Print("Enter User ID: ")
	var userID string
	fmt.Scanln(&userID)

	fmt.Print("Enter Password: ")
	var password string
	fmt.Scanln(&password)

	user, err := models.GetUserByID(userID)
	if err != nil {
		fmt.Println("Login failed:", err)
		return models.User{}
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		fmt.Println("Invalid password.")
		return models.User{}
	}

	return user
}
