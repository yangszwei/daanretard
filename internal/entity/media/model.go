package media

import "time"

// Media media object used inside the package, for the exposed object,
// refer to object.Media. This entity is used to persist information of
// user uploaded file
type Media struct {
	ID         uint32 `gorm:"autoIncrement"`
	UserID     string `gorm:"type:varchar(128);index"`
	Name       string `gorm:"type:varchar(200);index"`
	FacebookID string `gorm:"index"`
	CreatedAt  time.Time
}

// TableName set table name of Media with gorm
func (m *Media) TableName() string {
	return "media"
}
