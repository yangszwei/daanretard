package persistence

import (
	entity "daanretard/internal/domain/user"
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

// AutoMigrate setup table schema
func (u *UserRepository) AutoMigrate() error {
	return u.db.AutoMigrate(
		entity.User{},
		entity.Profile{},
	)
}

// InsertOne insert a user
func (u *UserRepository) InsertOne(user *entity.User) error {
	result := u.db.Create(user)
	return result.Error
}

// FindOne find a user with user.Query
func (u *UserRepository) FindOne(query entity.Query) (*entity.User, error) {
	var user entity.User
	result := u.db.Where(&entity.User{
		ID: query.ID,
		Name: query.Name,
		Email: query.Email,
	}).Preload("Profile").First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// FindAll find a user with user.Query
func (u *UserRepository) FindAll(query entity.Query) ([]*entity.User, error) {
	var users []*entity.User
	result := u.db.Where(&entity.User{
		ID: query.ID,
		Name: query.Name,
		Email: query.Email,
	}).Preload("Profile").Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

// SaveOne save a user
func (u *UserRepository) SaveOne(user *entity.User) error {
	result := u.db.Save(user)
	return result.Error
}

// DeleteOne delete a user
func (u *UserRepository) DeleteOne(user *entity.User) error {
	result := u.db.Select("Profile").Delete(user)
	return result.Error
}