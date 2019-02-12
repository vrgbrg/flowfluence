package models

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/rs/xid"
)

type PostHandler struct {
	db *sql.DB
}

func NewPostHandler(db *sql.DB) *PostHandler {
	return &PostHandler{
		db: db,
	}
}

func (h *PostHandler) Get(ctx context.Context, p Post) (Post, error) {
	var post Post

	err := h.db.QueryRowContext(ctx, "SELECT id, title, status, author, createdAt, modifiedAt, deletedAt FROM posts WHERE id = $1", p.ID).
		Scan(&post.ID, &post.Title, &post.Status, &post.Author, &post.CreatedAt, &post.ModifiedAt, &post.DeletedAt)

	if err != nil {
		return post, err
	}

	return post, nil
}

func (h *PostHandler) Create(ctx context.Context, p Post) (Post, error) {
	stmt, err := h.db.PrepareContext(ctx, "INSERT INTO posts(id, title, status, author, createdAt, modifiedAt, deletedAt) VALUES($1,$2,$3,$4,$5,$6,$7)")

	if err != nil {
		return Post{}, err
	}

	if p.ID == "" {
		p.ID = xid.New().String()
	}

	p.CreatedAt = time.Now()

	res, err := stmt.ExecContext(ctx, p.ID, p.Title, p.Status, p.Author, p.CreatedAt, p.ModifiedAt, p.DeletedAt)

	if err != nil {
		return Post{}, err
	}

	num, err := res.RowsAffected()
	if err != nil {
		return Post{}, err
	}
	if num != 1 {
		return Post{}, errors.New("affected rows invalid")
	}

	return p, nil
}

func (h *PostHandler) Update(ctx context.Context, p Post) (Post, error) {
	oldPost, err := h.Get(ctx, p)

	if err != nil {
		return Post{}, err
	}

	if p.Title == "" {
		p.Title = oldPost.Title
	}

	if p.Status == "" {
		p.Status = oldPost.Status
	}

	if p.Author == "" {
		p.Author = oldPost.Author
	}

	if p.CreatedAt.IsZero() {
		p.CreatedAt = oldPost.CreatedAt
	}

	if p.ModifiedAt.IsZero() {
		p.ModifiedAt = oldPost.ModifiedAt
	}

	if p.DeletedAt.IsZero() {
		p.DeletedAt = oldPost.DeletedAt
	}

	stmt, err := h.db.PrepareContext(ctx, "UPDATE posts SET id=$1, title=$2, status=$3, author=$4, createdAt=$5, modifiedAt=$6, deletedAt=$7 WHERE id = $8;")
	if err != nil {
		return Post{}, err
	}

	if p.ID == "" {
		p.ID = xid.New().String()
	}

	res, err := stmt.ExecContext(ctx, p.ID, p.Title, p.Status, p.Author, p.CreatedAt, p.ModifiedAt, p.DeletedAt, p.ID)
	if err != nil {
		return Post{}, err
	}
	num, err := res.RowsAffected()
	if err != nil {
		return Post{}, err
	}
	if num != 1 {
		return Post{}, errors.New("affected rows invalid")
	}

	return p, nil
}

func (h *PostHandler) Delete(ctx context.Context, p Post) error {
	stmt, err := h.db.PrepareContext(ctx, "DELETE FROM posts WHERE id = $1;")
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, p.ID)
	if err != nil {
		return err
	}
	num, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if num != 1 {
		return errors.New("affected rows invalid")
	}

	return nil
}

type PostElementHandler struct {
	db *sql.DB
}

func newPostElementHandler(db *sql.DB) *PostElementHandler {
	return &PostElementHandler{
		db: db,
	}
}

func (h *PostElementHandler) Get(ctx context.Context, pE PostElement) (PostElement, error) {
	var postElement PostElement

	err := h.db.QueryRowContext(ctx, "SELECT id, postid, type, content FROM postelements WHERE id = $1", pE.ID).
		Scan(&postElement.ID, &postElement.PostID, &postElement.Type, &postElement.Content)
	if err != nil {
		return postElement, err
	}

	return postElement, nil
}

func (h *PostElementHandler) Create(ctx context.Context, pE PostElement) (PostElement, error) {
	stmt, err := h.db.PrepareContext(ctx, "INSERT INTO postelements(id, postid, type, content) VALUES($1,$2,$3,$4)")

	if err != nil {
		return PostElement{}, err
	}

	if pE.ID == "" {
		pE.ID = xid.New().String()
	}

	res, err := stmt.ExecContext(ctx, pE.ID, pE.PostID, pE.Type, pE.Content)

	if err != nil {
		return PostElement{}, err
	}

	num, err := res.RowsAffected()
	if err != nil {
		return PostElement{}, err
	}
	if num != 1 {
		return PostElement{}, errors.New("affected rows invalid")
	}

	return pE, nil
}

func (h *PostElementHandler) Update(ctx context.Context, pE PostElement) (PostElement, error) {
	oldPostElement, err := h.Get(ctx, pE)

	if err != nil {
		return PostElement{}, err
	}

	if pE.PostID == "" {
		pE.PostID = oldPostElement.PostID
	}

	if pE.Type == "" {
		pE.Type = oldPostElement.Type
	}

	if pE.Content == "" {
		pE.Content = oldPostElement.Content
	}

	stmt, err := h.db.PrepareContext(ctx, "UPDATE postelements SET id=$1, postid=$2, type=$3, content=$4 WHERE id = $5;")
	if err != nil {
		return PostElement{}, err
	}

	if pE.ID == "" {
		pE.ID = xid.New().String()
	}

	res, err := stmt.ExecContext(ctx, pE.ID, pE.PostID, pE.Type, pE.Content, pE.ID)
	if err != nil {
		return PostElement{}, err
	}
	num, err := res.RowsAffected()
	if err != nil {
		return PostElement{}, err
	}
	if num != 1 {
		return PostElement{}, errors.New("affected rows invalid")
	}

	return pE, nil
}

func (h *PostElementHandler) Delete(ctx context.Context, pE PostElement) error {
	stmt, err := h.db.PrepareContext(ctx, "DELETE FROM postelements WHERE id = $1;")
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, pE.ID)
	if err != nil {
		return err
	}
	num, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if num != 1 {
		return errors.New("affected rows invalid")
	}

	return nil
}

type UserHandler struct {
	db *sql.DB
}

func NewUserHandler(db *sql.DB) *UserHandler {
	return &UserHandler{
		db: db,
	}
}

func (h *UserHandler) Get(ctx context.Context, u User) (User, error) {
	var user User

	err := h.db.QueryRowContext(ctx, "SELECT id, type, name, password, email FROM users WHERE id = $1", u.ID).
		Scan(&user.ID, &user.Type, &user.Name, &user.Password, &user.Email)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (h *UserHandler) Create(ctx context.Context, u User) (User, error) {
	stmt, err := h.db.PrepareContext(ctx, "INSERT INTO users(id, type, name, password, email) VALUES($1,$2,$3,$4,$5)")

	if err != nil {
		return User{}, err
	}

	if u.ID == "" {
		u.ID = xid.New().String()
	}

	res, err := stmt.ExecContext(ctx, u.ID, u.Type, u.Name, u.Password, u.Email)

	if err != nil {
		return User{}, err
	}

	num, err := res.RowsAffected()
	if err != nil {
		return User{}, err
	}
	if num != 1 {
		return User{}, errors.New("affected rows invalid")
	}

	return u, nil
}

func (h *UserHandler) Update(ctx context.Context, u User) (User, error) {
	oldUser, err := h.Get(ctx, u)

	if err != nil {
		return User{}, err
	}

	if u.Type == "" {
		u.Type = oldUser.Type
	}

	if u.Name == "" {
		u.Name = oldUser.Name
	}

	if u.Password == "" {
		u.Password = oldUser.Password
	}

	if u.Email == "" {
		u.Email = oldUser.Email
	}

	stmt, err := h.db.PrepareContext(ctx, "UPDATE users SET id=$1, type=$2, name=$3, password=$4, email=$5 WHERE id = $6;")
	if err != nil {
		return User{}, err
	}

	if u.ID == "" {
		u.ID = xid.New().String()
	}

	res, err := stmt.ExecContext(ctx, u.ID, u.Type, u.Name, u.Password, u.Email, u.ID)
	if err != nil {
		return User{}, err
	}
	num, err := res.RowsAffected()
	if err != nil {
		return User{}, err
	}
	if num != 1 {
		return User{}, errors.New("affected rows invalid")
	}

	return u, nil
}

func (h *UserHandler) Delete(ctx context.Context, u User) error {
	stmt, err := h.db.PrepareContext(ctx, "DELETE FROM users WHERE id = $1;")
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, u.ID)
	if err != nil {
		return err
	}
	num, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if num != 1 {
		return errors.New("affected rows invalid")
	}

	return nil
}
