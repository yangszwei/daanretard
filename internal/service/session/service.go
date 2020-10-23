package session

import (
	"daanretard/internal/object"
	"time"
)

// NewService create a service
func NewService(r IRepository) *Service {
	s := new(Service)
	s.r = r
	return s
}

// Service implement IUsecase
type Service struct {
	r IRepository
}

// Open open a session
func (s *Service) Open(id ...uint32) (object.SessionProps, error) {
	session := New()
	if len(id) >= 1 {
		session.UserID = id[0]
	}
	session.ExpiresAt = time.Now().Add(14 * 24 * time.Hour) // expires after 2 weeks
	err := s.r.InsertOne(session)
	if err != nil {
		return object.SessionProps{}, err
	}
	return object.SessionProps{
		ID: session.ID,
		UserID: session.UserID,
		CreatedAt: session.CreatedAt,
		ExpiresAt: session.ExpiresAt,
	}, err
}

// Extend set session expire time to 2 weeks later
func (s *Service) Extend(id uint64) error {
	session, err := s.r.FindOneByID(id)
	if err != nil {
		return err
	}
	session.ExpiresAt = time.Now().Add(14 * 24 * time.Hour) // expires after 2 weeks
	return s.r.SaveOne(session)
}

// Close close a session
func (s *Service) Close(id uint64) error {
	session, err := s.r.FindOneByID(id)
	if err != nil {
		return err
	}
	return s.r.DeleteOne(session)
}