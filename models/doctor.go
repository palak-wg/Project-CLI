package models

import (
	"database/sql"
	"doctor-patient-cli/utils"
	"fmt"
	"time"
)

func GetDoctorByID(userID string) (Doctor, error) {
	db := utils.GetDB()
	doctor := Doctor{}
	err := db.QueryRow("SELECT user_id, specialization, experience, rating FROM doctors WHERE user_id = ?", userID).
		Scan(&doctor.UserID, &doctor.Specialization, &doctor.Experience, &doctor.Rating)
	if err != nil {
		return Doctor{}, err
	}
	return doctor, nil
}

func GetAllDoctors() ([]Doctor, error) {
	db := utils.GetDB()
	rows, err := db.Query("SELECT user_id, specialization, experience, rating FROM doctors")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var doctors []Doctor
	for rows.Next() {
		var doctor Doctor
		err = rows.Scan(&doctor.UserID, &doctor.Specialization, &doctor.Experience, &doctor.Rating)
		if err != nil {
			return nil, err
		}
		doctors = append(doctors, doctor)
	}
	return doctors, nil
}

func UpdateDoctorExperience(userID string, experience int) error {
	db := utils.GetDB()
	_, err := db.Exec("UPDATE doctors SET experience = ? WHERE user_id = ?", experience, userID)
	return err
}

func UpdateDoctorSpecialization(userID, specialization string) error {
	db := utils.GetDB()
	_, err := db.Exec("UPDATE doctors SET specialization = ? WHERE user_id = ?", specialization, userID)
	return err
}

var db *sql.DB

// RespondToPatientRequest allows a doctor to respond to a patient request.
func RespondToPatientRequest(doctorID, patientID, response string) error {
	// Check if doctor is approved
	doctor, err := GetDoctorByID(doctorID)
	if err != nil {
		return fmt.Errorf("error fetching doctor: %v", err)
	}
	if !doctor.IsApproved {
		return fmt.Errorf("doctor is not approved")
	}

	// Check if patient request exists
	var requestID string
	err = db.QueryRow("SELECT id FROM requests WHERE doctor_id = ? AND patient_id = ? AND status = 'pending'", doctorID, patientID).Scan(&requestID)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("no pending request found from patient %s to doctor %s", patientID, doctorID)
		}
		return fmt.Errorf("error checking patient request: %v", err)
	}

	// Update request with doctor's response
	_, err = db.Exec("UPDATE requests SET response = ?, response_timestamp = ?, status = 'responded' WHERE id = ?", response, time.Now(), requestID)
	if err != nil {
		return fmt.Errorf("error updating patient request: %v", err)
	}

	// Create a notification for the patient
	_, err = db.Exec("INSERT INTO notifications (user_id, content, timestamp) VALUES (?, ?, ?)",
		patientID, fmt.Sprintf("Doctor %s has responded to your request: %s", doctorID, response), time.Now())
	if err != nil {
		return fmt.Errorf("error creating notification: %v", err)
	}

	return nil
}

// SuggestPrescription allows a doctor to suggest a prescription to a patient.
func SuggestPrescription(doctorID, patientID, prescription string) error {
	// Check if doctor is approved
	doctor, err := GetDoctorByID(doctorID)
	if err != nil {
		return fmt.Errorf("error fetching doctor: %v", err)
	}
	if !doctor.IsApproved {
		return fmt.Errorf("doctor is not approved")
	}

	// Create a new prescription record
	prescriptionID := fmt.Sprintf("presc-%d", time.Now().UnixNano())
	_, err = db.Exec("INSERT INTO prescriptions (id, doctor_id, patient_id, prescription) VALUES (?, ?, ?, ?)",
		prescriptionID, doctorID, patientID, prescription)
	if err != nil {
		return fmt.Errorf("error inserting prescription: %v", err)
	}

	// Create a notification for the patient
	_, err = db.Exec("INSERT INTO notifications (user_id, content, timestamp) VALUES (?, ?, ?)",
		patientID, fmt.Sprintf("Doctor %s has suggested a prescription for you: %s", doctorID, prescription), time.Now())
	if err != nil {
		return fmt.Errorf("error creating notification: %v", err)
	}

	return nil
}

func RequestForDoctorSignup(doctorID string) error {
	_, err := db.Exec("INSERT INTO notifications (user_id, content) VALUES (?, ?)",
		"admin", "please approve")
	return err
}
