package persistence_test

import (
	"daanretard/internal/infrastructure/persistence"
	"testing"
)

func NewAdministratorRepository() *persistence.AdministratorRepository {
	return persistence.NewAdministratorRepository(DB)
}

func TestNewAdministratorRepository(t *testing.T) {
	err := NewAdministratorRepository().AutoMigrate()
	if err != nil {
		t.Error(err)
	}
}

func TestAdministratorRepository_InsertOne(t *testing.T) {
	err := NewAdministratorRepository().InsertOne(1)
	if err != nil {
		t.Error(err)
	}
}

func TestAdministratorRepository_FindOneByID(t *testing.T) {
	_, err := NewAdministratorRepository().FindOneByID(1)
	if err != nil {
		t.Error(err)
	}
}

func TestAdministratorRepository_UpdateOne(t *testing.T) {
	admin, err := NewAdministratorRepository().FindOneByID(1)
	if err != nil {
		t.Error(err)
	}
	err = NewAdministratorRepository().UpdateOne(admin)
	if err != nil {
		t.Error(err)
	}
}

func TestAdministratorRepository_DeleteOne(t *testing.T) {
	admin, err := NewAdministratorRepository().FindOneByID(1)
	if err != nil {
		t.Error(err)
	}
	err = NewAdministratorRepository().DeleteOne(admin)
	if err != nil {
		t.Error(err)
	}
}
