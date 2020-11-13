package attachment_test

import (
	"daanretard/internal/infrastructure/mock_persistence"
	"daanretard/internal/infrastructure/utilities/attachment"
	"testing"
)

var (
	id      uint32
	service = attachment.NewService(mock_persistence.NewAttachmentRepository())
)

func TestService_Add(t *testing.T) {
	var err error
	id, err = service.Add("test name")
	if err != nil {
		t.Error(err)
	}
}

func TestService_GetOne(t *testing.T) {
	name, err := service.GetOne(id)
	if err != nil {
		t.Error(err)
	}
	if name != "test name" {
		t.Error(name)
	}
}

func TestService_Delete(t *testing.T) {
	err := service.Delete(id)
	if err != nil {
		t.Error(err)
	}
}
