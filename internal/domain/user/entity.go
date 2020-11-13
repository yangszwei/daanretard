package user

import "time"

// User aggregate root
type User struct {
	ID         uint32 `gorm:"autoIncrement"`
	Email      string `gorm:"unique;size:254"`
	Password   []byte
	Profile    Profile `gorm:"foreignKey:UserID"`
	IsVerified bool
	CreatedAt  time.Time
}
