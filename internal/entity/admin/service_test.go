package admin_test

import (
	"daanretard/internal/entity/admin"
	"daanretard/internal/infra/mock_fbgraph"
	"daanretard/internal/infra/persistence"
	"daanretard/internal/infra/validator"
	"github.com/joho/godotenv"
	"os"
	"testing"
)

var (
	fb      *mock_fbgraph.App
	uid     string
	service *admin.Service
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
	a := persistence.NewAdminRepo(db)
	fb = mock_fbgraph.New("test", "test", "test")
	v := validator.New()
	service = admin.NewService(a, fb, v)
}

func TestService_Authenticate(t *testing.T) {
	token := fb.NewUser("test user")
	uid, _, _ = fb.GetUserProfile(token)
	if err := service.Authenticate(fb.NewPage("test page", uid)); err != nil {
		t.Error(err)
	}
}

func TestService_IsAdmin(t *testing.T) {
	if err := service.IsAdmin(uid); err != nil {
		t.Error(err)
	}
	t.Cleanup(func() {
		db, err := persistence.Open(os.Getenv("DB_DSN"))
		if err != nil {
			panic(err)
		}
		a := persistence.NewAdminRepo(db)
		record, _ := a.FindOneByUserID(uid)
		_ = a.DeleteOne(&record)
	})
}
