package mock_persistence_test

import (
	"daanretard/internal/infrastructure/mock_persistence"
	"testing"
)

var admins = mock_persistence.NewAdministratorRepository()

func TestAdministratorRepository_InsertOne(t *testing.T) {
	err := admins.InsertOne(1)
	if err != nil {
		t.Error(err)
	}
}

func TestAdministratorRepository_FindOneByID(t *testing.T) {
	_, err := admins.FindOneByID(1)
	if err != nil {
		t.Error(err)
	}
}

func TestAdministratorRepository_UpdateOne(t *testing.T) {
	admin, err := admins.FindOneByID(1)
	if err != nil {
		t.Error(err)
	}
	err = admins.UpdateOne(admin)
	if err != nil {
		t.Error(err)
	}
}

func TestAdministratorRepository_DeleteOne(t *testing.T) {
	admin, err := admins.FindOneByID(1)
	if err != nil {
		t.Error(err)
	}
	err = admins.DeleteOne(admin)
	if err != nil {
		t.Error(err)
	}
}
