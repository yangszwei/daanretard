package mock_persistence_test

import (
	entity "daanretard/internal/domain/post"
	"daanretard/internal/infrastructure/mock_persistence"
	"net"
	"testing"
)

var (
	posts *mock_persistence.PostRepository
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
	posts = mock_persistence.NewPostRepository()
}

func TestPostRepository_InsertOne(t *testing.T) {
	err := posts.InsertOne(&testPost)
	if err != nil {
		t.Error(err)
	}
}

func TestPostRepository_FindOne(t *testing.T) {
	t.Run("find by ID", func(t *testing.T) {
		_, err := posts.FindOne(entity.Query{ ID: testPost.ID })
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("find by other fields", func(t *testing.T) {
		post, err := posts.FindOne(entity.Query{
			Status: testPost.Status,
			SubmitterID: testPost.Submission.SubmitterID,
			ReviewerID: testPost.Review.ReviewerID,
		})
		if err != nil {
			t.Error(err)
		}
		if post.ID != testPost.ID {
			t.Error(post)
		}
	})
	t.Run("not found", func(t *testing.T) {
		_, err := posts.FindOne(entity.Query{})
		if err == nil || err.Error() != "record not found" {
			t.Error(err)
		}
	})
}

func TestPostRepository_FindAll(t *testing.T) {
	t.Run("find by ID", func(t *testing.T) {
		_, err := posts.FindAll(entity.Query{ ID: testPost.ID })
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("find by other fields", func(t *testing.T) {
		posts, err := posts.FindAll(entity.Query{
			Status: testPost.Status,
			SubmitterID: testPost.Submission.SubmitterID,
			ReviewerID: testPost.Review.ReviewerID,
		})
		if err != nil {
			t.Error(err)
		}
		if posts[0].ID != testPost.ID {
			t.Error(posts)
		}
	})
	t.Run("not found", func(t *testing.T) {
		_, err := posts.FindAll(entity.Query{})
		if err == nil || err.Error() != "record not found" {
			t.Error(err)
		}
	})
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