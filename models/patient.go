package models

import (
	"doctor-patient-cli/utils"
	"fmt"
	"time"
)

func GetPatientByID(userID string) (Patient, error) {
	db := utils.GetDB()
	user := User{}
	db.QueryRow("SELECT user_id, username, age, gender,email, phone_number  FROM users WHERE user_id = ?", userID).
		Scan(&user.UserID, &user.Username, &user.Age, &user.Gender, &user.Email, &user.PhoneNumber)
	printProfile(user)
	patient := Patient{}
	err := db.QueryRow("SELECT user_id, medical_history FROM patients WHERE user_id = ?", userID).
		Scan(&patient.UserID, &patient.MedicalHistory)
	if err != nil {
		return Patient{}, err
	}
	return patient, nil
}

func SendMessageToDoctor(patientID, doctorID, message string) error {
	// Create a new message record
	messageID := fmt.Sprintf("msg-%d", time.Now().UnixNano())
	_, err := db.Exec("INSERT INTO messages (id, sender_id, receiver_id, message) VALUES (?, ?, ?, ?)",
		messageID, patientID, doctorID, message)
	if err != nil {
		return fmt.Errorf("error inserting message: %v", err)
	}

	// Create a notification for the doctor
	_, err = db.Exec("INSERT INTO notifications (user_id, content, timestamp) VALUES (?, ?, ?)",
		doctorID, fmt.Sprintf("You have a new message from patient %s: %s", patientID, message), time.Now())
	if err != nil {
		return fmt.Errorf("error creating notification: %v", err)
	}

	return nil
}

func printProfile(user User) {
	fmt.Println("\n----------YOUR PROFILE-----------")
	fmt.Println("User ID: ", user.UserID)
	fmt.Println("User Name: ", user.Username)
	fmt.Println("Age: ", user.Age)
	fmt.Println("Gender: ", user.Gender)
	fmt.Println("Email: ", user.Email)
	fmt.Println("PhoneNumber: ", user.PhoneNumber)
}
