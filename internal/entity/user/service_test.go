package user_test

import (
	"daanretard/internal/entity/user"
	"daanretard/internal/infra/mock_fbgraph"
	"daanretard/internal/infra/persistence"
	"github.com/joho/godotenv"
	"os"
	"testing"
)

var (
	fb      *mock_fbgraph.App
	uid     string
	service *user.Service
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
	u := persistence.NewUserRepo(db)
	fb = mock_fbgraph.New("test", "test", "test")
	service = user.NewService(u, fb)
}

func TestService_ContinueWithFacebook(t *testing.T) {
	token := fb.NewUser("Test User")
	var err error
	uid, err = service.ContinueWithFacebook(token)
	if err != nil {
		t.Error(err)
	}
}

func TestService_GetProfile(t *testing.T) {
	u, err := service.GetProfile(uid)
	if err != nil {
		t.Fatal(err)
	}
	if u.Name != "Test User" {
		t.Error(u)
	}
}

func TestService_Delete(t *testing.T) {
	if err := service.Delete(uid); err != nil {
		t.Error(err)
	}
}
