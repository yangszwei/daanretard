package user

// IRepository repository interface of user, implemented at
// infra/persistence & infra/mock_persistence, services depend
// on this interface instead of the implementation
type IRepository interface {
	InsertOne(user *User) error
	FindOneByID(id string) (User, error)
	FindMany(limit, offset int) ([]User, error)
	UpdateOne(user *User) error
	DeleteOne(user *User) error
}
