package post

import "time"

// Review review object used inside the package, for the exposed object,
// refer to object.PostReview
type Review struct {
	PostID    uint32 `gorm:"primaryKey"`
	UserID    string `gorm:"type:varchar(128);index"`
	Result    rune
	Message   string `gorm:"type:text;size:200"`
	CreatedAt time.Time
}

// Codes for Review.Result
const (
	ReviewApproved = 'a'
	ReviewRejected = 'r'
)

// TableName set table name of Review with gorm
func (r *Review) TableName() string {
	return "post_reviews"
}
