package persistence_test

import (
	"daanretard/internal/entity/media"
	"daanretard/internal/infra/persistence"
	"testing"
)

var (
	m = media.Media{
		UserID:     "test_user",
		Name:       "Test File",
		FacebookID: "test fbID",
	}
	mr *persistence.MediaRepo
)

func TestNewMediaRepo(t *testing.T) {
	mr = persistence.NewMediaRepo(newDB())
}

func TestMediaRepo_AutoMigrate(t *testing.T) {
	if err := mr.AutoMigrate(); err != nil {
		t.Error(err)
	}
}

func TestMediaRepo_InsertOne(t *testing.T) {
	if err := mr.InsertOne(&m); err != nil {
		t.Error(err)
	}
}

func TestMediaRepo_FindOneByID(t *testing.T) {
	if r, err := mr.FindOneByID(m.ID); err != nil || r.ID != m.ID {
		t.Error(r, err)
	}
}

func TestMediaRepo_UpdateOne(t *testing.T) {
	if err := mr.UpdateOne(&m); err != nil {
		t.Error(err)
	}
}

func TestMediaRepo_DeleteOne(t *testing.T) {
	if err := mr.DeleteOne(&m); err != nil {
		t.Error(err)
	}
}
