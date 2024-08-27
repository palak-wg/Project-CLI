package tests

import (
	"regexp"
	"testing"
	"time"

	"doctor-patient-cli/models"
	"doctor-patient-cli/utils"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestSendMessageToDoctor(t *testing.T) {
	MockInitDB(t)
	defer utils.CloseDB()

	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO messages (sender_id, receiver_id, message) VALUES (?, ?, ?)")).
		WithArgs("patient1", "doctor1", "Hello Doctor").
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO notifications (user_id, content) VALUES (?, ?)")).
		WithArgs("doctor1", "You have a new message from patient patient1: Hello Doctor").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := models.SendMessageToDoctor("patient1", "doctor1", "Hello Doctor")
	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUnreadMessagesByUserID(t *testing.T) {
	MockInitDB(t)
	defer utils.CloseDB()

	rows := sqlmock.NewRows([]string{"message", "timestamp"}).
		AddRow("Hello Doctor", time.Now())

	mock.ExpectQuery(regexp.QuoteMeta("SELECT message, timestamp FROM messages WHERE receiver_id = ? AND sender_id=? AND status = 'pending'")).
		WithArgs("doctor1", "patient1").
		WillReturnRows(rows)

	mock.ExpectExec(regexp.QuoteMeta("UPDATE messages SET status = 'read' WHERE receiver_id = ? AND sender_id=? AND status = 'pending'")).
		WithArgs("doctor1", "patient1").
		WillReturnResult(sqlmock.NewResult(0, 1))

	messages, err := models.GetUnreadMessagesByUserID("patient1", "doctor1")
	assert.NoError(t, err)
	assert.NotEmpty(t, messages)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUnreadMessage(t *testing.T) {
	MockInitDB(t)
	defer utils.CloseDB()

	rows := sqlmock.NewRows([]string{"sender_id", "message", "timestamp"}).
		AddRow("patient1", "Hello Doctor", time.Now())

	mock.ExpectQuery(regexp.QuoteMeta("SELECT sender_id, message, timestamp FROM messages WHERE (receiver_id = ? AND status = ?)")).
		WithArgs("doctor1", "pending").
		WillReturnRows(rows)

	mock.ExpectExec(regexp.QuoteMeta("UPDATE messages SET status = 'read' WHERE receiver_id = ? AND status = 'pending'")).
		WithArgs("doctor1").
		WillReturnResult(sqlmock.NewResult(0, 1))

	messages, err := models.GetUnreadMessage("doctor1")
	assert.NoError(t, err)
	assert.NotEmpty(t, messages)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRespondToPatientRequest(t *testing.T) {
	MockInitDB(t)
	defer utils.CloseDB()

	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO messages (sender_id, receiver_id, message) VALUES (?, ?, ?)")).
		WithArgs("doctor1", "patient1", "Response to your request").
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO notifications (user_id, content, timestamp) VALUES (?, ?, ?)")).
		WithArgs("patient1", "Doctor doctor1 has responded to your request: Response to your request", sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := models.RespondToPatientRequest("doctor1", "patient1", "Response to your request")
	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSuggestPrescription(t *testing.T) {
	MockInitDB(t)
	defer utils.CloseDB()

	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO messages (sender_id, receiver_id, message) VALUES (?, ?, ?)")).
		WithArgs("doctor1", "patient1", "Take 2 pills daily").
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO notifications (user_id, content) VALUES (?, ?)")).
		WithArgs("patient1", "Doctor doctor1 has suggested a prescription for you: Take 2 pills daily").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := models.SuggestPrescription("doctor1", "patient1", "Take 2 pills daily")
	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}
