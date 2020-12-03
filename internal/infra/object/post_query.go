package object

import "time"

// PostQuery query object for post.IRepository, Message field is used for
// searching with sql " message LIKE ?"
type PostQuery struct {
	Message       string
	Status        rune
	UserID        string
	CreatedAfter  time.Time
	CreatedBefore time.Time
	ReviewerID    string
	ReviewResult  rune
	FacebookID    string
	Limit         int
	Offset        int
}
