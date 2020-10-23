// Package user define User entity
package user

import "time"

// User aggregate root
type User struct {
	ID         uint32 `gorm:"autoIncrement"`
	Name       string `gorm:"unique"`
	Email      string `gorm:"unique;size:254"`
	Password   []byte
	Profile    Profile
	IsVerified bool
	CreatedAt  time.Time
}