package utils

import (
	"database/sql"
	"log"

	"github.com/fatih/color"
	_ "github.com/go-sql-driver/mysql"
)

// db is a global variable that holds the database connection instance
var db *sql.DB

// InitDB initiates database connection
func InitDB() {
	var err error
	db, err = sql.Open("mysql", "root:mysql@tcp(localhost:3306)/doctor_patient_db2")
	if err != nil {
		color.Red("Failed to connect to database: %v", err)
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		color.Red("Failed to ping database: %v", err)
		log.Fatal(err)
	}
}

// GetDB returns the current database connection instance
func GetDB() *sql.DB {
	return db
}

// CloseDB closes the database connection if it is open
func CloseDB() {
	if db != nil {
		db.Close()
	}
}
