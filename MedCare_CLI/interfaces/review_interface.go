package interfaces

import "doctor-patient-cli/models"

// ReviewRepository defines the methods related to reviews
type ReviewRepository interface {
	AddReview(review *models.Review) error
	GetAllReviews() ([]models.Review, error)
	GetReviewsByDoctorID(doctorID int) ([]models.Review, error)
}

// ReviewService defines the methods for review-related operations
type ReviewService interface {
	AddReview(review *models.Review) error
	GetAllReviews() ([]models.Review, error)
	GetReviewsByDoctorID(doctorID int) ([]models.Review, error)
}
