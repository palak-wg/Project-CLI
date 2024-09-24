package repositories

import (
	"database/sql"
	"doctor-patient-cli/interfaces"
	"doctor-patient-cli/models"
	"log"
)

type AdminRepositoryImpl struct {
	db *sql.DB
}

func NewAdminRepository(db *sql.DB) interfaces.AdminRepository {
	return &AdminRepositoryImpl{db}
}

// ApproveDoctorSignup approves a doctor signup request and inserts a record into the doctors table
func (repo *AdminRepositoryImpl) ApproveDoctorSignup(userID string) error {
	// Update the user record to set IsApproved to true
	_, err := repo.db.Exec("UPDATE users SET is_approved = ? WHERE user_id = ?", true, userID)
	if err != nil {
		log.Printf("Repository: Error updating user approval status: %v", err)
		return err
	}

	// Make entry to doctors table
	_, err = repo.db.Exec("INSERT INTO doctors (user_id, specialization, experience, rating) VALUES (?, ?, ?, ?)",
		userID, "xxx", 0, 2)
	if err != nil {
		log.Printf("Repository: Error inserting doctor record: %v", err)
		return err
	}

	return nil
}

// CreateNotificationForUser creates a notification for the specified user
func (repo *AdminRepositoryImpl) CreateNotificationForUser(userID string, content string) error {
	_, err := repo.db.Exec("INSERT INTO notifications (user_id, content) VALUES (?, ?)",
		userID, content)
	if err != nil {
		log.Printf("Repository: Error creating notification: %v", err)
		return err
	}
	return nil
}

func (repo *AdminRepositoryImpl) GetPendingDoctorRequests() ([]models.Doctor, error) {
	query := "SELECT user_id, username FROM users WHERE user_type = 'doctor' AND is_approved = 0"
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var doctors []models.Doctor
	for rows.Next() {
		doctor := models.Doctor{}
		err := rows.Scan(&doctor.UserID, &doctor.Name)
		if err != nil {
			return nil, err
		}
		doctors = append(doctors, doctor)
	}
	return doctors, nil
}

func (repo *AdminRepositoryImpl) GetAllUsers() ([]models.User, error) {
	query := "SELECT * FROM users"
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		user := models.User{}
		err := rows.Scan(&user.UserID, &user.Password, &user.Name, &user.Age, &user.Gender, &user.Email, &user.PhoneNumber, &user.UserType, &user.IsApproved)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
