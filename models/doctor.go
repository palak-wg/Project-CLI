package models

import (
	"doctor-patient-cli/utils"
	"fmt"
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

func ViewDoctorSpecificProfile(userID string) {
	db := utils.GetDB()
	doctor := Doctor{}
	_ = db.QueryRow("SELECT specialization, experience, rating FROM doctors WHERE user_id = ?", userID).
		Scan(&doctor.Specialization, &doctor.Experience, &doctor.Rating)

	fmt.Println("Specialization: ", doctor.Specialization)
	fmt.Println("Experience: ", doctor.Experience)
	fmt.Println("Rating: ", doctor.Rating)
}
