package mock_persistence_test

import (
	"daanretard/internal/domain/user"
	"daanretard/internal/infrastructure/mock_persistence"
	"testing"
)

var (
	testUserRepo *mock_persistence.UserRepository
	testUser = user.User{
		Email:      "mock_persistence.user@example.com",
		Password:   []byte("test password"),
		Profile:    user.Profile{
			DisplayName: "persistence_user",
			FirstName:   "Persistence",
			LastName:    "User",
		},
	}
)

func init() {
	testUserRepo = mock_persistence.NewUserRepository()
}

func TestUserRepository_InsertOne(t *testing.T) {
	err := testUserRepo.InsertOne(&testUser)
	if err != nil {
		t.Error(err)
	}
}

func TestUserRepository_FindOneByID(t *testing.T) {
	t.Run("should succeed", func(t *testing.T) {
		u, err := testUserRepo.FindOneByID(testUser.ID)
		if err != nil {
			t.Fatal(err)
		}
		if u.ID != testUser.ID {
			t.Error(u)
		}
	})
	t.Run("should fail with: record not found", func(t *testing.T) {
		_, err := testUserRepo.FindOneByID(0)
		if err == nil || err.Error() != "record not found" {
			t.Error(err)
		}
	})
}

func TestUserRepository_FindOneByEmail(t *testing.T) {
	t.Run("should succeed", func(t *testing.T) {
		u, err := testUserRepo.FindOneByEmail(testUser.Email)
		if err != nil {
			t.Fatal(err)
		}
		if u.ID != testUser.ID {
			t.Error(u)
		}
	})
	t.Run("should fail with: record not found", func(t *testing.T) {
		_, err := testUserRepo.FindOneByEmail("")
		if err == nil || err.Error() != "record not found" {
			t.Error(err)
		}
	})
}

func TestUserRepository_UpdateOne(t *testing.T) {
	err := testUserRepo.UpdateOne(&testUser)
	if err != nil {
		t.Error(err)
	}
}

func TestUserRepository_DeleteOne(t *testing.T) {
	err := testUserRepo.DeleteOne(&testUser)
	if err != nil {
		t.Error(err)
	}
}