//+build !github

package facebook_test

import (
	"daanretard/internal/infrastructure/utilities/facebook"
	"github.com/joho/godotenv"
	"os"
	"testing"
)

var (
	testPostID string
	service    *facebook.Service
)

func init() {
	if err := godotenv.Load("../../../../.env"); err != nil {
		panic(err)
	}
	service = facebook.NewService(os.Getenv("FB_APP_ID"), os.Getenv("FB_APP_SECRET"), "daanretard")
}

func TestService_PublishPost(t *testing.T) {
	t.Run("should succeed", func(t *testing.T) {
		var err error
		testPostID, err = service.PublishPost("test message", nil, os.Getenv("FB_TEST_ACCESS_TOKEN"))
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("should fail with: invalid access token", func(t *testing.T) {
		_, err := service.PublishPost("test message", nil, "invalid token")
		if err == nil || err.Error() != "invalid access token" {
			t.Error(err)
		}
	})
}

func TestService_PublishPhoto(t *testing.T) {
	t.Run("should succeed", func(t *testing.T) {
		_, err := service.PublishPhoto("../../../../ui/public/images/icon.png", os.Getenv("FB_TEST_ACCESS_TOKEN"))
		if err != nil {
			t.Error(err)
		}
	})
}

func TestService_DeletePost(t *testing.T) {
	t.Run("should succeed", func(t *testing.T) {
		err := service.DeletePost(testPostID, os.Getenv("FB_TEST_ACCESS_TOKEN"))
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("should fail with: not exist", func(t *testing.T) {
		err := service.DeletePost(testPostID, os.Getenv("FB_TEST_ACCESS_TOKEN"))
		if err == nil || err.Error() != "not exist" {
			t.Error(err)
		}
	})
	t.Run("should fail with: invalid access token", func(t *testing.T) {
		err := service.DeletePost(testPostID, "invalid token")
		if err == nil || err.Error() != "invalid access token" {
			t.Error(err)
		}
	})
}

func TestService_PublishPostWithPhoto(t *testing.T) {
	t.Run("should succeed", func(t *testing.T) {
		accessToken := os.Getenv("FB_TEST_ACCESS_TOKEN")
		photo, err := service.PublishPhoto("../../../../ui/public/images/icon.png", accessToken)
		if err != nil {
			t.Fatal(err)
		}
		post, err := service.PublishPost("test message", []string{photo}, accessToken)
		if err != nil {
			t.Error(err)
		}
		err = service.DeletePost(post, accessToken)
		if err != nil {
			t.Error(err)
		}
	})
}

func TestService_ExchangeToken(t *testing.T) {
	t.Run("should succeed", func(t *testing.T) {
		_, err := service.ExchangeToken(os.Getenv("FB_TEST_ACCESS_TOKEN"))
		if err != nil {
			t.Error(err)
		}
	})
}
