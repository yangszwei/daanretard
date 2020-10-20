package post

// Submission child entity of Post
type Submission struct {
	PostID      uint32 `gorm:"primaryKey"`
	SubmitterID uint32 `gorm:"index"`
	Message     string
	Attachments []string
	IPAddr      []byte
	UserAgent   string
}