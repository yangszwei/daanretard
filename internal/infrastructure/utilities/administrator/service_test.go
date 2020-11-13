package administrator_test

import (
	"daanretard/internal/infrastructure/mock_persistence"
	"daanretard/internal/infrastructure/utilities/administrator"
	"testing"
)

var service = administrator.NewService(mock_persistence.NewAdministratorRepository())

func TestService_Add(t *testing.T) {
	err := service.Add(2)
	if err != nil {
		t.Error(err)
	}
}

func TestService_SetFbAccessToken(t *testing.T) {
	err := service.SetFbAccessToken(2, "test access token")
	if err != nil {
		t.Error(err)
	}
}

func TestService_GetFbAccessToken(t *testing.T) {
	token, err := service.GetFbAccessToken(2)
	if err != nil {
		t.Error(err)
	}
	if token != "test access token" {
		t.Error(token)
	}
}

func TestService_Delete(t *testing.T) {
	err := service.Delete(2)
	if err != nil {
		t.Error(err)
	}
}
