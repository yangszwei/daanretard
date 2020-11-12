package mock_persistence_test

import (
	"daanretard/internal/infrastructure/mock_persistence"
	"testing"
)

var (
	attachments = mock_persistence.NewAttachmentRepository()
	id uint32
)

func TestAttachmentRepository_InsertOne(t *testing.T) {
	var err error
	id, err = attachments.InsertOne("test name")
	if err != nil {
		t.Error(err)
	}
}

func TestAttachmentRepository_FindOneByID(t *testing.T) {
	name, err := attachments.FindOneByID(id)
	if err != nil {
		t.Error(err)
	}
	if name != "test name" {
		t.Error(name)
	}
}

func TestAttachmentRepository_DeleteOne(t *testing.T) {
	err := attachments.DeleteOne(id)
	if err != nil {
		t.Error(err)
	}
}