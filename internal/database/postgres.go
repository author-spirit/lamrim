package database

import (
	"database/sql"
	"log"
)

func Connect() *sql.DB {
	db, err := sql.Open("postgres", "user=postgres password=postgres dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
