package persistence_test

import (
	entity "daanretard/internal/domain/post"
	"daanretard/internal/infrastructure/persistence"
	"net"
	"testing"
)

var (
	posts *persistence.PostRepository
	testPost = entity.Post{
		Status: entity.Submitted,
		Submission: entity.Submission{
			SubmitterID: 1,
			Message:     "test message",
			Attachments: "a,b,c",
			IPAddr:      net.ParseIP("127.0.0.1"),
			UserAgent:   "test user agent",
		},
		Review: entity.Review{
			ReviewerID: 2,
			Result:     1,
			Message:    "test message",
		},
	}
)


func TestNewPostRepository(t *testing.T) {
	posts = persistence.NewPostRepository(DB.Conn)
}

func TestPostRepository_AutoMigrate(t *testing.T) {
	err := posts.AutoMigrate()
	if err != nil {
		t.Error(err)
	}
}

func TestPostRepository_InsertOne(t *testing.T) {
	err := posts.InsertOne(&testPost)
	if err != nil {
		t.Error(err)
	}
}

func TestPostRepository_FindOne(t *testing.T) {
	_, err := posts.FindOne(entity.Query{ ID: testPost.ID })
	if err != nil {
		t.Error(err)
	}
}

func TestPostRepository_FindAll(t *testing.T) {
	_, err := posts.FindAll(entity.Query{ ID: testPost.ID })
	if err != nil {
		t.Error(err)
	}
}

func TestPostRepository_SaveOne(t *testing.T) {
	err := posts.SaveOne(&entity.Post{ ID: 1000 })
	if err != nil {
		t.Error(err)
	}
}

func TestPostRepository_DeleteOne(t *testing.T) {
	err := posts.DeleteOne(&testPost)
	if err != nil {
		t.Error(err)
	}
}