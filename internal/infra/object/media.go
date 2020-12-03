package object

import "time"

// Media media object for usage outside of the entity
type Media struct {
	ID         uint32
	UserID     string `validate:"max=128"`
	Name       string `validate:"max=200"`
	FacebookID string `validate:"max=128"`
	CreatedAt  time.Time
}
