package persistence

import (
	"daanretard/internal/infrastructure/utilities/administrator"
	"gorm.io/gorm"
)

// NewAdministratorRepository create an AdministratorRepository
func NewAdministratorRepository(db *DB) *AdministratorRepository {
	a := new(AdministratorRepository)
	a.db = db.Conn
	return a
}

// AdministratorRepository implement administrator.IRepository
type AdministratorRepository struct {
	db *gorm.DB
}

// AutoMigrate set table schema
func (a *AdministratorRepository) AutoMigrate() error {
	return a.db.AutoMigrate(&administrator.Administrator{})
}

// InsertOne add an administrator
func (a *AdministratorRepository) InsertOne(userID uint32) error {
	return a.db.Create(&administrator.Administrator{
		UserID: userID,
	}).Error
}

// FindOneByID find an administrator
func (a *AdministratorRepository) FindOneByID(id uint32) (*administrator.Administrator, error) {
	var admin administrator.Administrator
	result := a.db.First(&admin, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &admin, nil
}

// UpdateOne update an administrator
func (a *AdministratorRepository) UpdateOne(admin *administrator.Administrator) error {
	return a.db.Save(admin).Error
}

// DeleteOne delete an administrator
func (a *AdministratorRepository) DeleteOne(admin *administrator.Administrator) error {
	return a.db.Delete(&admin).Error
}
