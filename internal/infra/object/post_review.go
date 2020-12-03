package object

import "time"

// PostReview post.Review object for usage outside of the entity
type PostReview struct {
	PostID    uint32
	UserID    string `validate:"max=128"`
	Result    rune
	Message   string `validate:"max=200"`
	CreatedAt time.Time
}
