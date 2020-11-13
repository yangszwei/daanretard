package object

import "time"

// Query used in IRepository for selecting posts
type PostQuery struct {
	Status        uint8
	UserID        uint32
	IPAddr        []byte
	CreatedAfter  time.Time
	CreatedBefore time.Time
	ReviewerID    uint32
	ReviewResult  uint8
	Limit         int
	Offset        int
}
