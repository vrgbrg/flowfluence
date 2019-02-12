package models

import (
	"time"
)

type Post struct {
	ID         string
	Title      string
	Status     string
	Author     string
	CreatedAt  time.Time
	ModifiedAt time.Time
	DeletedAt  time.Time
}

type PostElement struct {
	ID      string
	PostID  string
	Type    string
	Content string
}

type User struct {
	ID       string
	Type     string
	Name     string
	Password string
	Email    string
}
