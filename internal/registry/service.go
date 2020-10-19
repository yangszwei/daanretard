package registry

import (
	"daanretard/internal/domain/user"
	"daanretard/internal/infrastructure/persistence"
)

// Services services container
type Services struct {
	User *user.Service
}

func prepareServices(db *persistence.DB) (*Services, error) {
	s := new(Services)
	s.User = user.NewService(persistence.NewUserRepository(db.Conn))
	return s, nil
}