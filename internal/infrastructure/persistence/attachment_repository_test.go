package persistence_test

import (
	"daanretard/internal/infrastructure/persistence"
	"testing"
)

var id uint32

func NewAttachmentRepository() *persistence.AttachmentRepository {
	return persistence.NewAttachmentRepository(DB)
}

func TestAttachmentRepository_AutoMigrate(t *testing.T) {
	err := NewAttachmentRepository().AutoMigrate()
	if err != nil {
		t.Error(err)
	}
}

func TestAttachmentRepository_InsertOne(t *testing.T) {
	var err error
	id, err = NewAttachmentRepository().InsertOne("test name")
	if err != nil {
		t.Error(err)
	}
}

func TestAttachmentRepository_FindOneByID(t *testing.T) {
	name, err := NewAttachmentRepository().FindOneByID(id)
	if err != nil {
		t.Error(err)
	}
	if name != "test name" {
		t.Error(name)
	}
}

func TestAttachmentRepository_DeleteOne(t *testing.T) {
	err := NewAttachmentRepository().DeleteOne(id)
	if err != nil {
		t.Error(err)
	}
}
