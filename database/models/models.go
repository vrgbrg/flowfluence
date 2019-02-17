package models

import (
	"time"
)

type Post struct {
	ID         string    `json:"id"`
	Title      string    `json:"title"`
	Status     string    `json:"status"`
	Author     string    `json:"author"`
	CreatedAt  time.Time `json:"createdAt"`
	ModifiedAt time.Time `json:"modifiedAt"`
	DeletedAt  time.Time `json:"deletedAt"`
}

type PostElement struct {
	ID      string `json:"id"`
	PostID  string `json:"postId"`
	Type    string `json:"type"`
	Content string `json:"content"`
}

type User struct {
	ID       string `json:"id"`
	Type     string `json:"type"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Article struct {
	ID       string        `json:"id"`
	Title    string        `json:"title"`
	Author   string        `json:"author"`
	Sections []PostElement `json:"sections"`
}
