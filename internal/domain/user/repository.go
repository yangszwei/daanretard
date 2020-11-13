package user

// IRepository interface
type IRepository interface {
	InsertOne(user *User) error
	FindOneByID(id uint32) (*User, error)
	FindOneByEmail(email string) (*User, error)
	UpdateOne(user *User) error
	DeleteOne(user *User) error
}
