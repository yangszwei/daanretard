package media_test

import (
	"daanretard/internal/entity/media"
	"daanretard/internal/infra/errors"
	"daanretard/internal/infra/object"
	"daanretard/internal/infra/persistence"
	"daanretard/internal/infra/validator"
	"github.com/joho/godotenv"
	"os"
	"testing"
)

var (
	id      uint32
	service *media.Service
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
	m := persistence.NewMediaRepo(db)
	v := validator.New()
	service = media.NewService(m, v)
}

func TestService_Add(t *testing.T) {
	var err error
	id, err = service.Add(object.Media{
		UserID: "test user",
		Name:   "Test Photo 1",
	})
	if err != nil {
		t.Error(err)
	}
}

func TestService_GetOne(t *testing.T) {
	t.Run("should succeed", func(t *testing.T) {
		if _, err := service.GetOne(id); err != nil {
			t.Error(err)
		}
	})
	t.Run("should fail: record not found", func(t *testing.T) {
		if _, err := service.GetOne(100); !errors.Is(err.(errors.Error), errors.ErrRecordNotFound) {
			t.Error(err)
		}
	})
}

func TestService_Delete(t *testing.T) {
	if err := service.Delete(id); err != nil {
		t.Error(err)
	}
}
