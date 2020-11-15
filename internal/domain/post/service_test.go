package post_test

import (
	"daanretard/internal/domain/post"
	"daanretard/internal/infrastructure/mock_persistence"
	"daanretard/internal/object"
	"net"
	"testing"
)

var (
	service  = post.NewService(mock_persistence.NewPostRepository())
	testPost = object.Post{
		UserID:      1,
		IPAddr:      net.ParseIP("127.0.0.1"),
		UserAgent:   "test ua",
		Message:     "test message",
		Attachments: []string{"1", "2", "3"},
		FacebookID:  "test id",
	}
)

func TestService_Submit(t *testing.T) {
	t.Run("should succeed", func(t *testing.T) {
		id, err := service.Submit(testPost)
		if err != nil {
			t.Error(err)
		}
		testPost.ID = id
	})
	t.Run("should fail with: invalid credentials", func(t *testing.T) {
		_, err := service.Submit(object.Post{})
		if err == nil || err.Error() != "invalid credentials" {
			t.Error(err)
		}
	})
	t.Run("should fail with: invalid credentials #2", func(t *testing.T) {
		_, err := service.Submit(object.Post{
			UserID:    1,
			IPAddr:    net.ParseIP("127.0.0.1"),
			UserAgent: "test ua",
			Message:   string([]byte{60000: 1}), // len() == 60001
		})
		if err == nil || err.Error() != "invalid credentials" {
			t.Error(err)
		}
	})
}

func TestService_GetOne(t *testing.T) {
	t.Run("should succeed", func(t *testing.T) {
		post, err := service.GetOne(testPost.ID)
		if err != nil {
			t.Error(err)
		}
		if post.ID != testPost.ID {
			t.Error(err)
		}
	})
}

func TestService_GetManyByUserID(t *testing.T) {
	t.Run("should succeed", func(t *testing.T) {
		posts, err := service.GetManyByUserID(testPost.UserID, 0, 10)
		if err != nil {
			t.Error(err)
		}
		if posts[0].ID != testPost.ID {
			t.Error(posts)
		}
	})
	t.Run("should fail with: record not found", func(t *testing.T) {
		_, err := service.GetManyByUserID(0, 0, 10)
		if err == nil || err.Error() != "record not found" {
			t.Error(err)
		}
	})
}

func TestService_GetManyNotReviewed(t *testing.T) {
	t.Run("should succeed", func(t *testing.T) {
		posts, err := service.GetManyNotReviewed(0, 10)
		if err != nil {
			t.Error(err)
		}
		if posts[0].ID != testPost.ID {
			t.Error(posts)
		}
	})
}

func TestService_Review(t *testing.T) {
	t.Run("should succeed", func(u *testing.T) {
		err := service.Review(testPost.ID, object.PostReview{
			UserID:  2,
			Result:  3,
			Message: "test message",
		})
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("should fail with: record not found", func(t *testing.T) {
		err := service.Review(0, object.PostReview{})
		if err == nil || err.Error() != "record not found" {
			t.Error(err)
		}
	})
}

func TestService_Publish(t *testing.T) {
	t.Run("should succeed", func(t *testing.T) {
		err := service.MarkAsPublished(testPost.ID, "fb post id")
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("should fail with: record not found", func(t *testing.T) {
		err := service.MarkAsPublished(0, "fb post id")
		if err == nil || err.Error() != "record not found" {
			t.Error(err)
		}
	})
}

func TestService_GetManyPublished(t *testing.T) {
	t.Run("should succeed", func(t *testing.T) {
		posts, err := service.GetManyPublished(0, 10)
		if err != nil {
			t.Error(err)
		}
		if posts[0].ID != testPost.ID {
			t.Error(posts)
		}
	})
}

func TestService_Delete(t *testing.T) {
	t.Run("should succeed", func(t *testing.T) {
		err := service.Delete(testPost.ID)
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("should fail with: record not found", func(t *testing.T) {
		err := service.Delete(0)
		if err == nil || err.Error() != "record not found" {
			t.Error(err)
		}
	})
}
