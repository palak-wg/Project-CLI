package tests

import (
	//"doctor-patient-cli/controllers"
	"doctor-patient-cli/models"
	"testing"
)

func TestAddReview(t *testing.T) {
	// Example test case for adding a review
	err := models.AddReview("patient1", "doctor1", "Great doctor!", 4)
	if err != nil {
		t.Errorf("Error approving doctor signup: %v", err)
	}
}
