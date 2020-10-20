package post

// Query model
type Query struct {
	ID           uint32
	Status       uint8
	SubmitterID  uint32
	ReviewerID   uint32
}