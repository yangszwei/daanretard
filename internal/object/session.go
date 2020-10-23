package object

import "time"

// SessionProps session props
type SessionProps struct {
	ID        uint64
	UserID    uint32
	CreatedAt time.Time
	ExpiresAt time.Time
	IsExpired bool
	IsAuthenticatedUser bool
}