package persistence

import "gorm.io/gorm"

// NewAttachmentRepository create a AttachmentRepository
func NewAttachmentRepository(db *DB) *AttachmentRepository {
	a := new(AttachmentRepository)
	a.db = db.Conn
	return a
}

// AttachmentSchema schema of AttachmentRepository
type AttachmentSchema struct {
	ID   uint32 `gorm:"autoIncrement"`
	Name string
}

// TableName set table name of AttachmentSchema
func (s *AttachmentSchema) TableName() string {
	return "attachments"
}

// AttachmentRepository implement attachment.IRepository
type AttachmentRepository struct {
	db *gorm.DB
}

// AutoMigrate set table schema
func (a *AttachmentRepository) AutoMigrate() error {
	return a.db.AutoMigrate(&AttachmentSchema{})
}

// InsertOne add an attachment
func (a *AttachmentRepository) InsertOne(name string) (uint32, error) {
	attachment := AttachmentSchema{Name: name}
	result := a.db.Create(&attachment)
	return attachment.ID, result.Error
}

// FindOneByID find an attachment
func (a *AttachmentRepository) FindOneByID(id uint32) (string, error) {
	var attachment AttachmentSchema
	result := a.db.First(&attachment, id)
	if result.Error != nil {
		return "", result.Error
	}
	return attachment.Name, nil
}

// DeleteOne delete an attachment
func (a *AttachmentRepository) DeleteOne(id uint32) error {
	var attachment AttachmentSchema
	attachment.ID = id
	return a.db.Delete(&attachment).Error
}
