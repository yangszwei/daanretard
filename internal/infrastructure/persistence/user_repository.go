package persistence

import (
	entity "daanretard/internal/domain/user"
	"errors"
	"gorm.io/gorm"
)

// NewUserRepository create a UserRepository
func NewUserRepository(db *gorm.DB) *UserRepository {
	u := new(UserRepository)
	u.DB = db
	u.ids = make(map[uint32]*entity.User)
	u.names = make(map[string]*entity.User)
	u.emails = make(map[string]*entity.User)
	return u
}

// UserRepository implement user.IRepository
type UserRepository struct {
	DB     *gorm.DB
	ids    map[uint32]*entity.User
	names  map[string]*entity.User
	emails map[string]*entity.User
}

// AutoMigrate setup table
func (u *UserRepository) AutoMigrate() error {
	err := u.DB.AutoMigrate(
		&entity.User{},
		&entity.Profile{},
	)
	return err
}

// InsertOne insert a user
func (u *UserRepository) InsertOne(user *entity.User) error {
	if _, exist := u.names[user.Name] ; exist {
		return errors.New("name already exist")
	}
	if _, exist := u.emails[user.Email] ; exist {
		return errors.New("email already exist")
	}
	result := u.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}
	u.ids[user.ID] = user
	u.names[user.Name] = user
	u.emails[user.Email] = user
	return nil
}

// FindOneByID find a user by ID
func (u *UserRepository) FindOneByID(id uint32) (*entity.User, error) {
	if _, exist := u.ids[id] ; !exist {
		result := u.DB.First(&entity.User{}, id)
		if result.Error != nil {
			return nil, result.Error
		}
		var user entity.User
		result.Scan(&user)
		u.ids[id] = &user
		u.names[user.Name] = &user
		u.emails[user.Email] = &user
	}
	return u.ids[id], nil
}

// FindOneByName find a user by name
func (u *UserRepository) FindOneByName(name string) (*entity.User, error) {
	if _, exist := u.names[name] ; !exist {
		var user entity.User
		result := u.DB.Where("name = ?", name).First(&user)
		if result.Error != nil {
			return nil, result.Error
		}
		result.Scan(&user)
		u.ids[user.ID] = &user
		u.names[user.Name] = &user
		u.emails[user.Email] = &user
	}
	return u.names[name], nil
}

// FindOneByEmail find a user by email
func (u *UserRepository) FindOneByEmail(email string) (*entity.User, error) {
	if _, exist := u.emails[email] ; !exist {
		var user entity.User
		result := u.DB.Where("email = ?", email).First(&user)
		if result.Error != nil {
			return nil, result.Error
		}
		result.Scan(&user)
		u.ids[user.ID] = &user
		u.names[user.Name] = &user
		u.emails[user.Email] = &user
	}
	return u.emails[email], nil
}

// SaveOne save a user
func (u *UserRepository) SaveOne(id uint32) error {
	if u.ids[id] == nil {
		return errors.New("record not found")
	}
	result := u.DB.Save(u.ids[id])
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DeleteOne delete a user
func (u *UserRepository) DeleteOne(id uint32) error {
	user, err := u.FindOneByID(id)
	if err != nil {
		return err
	}
	result := u.DB.Select("Profile").Delete(user)
	if result.Error != nil {
		return result.Error
	}
	if u.ids[id] != nil {
		delete(u.ids, id)
		delete(u.names, user.Name)
		delete(u.emails, user.Email)
	}
	return nil
}