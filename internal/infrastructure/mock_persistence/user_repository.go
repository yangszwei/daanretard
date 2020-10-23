package mock_persistence

import (
	entity "daanretard/internal/domain/user"
	"errors"
)

// NewUserRepository create a UserRepository
func NewUserRepository() *UserRepository {
	u := new(UserRepository)
	u.count = 0
	u.ids = make(map[uint32]*entity.User)
	return u
}

// UserRepository implement user.IRepository
type UserRepository struct {
	count uint32
	ids map[uint32]*entity.User
}

// InsertOne insert a user
func (u *UserRepository) InsertOne(user *entity.User) error {
	if _, exist := u.ids[user.ID] ; exist {
		return errors.New("duplicate entry")
	}
	u.count += 1
	user.ID = u.count
	u.ids[user.ID] = user
	return nil
}

// FindOne find a user with user.Query
func (u *UserRepository) FindOne(query entity.Query) (*entity.User, error) {
	for _, user := range u.ids {
		if query.ID != 0 && query.ID == user.ID {
			return user, nil
		}
		if query.Name == "" || query.Name != user.Name {
			continue
		}
		if query.Email == "" || query.Email != user.Email {
			continue
		}
		return user, nil
	}
	return nil, errors.New("record not found")
}

// FindAll find a user with user.Query
func (u *UserRepository) FindAll(query entity.Query) ([]*entity.User, error) {
	var users []*entity.User
	for _, user := range u.ids {
		if query.ID != 0 && query.ID == user.ID {
			return []*entity.User{ user }, nil
		}
		if query.Name != "" && query.Name != user.Name {
			continue
		}
		if query.Email != "" && query.Email != user.Email {
			continue
		}
		if query.Name == "" && query.Email == "" {
			continue
		}
		users = append(users, user)
	}
	if len(users) == 0 {
		return nil, errors.New("record not found")
	}
	return users, nil
}

// SaveOne save a user
func (u *UserRepository) SaveOne(user *entity.User) error {
	if _, exist := u.ids[user.ID] ; !exist {
		return u.InsertOne(user)
	}
	u.ids[user.ID] = user
	return nil
}

// DeleteOne delete a user
func (u *UserRepository) DeleteOne(user *entity.User) error {
	if _, exist := u.ids[user.ID] ; !exist {
		return errors.New("record not found")
	}
	delete(u.ids, user.ID)
	return nil
}