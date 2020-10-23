package session_test

import (
	"daanretard/internal/infrastructure/mock_persistence"
	"daanretard/internal/object"
	"daanretard/internal/service/session"
	"testing"
)

var (
	service *session.Service
	testProps object.SessionProps
)

func TestNewService(t *testing.T) {
	service = session.NewService(mock_persistence.NewSessionRepository())
}

func TestService_Open(t *testing.T) {
	var err error
	testProps, err = service.Open(1)
	if err != nil {
		t.Error(err)
	}
}

func TestService_Extend(t *testing.T) {
	err := service.Extend(testProps.ID)
	if err != nil {
		t.Error(err)
	}
	t.Run("not found", func(t *testing.T) {
		err := service.Extend(0)
		if err == nil || err.Error() != "record not found" {
			t.Error(err)
		}
	})
}

func TestService_Close(t *testing.T) {
	err := service.Close(testProps.ID)
	if err != nil {
		t.Error(err)
	}
	t.Run("not found", func(t *testing.T) {
		err := service.Close(0)
		if err == nil || err.Error() != "record not found" {
			t.Error(err)
		}
	})
}