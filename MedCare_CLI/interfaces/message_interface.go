package interfaces

import "doctor-patient-cli/models"

// MessageRepository defines the methods related to messaging between users
type MessageRepository interface {
	SendMessage(fromID, toID, message string) error
	GetUnreadMessages(userID string) ([]models.Message, error)
	RespondToPatient(doctorID, patientID, response string) error
	GetUnreadMessagesById(patientID, doctorID string) ([]models.Message, error)
}

// MessageService defines the methods for message-related operations
type MessageService interface {
	SendMessage(fromID, toID, message string) error
	GetUnreadMessages(userID string) ([]models.Message, error)
	RespondToPatient(doctorID, patientID, response string) error
	GetUnreadMessagesById(patientID, doctorID string) ([]models.Message, error)
}
