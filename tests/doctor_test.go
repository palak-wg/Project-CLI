package tests

import (
	"doctor-patient-cli/models"
	"testing"
)

func TestUpdateDoctorExperience(t *testing.T) {
	// Example test case for updating doctor experience
	err := models.UpdateDoctorExperience("doctor1", 10)
	if err != nil {
		t.Errorf("Error updating doctor experience: %v", err)
	}

	doctor, _ := models.GetDoctorByID("doctor1")
	if doctor.Experience != 10 {
		t.Errorf("Expected experience to be 10, got %d", doctor.Experience)
	}
}
