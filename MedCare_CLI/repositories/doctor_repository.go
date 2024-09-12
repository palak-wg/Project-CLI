package repositories

import (
	"database/sql"
	"doctor-patient-cli/interfaces"
	"doctor-patient-cli/models"
)

type DoctorRepositoryImpl struct {
	db *sql.DB
}

// Ensure DoctorRepositoryImpl implements the DoctorRepository interface
var _ interfaces.DoctorRepository = &DoctorRepositoryImpl{}

// NewDoctorRepository creates a new instance of DoctorRepositoryImpl
func NewDoctorRepository(db *sql.DB) interfaces.DoctorRepository {
	return &DoctorRepositoryImpl{db: db}
}

func (repo *DoctorRepositoryImpl) GetDoctorByID(doctorID string) (*models.Doctor, error) {
	query := "SELECT user_id, specialization, experience, rating FROM doctors WHERE user_id = ?"
	row := repo.db.QueryRow(query, doctorID)

	doctor := &models.Doctor{}
	err := row.Scan(&doctor.UserID, &doctor.Specialization, &doctor.Experience, &doctor.Rating)
	if err != nil {
		return nil, err
	}
	return doctor, nil
}

func (repo *DoctorRepositoryImpl) GetAllDoctors() ([]models.Doctor, error) {
	query := "SELECT user_id, specialization, experience, rating FROM doctors"
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var doctors []models.Doctor
	for rows.Next() {
		doctor := models.Doctor{}
		_ = rows.Scan(&doctor.UserID, &doctor.Specialization, &doctor.Experience, &doctor.Rating)
		doctors = append(doctors, doctor)
	}
	return doctors, nil
}

func (repo *DoctorRepositoryImpl) UpdateDoctorExperience(doctorID string, experience int) error {
	query := "UPDATE doctors SET experience = ? WHERE user_id = ?"
	_, err := repo.db.Exec(query, experience, doctorID)
	return err
}

func (repo *DoctorRepositoryImpl) UpdateDoctorSpecialization(doctorID string, specialization string) error {
	query := "UPDATE doctors SET specialization = ? WHERE user_id = ?"
	_, err := repo.db.Exec(query, specialization, doctorID)
	return err
}

func (repo *DoctorRepositoryImpl) ViewDoctorSpecificProfile(doctorID string) (*models.Doctor, error) {
	return repo.GetDoctorByID(doctorID)
}
