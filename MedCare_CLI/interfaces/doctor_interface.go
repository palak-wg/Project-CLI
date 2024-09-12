package interfaces

import "doctor-patient-cli/models"

// DoctorRepository defines the methods related to doctor operations
type DoctorRepository interface {
	GetDoctorByID(doctorID string) (*models.Doctor, error)
	GetAllDoctors() ([]models.Doctor, error)
	UpdateDoctorExperience(doctorID string, experience int) error
	UpdateDoctorSpecialization(doctorID string, specialization string) error
	ViewDoctorSpecificProfile(doctorID string) (*models.Doctor, error)
}

// DoctorService defines the methods for doctor-related operations
type DoctorService interface {
	GetDoctorByID(doctorID string) (*models.Doctor, error)
	GetAllDoctors() ([]models.Doctor, error)
	UpdateDoctorExperience(doctorID string, experience int) error
	UpdateDoctorSpecialization(doctorID string, specialization string) error
	ViewDoctorSpecificProfile(doctorID string) (*models.Doctor, error)
}
