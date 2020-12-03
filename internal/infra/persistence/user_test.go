package persistence_test

import (
	"daanretard/internal/entity/user"
	"daanretard/internal/infra/persistence"
	"testing"
	"time"
)

var (
	u = user.User{
		ID:          "test_user_repo",
		Name:        "Test User",
		AccessToken: "example_access_token",
		ExpiresAt:   time.Now().Add(2 * time.Hour),
	}
	ur *persistence.UserRepo
)

func TestNewUserRepo(t *testing.T) {
	ur = persistence.NewUserRepo(newDB())
}

func TestUserRepo_AutoMigrate(t *testing.T) {
	if err := ur.AutoMigrate(); err != nil {
		t.Error(err)
	}
}

func TestUserRepo_InsertOne(t *testing.T) {
	if err := ur.InsertOne(&u); err != nil {
		t.Error(err)
	}
}

func TestUserRepo_FindOneByID(t *testing.T) {
	if r, err := ur.FindOneByID(u.ID); err != nil || r.ID != u.ID {
		t.Error(r, err)
	}
}

func TestUserRepo_FindMany(t *testing.T) {
	if r, err := ur.FindMany(10, 0); err != nil || len(r) != 1 {
		t.Error(r, err)
	}
}

func TestUserRepo_UpdateOne(t *testing.T) {
	if err := ur.UpdateOne(&u); err != nil {
		t.Error(err)
	}
}

func TestUserRepo_DeleteOne(t *testing.T) {
	if err := ur.DeleteOne(&u); err != nil {
		t.Error(err)
	}
}
