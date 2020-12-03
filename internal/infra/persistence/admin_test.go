package persistence_test

import (
	"daanretard/internal/entity/admin"
	"daanretard/internal/infra/persistence"
	"testing"
)

var (
	a = admin.Admin{
		UserID:      "test_admin",
		AccessToken: "test_access_token",
	}
	ar *persistence.AdminRepo
)

func TestNewAdminRepo(t *testing.T) {
	ar = persistence.NewAdminRepo(newDB())
}

func TestAdminRepo_AutoMigrate(t *testing.T) {
	if err := ar.AutoMigrate(); err != nil {
		t.Error(err)
	}
}

func TestAdminRepo_InsertOne(t *testing.T) {
	if err := ar.InsertOne(&a); err != nil {
		t.Error(err)
	}
}

func TestAdminRepo_FindOneByUserID(t *testing.T) {
	if r, err := ar.FindOneByUserID(a.UserID); err != nil || r.UserID != a.UserID {
		t.Error(r, err)
	}
}

func TestAdminRepo_FindMany(t *testing.T) {
	if r, err := ar.FindMany(10, 0); err != nil || len(r) != 1 {
		t.Error(r, err)
	}
}

func TestAdminRepo_UpdateOne(t *testing.T) {
	if err := ar.UpdateOne(&a); err != nil {
		t.Error(err)
	}
}

func TestAdminRepo_DeleteOne(t *testing.T) {
	if err := ar.DeleteOne(&a); err != nil {
		t.Error(err)
	}
}
