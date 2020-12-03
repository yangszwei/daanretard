package admin

import (
	"daanretard/internal/infra/errors"
	"daanretard/internal/infra/fbgraph"
	"daanretard/internal/infra/validator"
	"log"
	"time"
)

type IUsecase interface {
	Authenticate(accessToken string) error
	IsAdmin(id string) error
}

// NewService create a Service
func NewService(r IRepository, fb fbgraph.IUsecase, v *validator.Validator) *Service {
	s := new(Service)
	s.r = r
	s.fb = fb
	s.v = v
	return s
}

// Service implement IUsecase
type Service struct {
	r  IRepository
	fb fbgraph.IUsecase
	v  *validator.Validator
}

// Authenticate verify accessToken and add an record to repository or refresh if record already exist
func (s *Service) Authenticate(accessToken string) error {
	info, err := s.fb.DebugAccessToken(accessToken)
	if err != nil {
		return err
	}
	log.Println(info)
	if !info.IsValid || info.Type != fbgraph.Page || info.ExpiresAt.Before(time.Now()) {
		return errors.New("invalid access token")
	}
	token, exp, err := s.fb.ExchangeAccessToken(accessToken)
	if err != nil {
		return err
	}
	return s.r.InsertOne(&Admin{
		UserID:      info.UserID,
		AccessToken: token,
		ExpiresAt:   exp,
	})
}

// IsAdmin find and verify an admin record from repository
func (s *Service) IsAdmin(id string) error {
	a, err := s.r.FindOneByUserID(id)
	if err != nil {
		return err
	}
	if a.AccessToken == "" || a.ExpiresAt.Before(time.Now()) {
		return errors.New("invalid access token")
	}
	return nil
}
