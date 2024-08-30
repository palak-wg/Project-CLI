package services

import (
	"doctor-patient-cli/models"
	"doctor-patient-cli/utils"
	"fmt"
)

func CreateUser(user models.User) error {
	db := utils.GetDB()

	_, err := db.Exec("INSERT INTO users (user_id, password, username, age, gender, email, phone_number, user_type, is_approved) VALUES (?, ?, ?, ?, ?, ?, ?, ?,?)",
		user.UserID, user.Password, user.Username, user.Age, user.Gender, user.Email, user.PhoneNumber, user.UserType, 0)

	if user.UserType == "doctor" {
		fmt.Println("Your signup request has been submitted for approval.")

		_, err = db.Exec("INSERT INTO notifications (user_id, content) VALUES (?, ?)",
			"admin", fmt.Sprintf("Please approve %s signup request for doctor role.", user.UserID))

		if err != nil {
			fmt.Println("Error requesting doctor signup:", err)
		}
	} else {
		_, _ = db.Exec("INSERT INTO patients (user_id, medical_history) VALUES (?,?)", user.UserID, "No History")
		fmt.Println("pat saved in patTab")
		fmt.Println("Signup successful. You can now log in.")
		_, _ = db.Exec("INSERT INTO notifications (user_id, content) VALUES (?, ?)",
			user.UserID, fmt.Sprintf("welcome %s to the application.", user.UserID))
	}

	return err
}

func GetUserByID(userID string) (models.User, error) {
	db := utils.GetDB()
	user := models.User{}
	err := db.QueryRow("SELECT user_id, password, username, age, gender, email, phone_number, user_type, is_approved FROM users WHERE user_id = ?", userID).
		Scan(&user.UserID, &user.Password, &user.Username, &user.Age, &user.Gender, &user.Email, &user.PhoneNumber, &user.UserType, &user.IsApproved)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func GetAllUserIDs() ([]string, error) {
	db := utils.GetDB()
	rows, err := db.Query("SELECT user_id FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userIDs []string
	for rows.Next() {
		var userID string
		_ = rows.Scan(&userID)
		userIDs = append(userIDs, userID)
	}
	return userIDs, nil
}

func UpdateUsername(userID, username string) error {
	db := utils.GetDB()
	_, err := db.Exec("UPDATE users SET username = ? WHERE user_id = ?", username, userID)
	return err
}

func UpdateAge(userID string, age int) error {
	db := utils.GetDB()
	_, err := db.Exec("UPDATE users SET age = ? WHERE user_id = ?", age, userID)
	return err
}

func UpdateGender(userID, gender string) error {
	db := utils.GetDB()
	_, err := db.Exec("UPDATE users SET gender = ? WHERE user_id = ?", gender, userID)
	return err
}

func UpdateEmail(userID, email string) error {
	db := utils.GetDB()
	_, err := db.Exec("UPDATE users SET email = ? WHERE user_id = ?", email, userID)
	return err
}

func UpdatePhoneNumber(userID, phoneNumber string) error {
	db := utils.GetDB()
	_, err := db.Exec("UPDATE users SET phone_number = ? WHERE user_id = ?", phoneNumber, userID)
	return err
}

func UpdatePassword(userID, password string) error {
	db := utils.GetDB()
	_, err := db.Exec("UPDATE users SET password = ? WHERE user_id = ?", password, userID)
	return err
}

func ViewProfile(user models.User) {
	db := utils.GetDB()
	db.QueryRow("SELECT user_id, username, age, gender,email, phone_number, user_type  FROM users WHERE user_id = ?", user.UserID).
		Scan(&user.UserID, &user.Username, &user.Age, &user.Gender, &user.Email, &user.PhoneNumber, &user.UserType)

	fmt.Printf("User ID: %v\nFirst Name: %v\nAge: %v\nGender: %v\nEmail: %v\nPhoneNumber: %v\n",
		user.UserID, user.Username, user.Age, user.Gender, user.Email, user.PhoneNumber)

	if user.UserType == "doctor" && user.IsApproved == true || user.UserType == "admin" {
		ViewDoctorSpecificProfile(user.UserID)
	} else if user.UserType == "patient" || user.UserType == "admin" {
		ViewPatientDetails(user.UserID)
	}
}
