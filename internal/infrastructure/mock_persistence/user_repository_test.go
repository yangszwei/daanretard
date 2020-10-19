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
	}
)

func TestUserRepository_Setup(t *testing.T) {
	users = mock_persistence.NewUserRepository()
}

func TestUserRepository_InsertOne(t *testing.T) {
	err := users.InsertOne(&testUser)
	if err != nil {
		t.Error(err)
	}
	t.Run("used name", func(t *testing.T) {
		err = users.InsertOne(&entity.User{ Name: testUser.Name })
		if err == nil || err.Error() != "name already exist" {
			t.Error(err)
		}
	})
	t.Run("used email", func(t *testing.T) {
		err = users.InsertOne(&entity.User{ Email: testUser.Email })
		if err == nil || err.Error() != "email already exist" {
			t.Error(err)
		}
	})
}

func TestUserRepository_FindOneByID(t *testing.T) {
	t.Run("Local", func(t *testing.T) {
		u, err := users.FindOneByID(testUser.ID)
		if err != nil {
			t.Error(err)
		}
		if u.ID != testUser.ID {
			t.Error(u)
		}
	})
	t.Run("Database", func(t *testing.T) {
		_, err := mock_persistence.NewUserRepository().FindOneByID(testUser.ID)
		if err == nil || err.Error() != "record not found" {
			t.Error(err)
		}
	})
}

func TestUserRepository_FindOneByName(t *testing.T) {
	t.Run("Local", func(t *testing.T) {
		u, err := users.FindOneByName(testUser.Name)
		if err != nil {
			t.Error(err)
		}
		if u.ID != testUser.ID {
			t.Error(u)
		}
	})
	t.Run("Database", func(t *testing.T) {
		_, err := mock_persistence.NewUserRepository().FindOneByName(testUser.Name)
		if err == nil || err.Error() != "record not found" {
			t.Error(err)
		}
	})
}

func TestUserRepository_FindOneByEmail(t *testing.T) {
	t.Run("Local", func(t *testing.T) {
		u, err := users.FindOneByEmail(testUser.Email)
		if err != nil {
			t.Error(err)
		}
		if u.ID != testUser.ID {
			t.Error(u)
		}
	})
	t.Run("Database", func(t *testing.T) {
		_, err := mock_persistence.NewUserRepository().FindOneByEmail(testUser.Email)
		if err == nil || err.Error() != "record not found" {
			t.Error(err)
		}
	})
}

func TestUserRepository_SaveOne(t *testing.T) {
	err := users.SaveOne(testUser.ID)
	if err != nil {
		t.Error(err)
	}
}

func TestUserRepository_DeleteOne(t *testing.T) {
	err := users.DeleteOne(testUser.ID)
	if err != nil {
		t.Error(err)
	}
}