package interfaces

import "doctor-patient-cli/models"

// AdminRepository defines the methods related to admin operations
type AdminRepository interface {
	ApproveDoctorSignup(userID string) error
	GetPendingDoctorRequests() ([]models.Doctor, error)
	GetAllUsers() ([]models.User, error)
	CreateNotificationForUser(userID string, content string) error
}

// AdminService defines the methods available for admin-related operations
type AdminService interface {
	ApproveDoctorSignup(userID string) error
	GetPendingDoctorRequests() ([]models.Doctor, error)
	GetAllUsers() ([]models.User, error)
	CreateNotificationForUser(userID string, content string) error
}
