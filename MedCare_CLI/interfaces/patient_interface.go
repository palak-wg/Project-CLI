package interfaces

import "doctor-patient-cli/models"

// PatientRepository defines the methods related to patient operations
type PatientRepository interface {
	GetPatientByID(patientID string) (*models.Patient, error)
	UpdatePatientDetails(patient *models.Patient) error
}

// PatientService defines the methods available for patient-related operations
type PatientService interface {
	GetPatientByID(patientID string) (*models.Patient, error)
	UpdatePatientDetails(patient *models.Patient) error
}
