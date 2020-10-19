package user

// IRepository interface
type IRepository interface {
	InsertOne(user *User) error
	FindOneByID(id uint32) (*User, error)
	FindOneByName(name string) (*User, error)
	FindOneByEmail(email string) (*User, error)
	SaveOne(id uint32) error
	DeleteOne(id uint32) error
}