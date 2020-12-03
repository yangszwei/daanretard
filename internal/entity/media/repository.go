package media

// IRepository repository interface of media, implemented at
// infra/persistence & infra/mock_persistence, services depend
// on this interface instead of the implementation
type IRepository interface {
	InsertOne(media *Media) error
	FindOneByID(id uint32) (Media, error)
	UpdateOne(media *Media) error
	DeleteOne(media *Media) error
}
