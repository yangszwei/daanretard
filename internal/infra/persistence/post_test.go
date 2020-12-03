package persistence_test

import (
	"daanretard/internal/entity/post"
	"daanretard/internal/infra/object"
	"daanretard/internal/infra/persistence"
	"net"
	"testing"
	"time"
)

var (
	p = post.Post{
		Status:      post.StatusPublished,
		UserID:      "test_user",
		IPAddr:      net.ParseIP("::"),
		UserAgent:   "test ua",
		Message:     "test message",
		Attachments: "1,2,3",
		Review: post.Review{
			UserID:  "test_admin",
			Result:  post.ReviewApproved,
			Message: "",
		},
		FacebookID: "test fbID",
	}
	pr *persistence.PostRepo
)

func TestNewPostRepo(t *testing.T) {
	pr = persistence.NewPostRepo(newDB())
}

func TestPostRepo_AutoMigrate(t *testing.T) {
	if err := pr.AutoMigrate(); err != nil {
		t.Error(err)
	}
}

func TestPostRepo_InsertOne(t *testing.T) {
	if err := pr.InsertOne(&p); err != nil {
		t.Error(err)
	}
}

func TestPostRepo_FindOneByID(t *testing.T) {
	if r, err := pr.FindOneByID(p.ID); err != nil || r.ID != p.ID {
		t.Error(r, err)
	}
}

func TestPostRepo_FindMany(t *testing.T) {
	records, err := pr.FindMany(object.PostQuery{
		Message:       "message",
		Status:        p.Status,
		UserID:        p.UserID,
		CreatedAfter:  time.Now().Add(-1 * time.Minute),
		CreatedBefore: time.Now().Add(1 * time.Minute),
		ReviewerID:    p.Review.UserID,
		ReviewResult:  p.Review.Result,
		FacebookID:    p.FacebookID,
		Limit:         10,
		Offset:        0,
	})
	if err != nil {
		t.Error(err)
	}
	if len(records) != 1 {
		t.Error(records)
	}
}

func TestPostRepo_UpdateOne(t *testing.T) {
	if err := pr.UpdateOne(&p); err != nil {
		t.Error(err)
	}
}

func TestPostRepo_DeleteOne(t *testing.T) {
	if err := pr.DeleteOne(&p); err != nil {
		t.Error(err)
	}
}
