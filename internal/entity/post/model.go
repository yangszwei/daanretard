package post

import "time"

// Post post object used inside the package, for the exposed object,
// refer to object.Post
type Post struct {
	ID          uint32 `gorm:"autoIncrement"`
	Status      rune   `gorm:"index"`
	UserID      string `gorm:"type:varchar(128);index"`
	IPAddr      []byte
	UserAgent   string `gorm:"type:text"`
	Message     string `gorm:"type:text;size:50000"`
	Attachments string
	Review      Review `gorm:"foreignKey:PostID"`
	FacebookID  string `gorm:"type:varchar(40)"`
	CreatedAt   time.Time
}

// Codes for Post.Status
const (
	StatusSubmitted = 's'
	StatusReviewed  = 'r'
	StatusPublished = 'p'
)
