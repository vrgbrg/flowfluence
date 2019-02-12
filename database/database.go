package database

import (
	"database/sql"
	"log"
)

func Connect(connection string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connection)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println("Successfully connected!")

	return db, err
}
