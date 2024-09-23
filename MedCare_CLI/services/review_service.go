package services

import (
	"doctor-patient-cli/interfaces"
	"doctor-patient-cli/models"
	"log"
)

type ReviewService struct {
	repo interfaces.ReviewRepository
}

// NewReviewService creates a new ReviewService instance
func NewReviewService(repo interfaces.ReviewRepository) *ReviewService {
	return &ReviewService{repo: repo}
}

// AddReview adds a new review
func (service *ReviewService) AddReview(review *models.Review) error {
	err := service.repo.AddReview(review)
	if err != nil {
		log.Printf("Service: Error adding review for doctorID %s: %v", review.DoctorID, err)
		return err
	}
	return nil
}

// GetAllReviews retrieves all reviews
func (service *ReviewService) GetAllReviews() ([]models.Review, error) {
	reviews, err := service.repo.GetAllReviews()
	if err != nil {
		log.Printf("Service: Error retrieving all reviews: %v", err)
		return nil, err
	}
	return reviews, nil
}

// GetReviewsByDoctorID retrieves reviews for a specific doctor
func (service *ReviewService) GetReviewsByDoctorID(doctorID string) ([]models.Review, error) {
	reviews, err := service.repo.GetReviewsByDoctorID(doctorID)
	if err != nil {
		log.Printf("Service: Error retrieving reviews for doctorID %s: %v", doctorID, err)
		return nil, err
	}
	return reviews, nil
}
