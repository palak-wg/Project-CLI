package utils

import (
	"database/sql"
	"fmt"
	"github.com/fatih/color"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
	"sync"
)

// DB is a global variable that holds the database connection instance
var DB *sql.DB
var once sync.Once

// InitDB initiates the database connection
func InitDB() {
	var err error

	err = godotenv.Load("config.env")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DBUser"),
		os.Getenv("DBPassword"),
		os.Getenv("DBHost"),
		os.Getenv("DBPort"),
		os.Getenv("DBName"),
	)
	DB, err = sql.Open("mysql", dsn)
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

//func GetSQLClient() *sql.DB {
//	once.Do(func() {
//		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
//			os.Getenv("DBUser"),
//			os.Getenv("DBPassword"),
//			os.Getenv("DBHost"),
//			os.Getenv("DBPort"),
//			os.Getenv("DBName"),
//		)
//		db, err := sql.Open("mysql", dsn)
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		if err := db.Ping(); err != nil {
//			log.Fatal(err)
//		}
//
//		dbClient = db
//	})
//	dbClient.SetConnMaxLifetime(time.Hour * 1)
//	return dbClient
//}

// GetDB returns the current database connection instance
func GetDB() *sql.DB {
	return DB
}

func CloseDB() {
	if DB != nil {
		err := DB.Close()
		if err != nil {
			log.Printf("Error closing DB: %v", err)
		}
		DB = nil // Set DB to nil after closing
	}
}
