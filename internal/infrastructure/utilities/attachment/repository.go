package attachment

// IRepository interface
type IRepository interface {
	InsertOne(name string) (uint32, error)
	FindOneByID(id uint32) (string, error)
	DeleteOne(id uint32) error
}
