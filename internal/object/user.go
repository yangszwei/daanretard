package object

import "time"

// UserProps data object of user.User
type UserProps struct {
	ID         uint32
	Email      string `validate:"omitempty,email,max=254"`
	Password   string `validate:"omitempty,min=8,max=254"`
	Profile    UserProfileProps
	IsVerified bool
	CreatedAt  time.Time
}
