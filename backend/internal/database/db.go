package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() error {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	if sslmode == "" {
		sslmode = "disable"
	}

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	// Retry connection since DB might take a few seconds to start in Docker
	for i := 0; i < 5; i++ {
		err = DB.Ping()
		if err == nil {
			log.Println("Successfully connected to database!")
			return nil
		}
		log.Printf("Failed to ping database, retrying in 2 seconds... (%d/5)\n", i+1)
		time.Sleep(2 * time.Second)
	}

	return fmt.Errorf("could not connect to database after retries: %v", err)
}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}
