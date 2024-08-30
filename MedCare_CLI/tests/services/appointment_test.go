package services

import (
	"doctor-patient-cli/services"
	"doctor-patient-cli/tests/mockDB"
	"doctor-patient-cli/utils"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAppointmentsByDoctorID(t *testing.T) {
	mockDB.MockInitDB(t)
	defer utils.CloseDB()

	t.Run("GetAppointmentsByDoctorID Success", func(t *testing.T) {

		rows := sqlmock.NewRows([]string{"appointment_id", "doctor_id", "patient_id", "timestamp", "is_approved"}).
			AddRow(1, "doctor1", "patient1", "2024-08-26 10:00:00", true).
			AddRow(2, "doctor1", "patient2", "2024-08-27 11:00:00", false)

		// Set up the expectation for the Query to return rows
		mockDB.Mock.ExpectQuery("SELECT appointment_id,doctor_id, patient_id, timestamp,is_approved FROM appointments WHERE doctor_id = ?").
			WithArgs("doctor1").
			WillReturnRows(rows)

		// Call the GetAppointmentsByDoctorID function
		appointments, err := services.GetAppointmentsByDoctorID("doctor1")

		// Assert that an error is returned and appointments is nil
		assert.NoError(t, err)
		assert.Len(t, appointments, 2)
		assert.Equal(t, 1, appointments[0].AppointmentID) // Expecting integer
		assert.Equal(t, "doctor1", appointments[0].DoctorID)

		// Ensure all expectations are met
		if err := mockDB.Mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})

	t.Run("GetAppointmentsByDoctorID Failure", func(t *testing.T) {
		// Set up the expectation for the Query to return an error
		mockDB.Mock.ExpectQuery("SELECT appointment_id,doctor_id, patient_id, timestamp,is_approved FROM appointments WHERE doctor_id = ?").
			WithArgs("doctor1").
			WillReturnError(fmt.Errorf("query error"))

		// Call the GetAppointmentsByDoctorID function
		appointments, err := services.GetAppointmentsByDoctorID("doctor1")

		// Assert that an error is returned and appointments is nil
		assert.Error(t, err)
		assert.Nil(t, appointments)
		assert.EqualError(t, err, "query error")

		// Ensure all expectations are met
		if err := mockDB.Mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})
}

func TestApproveAppointment(t *testing.T) {
	mockDB.MockInitDB(t)
	defer utils.CloseDB()

	query := regexp.QuoteMeta("UPDATE appointments SET is_approved = ? WHERE appointment_id = ?")

	mockDB.Mock.ExpectExec(query).
		WithArgs(true, "appointment1").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := services.ApproveAppointment("appointment1")
	assert.NoError(t, err)
}

func TestSendAppointmentRequest(t *testing.T) {
	mockDB.MockInitDB(t)
	defer utils.CloseDB()

	t.Run("SendAppointmentRequest Success", func(t *testing.T) {

		// Mock the Appointment request result
		mockDB.Mock.ExpectExec("INSERT INTO appointments \\(patient_id, doctor_id\\)VALUES \\(\\?, \\?\\)").
			WithArgs("patient1", "doctor1").
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := services.SendAppointmentRequest("patient1", "doctor1")
		assert.NoError(t, err)
	})

	t.Run("SendAppointmentRequest Failure", func(t *testing.T) {
		// Set up the expectation for the Exec query to return an error
		mockDB.Mock.ExpectExec("INSERT INTO appointments \\(patient_id, doctor_id\\)VALUES \\(\\?, \\?\\)").
			WithArgs("patient1", "doctor1").
			WillReturnError(fmt.Errorf("database error"))

		// Call the SendAppointmentRequest function
		err := services.SendAppointmentRequest("patient1", "doctor1")

		// Assert that an error is returned
		assert.Error(t, err)
		assert.EqualError(t, err, "error sending appointment request: database error")

		// Ensure all expectations are met
		if err := mockDB.Mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})
}
