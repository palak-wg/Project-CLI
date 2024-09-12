package interfaces

import "doctor-patient-cli/models"

// AppointmentRepository defines the methods related to appointment management
type AppointmentRepository interface {
	SendAppointmentRequest(patientID, doctorID string) error
	GetAppointmentsByDoctorID(doctorID string) ([]models.Appointment, error)
	ApproveAppointment(appointmentID int) error
	GetPendingAppointmentsByDoctorID(doctorID string) ([]models.Appointment, error)
}

// AppointmentService defines methods for managing appointments
type AppointmentService interface {
	SendAppointmentRequest(patientID, doctorID string) error
	GetAppointmentsByDoctorID(doctorID string) ([]models.Appointment, error)
	ApproveAppointment(appointmentID int) error
	GetPendingAppointmentsByDoctorID(doctorID string) ([]models.Appointment, error)
}
