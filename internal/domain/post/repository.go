package post

// IRepository interface
type IRepository interface {
	InsertOne(post *Post) error
	FindOne(query Query) (*Post, error)
	FindAll(query Query) ([]*Post, error)
	SaveOne(post *Post) error
	DeleteOne(post *Post) error
}
