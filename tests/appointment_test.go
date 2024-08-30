package tests

import (
	"doctor-patient-cli/models"
	"doctor-patient-cli/utils"
	"github.com/DATA-DOG/go-sqlmock"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAppointmentsByDoctorID(t *testing.T) {
	MockInitDB(t)
	defer utils.CloseDB()

	rows := sqlmock.NewRows([]string{"appointment_id", "doctor_id", "patient_id", "timestamp", "is_approved"}).
		AddRow(1, "doctor1", "patient1", "2024-08-26 10:00:00", true).
		AddRow(2, "doctor1", "patient2", "2024-08-27 11:00:00", false)

	mock.ExpectQuery("SELECT appointment_id,doctor_id, patient_id, timestamp,is_approved FROM appointments WHERE doctor_id = ?").
		WithArgs("doctor1").
		WillReturnRows(rows)

	appointments, err := models.GetAppointmentsByDoctorID("doctor1")
	assert.NoError(t, err)
	assert.Len(t, appointments, 2)
	assert.Equal(t, 1, appointments[0].AppointmentID) // Expecting integer
	assert.Equal(t, "doctor1", appointments[0].DoctorID)
}

func TestApproveAppointment(t *testing.T) {
	MockInitDB(t)
	defer utils.CloseDB()

	query := regexp.QuoteMeta("UPDATE appointments SET is_approved = ? WHERE appointment_id = ?")

	mock.ExpectExec(query).
		WithArgs(true, "appointment1").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := models.ApproveAppointment("appointment1")
	assert.NoError(t, err)
}

func TestSendAppointmentRequest(t *testing.T) {
	MockInitDB(t)
	defer utils.CloseDB()

	mock.ExpectExec("INSERT INTO appointments \\(patient_id, doctor_id\\)VALUES \\(\\?, \\?\\)").
		WithArgs("patient1", "doctor1").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := models.SendAppointmentRequest("patient1", "doctor1")
	assert.NoError(t, err)
}
