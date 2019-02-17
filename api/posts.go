package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/vrgbrg/flowfluence/database/models"

	"github.com/go-chi/chi"
)

type PostsSource struct {
	db *sql.DB
}

func NewPostsSource(db *sql.DB) *PostsSource {
	return &PostsSource{
		db: db,
	}
}

func (pS PostsSource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", pS.GetAll)
	r.Post("/", pS.Create)

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", pS.Get)
		r.Put("/", pS.Update)
		r.Delete("/", pS.Delete)
	})

	return r
}

func (pS PostsSource) GetAll(w http.ResponseWriter, r *http.Request) {
	postHandler := models.NewPostHandler(pS.db)

	getPosts, err := postHandler.GetAll(context.Background())

	if err != nil {
		log.Print("Posts get error is: ", err)
		respondWithError(w, http.StatusInternalServerError, "Server Error Post")
		return
	}

	respondwithJSON(w, http.StatusOK, getPosts)
}

func (pS PostsSource) Get(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	var article models.Article

	id := chi.URLParam(r, "id")

	json.NewDecoder(r.Body).Decode(&article)

	post.ID = id
	post.Title = article.Title
	post.Author = article.Author

	postElements := article.Sections[:]

	postHandler := models.NewPostHandler(pS.db)

	getPost, err := postHandler.Get(context.Background(), post)

	if err != nil {
		log.Print("Post get error is: ", err)
		respondWithError(w, http.StatusInternalServerError, "Server Error Post")
		return
	}

	var getArticle models.Article

	getArticle.ID = getPost.ID
	getArticle.Title = getPost.Title
	getArticle.Author = getPost.Author

	postElementHandler := models.NewPostElementHandler(pS.db)

	for _, postElement := range postElements {
		postElement.PostID = getPost.ID
		getPostElement, err := postElementHandler.Get(context.Background(), postElement)

		if err != nil {
			log.Print("Post element get error is: ", err)
			respondWithError(w, http.StatusInternalServerError, "Server Error")
			return
		}

		getArticle.Sections = append(getArticle.Sections, getPostElement)
	}

	respondwithJSON(w, http.StatusOK, getArticle)
}

func (pS PostsSource) Create(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	var article models.Article

	json.NewDecoder(r.Body).Decode(&article)

	post.ID = article.ID
	post.Title = article.Title
	post.Author = article.Author

	postElements := article.Sections[:]

	postHandler := models.NewPostHandler(pS.db)

	savedPost, err := postHandler.Create(context.Background(), post)

	if err != nil {
		log.Print("Post save error is: ", err)
		respondWithError(w, http.StatusInternalServerError, "Server Error Post")
		return
	}

	var savedArticle models.Article

	savedArticle.ID = savedPost.ID
	savedArticle.Title = savedPost.Title
	savedArticle.Author = savedPost.Author

	postElementHandler := models.NewPostElementHandler(pS.db)

	for _, postElement := range postElements {
		postElement.PostID = savedPost.ID
		savedPostElement, err := postElementHandler.Create(context.Background(), postElement)

		if err != nil {
			log.Print("Post element update error is: ", err)
			respondWithError(w, http.StatusInternalServerError, "Server Error")
			return
		}

		savedArticle.Sections = append(savedArticle.Sections, savedPostElement)
	}

	respondwithJSON(w, http.StatusOK, savedArticle)
}

func (pS PostsSource) Update(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	var article models.Article

	id := chi.URLParam(r, "id")

	json.NewDecoder(r.Body).Decode(&article)

	post.ID = id
	post.Title = article.Title
	post.Author = article.Author

	postElements := article.Sections[:]

	postHandler := models.NewPostHandler(pS.db)

	updatedPost, err := postHandler.Update(context.Background(), post)

	if err != nil {
		log.Print("Post update error is: ", err)
		respondWithError(w, http.StatusInternalServerError, "Server Error Post")
		return
	}

	var updatedArticle models.Article

	updatedArticle.ID = updatedPost.ID
	updatedArticle.Title = updatedPost.Title
	updatedArticle.Author = updatedPost.Author

	postElementHandler := models.NewPostElementHandler(pS.db)

	for _, postElement := range postElements {
		postElement.PostID = updatedPost.ID
		updatedPostElement, err := postElementHandler.Update(context.Background(), postElement)

		if err != nil {
			log.Print("Post element update error is: ", err)
			respondWithError(w, http.StatusInternalServerError, "Server Error")
			return
		}

		updatedArticle.Sections = append(updatedArticle.Sections, updatedPostElement)
	}

	respondwithJSON(w, http.StatusOK, updatedArticle)

}

func (pS PostsSource) Delete(w http.ResponseWriter, r *http.Request) {
	var post models.Post

	id := chi.URLParam(r, "id")

	post.ID = id

	postElementHandler := models.NewPostElementHandler(pS.db)

	postElementHandler.DeleteByPostId(context.Background(), id)

	postHandler := models.NewPostHandler(pS.db)

	postHandler.Delete(context.Background(), post)

	respondwithJSON(w, http.StatusMovedPermanently, map[string]string{"message": "Delete Successfully"})
}
