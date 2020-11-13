package object

import "time"

// Post object
type Post struct {
	ID          uint32
	Status      uint8
	UserID      uint32
	IPAddr      []byte
	UserAgent   string
	Message     string `validate:"max=60000"`
	Attachments []string
	FacebookID  string `validate:"max=40"`
	Review      PostReview
	CreatedAt   time.Time
}
