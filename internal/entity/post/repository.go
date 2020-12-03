package post

import "daanretard/internal/infra/object"

// IRepository repository interface of post, implemented at
// infra/persistence & infra/mock_persistence, services depend
// on this interface instead of the implementation
type IRepository interface {
	InsertOne(post *Post) error
	FindOneByID(id uint32) (Post, error)
	FindMany(query object.PostQuery) ([]Post, error)
	UpdateOne(post *Post) error
	DeleteOne(post *Post) error
}
