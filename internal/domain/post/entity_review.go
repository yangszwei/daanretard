package post

// Review child entity of Post
type Review struct {
	PostID     uint32 `gorm:"primaryKey"`
	ReviewerID uint32 `gorm:"index"`
	Result     uint8  `gorm:"index"`
	Message    string `gorm:"size:200"`
}

// TableName set table name
func (r *Review) TableName() string {
	return "post_reviews"
}