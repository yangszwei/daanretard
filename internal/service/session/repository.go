package session

// IRepository interface
type IRepository interface {
	InsertOne(session *Session) error
	FindOneByID(uint64) (*Session, error)
	FindAllByUserID(uint32) ([]*Session, error)
	SaveOne(session *Session) error
	DeleteOne(session *Session) error
}