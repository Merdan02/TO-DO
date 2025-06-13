package database

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var DB *sql.DB

func ConnectDB() (*sql.DB, error) {
	err2 := godotenv.Load()
	if err2 != nil {
		log.Fatal("Error loading .env file")
		return nil, err2
	}
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_NAME")

	ConStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, database)
	db, err := sql.Open("postgres", ConStr)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	DB = db
	return DB, nil
}
