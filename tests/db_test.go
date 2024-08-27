package tests

import (
	"fmt"
	"testing"

	"doctor-patient-cli/utils" // Adjust the import path according to your project structure
	"github.com/DATA-DOG/go-sqlmock"
)

var mock sqlmock.Sqlmock

// MockInitDB sets up the mocked database
func MockInitDB(t *testing.T) {
	var err error
	utils.DB, mock, err = sqlmock.New() // Set the mock database to the DB variable in the utils package
	if err != nil {
		t.Fatalf("failed to open sqlmock database: %v", err)
	}
	fmt.Println("yes it worked")
}

func TestInitDB(t *testing.T) {
	t.Run("Valid Connection", func(t *testing.T) {
		MockInitDB(t)
		mock.ExpectPing()

		err := utils.DB.Ping()
		if err != nil {
			t.Errorf("InitDB() failed with error = %v", err)
		}

		// Ensure all expectations are met
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})
}

func TestGetDB(t *testing.T) {
	t.Run("GetDB After InitDB", func(t *testing.T) {
		MockInitDB(t)

		if got := utils.GetDB(); got != utils.DB {
			t.Errorf("GetDB() = %v, want %v", got, utils.DB)
		}
	})

	t.Run("GetDB Without InitDB", func(t *testing.T) {
		utils.DB = nil // Simulate not initializing the DB
		if got := utils.GetDB(); got != nil {
			t.Errorf("GetDB() = %v, want nil", got)
		}
	})
}

func TestCloseDB(t *testing.T) {
	t.Run("CloseDB After InitDB", func(t *testing.T) {
		MockInitDB(t) // Initialize the mock DB
		mock.ExpectClose()

		// Check if DB is set to nil
		if utils.CloseDB(); utils.DB != nil {
			t.Errorf("CloseDB() did not set DB to nil, got: %v", utils.DB)
		}

		// Ensure all expectations are met
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})

	t.Run("CloseDB Without InitDB", func(t *testing.T) {
		utils.DB = nil // Simulate not initializing the DB
		utils.CloseDB()

		// Check if DB is still nil
		if utils.DB != nil {
			t.Errorf("CloseDB() did not set DB to nil")
		}
	})
}
