package controllers

import (
	"doctor-patient-cli/models"
	"doctor-patient-cli/services"
	"doctor-patient-cli/utils"
	"fmt"
	"github.com/fatih/color"
	"golang.org/x/crypto/ssh/terminal"
	"os"
)

func Signup() {
	color.Cyan("\n========== Enter Your Details ==========")
	user := models.User{}

	for {
		color.Magenta("Enter Role (doctor/patient): ")
		fmt.Scanln(&user.UserType)

		if !(utils.ValidateRole(user.UserType)) {
			color.Red("🚨 Invalid Role. It must be doctor or patient.")
			continue
		}
		break
	}

	for {
		color.Magenta("Enter UserID: ")
		fmt.Scanln(&user.UserID)
		if !utils.ValidateUserID(user.UserID) {
			color.Red("🚨 Invalid UserID")
			continue
		}
		break
	}

	for {
		color.Magenta("Enter Password: ")
		passwordBytes, _ := terminal.ReadPassword(int(os.Stdin.Fd()))
		user.Password = string(passwordBytes)
		if !utils.ValidatePassword(user.Password) {
			color.Red("🚨 Password criteria doesn't match")
			continue
		}
		break
	}

	for {
		color.Magenta("Enter First Name: ")
		fmt.Scanln(&user.Username)
		if !utils.ValidateUsername(user.Username) {
			color.Red("🚨 Invalid Username")
			continue
		}
		break
	}

	for {
		color.Magenta("Enter Age: ")
		fmt.Scanln(&user.Age)
		if !utils.ValidateAge(user.Age) {
			color.Red("🚨 Invalid Age")
			continue
		}
		break
	}

	for {
		color.Magenta("Enter Gender: ")
		fmt.Scanln(&user.Gender)
		if !utils.ValidateGender(user.Gender) {
			color.Red("🚨 Invalid Gender")
			continue
		}
		break
	}

	for {
		color.Magenta("Enter Email: ")
		fmt.Scanln(&user.Email)
		if !utils.ValidateEmail(user.Email) {
			color.Red("🚨 Invalid Email")
			continue
		}
		break
	}

	for {
		color.Magenta("Enter Phone Number (10 digits): ")
		fmt.Scanln(&user.PhoneNumber)
		if !utils.ValidatePhoneNumber(user.PhoneNumber) {
			color.Red("🚨 Invalid Phone Number")
			continue
		}
		break
	}

	user.Password = utils.HashPassword(user.Password)

	err := services.CreateUser(user)
	if err != nil {
		color.Red("🚨 Error creating user: %v", err)
		return
	}

	color.Green("✅ User created successfully!")
}

func Login() models.User {
	color.Cyan("\n========== Enter Your Details ==========")
	color.Magenta("Enter User ID: ")
	var userID string
	fmt.Scanln(&userID)

	color.Magenta("Enter Password: ")
	passwordBytes, _ := terminal.ReadPassword(int(os.Stdin.Fd()))
	password := string(passwordBytes)

	user, err := services.GetUserByID(userID)
	if err != nil {
		color.Red("🚨 Login failed: %v", err)
		return models.User{}
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		color.Red("🚨 Invalid password.")
		return models.User{}
	}

	color.Green("✅ Login successful!")
	return user
}
