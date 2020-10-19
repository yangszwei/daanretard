package user

// Profile child entity of User
type Profile struct {
	UserID    uint32 `gorm:"primaryKey"`
	FirstName string `gorm:"size:50"`
	LastName  string `gorm:"size:50"`
}

// TableName set table name
func (p *Profile) TableName() string {
	return "user_profiles"
}