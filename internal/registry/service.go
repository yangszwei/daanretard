package registry

import (
	"daanretard/internal/domain/user"
	"daanretard/internal/infrastructure/application"
	"daanretard/internal/infrastructure/persistence"
)

// SetupService setup services
func SetupService(db *persistence.DB) (*application.Services, error) {
	s := new(application.Services)
	users := persistence.NewUserRepository(db)
	var err error
	err = users.AutoMigrate()
	if err != nil {
		return nil, err
	}
	s.User = user.NewService(users)
	return s, nil
}