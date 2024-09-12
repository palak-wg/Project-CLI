package repositories_test

import (
	"database/sql"
	"doctor-patient-cli/repositories"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestSendAppointmentRequest(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := repositories.NewAppointmentRepository(db)

	t.Run("success case", func(t *testing.T) {
		mock.ExpectExec("INSERT INTO appointments").WithArgs("patient1", "doctor1").
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := repo.SendAppointmentRequest("patient1", "doctor1")
		assert.NoError(t, err)
		mock.ExpectationsWereMet()
	})

	t.Run("failure case", func(t *testing.T) {
		mock.ExpectExec("INSERT INTO appointments").WithArgs("patient1", "doctor1").
			WillReturnError(sql.ErrConnDone)

		err := repo.SendAppointmentRequest("patient1", "doctor1")
		assert.Error(t, err)
		mock.ExpectationsWereMet()
	})
}

func TestGetAppointmentsByDoctorID(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := repositories.NewAppointmentRepository(db)

	t.Run("success with multiple appointments", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"appointment_id", "doctor_id", "patient_id", "timestamp", "is_approved"}).
			AddRow(1, "doctor1", "patient1", "2024-01-01 10:00:00", true).
			AddRow(2, "doctor1", "patient2", "2024-01-02 11:00:00", false)

		mock.ExpectQuery("SELECT appointment_id,doctor_id, patient_id, timestamp,is_approved FROM appointments WHERE doctor_id = ?").
			WithArgs("doctor1").
			WillReturnRows(rows)

		appointments, err := repo.GetAppointmentsByDoctorID("doctor1")
		assert.NoError(t, err)
		assert.Len(t, appointments, 2)
		mock.ExpectationsWereMet()
	})

	t.Run("success with no appointments", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"appointment_id", "doctor_id", "patient_id", "timestamp", "is_approved"})
		mock.ExpectQuery("SELECT appointment_id,doctor_id, patient_id, timestamp,is_approved FROM appointments WHERE doctor_id = ?").
			WithArgs("doctor1").
			WillReturnRows(rows)

		appointments, err := repo.GetAppointmentsByDoctorID("doctor1")
		assert.NoError(t, err)
		assert.Len(t, appointments, 0)
		mock.ExpectationsWereMet()
	})

	t.Run("failure case", func(t *testing.T) {
		mock.ExpectQuery("SELECT appointment_id,doctor_id, patient_id, timestamp,is_approved FROM appointments WHERE doctor_id = ?").
			WithArgs("doctor1").
			WillReturnError(sql.ErrConnDone)

		appointments, err := repo.GetAppointmentsByDoctorID("doctor1")
		assert.Error(t, err)
		assert.Nil(t, appointments)
		mock.ExpectationsWereMet()
	})
}

func TestApproveAppointment(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := repositories.NewAppointmentRepository(db)

	t.Run("success case", func(t *testing.T) {
		mock.ExpectExec("UPDATE appointments SET is_approved").WithArgs(true, 1).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := repo.ApproveAppointment(1)
		assert.NoError(t, err)
		mock.ExpectationsWereMet()
	})

	t.Run("failure case", func(t *testing.T) {
		mock.ExpectExec("UPDATE appointments SET is_approved").WithArgs(true, 1).
			WillReturnError(sql.ErrConnDone)

		err := repo.ApproveAppointment(1)
		assert.Error(t, err)
		mock.ExpectationsWereMet()
	})
}

func TestGetPendingAppointmentsByDoctorID(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := repositories.NewAppointmentRepository(db)

	t.Run("success with pending appointments", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"appointment_id", "patient_id"}).
			AddRow(1, "patient1").
			AddRow(2, "patient2")

		mock.ExpectQuery("^SELECT appointment_id, patient_id FROM appointments WHERE doctor_id = \\? AND is_approved = \\?$").
			WithArgs("doctor1", false).
			WillReturnRows(rows)

		appointments, err := repo.GetPendingAppointmentsByDoctorID("doctor1")
		assert.NoError(t, err)
		assert.Len(t, appointments, 2)
		mock.ExpectationsWereMet()
	})

	t.Run("success with no pending appointments", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"appointment_id", "patient_id"})
		mock.ExpectQuery("^SELECT appointment_id, patient_id FROM appointments WHERE doctor_id = \\? AND is_approved = \\?$").
			WithArgs("doctor1", false).
			WillReturnRows(rows)

		appointments, err := repo.GetPendingAppointmentsByDoctorID("doctor1")
		assert.NoError(t, err)
		assert.Len(t, appointments, 0)
		mock.ExpectationsWereMet()
	})

	t.Run("failure case", func(t *testing.T) {
		mock.ExpectQuery("SELECT appointment_id, patient_id FROM appointments WHERE doctor_id = ? AND is_approved = ?").
			WithArgs("doctor1", false).
			WillReturnError(sql.ErrConnDone)

		appointments, err := repo.GetPendingAppointmentsByDoctorID("doctor1")
		assert.Error(t, err)
		assert.Nil(t, appointments)
		mock.ExpectationsWereMet()
	})
}
