package post_test

import (
	"daanretard/internal/entity/post"
	"daanretard/internal/infra/mock_fbgraph"
	"daanretard/internal/infra/object"
	"daanretard/internal/infra/persistence"
	"daanretard/internal/infra/validator"
	"github.com/joho/godotenv"
	"net"
	"os"
	"testing"
)

var (
	fb      *mock_fbgraph.App
	id      uint32
	token   string
	service *post.Service
)

func init() {
	if err := godotenv.Load("../../../.env"); err != nil {
		panic(err)
	}
}

func TestNewService(t *testing.T) {
	db, err := persistence.Open(os.Getenv("DB_DSN"))
	if err != nil {
		panic(err)
	}
	p := persistence.NewPostRepo(db)
	fb = mock_fbgraph.New("test", "test", "test")
	v := validator.New()
	service = post.NewService(v, p, fb)
	uid, _, _ := fb.GetUserProfile(fb.NewUser("test user"))
	token = fb.NewPage("test page", uid)
}

func TestService_Submit(t *testing.T) {
	var err error
	id, err = service.Submit(object.Post{
		UserID:      "test user",
		IPAddr:      net.ParseIP("::"),
		UserAgent:   "test ua",
		Message:     "test message",
		Attachments: []uint32{1, 2, 3},
	})
	if err != nil {
		t.Error(err)
	}
}

func TestService_Search(t *testing.T) {
	if _, err := service.Search(object.PostQuery{Message: "message"}); err != nil {
		t.Error(err)
	}
}

func TestService_Review(t *testing.T) {
	if err := service.Review(id, object.PostReview{
		UserID:  "test admin",
		Result:  post.ReviewApproved,
		Message: "",
	}); err != nil {
		t.Error(err)
	}
}

func TestService_Publish(t *testing.T) {
	if _, err := service.Publish(id, nil, token); err != nil {
		t.Error(err)
	}
}

func TestService_Delete(t *testing.T) {
	if err := service.Delete(id, token); err != nil {
		t.Error(err)
	}
}
