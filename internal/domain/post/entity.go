package post

import "time"

// Post aggregate root
type Post struct {
	ID          uint32 `gorm:"autoIncrement"`
	Status      uint8  `gorm:"index"`
	UserID      uint32 `gorm:"index"`
	IPAddr      []byte
	UserAgent   string `gorm:"type:text"`
	Message     string `gorm:"type:text;size:60000"`
	Attachments string
	Review      Review `gorm:"foreignKey:PostID"`
	FacebookID  string `gorm:"type:varchar(40)"`
	CreatedAt   time.Time
}