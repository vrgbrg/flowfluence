package models

import (
	"context"
	"fmt"
	"testing"

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

func TestPostHandler(t *testing.T) {
	connection := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := database.Connect(connection)

	if err != nil {
		t.Error(err)
		return
	}

	userHandler := NewUserHandler(db)

	newUser, err := userHandler.Create(context.Background(), User{
		Type:     "Mentor",
		Name:     "Feri",
		Password: "Go",
		Email:    "feri@feri.com",
	})
	if err != nil {
		t.Error(err)
		return
	}

	postHandler := NewPostHandler(db)

	newPost, err := postHandler.Create(context.Background(), Post{
		Title:  "Szegedi Attila",
		Status: "Király",
		Author: newUser.ID,
	})
	if err != nil {
		t.Error(err)
		return
	}

	oldPost, err := postHandler.Get(context.Background(), newPost)
	if err != nil {
		t.Error(err)
		return
	}

	if newPost.ID != oldPost.ID {
		t.Errorf("Post id should be %s, instead of %s", newPost.ID, oldPost.ID)
	}

	postUpdate, err := postHandler.Update(context.Background(), Post{
		ID:     oldPost.ID,
		Title:  "Szegedi Attila",
		Status: "Nagyon király",
		Author: newUser.ID,
	})
	if err != nil {
		t.Error(err)
		return
	}

	if postUpdate.ID != oldPost.ID {
		t.Errorf("Post id should be %s, instead of %s", oldPost.ID, postUpdate.ID)
	}

	postElementHandler := newPostElementHandler(db)

	newPostElement, err := postElementHandler.Create(context.Background(), PostElement{
		PostID:  newPost.ID,
		Type:    "Paragraph",
		Content: "Test",
	})
	if err != nil {
		t.Error(err)
		return
	}

	oldPostElement, err := postElementHandler.Get(context.Background(), newPostElement)
	if err != nil {
		t.Error(err)
		return
	}

	if newPostElement.ID != oldPostElement.ID {
		t.Errorf("Post element id should be %s, instead of %s", newPostElement.ID, oldPostElement.ID)
	}

	postElementUpdate, err := postElementHandler.Update(context.Background(), PostElement{
		ID:      oldPostElement.ID,
		PostID:  newPost.ID,
		Type:    "Image",
		Content: "Test",
	})
	if err != nil {
		t.Error(err)
		return
	}

	if postElementUpdate.ID != oldPostElement.ID {
		t.Errorf("Post element id should be %s, instead of %s", oldPostElement.ID, postElementUpdate.ID)
	}

	// err = postHandler.Delete(context.Background(), Post{
	// 	ID: oldPost.ID,
	// })
	// if err != nil {
	// 	t.Error(err)
	// 	return
	// }

	// err = postElementHandler.Delete(context.Background(), PostElement{
	// 	ID: oldPostElement.ID,
	// })
	// if err != nil {
	// 	t.Error(err)
	// 	return
	// }

	// err = db.Close()
	// if err != nil {
	// 	t.Error(err)
	// 	return
	// }
}

func TestUserHandler(t *testing.T) {
	connection := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := database.Connect(connection)

	if err != nil {
		t.Error(err)
		return
	}

	handler := NewUserHandler(db)

	newUser, err := handler.Create(context.Background(), User{
		Type:     "Mentor",
		Name:     "Feri",
		Password: "Go",
		Email:    "feri@feri.com",
	})
	if err != nil {
		t.Error(err)
		return
	}

	oldUser, err := handler.Get(context.Background(), newUser)
	if err != nil {
		t.Error(err)
		return
	}

	if newUser.ID != oldUser.ID {
		t.Errorf("User id should be %s, instead of %s", newUser.ID, oldUser.ID)
	}

	userUpdate, err := handler.Update(context.Background(), User{
		ID:       oldUser.ID,
		Type:     "Mentor",
		Name:     "Máté",
		Password: "Go",
		Email:    "mate@mate.com",
	})
	if err != nil {
		t.Error(err)
		return
	}

	if userUpdate.ID != oldUser.ID {
		t.Errorf("User id should be %s, instead of %s", oldUser.ID, userUpdate.ID)
	}

	err = handler.Delete(context.Background(), User{
		ID: oldUser.ID,
	})
	if err != nil {
		t.Error(err)
		return
	}

	// err = db.Close()
	// if err != nil {
	// 	t.Error(err)
	// 	return
	// }
}
