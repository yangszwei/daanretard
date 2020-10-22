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
		Sessions: []entity.Session{
			{},
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