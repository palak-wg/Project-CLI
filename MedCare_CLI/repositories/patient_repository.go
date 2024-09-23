package repositories

import (
	"database/sql"
	"doctor-patient-cli/interfaces"
	"doctor-patient-cli/models"
	"errors"
	"log"
)

// PatientRepositoryImpl is the concrete implementation of the PatientRepository interface
type PatientRepositoryImpl struct {
	db *sql.DB
}

// NewPatientRepository returns a new instance of PatientRepositoryImpl
func NewPatientRepository(db *sql.DB) interfaces.PatientRepository {
	return &PatientRepositoryImpl{db}
}

// GetPatientByID returns a patient by their ID from the database
func (repo *PatientRepositoryImpl) GetPatientByID(patientID string) (*models.Patient, error) {
	var patient models.Patient
	query := `SELECT user_id, username, age, gender,email, phone_number  FROM users WHERE user_id = ?`
	err := repo.db.QueryRow(query, patientID).Scan(
		&patient.UserID, &patient.Name, &patient.Age, &patient.Gender, &patient.Email,
		&patient.PhoneNumber)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("patient not found")
		}
		log.Printf("Error fetching patient: %v", err)
		return nil, err
	}

	return &patient, nil
}

// UpdatePatientDetails updates the patient's profile in the database
func (repo *PatientRepositoryImpl) UpdatePatientDetails(patient *models.Patient) error {
	query := `UPDATE users SET name = ?, age = ?, gender = ?, email = ?, phone_number = ? WHERE user_id = ?`
	_, err := repo.db.Exec(query, patient.Name, patient.Age, patient.Gender, patient.Email, patient.PhoneNumber, patient.UserID)

	if err != nil {
		log.Printf("Error updating patient details: %v", err)
		return err
	}
	return nil
}
