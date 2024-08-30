package mockDB

import (
	"doctor-patient-cli/utils"
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
)

var Mock sqlmock.Sqlmock

// MockInitDB sets up the mocked database
func MockInitDB(t *testing.T) {
	utils.DB, Mock, _ = sqlmock.New(sqlmock.MonitorPingsOption(true)) // Set the mock database to the DB variable in the utils package
}
