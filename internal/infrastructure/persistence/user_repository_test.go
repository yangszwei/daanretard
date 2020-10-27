package persistence_test

import (
	"daanretard/internal/domain/user"
	"daanretard/internal/infrastructure/persistence"
	"testing"
)

var testUser = user.User{
	Email:      "persistence.user@example.com",
	Password:   []byte("test password"),
	Profile:    user.Profile{
		DisplayName: "persistence_user",
		FirstName:   "Persistence",
		LastName:    "User",
	},
}

// NewUserRepository return a UserRepository for testing
func NewUserRepository() *persistence.UserRepository {
	return persistence.NewUserRepository(DB)
}

func TestUserRepository_AutoMigrate(t *testing.T) {
	err := NewUserRepository().AutoMigrate()
	if err != nil {
		t.Error(err)
	}
}

func TestUserRepository_InsertOne(t *testing.T) {
	err := NewUserRepository().InsertOne(&testUser)
	if err != nil {
		t.Error(err)
	}
}

func TestUserRepository_FindOneByID(t *testing.T) {
	t.Run("should succeed", func(t *testing.T) {
		u, err := NewUserRepository().FindOneByID(testUser.ID)
		if err != nil {
			t.Fatal(err)
		}
		if u.ID != testUser.ID {
			t.Error(u)
		}
	})
	t.Run("should fail with: record not found", func(t *testing.T) {
		_, err := NewUserRepository().FindOneByID(0)
		if err == nil || err.Error() != "record not found" {
			t.Error(err)
		}
	})
}

func TestUserRepository_FindOneByEmail(t *testing.T) {
	t.Run("should succeed", func(t *testing.T) {
		u, err := NewUserRepository().FindOneByEmail(testUser.Email)
		if err != nil {
			t.Fatal(err)
		}
		if u.ID != testUser.ID {
			t.Error(u)
		}
	})
	t.Run("should fail with: record not found", func(t *testing.T) {
		_, err := NewUserRepository().FindOneByEmail("")
		if err == nil || err.Error() != "record not found" {
			t.Error(err)
		}
	})
}

func TestUserRepository_UpdateOne(t *testing.T) {
	err := NewUserRepository().UpdateOne(&testUser)
	if err != nil {
		t.Error(err)
	}
}

func TestUserRepository_DeleteOne(t *testing.T) {
	err := NewUserRepository().DeleteOne(&testUser)
	if err != nil {
		t.Error(err)
	}
}