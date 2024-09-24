package repositories

import (
	"database/sql"
	"doctor-patient-cli/interfaces"
	"doctor-patient-cli/models"
	"log"
)

type ReviewRepositoryImpl struct {
	db *sql.DB
}

// NewReviewRepository creates a new ReviewRepositoryImpl instance
func NewReviewRepository(db *sql.DB) interfaces.ReviewRepository {
	return &ReviewRepositoryImpl{db: db}
}

// AddReview inserts a new review into the database
func (repo *ReviewRepositoryImpl) AddReview(review *models.Review) error {
	_, err := repo.db.Exec("INSERT INTO reviews (patient_id, doctor_id, content, rating) VALUES (?, ?, ?, ?)",
		review.PatientID, review.DoctorID, review.Content, review.Rating)
	if err != nil {
		log.Printf("Repository: Error adding review for doctorID %s: %v", review.DoctorID, err)
		return err
	}
	return nil
}

// GetAllReviews retrieves all reviews from the database
func (repo *ReviewRepositoryImpl) GetAllReviews() ([]models.Review, error) {
	rows, err := repo.db.Query("SELECT patient_id, doctor_id, content, rating, timestamp FROM reviews")
	if err != nil {
		log.Printf("Repository: Error fetching all reviews: %v", err)
		return nil, err
	}
	defer rows.Close()

	var reviews []models.Review
	for rows.Next() {
		var review models.Review
		_ = rows.Scan(&review.PatientID, &review.DoctorID, &review.Content, &review.Rating, &review.Timestamp)

		reviews = append(reviews, review)
	}
	return reviews, nil
}

// GetReviewsByDoctorID retrieves reviews for a specific doctor from the database
func (repo *ReviewRepositoryImpl) GetReviewsByDoctorID(doctorID string) ([]models.Review, error) {
	rows, err := repo.db.Query("SELECT patient_id, doctor_id, content, rating, timestamp FROM reviews WHERE doctor_id = ?", doctorID)
	if err != nil {
		log.Printf("Repository: Error fetching reviews for doctorID %s: %v", doctorID, err)
		return nil, err
	}
	defer rows.Close()

	var reviews []models.Review
	for rows.Next() {
		var review models.Review
		_ = rows.Scan(&review.PatientID, &review.DoctorID, &review.Content, &review.Rating, &review.Timestamp)
		reviews = append(reviews, review)
	}
	return reviews, nil
}
