package services

import (
	"doctor-patient-cli/models"
	"doctor-patient-cli/utils"
	"fmt"
	"time"
)

func SendMessageToDoctor(patientID, doctorID, message string) error {
	db := utils.GetDB()

	// Create a new message record
	_, err := db.Exec("INSERT INTO messages (sender_id, receiver_id, message) VALUES (?, ?, ?)",
		patientID, doctorID, message)
	if err != nil {
		return fmt.Errorf("error inserting message: %v", err)
	}

	// Create a notification for the doctor
	_, err = db.Exec("INSERT INTO notifications (user_id, content) VALUES (?, ?)",
		doctorID, fmt.Sprintf("You have a new message from patient %s: %s", patientID, message))
	if err != nil {
		return fmt.Errorf("error creating notification: %v", err)
	}

	return nil
}

func GetUnreadMessagesByUserID(patientID, doctorID string) ([]models.Message, error) {
	db := utils.GetDB()
	rows, err := db.Query("SELECT message, timestamp FROM messages WHERE receiver_id = ? AND sender_id=? AND status = 'pending'", doctorID, patientID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []models.Message
	for rows.Next() {
		var message models.Message
		if err = rows.Scan(&message.Content, &message.Timestamp); err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}

	// Return immediately if no messages found, or if there's a scan error
	if len(messages) == 0 {
		return messages, nil
	}

	// Update unread messages status to read
	if _, err = db.Exec("UPDATE messages SET status = 'read' WHERE receiver_id = ? AND sender_id=? AND status = 'pending'", doctorID, patientID); err != nil {
		return nil, err
	}

	return messages, nil
}

func GetUnreadMessage(doctorID string) ([]models.Message, error) {
	db := utils.GetDB()
	rows, err := db.Query("SELECT sender_id, message, timestamp FROM messages WHERE receiver_id = ? AND status = ?",
		doctorID, "pending")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []models.Message
	for rows.Next() {
		var message models.Message
		rows.Scan(&message.Sender, &message.Content, &message.Timestamp)
		messages = append(messages, message)
	}

	if len(messages) == 0 {
		return messages, nil
	}

	_, err = db.Exec("UPDATE messages SET status = 'read' WHERE receiver_id = ? AND status = 'pending'", doctorID)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

// RespondToPatientRequest allows a doctor to respond to a patient request.
func RespondToPatientRequest(doctorID, patientID, response string) error {
	db := utils.GetDB()

	// Update message with doctor's response
	_, err := db.Exec("INSERT INTO messages (sender_id, receiver_id, message) VALUES (?, ?, ?)",
		doctorID, patientID, response)
	if err != nil {
		return fmt.Errorf("error responding patient request: %v", err)
	}

	// Create a notification for the patient
	_, err = db.Exec("INSERT INTO notifications (user_id, content, timestamp) VALUES (?, ?, ?)",
		patientID, fmt.Sprintf("Doctor %s has responded to your request: %s", doctorID, response), time.Now())
	if err != nil {
		return fmt.Errorf("error creating notification: %v", err)
	}

	return nil
}

// SuggestPrescription allows a doctor to suggest a prescription to a patient.
func SuggestPrescription(doctorID, patientID, prescription string) error {
	db := utils.GetDB()

	// Create a new prescription record
	_, err := db.Exec("INSERT INTO messages (sender_id, receiver_id, message) VALUES (?, ?, ?)",
		doctorID, patientID, prescription)
	if err != nil {
		return fmt.Errorf("error inserting prescription: %v", err)
	}

	// Create a notification for the patient
	_, err = db.Exec("INSERT INTO notifications (user_id, content) VALUES (?, ?)",
		patientID, fmt.Sprintf("Doctor %s has suggested a prescription for you: %s", doctorID, prescription))
	if err != nil {
		return fmt.Errorf("error creating notification: %v", err)
	}

	return nil
}
