package repositories

import (
	"database/sql"
	"doctor-patient-cli/interfaces"
	"doctor-patient-cli/models"
	"errors"
	"fmt"
)

// UserRepositoryImpl is the concrete implementation of the UserRepository interface
type UserRepositoryImpl struct {
	db *sql.DB
}

var _ interfaces.UserRepository = &UserRepositoryImpl{}

// NewUserRepository returns a new instance of UserRepositoryImpl
func NewUserRepository(db *sql.DB) interfaces.UserRepository {
	return &UserRepositoryImpl{db}
}

// CreateUser adds a new user to the database
func (repo *UserRepositoryImpl) CreateUser(user *models.User) error {
	query := "INSERT INTO users (user_id, password, username, age, gender, email, phone_number, user_type, is_approved) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_, err := repo.db.Exec(query, user.UserID, user.Password, user.Name, user.Age, user.Gender, user.Email, user.PhoneNumber, user.UserType, 0)
	if user.UserType == "doctor" {
		fmt.Println("Your signup request has been submitted for approval.")

		_, err = repo.db.Exec("INSERT INTO notifications (user_id, content) VALUES (?, ?)",
			"admin", fmt.Sprintf("Please approve %s signup request for doctor role.", user.UserID))

		if err != nil {
			fmt.Println("Error sending doctor signup notification:", err)
		}
	} else {
		_, _ = repo.db.Exec("INSERT INTO patients (user_id, medical_history) VALUES (?,?)", user.UserID, "No History")
		fmt.Println("pat saved in patTab")
		fmt.Println("Signup successful. You can now log in.")
		_, _ = repo.db.Exec("INSERT INTO notifications (user_id, content) VALUES (?, ?)",
			user.UserID, fmt.Sprintf("welcome %s to the application.", user.UserID))
	}

	return err
}

// GetUserByID returns a user by their ID from the database
func (repo *UserRepositoryImpl) GetUserByID(userID string) (*models.User, error) {
	query := "SELECT user_id, password, username, age, gender, email, phone_number, user_type, is_approved FROM users WHERE user_id = ?"
	row := repo.db.QueryRow(query, userID)

	user := &models.User{}
	err := row.Scan(&user.UserID, &user.Password, &user.Name, &user.Age, &user.Gender, &user.Email, &user.PhoneNumber, &user.UserType, &user.IsApproved)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errors.New("user not found")
	}
	return user, err
}

func (repo *UserRepositoryImpl) UpdateName(userID string, name string) error {
	query := "UPDATE users SET Name = ? WHERE UserID = ?"
	_, err := repo.db.Exec(query, name, userID)
	return err
}

func (repo *UserRepositoryImpl) UpdateAge(userID string, age string) error {
	query := "UPDATE users SET Age = ? WHERE UserID = ?"
	_, err := repo.db.Exec(query, age, userID)
	return err
}

func (repo *UserRepositoryImpl) UpdateEmail(userID string, email string) error {
	query := "UPDATE users SET Email = ? WHERE UserID = ?"
	_, err := repo.db.Exec(query, email, userID)
	return err
}

func (repo *UserRepositoryImpl) UpdatePhoneNumber(userID string, phoneNumber string) error {
	query := "UPDATE users SET PhoneNumber = ? WHERE UserID = ?"
	_, err := repo.db.Exec(query, phoneNumber, userID)
	return err
}

func (repo *UserRepositoryImpl) UpdateGender(userID string, gender string) error {
	query := "UPDATE users SET Gender = ? WHERE UserID = ?"
	_, err := repo.db.Exec(query, gender, userID)
	return err
}

func (repo *UserRepositoryImpl) UpdatePassword(userID string, newPassword string) error {
	query := "UPDATE users SET Password = ? WHERE UserID = ?"
	_, err := repo.db.Exec(query, newPassword, userID)
	return err
}
