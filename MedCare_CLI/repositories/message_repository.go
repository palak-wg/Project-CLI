package repositories

import (
	"database/sql"
	"doctor-patient-cli/interfaces"
	"doctor-patient-cli/models"
	"fmt"
	"log"
)

type MessageRepositoryImpl struct {
	db *sql.DB
}

// NewMessageRepository creates a new MessageRepositoryImpl instance
func NewMessageRepository(db *sql.DB) interfaces.MessageRepository {
	return &MessageRepositoryImpl{db: db}
}

// SendMessage inserts a new message into the database
func (repo *MessageRepositoryImpl) SendMessage(fromID, toID, message string) error {
	_, err := repo.db.Exec("INSERT INTO messages (sender_id, receiver_id, message) VALUES (?, ?, ?)",
		fromID, toID, message)
	if err != nil {
		log.Printf("Repository: Error sending message from %v to %v: %v", fromID, toID, err)
		return err
	}

	_, err = repo.db.Exec("INSERT INTO notifications (user_id, content) VALUES (?, ?)",
		toID, fmt.Sprintf("You have a new message from patient %s: %s", fromID, message))
	if err != nil {
		return fmt.Errorf("error creating notification: %v", err)
	}
	return nil
}

// GetUnreadMessages retrieves unread messages for a specific user
func (repo *MessageRepositoryImpl) GetUnreadMessages(userID string) ([]models.Message, error) {
	rows, err := repo.db.Query("SELECT sender_id, message, timestamp FROM messages WHERE receiver_id = ? AND status = ?", userID, "pending")
	if err != nil {
		log.Printf("Repository: Error fetching unread messages for userID %d: %v", userID, err)
		return nil, err
	}
	defer rows.Close()

	var messages []models.Message
	for rows.Next() {
		var message models.Message
		_ = rows.Scan(&message.Sender, &message.Content, &message.Timestamp)
		messages = append(messages, message)
	}
	if len(messages) == 0 {
		return messages, nil
	}
	_, err = repo.db.Exec("UPDATE messages SET status = 'read' WHERE receiver_id = ? AND status = 'pending'", userID)
	if err != nil {
		return nil, err
	}
	return messages, nil
}

func (repo *MessageRepositoryImpl) GetUnreadMessagesById(patientID, doctorID string) ([]models.Message, error) {
	rows, err := repo.db.Query("SELECT message, timestamp FROM messages WHERE receiver_id = ? AND sender_id=? AND status = 'pending'", doctorID, patientID)
	if err != nil {
		log.Printf("Repository: Error fetching unread messages for userID %v: %v", doctorID, err)
		return nil, err
	}
	defer rows.Close()

	var messages []models.Message
	for rows.Next() {
		var message models.Message
		_ = rows.Scan(&message.Content, &message.Timestamp)
		messages = append(messages, message)
	}
	if len(messages) == 0 {
		return messages, nil
	}
	_, err = repo.db.Exec("UPDATE messages SET status = 'read' WHERE receiver_id = ? AND sender_id=? AND status = 'pending'", doctorID, patientID)
	if err != nil {
		return nil, err
	}
	return messages, nil
}

// RespondToPatient updates the message status and content as a response from the doctor
func (repo *MessageRepositoryImpl) RespondToPatient(doctorID, patientID, response string) error {
	_, err := repo.db.Exec("INSERT INTO messages (sender_id, receiver_id, message) VALUES (?, ?, ?)",
		doctorID, patientID, response)
	if err != nil {
		log.Printf("Repository: Error sending message from %v to %v: %v", doctorID, patientID, err)
		return err
	}

	_, err = repo.db.Exec("INSERT INTO notifications (user_id, content) VALUES (?, ?)",
		patientID, fmt.Sprintf("You have a new message from doctor %s: %s", doctorID, response))
	if err != nil {
		return fmt.Errorf("error creating notification: %v", err)
	}
	return nil
}
