package utils

import (
	"regexp"
)

func ValidateEmail(email string) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`)
	return re.MatchString(email)
}

func ValidatePhoneNumber(phone string) bool {
	re := regexp.MustCompile(`^\+?\d{10,15}$`)
	return re.MatchString(phone)
}

func ValidateUsername(username string) bool {
	re := regexp.MustCompile(`^[a-zA-Z\s]{3,}$`)
	return re.MatchString(username)
}

func ValidatePassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasDigit := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>]`).MatchString(password)

	return hasUpper && hasLower && hasDigit && hasSpecial

}

func ValidateUserID(userID string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9]{3,16}$`).MatchString(userID)
}

func ValidateAge(age int) bool {
	if age < 0 {
		return false
	}
	return true
}

func ValidateGender(gender string) bool {
	if gender != "male" && gender != "female" && gender != "other" {
		return false
	}
	return true
}

func ValidateRole(role string) bool {
	if role != "patient" && role != "doctor" {
		return false
	}
	return true
}
