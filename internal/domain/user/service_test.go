package user_test

import (
	"daanretard/internal/domain/user"
	"daanretard/internal/infrastructure/mock_persistence"
	"daanretard/internal/object"
	"testing"
)

var (
	service  = user.NewService(mock_persistence.NewUserRepository())
	testUser = object.UserProps{
		Email:    "user.service@example.com",
		Password: "12345678",
		Profile: object.UserProfileProps{
			DisplayName: "User Service Test",
			FirstName:   "User",
			LastName:    "Service",
		},
	}
)

func TestService_Register(t *testing.T) {
	t.Run("should succeed", func(t *testing.T) {
		id, err := service.Register(testUser)
		if err != nil {
			t.Error(err)
		}
		testUser.ID = id
	})
	t.Run("should fail with: invalid", func(t *testing.T) {
		_, err := service.Register(object.UserProps{})
		if err == nil || err.Error() != "invalid credentials" {
			t.Error(err)
		}
	})
	t.Run("should fail with: invalid (2)", func(t *testing.T) {
		_, err := service.Register(object.UserProps{
			Email:    "test",
			Password: "test",
			Profile: object.UserProfileProps{
				DisplayName: "test",
				FirstName:   "test",
				LastName:    "test",
			},
		})
		if err == nil || err.Error() != "invalid credentials" {
			t.Error(err)
		}
	})
	t.Run("should fail with: email already taken", func(t *testing.T) {
		_, err := service.Register(testUser)
		if err == nil || err.Error() != "email already taken" {
			t.Error(err)
		}
	})
}

func TestService_GetProps(t *testing.T) {
	props, err := service.GetProps(testUser.ID)
	if err != nil {
		t.Error(err)
	}
	if props.ID != testUser.ID {
		t.Error(err)
	}
}

func TestService_Authenticate(t *testing.T) {
	t.Run("should succeed", func(t *testing.T) {
		id, err := service.Authenticate(testUser.Email, testUser.Password)
		if err != nil {
			t.Error(err)
		}
		if id != testUser.ID {
			t.Error(id)
		}
	})
	t.Run("should fail with: invalid credentials (record not found)", func(t *testing.T) {
		_, err := service.Authenticate("", testUser.Password)
		if err == nil || err.Error() != "invalid credentials" {
			t.Error(err)
		}
	})
	t.Run("should fail with: invalid credentials", func(t *testing.T) {
		_, err := service.Authenticate(testUser.Email, "")
		if err == nil || err.Error() != "invalid credentials" {
			t.Error(err)
		}
	})
}

func TestService_AuthenticateWithID(t *testing.T) {
	t.Run("should succeed", func(t *testing.T) {
		err := service.AuthenticateWithID(testUser.ID, testUser.Password)
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("should fail with: invalid credentials (record not found)", func(t *testing.T) {
		err := service.AuthenticateWithID(0, testUser.Password)
		if err == nil || err.Error() != "invalid credentials" {
			t.Error(err)
		}
	})
	t.Run("should fail with: invalid credentials", func(t *testing.T) {
		err := service.AuthenticateWithID(testUser.ID, "")
		if err == nil || err.Error() != "invalid credentials" {
			t.Error(err)
		}
	})
}

func TestService_UpdateEmail(t *testing.T) {
	t.Run("should succeed", func(t *testing.T) {
		err := service.UpdateEmail(testUser.ID, "user.service2@example.com")
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("should fail with: record not found", func(t *testing.T) {
		err := service.UpdateEmail(0, "")
		if err == nil || err.Error() != "record not found" {
			t.Error(err)
		}
	})
	t.Run("should fail with: invalid credentials", func(t *testing.T) {
		err := service.UpdateEmail(testUser.ID, "invalid")
		if err == nil || err.Error() != "invalid credentials" {
			t.Error(err)
		}
	})
}

func TestService_UpdatePassword(t *testing.T) {
	t.Run("should succeed", func(t *testing.T) {
		err := service.UpdatePassword(testUser.ID, "new password")
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("should fail with: record not found", func(t *testing.T) {
		err := service.UpdatePassword(0, "")
		if err == nil || err.Error() != "record not found" {
			t.Error(err)
		}
	})
	t.Run("should fail with: invalid credentials", func(t *testing.T) {
		// shorter than 8 characters
		err := service.UpdatePassword(testUser.ID, "invalid")
		if err == nil || err.Error() != "invalid credentials" {
			t.Error(err)
		}
	})
}

func TestService_MarkAsVerified(t *testing.T) {
	t.Run("should succeed", func(t *testing.T) {
		err := service.MarkAsVerified(testUser.ID)
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("should fail with: record not found", func(t *testing.T) {
		err := service.MarkAsVerified(0)
		if err == nil || err.Error() != "record not found" {
			t.Error(err)
		}
	})
}

func TestService_AddAdministrator(t *testing.T) {
	t.Run("should succeed", func(t *testing.T) {
		err := service.AddAdministrator(testUser.ID)
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("should fail with: record not found", func(t *testing.T) {
		err := service.AddAdministrator(0)
		if err == nil || err.Error() != "record not found" {
			t.Error(err)
		}
	})
}

func TestService_IsAdministrator(t *testing.T) {
	t.Run("should succeed", func(t *testing.T) {
		isAdmin, err := service.IsAdministrator(testUser.ID)
		if err != nil {
			t.Error(err)
		}
		if !isAdmin {
			t.Error(isAdmin)
		}
	})
	t.Run("should fail with: record not found", func(t *testing.T) {
		_, err := service.IsAdministrator(0)
		if err == nil || err.Error() != "record not found" {
			t.Error(err)
		}
	})
}

func TestService_RemoveAdministrator(t *testing.T) {
	t.Run("should succeed", func(t *testing.T) {
		err := service.RemoveAdministrator(testUser.ID)
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("should fail with: record not found", func(t *testing.T) {
		err := service.RemoveAdministrator(0)
		if err == nil || err.Error() != "record not found" {
			t.Error(err)
		}
	})
}

func TestService_UpdateProfile(t *testing.T) {
	t.Run("should succeed", func(t *testing.T) {
		err := service.UpdateProfile(testUser.ID, object.UserProfileProps{
			DisplayName: "User Service Test 2",
			FirstName:   "User 2",
			LastName:    "Service 2",
		})
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("should fail with: record not found", func(t *testing.T) {
		err := service.UpdateProfile(0, object.UserProfileProps{})
		if err == nil || err.Error() != "record not found" {
			t.Error(err)
		}
	})
	t.Run("should fail with: invalid credentials", func(t *testing.T) {
		err := service.UpdateProfile(testUser.ID, object.UserProfileProps{
			DisplayName: "string longer than 50 characters string longer than 50 characters",
		})
		if err == nil || err.Error() != "invalid credentials" {
			t.Error(err)
		}
	})
}

func TestService_Delete(t *testing.T) {
	t.Run("should succeed", func(t *testing.T) {
		err := service.Delete(testUser.ID)
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("should fail with: record not found", func(t *testing.T) {
		err := service.Delete(testUser.ID)
		if err == nil || err.Error() != "record not found" {
			t.Error(err)
		}
	})
}
