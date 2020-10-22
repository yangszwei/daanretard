package mock_persistence_test

import (
	entity "daanretard/internal/domain/user"
	"daanretard/internal/infrastructure/mock_persistence"
	"testing"
)

var (
	users *mock_persistence.UserRepository
	testUser = entity.User{
		Name:     "user_repo",
		Email:    "user_repo@example.com",
		Password: []byte("12345678"),
		Profile: entity.Profile{
			FirstName: "Test",
			LastName:  "User",
		},
		Sessions: []entity.Session{
			{},
		},
	}
)

func TestNewUserRepository(t *testing.T) {
	users = mock_persistence.NewUserRepository()
}

func TestUserRepository_InsertOne(t *testing.T) {
	err := users.InsertOne(&testUser)
	if err != nil {
		t.Error(err)
	}
}

func TestUserRepository_FindOne(t *testing.T) {
	t.Run("find by ID", func(t *testing.T) {
		_, err := users.FindOne(entity.Query{ ID: testUser.ID })
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("find by other fields", func(t *testing.T) {
		users, err := users.FindAll(entity.Query{
			Name: testUser.Name,
			Email: testUser.Email,
		})
		if err != nil {
			t.Fatal(err)
		}
		if users[0].ID != testUser.ID {
			t.Error(users)
		}
	})
	t.Run("not found", func(t *testing.T) {
		_, err := users.FindOne(entity.Query{})
		if err == nil || err.Error() != "record not found" {
			t.Error(err)
		}
	})
}

func TestUserRepository_FindAll(t *testing.T) {
	_, err := users.FindAll(entity.Query{ ID: testUser.ID })
	if err != nil {
		t.Error(err)
	}
}

func TestUserRepository_FindOneBySessionID(t *testing.T) {
	if len(testUser.Sessions) >= 1 {
		_, err := users.FindOneBySessionID(testUser.Sessions[0].ID)
		if err != nil {
			t.Error(err)
		}
	} else {
		t.Error("sessions empty")
	}
}

func TestUserRepository_SaveOne(t *testing.T) {
	err := users.SaveOne(&testUser)
	if err != nil {
		t.Error(err)
	}
}

func TestUserRepository_DeleteOne(t *testing.T) {
	err := users.DeleteOne(&testUser)
	if err != nil {
		t.Error(err)
	}
}