package models

import (
	"doctor-patient-cli/utils"
	"fmt"
)

func GetPatientByID(userID string) (Patient, error) {
	db := utils.GetDB()
	user := User{}
	db.QueryRow("SELECT user_id, username, age, gender,email, phone_number  FROM users WHERE user_id = ?", userID).
		Scan(&user.UserID, &user.Username, &user.Age, &user.Gender, &user.Email, &user.PhoneNumber)

	patient := Patient{}
	err := db.QueryRow("SELECT user_id, medical_history FROM patients WHERE user_id = ?", userID).
		Scan(&patient.UserID, &patient.MedicalHistory)
	if err != nil {
		return Patient{}, err
	}
	return patient, nil
}

func ViewPatientDetails(userID string) {
	db := utils.GetDB()
	patient := Patient{}
	db.QueryRow("SELECT medical_history FROM patients WHERE user_id = ?", userID).
		Scan(&patient.MedicalHistory)

	fmt.Println("Medical History: ", patient.MedicalHistory)
}
