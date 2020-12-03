package object

import "time"

// User user object for usage outside of the entity
type User struct {
	ID          string `validate:"max=128"`
	Name        string `validate:"max=100"`
	AccessToken string `validate:"max=255"`
	ExpiresAt   time.Time
	CreatedAt   time.Time
}
