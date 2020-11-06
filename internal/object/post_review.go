package object

import "time"

// PostReview object
type PostReview struct {
	UserID uint32
	Result uint8
	Message string
	CreatedAt time.Time
}