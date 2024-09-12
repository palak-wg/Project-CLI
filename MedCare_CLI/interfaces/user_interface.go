package interfaces

import "doctor-patient-cli/models"

// UserRepository defines the methods related to user management
type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByID(userID string) (*models.User, error)
	//UpdateUser(user *models.User) error
	UpdateName(userID string, name string) error
	UpdateAge(userID string, age string) error
	UpdateEmail(userID string, email string) error
	UpdatePhoneNumber(userID string, phoneNumber string) error
	UpdateGender(userID string, gender string) error
	UpdatePassword(userID string, newPassword string) error
}

// UserService defines the methods available for user-related operations
type UserService interface {
	CreateUser(user *models.User) error
	GetUserByID(userID string) (*models.User, error)
	//UpdateUser(user *models.User) error
	UpdateName(userID string, name string) error
	UpdateAge(userID string, age string) error
	UpdateEmail(userID string, email string) error
	UpdatePhoneNumber(userID string, phoneNumber string) error
	UpdateGender(userID string, gender string) error
	UpdatePassword(userID string, newPassword string) error
}
