package user

import "time"

// Session object of User
type Session struct {
	ID        uint32 `gorm:"autoIncrement"`
	UserID    uint32 `gorm:"index"`
	CreatedAt time.Time
}

// TableName set table name
func (s *Session) TableName() string {
	return "user_sessions"
}