package post

import "time"

// Review child entity of Post
type Review struct {
	PostID    uint32 `gorm:"primaryKey"`
	UserID    uint32 `gorm:"index"`
	Result    uint8
	Message   string
	CreatedAt time.Time
}

// TableName set table name of Review
func (r *Review) TableName() string {
	return "post_reviews"
}
