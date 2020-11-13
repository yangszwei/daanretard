package mock_persistence

import (
	"daanretard/internal/infrastructure/utilities/administrator"
	"errors"
)

// NewAdministratorRepository create an AdministratorRepository
func NewAdministratorRepository() *AdministratorRepository {
	a := new(AdministratorRepository)
	a.list = make(map[uint32]*administrator.Administrator)
	return a
}

// AdministratorRepository implement administrator.IRepository
type AdministratorRepository struct {
	list map[uint32]*administrator.Administrator
}

// InsertOne add an administrator
func (a *AdministratorRepository) InsertOne(userID uint32) error {
	admin := administrator.New()
	admin.UserID = userID
	a.list[userID] = admin
	return nil
}

// FindOneByID find an administrator
func (a *AdministratorRepository) FindOneByID(id uint32) (*administrator.Administrator, error) {
	if admin, exist := a.list[id]; exist {
		return admin, nil
	}
	return nil, errors.New("record not found")
}

// UpdateOne update an administrator
func (a *AdministratorRepository) UpdateOne(admin *administrator.Administrator) error {
	a.list[admin.UserID] = admin
	return nil
}

// DeleteOne delete an administrator
func (a *AdministratorRepository) DeleteOne(admin *administrator.Administrator) error {
	delete(a.list, admin.UserID)
	return nil
}
