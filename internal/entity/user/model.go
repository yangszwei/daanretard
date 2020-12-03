package user

import "time"

// User user object used inside the package, for the exposed object,
// refer to object.User
type User struct {
	ID          string `gorm:"type:varchar(128)"`
	Name        string `gorm:"size:100"`
	AccessToken string `gorm:"size:255"`
	ExpiresAt   time.Time
	CreatedAt   time.Time
}
