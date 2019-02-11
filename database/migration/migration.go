package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/vrgbrg/flowfluence/database"
)

const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "flowfluence"
	password = "flowfluence"
	dbname   = "flowfluence"
)

func main() {

	connection := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := database.Connect(connection)

	if err != nil {
		panic(err)
	}

	for _, schema := range schemas {
		createSchema(db, schema)
	}
}

func createSchema(db *sql.DB, schema string) {
	_, err := db.Exec(schema)
	if err != nil {
		panic(err)
	}
}

var schemas = []string{
	`CREATE TABLE users (
		id char(20) PRIMARY KEY,
		type varchar(255),
		name varchar(255),
		password varchar(255),
		email varchar(255)
	)`,
	`CREATE TABLE posts (
		id char(20) PRIMARY KEY,
		title varchar(255),
		status varchar(255),
		author varchar(255) REFERENCES users (id),
		createdAt timestamp,
		modifiedAt timestamp,
		deletedAt timestamp
	)`,
	`CREATE TABLE postElements (
		id char(20) PRIMARY KEY,
		postId char(20) REFERENCES posts (id),
		type varchar(255),
		content varchar(2000)
	)`,
}
