package persistence

import (
	"daanretard/internal/domain/user"
	"gorm.io/gorm"
)

// NewUserRepository create a UserRepository
func NewUserRepository(db *DB) *UserRepository {
	u := new(UserRepository)
	u.db = db.Conn
	return u
}

// UserRepository implement user.IRepository
type UserRepository struct {
	db *gorm.DB
}

// AutoMigrate set table schema
func (u *UserRepository) AutoMigrate() error {
	return u.db.AutoMigrate(&user.User{}, &user.Profile{})
}

// InsertOne insert a user
func (u *UserRepository) InsertOne(user *user.User) error {
	result := u.db.Create(user)
	return result.Error
}

// FindOneByID find a user by its ID
func (u *UserRepository) FindOneByID(id uint32) (*user.User, error) {
	var record user.User
	result := u.db.Preload("Profile").First(&record, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &record, nil
}

// FindOneByEmail find a user by its email
func (u *UserRepository) FindOneByEmail(email string) (*user.User, error) {
	var record user.User
	result := u.db.Where("email = ?", email).First(&record)
	if result.Error != nil {
		return nil, result.Error
	}
	return &record, nil
}

// UpdateOne save a user
func (u *UserRepository) UpdateOne(user *user.User) error {
	result := u.db.Save(user)
	return result.Error
}

// DeleteOne delete a user
func (u *UserRepository) DeleteOne(user *user.User) error {
	result := u.db.Select("Profile").Delete(user)
	return result.Error
}
