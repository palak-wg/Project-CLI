package tests

import (
	"doctor-patient-cli/models"
	"testing"
)

func TestApproveDoctorSignup(t *testing.T) {
	// Example test case for approving doctor signup
	err := models.ApproveDoctorSignup("doctor1")
	if err != nil {
		t.Errorf("Error approving doctor signup: %v", err)
	}

	doctor, _ := models.GetUserByID("doctor1")
	if !doctor.IsApproved {
		t.Errorf("Expected doctor to be approved, but it is not.")
	}
}
