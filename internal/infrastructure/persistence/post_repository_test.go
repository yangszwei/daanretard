package persistence_test

import (
	"daanretard/internal/domain/post"
	"daanretard/internal/infrastructure/persistence"
	"daanretard/internal/object"
	"net"
	"testing"
	"time"
)

var testPost = post.Post{
	Status:      post.StatusPublished,
	UserID:      1,
	IPAddr:      net.ParseIP("::"),
	UserAgent:   "post_repository test",
	Message:     "test message",
	Attachments: "test_attachment_1,test_attachment_2",
	Review:      post.Review{
		UserID: 10,
		Result: 10,
	},
	FacebookID:  "12345",
}

func TestPostRepository_AutoMigrate(t *testing.T) {
	err := persistence.NewPostRepository(DB).AutoMigrate()
	if err != nil {
		t.Error(err)
	}
}

func TestPostRepository_InsertOne(t *testing.T) {
	err := persistence.NewPostRepository(DB).InsertOne(&testPost)
	if err != nil {
		t.Error(err)
	}
}

func TestPostRepository_FindOneByID(t *testing.T) {
	t.Run("should succeed", func(t *testing.T) {
		_, err := persistence.NewPostRepository(DB).FindOneByID(testPost.ID)
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("should fail with: record not found", func(t *testing.T) {
		_, err := persistence.NewPostRepository(DB).FindOneByID(0)
		if err == nil || err.Error() != "record not found" {
			t.Error(err)
		}
	})
}

func TestPostRepository_FindMany(t *testing.T) {
	t.Run("should succeed", func(t *testing.T) {
		_, err := persistence.NewPostRepository(DB).FindMany(object.PostQuery{
			Status:        testPost.Status,
			UserID:        testPost.UserID,
			IPAddr:        testPost.IPAddr,
			CreatedBefore: time.Now(),
			ReviewerID:    testPost.Review.UserID,
			ReviewResult:  testPost.Review.Result,
			Limit:         1,
			Offset:        0,
		})
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("should fail with: record not found", func(t *testing.T) {
		_, err := persistence.NewPostRepository(DB).FindMany(object.PostQuery{})
		if err == nil || err.Error() != "record not found" {
			t.Error(err)
		}
	})
}

func TestPostRepository_UpdateOne(t *testing.T) {
	err := persistence.NewPostRepository(DB).UpdateOne(&testPost)
	if err != nil {
		t.Error(err)
	}
}

func TestPostRepository_DeleteOne(t *testing.T) {
	err := persistence.NewPostRepository(DB).DeleteOne(&testPost)
	if err != nil {
		t.Error(err)
	}
}