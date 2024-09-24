package repositories

import (
	"database/sql"
	"doctor-patient-cli/interfaces"
	"doctor-patient-cli/models"
	"log"
)

type AppointmentRepositoryImpl struct {
	db *sql.DB
}

// NewAppointmentRepository creates a new instance of AppointmentRepositoryImpl.
func NewAppointmentRepository(db *sql.DB) interfaces.AppointmentRepository {
	return &AppointmentRepositoryImpl{db: db}
}

// SendAppointmentRequest sends an appointment request from a patient to a doctor
func (repo *AppointmentRepositoryImpl) SendAppointmentRequest(patientID, doctorID string) error {
	query := `INSERT INTO appointments (patient_id, doctor_id)VALUES (?, ?)`
	_, err := repo.db.Exec(query, patientID, doctorID)
	if err != nil {
		log.Printf("Error sending appointment request: %v", err)
		return err
	}
	return nil
}

// GetAppointmentsByDoctorID retrieves appointments for a specific doctor
func (repo *AppointmentRepositoryImpl) GetAppointmentsByDoctorID(doctorID string) ([]models.Appointment, error) {
	query := "SELECT appointment_id,doctor_id, patient_id, timestamp,is_approved FROM appointments WHERE doctor_id = ?"
	rows, err := repo.db.Query(query, doctorID)
	if err != nil {
		log.Printf("Error fetching appointments: %v", err)
		return nil, err
	}
	defer rows.Close()

	var appointments []models.Appointment
	for rows.Next() {
		var appointment models.Appointment
		_ = rows.Scan(&appointment.AppointmentID, &appointment.DoctorID, &appointment.PatientID, &appointment.DateTime, &appointment.IsApproved)
		appointments = append(appointments, appointment)
	}
	return appointments, nil
}

// ApproveAppointment approves a pending appointment by updating its status
func (repo *AppointmentRepositoryImpl) ApproveAppointment(appointmentID int) error {
	query := "UPDATE appointments SET is_approved = ? WHERE appointment_id = ?"
	_, err := repo.db.Exec(query, true, appointmentID)
	if err != nil {
		log.Printf("Error approving appointment: %v", err)
		return err
	}
	return nil
}

func (repo *AppointmentRepositoryImpl) GetPendingAppointmentsByDoctorID(doctorID string) ([]models.Appointment, error) {
	query := "SELECT appointment_id, patient_id, timestamp FROM appointments WHERE doctor_id = ? AND is_approved = ?"
	rows, err := repo.db.Query(query, doctorID, false)
	if err != nil {
		log.Printf("Error fetching appointments: %v", err)
		return nil, err
	}
	defer rows.Close()

	var appointments []models.Appointment
	for rows.Next() {
		var appointment models.Appointment
		_ = rows.Scan(&appointment.AppointmentID, &appointment.PatientID, &appointment.DateTime)
		appointments = append(appointments, appointment)
	}
	return appointments, nil
}
