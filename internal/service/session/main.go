// Package session implement session service
package session

import "time"

// Session model
type Session struct {
	ID        uint64 `gorm:"autoIncrement"`
	UserID    uint32 `gorm:"index"`
	CreatedAt time.Time
	ExpiresAt time.Time
}
