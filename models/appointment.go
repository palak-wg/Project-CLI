package models

import (
	"doctor-patient-cli/utils"
	"fmt"
)

func GetAppointmentsByDoctorID(doctorID string) ([]Appointment, error) {
	db := utils.GetDB()
	rows, err := db.Query("SELECT appointment_id,doctor_id, patient_id, timestamp,is_approved FROM appointments WHERE doctor_id = ?", doctorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var appointments []Appointment
	for rows.Next() {
		var appointment Appointment
		err = rows.Scan(&appointment.AppointmentID, &appointment.DoctorID, &appointment.PatientID, &appointment.DateTime, &appointment.IsApproved)
		if err != nil {
			return nil, err
		}
		appointments = append(appointments, appointment)
	}
	return appointments, nil
}

func ApproveAppointment(appointmentID string) error {
	db := utils.GetDB()
	_, err := db.Exec("UPDATE appointments SET is_approved = ? WHERE appointment_id = ?", true, appointmentID)
	return err
}

// SendAppointmentRequest allows a patient to send an appointment request to a doctor
func SendAppointmentRequest(patientID, doctorID string) error {
	db := utils.GetDB()

	// Insert the appointment request into the appointments table
	_, err := db.Exec(`INSERT INTO appointments (patient_id, doctor_id)VALUES (?, ?)`, patientID, doctorID)

	if err != nil {
		return fmt.Errorf("error sending appointment request: %v", err)
	}

	fmt.Println("Appointment request sent successfully.")
	return nil
}
