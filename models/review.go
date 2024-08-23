package models

import (
	"doctor-patient-cli/utils"
)

func AddReview(patientID, doctorID, content string, rating int) error {
	db := utils.GetDB()
	_, err := db.Exec("INSERT INTO reviews (patient_id, doctor_id, content, rating) VALUES (?, ?, ?, ?)", patientID, doctorID, content, rating)
	return err
}

func GetAllReviews() ([]Review, error) {
	db := utils.GetDB()
	rows, err := db.Query("SELECT patient_id, doctor_id, content, rating FROM reviews")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []Review
	for rows.Next() {
		var review Review
		err = rows.Scan(&review.PatientID, &review.DoctorID, &review.Content, &review.Rating)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, review)
	}
	return reviews, nil
}
