package user

// Profile non-critical information of User ; child entity
type Profile struct {
	UserID      uint32 `gorm:"primaryKey"`
	DisplayName string `gorm:"unique;size:50"`
	FirstName   string `gorm:"size:50"`
	LastName    string `gorm:"size:50"`
}

// TableName sets table name, used in gorm AutoMigrate
func (p *Profile) TableName() string {
	return "user_profiles"
}