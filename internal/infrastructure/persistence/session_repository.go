package persistence

import (
	entity "daanretard/internal/service/session"
	"gorm.io/gorm"
)

// NewSessionRepository create a SessionRepository
func NewSessionRepository(db *DB) *SessionRepository {
	s := new(SessionRepository)
	s.db = db.Conn
	s.ids = make(map[uint64]*entity.Session)
	return s
}

// SessionRepository implement session.IRepository
type SessionRepository struct {
	db *gorm.DB
	ids map[uint64]*entity.Session
}

// AutoMigrate setup table schema
func (s *SessionRepository) AutoMigrate() error {
	return s.db.AutoMigrate(entity.Session{})
}

// InsertOne insert a session
func (s *SessionRepository) InsertOne(session *entity.Session) error {
	result := s.db.Create(session)
	s.ids[session.ID] = session
	return result.Error
}

// FindOneByID find a session by ID
func (s *SessionRepository) FindOneByID(id uint64) (*entity.Session, error) {
	if _, exist := s.ids[id] ; !exist {
		var session entity.Session
		result := s.db.First(&session, id)
		if result.Error != nil {
			return nil, result.Error
		}
		result.Scan(&session)
		s.ids[id] = &session
	}
	return s.ids[id], nil
}

// FindAllByUserID find sessions by user ID
func (s *SessionRepository) FindAllByUserID(id uint32) ([]*entity.Session, error) {
	var sessions []*entity.Session
	result := s.db.Where("user_id = ?", id).Find(&entity.Session{})
	if result.Error != nil {
		return nil, result.Error
	}
	result.Scan(&sessions)
	for _, session := range sessions {
		if _, exist := s.ids[session.ID] ; !exist {
			s.ids[session.ID] = session
		}
	}
	return sessions, nil
}

// SaveOne save a session
func (s *SessionRepository) SaveOne(session *entity.Session) error {
	result := s.db.Save(session)
	return result.Error
}

// DeleteOne delete a session
func (s *SessionRepository) DeleteOne(session *entity.Session) error {
	result := s.db.Delete(session)
	delete(s.ids, session.ID)
	return result.Error
}