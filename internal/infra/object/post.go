package object

import (
	"net"
	"time"
)

// Post post object for usage outside of the entity
type Post struct {
	ID          uint32
	Status      rune
	UserID      string `validate:"max=128"`
	IPAddr      net.IP
	UserAgent   string
	Message     string `validate:"max=50000"`
	Attachments []uint32
	Review      PostReview
	FacebookID  string `validate:"max=40"`
	CreatedAt   time.Time
}
