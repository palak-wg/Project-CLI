package services

import (
	"doctor-patient-cli/interfaces"
	"doctor-patient-cli/models"
	"log"
)

type MessageService struct {
	repo interfaces.MessageRepository
}

// NewMessageService creates a new MessageService instance
func NewMessageService(repo interfaces.MessageRepository) *MessageService {
	return &MessageService{repo: repo}
}

// SendMessage sends a new message from one user to another
func (service *MessageService) SendMessage(fromID, toID, message string) error {
	err := service.repo.SendMessage(fromID, toID, message)
	if err != nil {
		log.Printf("Service: Error sending message from %v to %v: %v", fromID, toID, err)
		return err
	}
	return nil
}

// GetUnreadMessages retrieves unread messages for a specific user
func (service *MessageService) GetUnreadMessages(userID string) ([]models.Message, error) {
	messages, err := service.repo.GetUnreadMessages(userID)
	if err != nil {
		log.Printf("Service: Error retrieving unread messages for userID %v: %v", userID, err)
		return nil, err
	}
	return messages, nil
}

// RespondToPatient updates the message status and content as a response from the doctor
func (service *MessageService) RespondToPatient(doctorID, patientID, response string) error {
	err := service.repo.RespondToPatient(doctorID, patientID, response)
	if err != nil {
		log.Printf("Service: Error responding to patientID %v from doctorID %v: %v", patientID, doctorID, err)
		return err
	}
	return nil
}

func (service *MessageService) GetUnreadMessagesById(patientID, doctorID string) ([]models.Message, error) {
	messages, err := service.repo.GetUnreadMessagesById(patientID, doctorID)
	if err != nil {
		log.Printf("Service: Error retrieving unread messages for userID %v: %v", doctorID, err)
		return nil, err
	}
	return messages, nil
}
