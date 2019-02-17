package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/vrgbrg/flowfluence/database/models"
)

type UsersSource struct {
	db *sql.DB
}

func NewUsersSource(db *sql.DB) *UsersSource {
	return &UsersSource{
		db: db,
	}
}

func (uS UsersSource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", uS.GetAll)
	r.Post("/", uS.Create)

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", uS.Get)
		r.Put("/", uS.Update)
		r.Delete("/", uS.Delete)
	})

	return r
}

func (uS UsersSource) GetAll(w http.ResponseWriter, r *http.Request) {
	userHandler := models.NewUserHandler(uS.db)

	getUsers, err := userHandler.GetAll(context.Background())

	if err != nil {
		log.Print("Users get error is: ", err)
		respondWithError(w, http.StatusInternalServerError, "Server Error Post")
		return
	}

	respondwithJSON(w, http.StatusOK, getUsers)
}

func (uS UsersSource) Get(w http.ResponseWriter, r *http.Request) {
	var user models.User

	id := chi.URLParam(r, "id")

	user.ID = id

	userHandler := models.NewUserHandler(uS.db)

	getUser, err := userHandler.Get(context.Background(), user)

	if err != nil {
		log.Print("User get error is: ", err)
		respondWithError(w, http.StatusNoContent, "User not found")
		return
	}

	respondwithJSON(w, http.StatusOK, getUser)
}

func (uS UsersSource) Create(w http.ResponseWriter, r *http.Request) {
	var user models.User

	json.NewDecoder(r.Body).Decode(&user)

	fmt.Println(user)

	userHandler := models.NewUserHandler(uS.db)

	savedUser, err := userHandler.Create(context.Background(), user)

	if err != nil {
		log.Print("User save error is: ", err)
		respondWithError(w, http.StatusInternalServerError, "Server Error")
		return
	}

	respondwithJSON(w, http.StatusOK, savedUser)
}

func (us UsersSource) Update(w http.ResponseWriter, r *http.Request) {
	var user models.User

	id := chi.URLParam(r, "id")

	json.NewDecoder(r.Body).Decode(&user)

	userHandler := models.NewUserHandler(us.db)

	user.ID = id

	updatedUser, err := userHandler.Update(context.Background(), user)

	if err != nil {
		log.Print("User update error is: ", err)
		respondWithError(w, http.StatusInternalServerError, "Server Error")
		return
	}

	respondwithJSON(w, http.StatusOK, updatedUser)
}

func (us UsersSource) Delete(w http.ResponseWriter, r *http.Request) {
	var user models.User

	id := chi.URLParam(r, "id")

	userHandler := models.NewUserHandler(us.db)

	user.ID = id

	err := userHandler.Delete(context.Background(), user)

	if err != nil {
		log.Print("User delete error is: ", err)
		respondWithError(w, http.StatusInternalServerError, "Server Error")
		return
	}

	respondwithJSON(w, http.StatusMovedPermanently, map[string]string{"message": "Delete Successfully"})
}

func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	fmt.Println(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondwithJSON(w, code, map[string]string{"message": msg})
}
