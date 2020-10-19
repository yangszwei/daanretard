package mock_persistence

import (
	entity "daanretard/internal/domain/user"
	"errors"
)

// NewUserRepository create a UserRepository
func NewUserRepository() *UserRepository {
	u := new(UserRepository)
	u.ids = make(map[uint32]*entity.User)
	u.names = make(map[string]*entity.User)
	u.emails = make(map[string]*entity.User)
	return u
}

// UserRepository implement user.IRepository
type UserRepository struct {
	counter uint32
	ids    map[uint32]*entity.User
	names  map[string]*entity.User
	emails map[string]*entity.User
}

// InsertOne insert a user
func (u *UserRepository) InsertOne(user *entity.User) error {
	if _, exist := u.names[user.Name] ; exist {
		return errors.New("name already exist")
	}
	if _, exist := u.emails[user.Email] ; exist {
		return errors.New("email already exist")
	}
	u.counter += 1
	user.ID = u.counter
	u.ids[user.ID] = user
	u.names[user.Name] = user
	u.emails[user.Email] = user
	return nil
}

// FindOneByID find a user by ID
func (u *UserRepository) FindOneByID(id uint32) (*entity.User, error) {
	if _, exist := u.ids[id] ; !exist {
		return nil, errors.New("record not found")
	}
	return u.ids[id], nil
}

// FindOneByName find a user by name
func (u *UserRepository) FindOneByName(name string) (*entity.User, error) {
	if _, exist := u.names[name] ; !exist {
		return nil, errors.New("record not found")
	}
	return u.names[name], nil
}

// FindOneByEmail find a user by email
func (u *UserRepository) FindOneByEmail(email string) (*entity.User, error) {
	if _, exist := u.emails[email] ; !exist {
		return nil, errors.New("record not found")
	}
	return u.emails[email], nil
}

// SaveOne save a user
func (u *UserRepository) SaveOne(id uint32) error {
	return nil
}

// DeleteOne delete a user
func (u *UserRepository) DeleteOne(id uint32) error {
	user, err := u.FindOneByID(id)
	if err != nil {
		return err
	}
	if u.ids[id] != nil {
		delete(u.ids, id)
		delete(u.names, user.Name)
		delete(u.emails, user.Email)
	}
	return nil
}