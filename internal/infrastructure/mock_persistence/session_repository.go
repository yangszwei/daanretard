package mock_persistence

import (
	entity "daanretard/internal/service/session"
	"errors"
)

// NewSessionRepository create a SessionRepository
func NewSessionRepository() *SessionRepository {
	s := new(SessionRepository)
	s.count = 0
	s.ids = make(map[uint64]*entity.Session)
	return s
}

// SessionRepository implement session.IRepository
type SessionRepository struct {
	count uint64
	ids map[uint64]*entity.Session
}

// InsertOne insert a session
func (s *SessionRepository) InsertOne(session *entity.Session) error {
	s.count++
	session.ID = s.count
	s.ids[session.ID] = session
	return nil
}

// FindOneByID find a session by ID
func (s *SessionRepository) FindOneByID(id uint64) (*entity.Session, error) {
	if _, exist := s.ids[id] ; !exist {
		return nil, errors.New("record not found")
	}
	return s.ids[id], nil
}

// FindAllByUserID find sessions by user ID
func (s *SessionRepository) FindAllByUserID(id uint32) ([]*entity.Session, error) {
	var sessions []*entity.Session
	for _, session := range s.ids {
		if session.UserID == id {
			sessions = append(sessions, session)
		}
	}
	return sessions, nil
}

// SaveOne save a session
func (s *SessionRepository) SaveOne(session *entity.Session) error {
	return nil
}

// DeleteOne delete a session
func (s *SessionRepository) DeleteOne(session *entity.Session) error {
	delete(s.ids, session.ID)
	return nil
}