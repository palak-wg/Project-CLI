package repositories_test

import (
	"doctor-patient-cli/models"
	"doctor-patient-cli/repositories"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSendMessage_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repositories.NewMessageRepository(db)

	fromID := "patient123"
	toID := "doctor123"
	message := "I have a question about my prescription."

	// Mocking the SQL queries
	mock.ExpectExec("INSERT INTO messages").WithArgs(fromID, toID, message).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO notifications").WithArgs(toID, "You have a new message from patient "+fromID+": "+message).WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.SendMessage(fromID, toID, message)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSendMessage_Failure_MessageInsert(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repositories.NewMessageRepository(db)

	fromID := "patient123"
	toID := "doctor123"
	message := "I have a question about my prescription."

	// Simulate error in message insert
	mock.ExpectExec("INSERT INTO messages").WithArgs(fromID, toID, message).WillReturnError(errors.New("insert error"))

	err = repo.SendMessage(fromID, toID, message)
	assert.EqualError(t, err, "insert error")
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSendMessage_Failure_NotificationInsert(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repositories.NewMessageRepository(db)

	fromID := "patient123"
	toID := "doctor123"
	message := "I have a question about my prescription."

	// Mock successful message insert, but notification fails
	mock.ExpectExec("INSERT INTO messages").WithArgs(fromID, toID, message).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO notifications").WithArgs(toID, "You have a new message from patient "+fromID+": "+message).WillReturnError(errors.New("notification error"))

	err = repo.SendMessage(fromID, toID, message)
	assert.EqualError(t, err, "error creating notification: notification error")
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUnreadMessages_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repositories.NewMessageRepository(db)

	userID := "user123"
	timestamp := []uint8("09-09-2024")
	expectedMessages := []models.Message{
		{Sender: "sender1", Content: "Message 1", Timestamp: timestamp},
		{Sender: "sender2", Content: "Message 2", Timestamp: timestamp},
	}

	// Mocking the SQL query to return unread messages
	rows := sqlmock.NewRows([]string{"sender_id", "message", "timestamp"}).
		AddRow(expectedMessages[0].Sender, expectedMessages[0].Content, expectedMessages[0].Timestamp).
		AddRow(expectedMessages[1].Sender, expectedMessages[1].Content, expectedMessages[1].Timestamp)
	mock.ExpectQuery("^SELECT sender_id, message, timestamp FROM messages WHERE receiver_id = \\? AND status = \\?$").
		WithArgs(userID, "pending").WillReturnRows(rows)

	// Mocking the update query to mark the messages as read
	mock.ExpectExec("^UPDATE messages SET status = 'read' WHERE receiver_id = \\? AND status = 'pending'$").
		WithArgs(userID).WillReturnResult(sqlmock.NewResult(1, 2))

	messages, err := repo.GetUnreadMessages(userID)
	assert.NoError(t, err)
	assert.Equal(t, expectedMessages, messages)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUnreadMessages_EmptyResult(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repositories.NewMessageRepository(db)

	userID := "user123"

	// Simulate no unread messages
	rows := sqlmock.NewRows([]string{"sender_id", "message", "timestamp"})
	mock.ExpectQuery("^SELECT sender_id, message, timestamp FROM messages WHERE receiver_id = \\? AND status = \\?$").
		WithArgs(userID, "pending").WillReturnRows(rows)

	messages, err := repo.GetUnreadMessages(userID)
	assert.NoError(t, err)
	assert.Empty(t, messages)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUnreadMessages_ErrorFetching(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repositories.NewMessageRepository(db)

	userID := "user123"

	// Simulate an error while fetching messages
	mock.ExpectQuery("^SELECT sender_id, message, timestamp FROM messages WHERE receiver_id = \\? AND status = \\?$").
		WithArgs(userID, "pending").WillReturnError(errors.New("query error"))

	messages, err := repo.GetUnreadMessages(userID)
	assert.Error(t, err)
	assert.Nil(t, messages)
	assert.EqualError(t, err, "query error")
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUnreadMessagesById_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repositories.NewMessageRepository(db)

	patientID := "patient123"
	doctorID := "doctor123"
	timestamp := []uint8("09-09-2024")
	expectedMessages := []models.Message{
		{Content: "Message 1", Timestamp: timestamp},
		{Content: "Message 2", Timestamp: timestamp},
	}

	// Mocking the SQL query to return unread messages between a specific patient and doctor
	rows := sqlmock.NewRows([]string{"message", "timestamp"}).
		AddRow(expectedMessages[0].Content, expectedMessages[0].Timestamp).
		AddRow(expectedMessages[1].Content, expectedMessages[1].Timestamp)
	mock.ExpectQuery("^SELECT message, timestamp FROM messages WHERE receiver_id = \\? AND sender_id=\\? AND status = 'pending'$").
		WithArgs(doctorID, patientID).WillReturnRows(rows)

	// Mocking the update query to mark the messages as read
	mock.ExpectExec("^UPDATE messages SET status = 'read' WHERE receiver_id = \\? AND sender_id=\\? AND status = 'pending'$").
		WithArgs(doctorID, patientID).WillReturnResult(sqlmock.NewResult(1, 2))

	messages, err := repo.GetUnreadMessagesById(patientID, doctorID)
	assert.NoError(t, err)
	assert.Equal(t, expectedMessages, messages)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUnreadMessagesById_EmptyResult(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repositories.NewMessageRepository(db)

	patientID := "patient123"
	doctorID := "doctor123"

	// Simulate no unread messages between the patient and doctor
	rows := sqlmock.NewRows([]string{"message", "timestamp"})
	mock.ExpectQuery("^SELECT message, timestamp FROM messages WHERE receiver_id = \\? AND sender_id=\\? AND status = 'pending'$").
		WithArgs(doctorID, patientID).WillReturnRows(rows)

	messages, err := repo.GetUnreadMessagesById(patientID, doctorID)
	assert.NoError(t, err)
	assert.Empty(t, messages)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUnreadMessagesById_ErrorFetching(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repositories.NewMessageRepository(db)

	patientID := "patient123"
	doctorID := "doctor123"

	// Simulate an error while fetching messages
	mock.ExpectQuery("^SELECT message, timestamp FROM messages WHERE receiver_id = \\? AND sender_id=\\? AND status = 'pending'$").
		WithArgs(doctorID, patientID).WillReturnError(errors.New("query error"))

	messages, err := repo.GetUnreadMessagesById(patientID, doctorID)
	assert.Error(t, err)
	assert.Nil(t, messages)
	assert.EqualError(t, err, "query error")
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUnreadMessages_UpdateQueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repositories.NewMessageRepository(db)

	userID := "user123"
	timestamp := []uint8("09-09-2024")
	expectedMessages := []models.Message{
		{Sender: "sender1", Content: "Message 1", Timestamp: timestamp},
	}

	// Mocking the SQL query to return unread messages
	rows := sqlmock.NewRows([]string{"sender_id", "message", "timestamp"}).
		AddRow(expectedMessages[0].Sender, expectedMessages[0].Content, expectedMessages[0].Timestamp)
	mock.ExpectQuery("^SELECT sender_id, message, timestamp FROM messages WHERE receiver_id = \\? AND status = \\?$").
		WithArgs(userID, "pending").WillReturnRows(rows)

	// Mocking the update query to fail
	mock.ExpectExec("^UPDATE messages SET status = 'read' WHERE receiver_id = \\? AND status = 'pending'$").
		WithArgs(userID).WillReturnError(errors.New("update query error"))

	messages, err := repo.GetUnreadMessages(userID)
	assert.Error(t, err)
	assert.Nil(t, messages)
	assert.EqualError(t, err, "update query error")
	assert.NoError(t, mock.ExpectationsWereMet())
}

//
//func TestGetUnreadMessages_QueryError(t *testing.T) {
//	db, mock, err := sqlmock.New()
//	assert.NoError(t, err)
//	defer db.Close()
//
//	repo := repositories.NewMessageRepository(db)
//
//	userID := "doctor123"
//
//	// Simulate query error
//	mock.ExpectQuery("^SELECT sender_id, message, timestamp FROM messages WHERE receiver_id = \\? AND status = \\?$").
//		WithArgs(userID, "pending").WillReturnError(errors.New("query error"))
//
//	messages, err := repo.GetUnreadMessages(userID)
//	assert.EqualError(t, err, "query error")
//	assert.Nil(t, messages)
//	assert.NoError(t, mock.ExpectationsWereMet())
//}

func TestGetUnreadMessagesById_UpdateQueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repositories.NewMessageRepository(db)

	patientID := "patient123"
	doctorID := "doctor123"
	timestamp := []uint8("09-09-2024")
	expectedMessages := []models.Message{
		{Content: "Message 1", Timestamp: timestamp},
	}

	// Mocking the SQL query to return unread messages
	rows := sqlmock.NewRows([]string{"message", "timestamp"}).
		AddRow(expectedMessages[0].Content, expectedMessages[0].Timestamp)
	mock.ExpectQuery("^SELECT message, timestamp FROM messages WHERE receiver_id = \\? AND sender_id=\\? AND status = 'pending'$").
		WithArgs(doctorID, patientID).WillReturnRows(rows)

	// Mocking the update query to fail
	mock.ExpectExec("^UPDATE messages SET status = 'read' WHERE receiver_id = \\? AND sender_id=\\? AND status = 'pending'$").
		WithArgs(doctorID, patientID).WillReturnError(errors.New("update query error"))

	messages, err := repo.GetUnreadMessagesById(patientID, doctorID)
	assert.Error(t, err)
	assert.Nil(t, messages)
	assert.EqualError(t, err, "update query error")
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRespondToPatient_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repositories.NewMessageRepository(db)

	doctorID := "doctor123"
	patientID := "patient123"
	response := "Your prescription has been updated."

	// Mocking the SQL queries
	mock.ExpectExec("INSERT INTO messages").WithArgs(doctorID, patientID, response).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO notifications").WithArgs(patientID, "You have a new message from doctor "+doctorID+": "+response).WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.RespondToPatient(doctorID, patientID, response)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRespondToPatient_MessageInsertFailure(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repositories.NewMessageRepository(db)

	doctorID := "doctor123"
	patientID := "patient123"
	response := "Your prescription has been updated."

	// Simulate message insert failure
	mock.ExpectExec("INSERT INTO messages").WithArgs(doctorID, patientID, response).WillReturnError(errors.New("message insert error"))

	err = repo.RespondToPatient(doctorID, patientID, response)
	assert.EqualError(t, err, "message insert error")
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRespondToPatient_NotificationFailure(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repositories.NewMessageRepository(db)

	doctorID := "doctor123"
	patientID := "patient123"
	response := "Your prescription has been updated."

	// Mock successful message insert, but notification fails
	mock.ExpectExec("INSERT INTO messages").WithArgs(doctorID, patientID, response).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO notifications").WithArgs(patientID, "You have a new message from doctor "+doctorID+": "+response).WillReturnError(errors.New("notification error"))

	err = repo.RespondToPatient(doctorID, patientID, response)
	assert.EqualError(t, err, "error creating notification: notification error")
	assert.NoError(t, mock.ExpectationsWereMet())
}
