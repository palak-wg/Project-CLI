package repositories_test

import (
	"database/sql"
	"doctor-patient-cli/models"
	"doctor-patient-cli/repositories"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetDoctorByID_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repositories.NewDoctorRepository(db)

	doctorID := "doctor123"
	expectedDoctor := &models.Doctor{
		User: models.User{
			UserID: doctorID,
		},
		Specialization: "Cardiology",
		Experience:     10,
		Rating:         4.5,
	}

	// Mocking the SQL query
	row := sqlmock.NewRows([]string{"user_id", "specialization", "experience", "rating"}).
		AddRow(expectedDoctor.UserID, expectedDoctor.Specialization, expectedDoctor.Experience, expectedDoctor.Rating)
	mock.ExpectQuery("SELECT user_id, specialization, experience, rating FROM doctors WHERE user_id = ?").
		WithArgs(doctorID).WillReturnRows(row)

	doctor, err := repo.GetDoctorByID(doctorID)
	assert.NoError(t, err)
	assert.Equal(t, expectedDoctor, doctor)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetDoctorByID_NotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repositories.NewDoctorRepository(db)

	doctorID := "doctor123"

	// Simulating no rows found
	mock.ExpectQuery("SELECT user_id, specialization, experience, rating FROM doctors WHERE user_id = ?").
		WithArgs(doctorID).WillReturnError(sql.ErrNoRows)

	doctor, err := repo.GetDoctorByID(doctorID)
	assert.Error(t, err)
	assert.Nil(t, doctor)
	assert.EqualError(t, err, sql.ErrNoRows.Error())
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetAllDoctors_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repositories.NewDoctorRepository(db)

	expectedDoctors := []models.Doctor{
		{User: models.User{UserID: "doctor1"}, Specialization: "Cardiology", Experience: 15, Rating: 4.3},
		{User: models.User{UserID: "doctor2"}, Specialization: "Neurology", Experience: 12, Rating: 4.6},
	}

	// Mocking the SQL query
	rows := sqlmock.NewRows([]string{"user_id", "specialization", "experience", "rating"}).
		AddRow(expectedDoctors[0].UserID, expectedDoctors[0].Specialization, expectedDoctors[0].Experience, expectedDoctors[0].Rating).
		AddRow(expectedDoctors[1].UserID, expectedDoctors[1].Specialization, expectedDoctors[1].Experience, expectedDoctors[1].Rating)
	mock.ExpectQuery("SELECT user_id, specialization, experience, rating FROM doctors").WillReturnRows(rows)

	doctors, err := repo.GetAllDoctors()
	assert.NoError(t, err)
	assert.Equal(t, expectedDoctors, doctors)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetAllDoctors_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repositories.NewDoctorRepository(db)

	// Simulating an error
	mock.ExpectQuery("SELECT user_id, specialization, experience, rating FROM doctors").
		WillReturnError(errors.New("query error"))

	doctors, err := repo.GetAllDoctors()
	assert.Error(t, err)
	assert.Nil(t, doctors)
	assert.EqualError(t, err, "query error")
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateDoctorExperience_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repositories.NewDoctorRepository(db)

	doctorID := "doctor123"
	newExperience := 12

	// Mocking the SQL query
	mock.ExpectExec("^UPDATE doctors SET experience = \\? WHERE user_id = \\?$").
		WithArgs(newExperience, doctorID).WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.UpdateDoctorExperience(doctorID, newExperience)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateDoctorExperience_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repositories.NewDoctorRepository(db)

	doctorID := "doctor123"
	newExperience := 12

	// Simulating an error
	mock.ExpectExec("^UPDATE doctors SET experience = \\? WHERE user_id = \\?$").
		WithArgs(newExperience, doctorID).WillReturnError(errors.New("update error"))

	err = repo.UpdateDoctorExperience(doctorID, newExperience)
	assert.Error(t, err)
	assert.EqualError(t, err, "update error")
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateDoctorSpecialization_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repositories.NewDoctorRepository(db)

	doctorID := "doctor123"
	newSpecialization := "Orthopedics"

	// Mocking the SQL query
	mock.ExpectExec("^UPDATE doctors SET specialization = \\? WHERE user_id = \\?$").
		WithArgs(newSpecialization, doctorID).WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.UpdateDoctorSpecialization(doctorID, newSpecialization)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateDoctorSpecialization_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repositories.NewDoctorRepository(db)

	doctorID := "doctor123"
	newSpecialization := "Orthopedics"

	// Simulating an error
	mock.ExpectExec("^UPDATE doctors SET specialization = \\? WHERE user_id = \\?$").
		WithArgs(newSpecialization, doctorID).WillReturnError(errors.New("update error"))

	err = repo.UpdateDoctorSpecialization(doctorID, newSpecialization)
	assert.Error(t, err)
	assert.EqualError(t, err, "update error")
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestViewDoctorSpecificProfile_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repositories.NewDoctorRepository(db)

	doctorID := "doctor123"
	expectedDoctor := &models.Doctor{
		User:           models.User{UserID: "doctor1"},
		Specialization: "Pediatrics",
		Experience:     5,
		Rating:         4.7,
	}

	// Mocking the SQL query
	row := sqlmock.NewRows([]string{"user_id", "specialization", "experience", "rating"}).
		AddRow(expectedDoctor.UserID, expectedDoctor.Specialization, expectedDoctor.Experience, expectedDoctor.Rating)
	mock.ExpectQuery("^SELECT user_id, specialization, experience, rating FROM doctors WHERE user_id = \\?$").
		WithArgs(doctorID).WillReturnRows(row)

	doctor, err := repo.ViewDoctorSpecificProfile(doctorID)
	assert.NoError(t, err)
	assert.Equal(t, expectedDoctor, doctor)
	assert.NoError(t, mock.ExpectationsWereMet())
}
