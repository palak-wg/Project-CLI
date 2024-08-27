package utils

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/fatih/color"
	_ "github.com/go-sql-driver/mysql"
)

// DB is a global variable that holds the database connection instance
var DB *sql.DB

// InitDB initiates the database connection
func InitDB() {
	var err error
	DB, err = sql.Open("mysql", "root:mysql@tcp(localhost:3306)/doctor_patient_db2")
	if err != nil {
		color.Red("Failed to connect to database: %v", err)
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		color.Red("Failed to ping database: %v", err)
		log.Fatal(err)
	}
}

// GetDB returns the current database connection instance
func GetDB() *sql.DB {
	return DB
}

// CloseDB closes the database connection if it is open
//func CloseDB() {
//	if DB != nil {
//		DB.Close()
//	}
//}

func CloseDB() {
	if DB != nil {
		err := DB.Close()
		if err != nil {
			fmt.Println("Failed to close database")
			log.Printf("Error closing DB: %v", err)
		}
		DB = nil // Set DB to nil after closing
	}
}
