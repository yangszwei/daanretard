package mock_persistence

import (
	"daanretard/internal/domain/user"
	"errors"
)

// NewUserRepository create a UserRepository
func NewUserRepository() *UserRepository {
	u := new(UserRepository)
	u.count = 0
	u.ids = make(map[uint32]*user.User)
	u.emails = make(map[string]*user.User)
	return u
}

// UserRepository implement user.IRepository
type UserRepository struct {
	count uint32
	ids map[uint32]*user.User
	emails map[string]*user.User
}

// InsertOne insert a user
func (u *UserRepository) InsertOne(user *user.User) error {
	_, idUsed := u.ids[user.ID]
	_, emailUsed := u.emails[user.Email]
	if idUsed || emailUsed {
		return errors.New("error 1062")
	}
	if user.ID == 0 {
		u.count++
		user.ID = u.count
	}
	u.ids[user.ID] = user
	u.emails[user.Email] = user
	return nil
}

// FindOneByID find a user by its ID
func (u *UserRepository) FindOneByID(id uint32) (*user.User, error) {
	if _, exist := u.ids[id] ; !exist {
		return nil, errors.New("record not found")
	}
	return u.ids[id], nil
}

// FindOneByEmail find a user by its email
func (u *UserRepository) FindOneByEmail(email string) (*user.User, error) {
	if _, exist := u.emails[email] ; !exist {
		return nil, errors.New("record not found")
	}
	return u.emails[email], nil
}

// UpdateOne save a user
func (u *UserRepository) UpdateOne(user *user.User) error {
	u.ids[user.ID] = user
	u.emails[user.Email] = user
	return nil
}

// DeleteOne delete a user
func (u *UserRepository) DeleteOne(user *user.User) error {
	delete(u.ids, user.ID)
	delete(u.emails, user.Email)
	return nil
}