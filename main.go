package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/lib/pq"
	"github.com/vrgbrg/flowfluence/api"
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
		log.Fatal(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	userRoutes := api.NewUsersSource(db)
	postRoutes := api.NewPostsSource(db)
	r.Mount("/users", userRoutes.Routes())
	r.Mount("/posts", postRoutes.Routes())

	http.ListenAndServe(":3333", r)
}
