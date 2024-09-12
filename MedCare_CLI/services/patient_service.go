package services

import (
	"doctor-patient-cli/interfaces"
	"doctor-patient-cli/models"
	"log"
)

// PatientService defines methods for managing patient-related operations
type PatientService struct {
	repo interfaces.PatientRepository
}

// NewPatientService creates a new PatientService instance
func NewPatientService(repo interfaces.PatientRepository) *PatientService {
	return &PatientService{repo: repo}
}

// GetPatientByID retrieves a patient by their ID and applies any business logic if needed.
func (service *PatientService) GetPatientByID(patientID string) (*models.Patient, error) {
	patient, err := service.repo.GetPatientByID(patientID)
	if err != nil {
		log.Printf("Service: Error retrieving patient with patientID %s: %v", patientID, err)
		return nil, err
	}
	return patient, nil
}

// UpdatePatientDetails updates a patient's details.
func (service *PatientService) UpdatePatientDetails(patient *models.Patient) error {
	err := service.repo.UpdatePatientDetails(patient)
	if err != nil {
		log.Printf("Service: Error updating details for patientID %s: %v", patient.UserID, err)
		return err
	}
	return nil
}
