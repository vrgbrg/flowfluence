package database

import (
	"database/sql"
	"fmt"
)

func Connect(connection string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connection)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	return db, err
}
