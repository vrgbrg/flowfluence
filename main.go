package main

import (
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

	_, err := database.Connect(connection)

	if err != nil {
		panic(err)
	}
}
