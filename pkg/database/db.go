package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func ConnectDB() (*sql.DB, error){
	connectionString := os.Getenv("DATABASE_URI")

	if connectionString == "" {
		return nil, fmt.Errorf("DATABASE_URI is not set in .env")
	}

	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		
		return nil, err
	}

	return db, nil 
}

func Close() {
	if db !=nil {
		db.Close()
	}
}

func GetDB() *sql.DB {
	return db
}