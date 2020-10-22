package user

// IRepository interface
type IRepository interface {
	InsertOne(user *User) error
	FindOne(query Query) (*User, error)
	FindOneBySessionID(id uint32) (*User, error)
	FindAll(query Query) ([]*User, error)
	SaveOne(user *User) error
	DeleteOne(user *User) error
}