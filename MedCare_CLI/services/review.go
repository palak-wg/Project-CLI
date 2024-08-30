package services

import (
	"doctor-patient-cli/models"
	"doctor-patient-cli/utils"
	"fmt"
)

func AddReview(patientID, doctorID, content string, rating int) error {
	db := utils.GetDB()
	_, err := db.Exec("INSERT INTO reviews (patient_id, doctor_id, content, rating) VALUES (?, ?, ?, ?)", patientID, doctorID, content, rating)
	return err
}

func GetAllReviews() ([]models.Review, error) {
	db := utils.GetDB()
	fmt.Println("All reviews:")
	rows, err := db.Query("SELECT patient_id, doctor_id, content, rating FROM reviews")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []models.Review
	for rows.Next() {
		var review models.Review
		_ = rows.Scan(&review.PatientID, &review.DoctorID, &review.Content, &review.Rating)
		reviews = append(reviews, review)
	}
	return reviews, nil
}
