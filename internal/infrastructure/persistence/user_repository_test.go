package persistence_test

import (
	entity "daanretard/internal/domain/user"
	"daanretard/internal/infrastructure/persistence"
	"testing"
)

var (
	users *persistence.UserRepository
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

func TestNewUserRepository(t *testing.T) {
	users = persistence.NewUserRepository(DB)
}

func TestUserRepository_AutoMigrate(t *testing.T) {
	err := users.AutoMigrate()
	if err != nil {
		t.Error(err)
	}
}

func TestUserRepository_InsertOne(t *testing.T) {
	err := users.InsertOne(&testUser)
	if err != nil {
		t.Error(err)
	}
}

func TestUserRepository_FindOne(t *testing.T) {
	_, err := users.FindOne(entity.Query{ ID: testUser.ID })
	if err != nil {
		t.Error(err)
	}
}

func TestUserRepository_FindAll(t *testing.T) {
	_, err := users.FindAll(entity.Query{ ID: testUser.ID })
	if err != nil {
		t.Error(err)
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