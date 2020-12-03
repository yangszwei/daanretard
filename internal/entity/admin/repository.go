package admin

// IRepository interface
type IRepository interface {
	InsertOne(admin *Admin) error
	FindOneByUserID(id string) (Admin, error)
	FindMany(limit, offset int) ([]Admin, error)
	UpdateOne(admin *Admin) error
	DeleteOne(admin *Admin) error
}
