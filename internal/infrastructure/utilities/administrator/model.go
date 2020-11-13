package administrator

import "time"

// Administrator model
type Administrator struct {
	UserID        uint32 `gorm:"primaryKey"`
	FbAccessToken string
	CreatedAt     time.Time
}
