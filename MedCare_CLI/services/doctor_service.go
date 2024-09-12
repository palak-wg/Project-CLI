package services

import (
	"doctor-patient-cli/interfaces"
	"doctor-patient-cli/models"
)

type DoctorService struct {
	repo interfaces.DoctorRepository
}

// NewDoctorService creates a new instance of DoctorService
func NewDoctorService(repo interfaces.DoctorRepository) *DoctorService {
	return &DoctorService{repo: repo}
}

// GetDoctorByID retrieves a doctor by their ID
func (s *DoctorService) GetDoctorByID(doctorID string) (*models.Doctor, error) {
	return s.repo.GetDoctorByID(doctorID)
}

// ViewDoctorSpecificProfile retrieves the specific profile of a doctor
func (s *DoctorService) ViewDoctorSpecificProfile(doctorID string) (*models.Doctor, error) {
	return s.repo.ViewDoctorSpecificProfile(doctorID)
}

// UpdateDoctorExperience updates the doctor's experience
func (s *DoctorService) UpdateDoctorExperience(doctorID string, experience int) error {
	return s.repo.UpdateDoctorExperience(doctorID, experience)
}

// UpdateDoctorSpecialization updates the doctor's specialization
func (s *DoctorService) UpdateDoctorSpecialization(doctorID string, specialization string) error {
	return s.repo.UpdateDoctorSpecialization(doctorID, specialization)
}

// GetAllDoctors retrieves all doctors
func (s *DoctorService) GetAllDoctors() ([]models.Doctor, error) {
	return s.repo.GetAllDoctors()
}
