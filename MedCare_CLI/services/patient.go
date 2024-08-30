package services

import (
	"doctor-patient-cli/models"
	"doctor-patient-cli/utils"
	"fmt"
)

func GetPatientByID(userID string) (models.Patient, error) {
	db := utils.GetDB()
	user := models.User{}
	db.QueryRow("SELECT user_id, username, age, gender,email, phone_number  FROM users WHERE user_id = ?", userID).
		Scan(&user.UserID, &user.Username, &user.Age, &user.Gender, &user.Email, &user.PhoneNumber)

	patient := models.Patient{}
	err := db.QueryRow("SELECT user_id, medical_history FROM patients WHERE user_id = ?", userID).
		Scan(&patient.UserID, &patient.MedicalHistory)
	if err != nil {
		return models.Patient{}, err
	}
	return patient, nil
}

func ViewPatientDetails(userID string) {
	db := utils.GetDB()
	patient := models.Patient{}
	db.QueryRow("SELECT medical_history FROM patients WHERE user_id = ?", userID).
		Scan(&patient.MedicalHistory)

	fmt.Println("Medical History: ", patient.MedicalHistory)
}
