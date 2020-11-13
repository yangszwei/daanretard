package post

import "daanretard/internal/object"

// IRepository interface
type IRepository interface {
	InsertOne(post *Post) error
	FindOneByID(id uint32) (*Post, error)
	FindMany(query object.PostQuery) ([]*Post, error)
	UpdateOne(post *Post) error
	DeleteOne(post *Post) error
}
