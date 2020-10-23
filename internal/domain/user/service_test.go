package user_test

import (
	"daanretard/internal/domain/user"
	"daanretard/internal/infrastructure/mock_persistence"
	"daanretard/internal/object"
	"testing"
)

var (
	s *user.Service
	id uint32
	props = object.UserProps{
		Name:     "testuser",
		Email:    "user.service@example.com",
		Password: "12345678",
	}
	profile = object.UserProfileProps{
		FirstName: "Test",
		LastName:  "User",
	}
)

func TestNewService(t *testing.T) {
	s = user.NewService(mock_persistence.NewUserRepository())
}

func TestService_Register(t *testing.T) {
	var err error
	id, err = s.Register(props, profile)
	if err != nil {
		t.Error(err)
	}
}

func TestService_Authenticate(t *testing.T) {
	i, err := s.Authenticate(props)
	if err != nil {
		t.Error(err)
	}
	if i != id {
		t.Error(i, id)
	}
}

func TestService_Delete(t *testing.T) {
	err := s.Delete(id)
	if err != nil {
		t.Error(err)
	}
}