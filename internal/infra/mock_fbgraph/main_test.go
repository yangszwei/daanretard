package mock_fbgraph_test

import (
	"daanretard/internal/infra/mock_fbgraph"
	"github.com/joho/godotenv"
	"os"
	"testing"
)

var token, photo, post string
var app *mock_fbgraph.App

func init() {
	if err := godotenv.Load("../../../.env"); err != nil {
		panic(err)
	}
}

func TestNew(t *testing.T) {
	pageID := "daanretard"
	appID := os.Getenv("FB_APP_ID")
	appSecret := os.Getenv("FB_APP_SECRET")
	app = mock_fbgraph.New(pageID, appID, appSecret)
	token = app.NewUser("Test User")
}

func TestApp_ExchangeAccessToken(t *testing.T) {
	var err error
	token, _, err = app.ExchangeAccessToken(token)
	if err != nil {
		t.Error(err)
	}
}

func TestApp_DebugAccessToken(t *testing.T) {
	t.Run("should succeed", func(t *testing.T) {
		res, err := app.DebugAccessToken(token)
		if err != nil {
			t.Error(err)
		}
		if !res.IsValid {
			t.Error(err)
		}
	})
	t.Run("should fail: invalid access token", func(t *testing.T) {
		_, err := app.DebugAccessToken("invalid access token")
		if err == nil || err.Error() != "invalid access token" {
			t.Error(err)
		}
	})
}

func TestApp_GetUserProfile(t *testing.T) {
	if _, _, err := app.GetUserProfile(token); err != nil {
		t.Error(err)
	}
}

func TestApp_NewPostAccessToken(t *testing.T) {
	i, _ := app.DebugAccessToken(token)
	token = app.NewPage("daanretard", i.UserID)
}

func TestApp_UploadPhoto(t *testing.T) {
	t.Run("should succeed", func(t *testing.T) {
		var err error
		photo, err = app.UploadPhoto("../../../ui/public/images/icon.png", false, token)
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("should fail: invalid access token", func(t *testing.T) {
		_, err := app.UploadPhoto("../../../ui/public/images/icon.png", false, "invalid token")
		if err == nil || err.Error() != "invalid access token" {
			t.Error(err)
		}
	})

}

func TestApp_PublishPost(t *testing.T) {
	t.Run("should succeed", func(t *testing.T) {
		var err error
		post, err = app.PublishPost("test message", []string{photo}, token)
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("should fail: invalid access token", func(t *testing.T) {
		_, err := app.PublishPost("test", nil, "invalid token")
		if err == nil || err.Error() != "invalid access token" {
			t.Error(err)
		}
	})
}

func TestApp_Delete(t *testing.T) {
	t.Run("should succeed", func(t *testing.T) {
		if err := app.Delete(post, token); err != nil {
			t.Error(err)
		}
	})
	t.Run("should fail: invalid access token", func(t *testing.T) {
		err := app.Delete(post, "invalid token")
		if err == nil || err.Error() != "invalid access token" {
			t.Error(err)
		}
	})
}
