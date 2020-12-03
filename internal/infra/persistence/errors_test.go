package persistence_test

import (
	"daanretard/internal/entity/user"
	"daanretard/internal/infra/errors"
	"daanretard/internal/infra/persistence"
	"testing"
)

func TestParseError(t *testing.T) {
	r := persistence.NewUserRepo(newDB())
	_ = r.AutoMigrate()
	_ = r.InsertOne(&u)
	if err := r.InsertOne(&user.User{ID: string([]byte{1000: 1})}); err == nil || err.Error() != "data too lang: id" {
		t.Error(err)
	}
	if err := r.InsertOne(&u).(errors.Error); !errors.Is(err, errors.ErrDuplicateEntry) {
		t.Error(err)
	}
	if _, err := r.FindOneByID("not exist"); !errors.Is(err.(errors.Error), errors.ErrRecordNotFound) {
		t.Error(err)
	}
	_ = r.DeleteOne(&u)
}
