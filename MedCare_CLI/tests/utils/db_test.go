package utils

import (
	"testing"

	"doctor-patient-cli/tests/mockDB"
	"doctor-patient-cli/utils"
)

func TestInitDB(t *testing.T) {
	t.Run("Valid Connection", func(t *testing.T) {
		mockDB.MockInitDB(t)
		mockDB.Mock.ExpectPing().WillReturnError(nil)

		err := utils.DB.Ping()
		if err != nil {
			t.Errorf("InitDB() failed with error = %v", err)
		}

		// Ensure all expectations are met
		if err := mockDB.Mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})
}

func TestGetDB(t *testing.T) {
	t.Run("GetDB After InitDB", func(t *testing.T) {
		mockDB.MockInitDB(t)

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
		mockDB.MockInitDB(t) // Initialize the mock DB
		mockDB.Mock.ExpectClose()

		// Check if DB is set to nil
		if utils.CloseDB(); utils.DB != nil {
			t.Errorf("CloseDB() did not set DB to nil, got: %v", utils.DB)
		}

		// Ensure all expectations are met
		if err := mockDB.Mock.ExpectationsWereMet(); err != nil {
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
