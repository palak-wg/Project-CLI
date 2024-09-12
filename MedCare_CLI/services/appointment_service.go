package services

import (
	"doctor-patient-cli/interfaces"
	"doctor-patient-cli/models"
	"log"
)

// AppointmentService defines methods for managing appointments
type AppointmentService struct {
	repo interfaces.AppointmentRepository
}

// NewAppointmentService creates a new instance of AppointmentService
func NewAppointmentService(repo interfaces.AppointmentRepository) *AppointmentService {
	return &AppointmentService{repo: repo}
}

// SendAppointmentRequest creates a new appointment request.
func (service *AppointmentService) SendAppointmentRequest(patientID, doctorID string) error {
	err := service.repo.SendAppointmentRequest(patientID, doctorID)
	if err != nil {
		log.Printf("Service: Error sending appointment request for patientID %s and doctorID %s: %v", patientID, doctorID, err)
		return err
	}
	return nil
}

// GetAppointmentsByDoctorID retrieves all appointments for a specific doctor
func (service *AppointmentService) GetAppointmentsByDoctorID(doctorID string) ([]models.Appointment, error) {
	appointments, err := service.repo.GetAppointmentsByDoctorID(doctorID)
	if err != nil {
		log.Printf("Service: Error fetching appointments for doctorID %s: %v", doctorID, err)
		return nil, err
	}
	return appointments, nil
}

// ApproveAppointment approves an appointment by its ID
func (service *AppointmentService) ApproveAppointment(appointmentID int) error {
	err := service.repo.ApproveAppointment(appointmentID)
	if err != nil {
		log.Printf("Service: Error approving appointment with appointmentID %d: %v", appointmentID, err)
		return err
	}
	return nil
}

func (service *AppointmentService) GetPendingAppointmentsByDoctorID(doctorID string) ([]models.Appointment, error) {
	appointments, err := service.repo.GetPendingAppointmentsByDoctorID(doctorID)
	if err != nil {
		log.Printf("Service: Error fetching appointments for doctorID %s: %v", doctorID, err)
		return nil, err
	}
	return appointments, nil
}
