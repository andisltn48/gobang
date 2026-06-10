package db

import (
	"database/sql"
	"log"
	"time"
)

func Connect() *sql.DB {
	connectionString := "postgres://postgres@localhost:5432/gobang?sslmode=disable"
	db, err := sql.Open("pgx",connectionString)
	if err != nil {
		log.Fatal("Failed connection to database", err)
	}
	
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5* time.Minute)
	
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping database", err)
	}
	
	return db
}