package services

import (
	"doctor-patient-cli/tests/mockDB"
	"fmt"
	"regexp"
	"testing"
	"time"

	"doctor-patient-cli/services"
	"doctor-patient-cli/utils"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestSendMessageToDoctor(t *testing.T) {
	mockDB.MockInitDB(t)
	defer utils.CloseDB()

	t.Run("SendMessageToDoctor Success", func(t *testing.T) {
		mockDB.Mock.ExpectExec(regexp.QuoteMeta("INSERT INTO messages (sender_id, receiver_id, message) VALUES (?, ?, ?)")).
			WithArgs("patient1", "doctor1", "Hello Doctor").
			WillReturnResult(sqlmock.NewResult(1, 1))

		mockDB.Mock.ExpectExec(regexp.QuoteMeta("INSERT INTO notifications (user_id, content) VALUES (?, ?)")).
			WithArgs("doctor1", "You have a new message from patient patient1: Hello Doctor").
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := services.SendMessageToDoctor("patient1", "doctor1", "Hello Doctor")
		assert.NoError(t, err)

		assert.NoError(t, mockDB.Mock.ExpectationsWereMet())
	})

	t.Run("SendMessageToDoctor Errors", func(t *testing.T) {
		// Scenario 1. Error during Insert Query for message
		// Setup expectation for the INSERT INTO messages query to fail
		mockDB.Mock.ExpectExec(regexp.QuoteMeta("INSERT INTO messages (sender_id, receiver_id, message) VALUES (?, ?, ?)")).
			WithArgs("patient1", "doctor1", "Hello Doctor").
			WillReturnError(fmt.Errorf("database error inserting message"))

		// Call the function under test
		err := services.SendMessageToDoctor("patient1", "doctor1", "Hello Doctor")

		// Validate results
		assert.Error(t, err, "Expected an error but got none")
		assert.Equal(t, "error inserting message: database error inserting message", err.Error(), "Error message does not match")

		// Ensure all expectations are met
		assert.NoError(t, mockDB.Mock.ExpectationsWereMet())

		// Scenario 2. Error during Insertion of message
		// Setup expectation for the INSERT INTO messages query to succeed
		mockDB.Mock.ExpectExec(regexp.QuoteMeta("INSERT INTO messages (sender_id, receiver_id, message) VALUES (?, ?, ?)")).
			WithArgs("patient1", "doctor1", "Hello Doctor").
			WillReturnResult(sqlmock.NewResult(1, 1)) // Mock successful insertion with one row affected

		// Setup expectation for the INSERT INTO notifications query to fail
		mockDB.Mock.ExpectExec(regexp.QuoteMeta("INSERT INTO notifications (user_id, content) VALUES (?, ?)")).
			WithArgs("doctor1", "You have a new message from patient patient1: Hello Doctor").
			WillReturnError(fmt.Errorf("database error creating notification"))

		// Call the function under test
		err = services.SendMessageToDoctor("patient1", "doctor1", "Hello Doctor")

		// Validate results
		assert.Error(t, err, "Expected an error but got none")
		assert.Equal(t, "error creating notification: database error creating notification", err.Error(), "Error message does not match")

		// Ensure all expectations are met
		assert.NoError(t, mockDB.Mock.ExpectationsWereMet())
	})
}

func TestGetUnreadMessagesByUserID(t *testing.T) {
	mockDB.MockInitDB(t)
	defer utils.CloseDB()

	t.Run("GetUnreadMessagesByUserID Success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"message", "timestamp"}).
			AddRow("Hello Doctor", time.Now())

		mockDB.Mock.ExpectQuery(regexp.QuoteMeta("SELECT message, timestamp FROM messages WHERE receiver_id = ? AND sender_id=? AND status = 'pending'")).
			WithArgs("doctor1", "patient1").
			WillReturnRows(rows)

		mockDB.Mock.ExpectExec(regexp.QuoteMeta("UPDATE messages SET status = 'read' WHERE receiver_id = ? AND sender_id=? AND status = 'pending'")).
			WithArgs("doctor1", "patient1").
			WillReturnResult(sqlmock.NewResult(0, 1))

		messages, err := services.GetUnreadMessagesByUserID("patient1", "doctor1")
		assert.NoError(t, err)
		assert.NotEmpty(t, messages)

		assert.NoError(t, mockDB.Mock.ExpectationsWereMet())
	})

	t.Run("GetUnreadMessagesByUserID Errors", func(t *testing.T) {
		// Scenario 1: Error during the SELECT query
		mockDB.Mock.ExpectQuery(regexp.QuoteMeta("SELECT message, timestamp FROM messages WHERE receiver_id = ? AND sender_id=? AND status = 'pending'")).
			WithArgs("doctor1", "patient1").
			WillReturnError(fmt.Errorf("query error"))

		messages, err := services.GetUnreadMessagesByUserID("patient1", "doctor1")
		assert.Error(t, err, "Expected query error but got none")
		assert.Nil(t, messages, "Expected messages to be nil due to query error")

		// Scenario 2: No messages found (covers len(messages) == 0)
		emptyRows := sqlmock.NewRows([]string{"message", "timestamp"}) // No rows added

		mockDB.Mock.ExpectQuery(regexp.QuoteMeta("SELECT message, timestamp FROM messages WHERE receiver_id = ? AND sender_id=? AND status = 'pending'")).
			WithArgs("doctor1", "patient1").
			WillReturnRows(emptyRows)

		messages, err = services.GetUnreadMessagesByUserID("patient1", "doctor1")
		assert.NoError(t, err, "Expected no error but got one")
		assert.Empty(t, messages, "Expected messages to be empty since no rows were returned")

		// Scenario 3: Error during the Scan operation
		rowsWithError := sqlmock.NewRows([]string{"message", "timestamp"}).
			AddRow(nil, time.Now()) // Invalid data to cause scan error

		mockDB.Mock.ExpectQuery(regexp.QuoteMeta("SELECT message, timestamp FROM messages WHERE receiver_id = ? AND sender_id=? AND status = 'pending'")).
			WithArgs("doctor1", "patient1").
			WillReturnRows(rowsWithError)

		messages, err = services.GetUnreadMessagesByUserID("patient1", "doctor1")
		assert.Error(t, err, "Expected scan error but got none")
		assert.Nil(t, messages, "Expected messages to be nil due to scan error")

		// Scenario 4: Error during the UPDATE operation
		validRows := sqlmock.NewRows([]string{"message", "timestamp"}).
			AddRow("Hello Doctor", time.Now())

		mockDB.Mock.ExpectQuery(regexp.QuoteMeta("SELECT message, timestamp FROM messages WHERE receiver_id = ? AND sender_id=? AND status = 'pending'")).
			WithArgs("doctor1", "patient1").
			WillReturnRows(validRows)

		mockDB.Mock.ExpectExec(regexp.QuoteMeta("UPDATE messages SET status = 'read' WHERE receiver_id = ? AND sender_id=? AND status = 'pending'")).
			WithArgs("doctor1", "patient1").
			WillReturnError(fmt.Errorf("update error"))

		messages, err = services.GetUnreadMessagesByUserID("patient1", "doctor1")
		assert.Error(t, err, "Expected update error but got none")
		assert.Nil(t, messages, "Expected messages to be nil due to update error")

		// Ensure all expectations are met
		assert.NoError(t, mockDB.Mock.ExpectationsWereMet())
	})
}

func TestGetUnreadMessage(t *testing.T) {
	mockDB.MockInitDB(t)
	defer utils.CloseDB()

	t.Run("GetUnreadMessage Success", func(t *testing.T) {

		// Define the expected rows to be returned by the query
		rows := sqlmock.NewRows([]string{"sender_id", "message", "timestamp"}).
			AddRow("patient1", "Hello Doctor", time.Now())

		// Setup expectation for the SELECT query
		mockDB.Mock.ExpectQuery(regexp.QuoteMeta("SELECT sender_id, message, timestamp FROM messages WHERE receiver_id = ? AND status = ?")).
			WithArgs("doctor1", "pending").
			WillReturnRows(rows)

		// Setup expectation for the UPDATE query
		mockDB.Mock.ExpectExec(regexp.QuoteMeta("UPDATE messages SET status = 'read' WHERE receiver_id = ? AND status = 'pending'")).
			WithArgs("doctor1").
			WillReturnResult(sqlmock.NewResult(0, 1)) // Mock successful execution with one row affected

		// Call the function under test
		messages, err := services.GetUnreadMessage("doctor1")

		// Validate results
		assert.NoError(t, err, "Expected no error but got one")
		assert.Len(t, messages, 1, "Expected one message to be returned")
		assert.Equal(t, "patient1", messages[0].Sender, "Expected sender to be 'patient1'")
		assert.Equal(t, "Hello Doctor", messages[0].Content, "Expected message content to be 'Hello Doctor'")

		// Ensure all expectations are met
		assert.NoError(t, mockDB.Mock.ExpectationsWereMet())
	})

	t.Run("GetUnreadMessage Errors", func(t *testing.T) {
		// Scenario 1: Error during the SELECT query
		mockDB.Mock.ExpectQuery(regexp.QuoteMeta("SELECT sender_id, message, timestamp FROM messages WHERE receiver_id = ? AND status = ?")).
			WithArgs("doctor1", "pending").
			WillReturnError(fmt.Errorf("query error"))

		messages, err := services.GetUnreadMessage("doctor1")
		assert.Error(t, err, "Expected query error but got none")
		assert.Nil(t, messages, "Expected messages to be nil due to query error")

		// Scenario 2: No messages found
		emptyRows := sqlmock.NewRows([]string{"sender_id", "message", "timestamp"}) // No rows added

		mockDB.Mock.ExpectQuery(regexp.QuoteMeta("SELECT sender_id, message, timestamp FROM messages WHERE receiver_id = ? AND status = ?")).
			WithArgs("doctor1", "pending").
			WillReturnRows(emptyRows)

		messages, err = services.GetUnreadMessage("doctor1")
		assert.NoError(t, err, "Expected no error but got one")
		assert.Empty(t, messages, "Expected messages to be empty since no rows were returned")

		// Scenario 3: Error during UPDATE operation
		validRows := sqlmock.NewRows([]string{"sender_id", "message", "timestamp"}).
			AddRow("patient1", "Hello Doctor", time.Now())

		mockDB.Mock.ExpectQuery(regexp.QuoteMeta("SELECT sender_id, message, timestamp FROM messages WHERE receiver_id = ? AND status = ?")).
			WithArgs("doctor1", "pending").
			WillReturnRows(validRows)

		mockDB.Mock.ExpectExec(regexp.QuoteMeta("UPDATE messages SET status = 'read' WHERE receiver_id = ? AND status = 'pending'")).
			WithArgs("doctor1").
			WillReturnError(fmt.Errorf("update error"))

		messages, err = services.GetUnreadMessage("doctor1")
		assert.Error(t, err, "Expected update error but got none")
		assert.Nil(t, messages, "Expected messages to be nil due to update error")
	})
}

func TestRespondToPatientRequest(t *testing.T) {
	mockDB.MockInitDB(t)
	defer utils.CloseDB()

	t.Run("RespondToPatientRequest Success", func(t *testing.T) {
		mockDB.Mock.ExpectExec(regexp.QuoteMeta("INSERT INTO messages (sender_id, receiver_id, message) VALUES (?, ?, ?)")).
			WithArgs("doctor1", "patient1", "Response to your request").
			WillReturnResult(sqlmock.NewResult(1, 1))

		mockDB.Mock.ExpectExec(regexp.QuoteMeta("INSERT INTO notifications (user_id, content, timestamp) VALUES (?, ?, ?)")).
			WithArgs("patient1", "Doctor doctor1 has responded to your request: Response to your request", sqlmock.AnyArg()).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := services.RespondToPatientRequest("doctor1", "patient1", "Response to your request")
		assert.NoError(t, err)

		assert.NoError(t, mockDB.Mock.ExpectationsWereMet())
	})

	t.Run("RespondToPatientRequest Errors", func(t *testing.T) {
		// Scenario 1. Error in INSERT INTO messages
		// Setup expectation for the INSERT INTO messages query to fail
		mockDB.Mock.ExpectExec(regexp.QuoteMeta("INSERT INTO messages (sender_id, receiver_id, message) VALUES (?, ?, ?)")).
			WithArgs("doctor1", "patient1", "Here is my response").
			WillReturnError(fmt.Errorf("database error inserting message"))

		// Call the function under test
		err := services.RespondToPatientRequest("doctor1", "patient1", "Here is my response")

		// Validate results
		assert.Error(t, err, "Expected an error but got none")
		assert.Equal(t, "error responding patient request: database error inserting message", err.Error(), "Error message does not match")

		// Ensure all expectations are met
		assert.NoError(t, mockDB.Mock.ExpectationsWereMet())

		// Scenario 2. Error in INSERT INTO notifications
		// Setup expectation for the INSERT INTO messages query to succeed
		mockDB.Mock.ExpectExec(regexp.QuoteMeta("INSERT INTO messages (sender_id, receiver_id, message) VALUES (?, ?, ?)")).
			WithArgs("doctor1", "patient1", "Here is my response").
			WillReturnResult(sqlmock.NewResult(1, 1)) // Mock successful insertion with one row affected

		// Setup expectation for the INSERT INTO notifications query to fail
		mockDB.Mock.ExpectExec(regexp.QuoteMeta("INSERT INTO notifications (user_id, content, timestamp) VALUES (?, ?, ?)")).
			WithArgs("patient1", "Doctor doctor1 has responded to your request: Here is my response", sqlmock.AnyArg()).
			WillReturnError(fmt.Errorf("database error creating notification"))

		// Call the function under test
		err = services.RespondToPatientRequest("doctor1", "patient1", "Here is my response")

		// Validate results
		assert.Error(t, err, "Expected an error but got none")
		assert.Equal(t, "error creating notification: database error creating notification", err.Error(), "Error message does not match")

		// Ensure all expectations are met
		assert.NoError(t, mockDB.Mock.ExpectationsWereMet())
	})

}

func TestSuggestPrescription(t *testing.T) {
	mockDB.MockInitDB(t)
	defer utils.CloseDB()

	t.Run("SuggestPrescription Success", func(t *testing.T) {
		mockDB.Mock.ExpectExec(regexp.QuoteMeta("INSERT INTO messages (sender_id, receiver_id, message) VALUES (?, ?, ?)")).
			WithArgs("doctor1", "patient1", "Take 2 pills daily").
			WillReturnResult(sqlmock.NewResult(1, 1))

		mockDB.Mock.ExpectExec(regexp.QuoteMeta("INSERT INTO notifications (user_id, content) VALUES (?, ?)")).
			WithArgs("patient1", "Doctor doctor1 has suggested a prescription for you: Take 2 pills daily").
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := services.SuggestPrescription("doctor1", "patient1", "Take 2 pills daily")
		assert.NoError(t, err)

		assert.NoError(t, mockDB.Mock.ExpectationsWereMet())
	})

	t.Run("SuggestPrescription Errors", func(t *testing.T) {
		// Scenario 1. Error in INSERT INTO messages
		// Setup expectation for the INSERT INTO messages query to fail
		mockDB.Mock.ExpectExec(regexp.QuoteMeta("INSERT INTO messages (sender_id, receiver_id, message) VALUES (?, ?, ?)")).
			WithArgs("doctor1", "patient1", "Prescription details").
			WillReturnError(fmt.Errorf("database error inserting prescription"))

		// Call the function under test
		err := services.SuggestPrescription("doctor1", "patient1", "Prescription details")

		// Validate results
		assert.Error(t, err, "Expected an error but got none")
		assert.Equal(t, "error inserting prescription: database error inserting prescription", err.Error(), "Error message does not match")

		// Ensure all expectations are met
		assert.NoError(t, mockDB.Mock.ExpectationsWereMet())

		// Scenario 2. Error in INSERT INTO notifications
		// Setup expectation for the INSERT INTO messages query to succeed
		mockDB.Mock.ExpectExec(regexp.QuoteMeta("INSERT INTO messages (sender_id, receiver_id, message) VALUES (?, ?, ?)")).
			WithArgs("doctor1", "patient1", "Prescription details").
			WillReturnResult(sqlmock.NewResult(1, 1)) // Mock successful insertion with one row affected

		// Setup expectation for the INSERT INTO notifications query to fail
		mockDB.Mock.ExpectExec(regexp.QuoteMeta("INSERT INTO notifications (user_id, content) VALUES (?, ?)")).
			WithArgs("patient1", "Doctor doctor1 has suggested a prescription for you: Prescription details").
			WillReturnError(fmt.Errorf("database error creating notification"))

		// Call the function under test
		err = services.SuggestPrescription("doctor1", "patient1", "Prescription details")

		// Validate results
		assert.Error(t, err, "Expected an error but got none")
		assert.Equal(t, "error creating notification: database error creating notification", err.Error(), "Error message does not match")

		// Ensure all expectations are met
		assert.NoError(t, mockDB.Mock.ExpectationsWereMet())
	})
}
