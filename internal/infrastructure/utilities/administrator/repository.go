package administrator

// IRepository interface
type IRepository interface {
	InsertOne(userID uint32) error
	FindOneByID(id uint32) (*Administrator, error)
	UpdateOne(admin *Administrator) error
	DeleteOne(admin *Administrator) error
}
