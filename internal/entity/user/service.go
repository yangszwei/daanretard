package user

import (
	"daanretard/internal/infra/fbgraph"
	"daanretard/internal/infra/object"
)

// IUsecase usecase interface of User, serve as the interface for other
// packages to work with user entity, other packages should depend on this
// interface instead of Service
type IUsecase interface {
	ContinueWithFacebook(accessToken string) (string, error)
	GetProfile(id string) (object.User, error)
	Delete(id string) error
}

// NewService create a Service instance
func NewService(r IRepository, fb fbgraph.IUsecase) *Service {
	s := new(Service)
	s.r = r
	s.fb = fb
	return s
}

// Service implement IUsecase
type Service struct {
	r  IRepository
	fb fbgraph.IUsecase
}

// ContinueWithFacebook create a new user if a record with same user id (the
// user_id of the access token) does not already exist, then it return the id,
// the token is exchanged to a long-lived access token before saving to
// database
func (s *Service) ContinueWithFacebook(accessToken string) (string, error) {
	id, name, err := s.fb.GetUserProfile(accessToken)
	if err != nil {
		return "", err
	}
	token, exp, err := s.fb.ExchangeAccessToken(accessToken)
	if err != nil {
		return "", err
	}
	u, err := s.r.FindOneByID(id)
	if err != nil && err.Error() == "record not found" {
		u = User{ID: id}
	} else if err != nil {
		return "", err
	}
	u.AccessToken = token
	u.Name = name
	u.ExpiresAt = exp
	if err := s.r.UpdateOne(&u); err != nil {
		return "", err
	}
	return id, nil
}

// GetProfile return a User object that contains all user data
func (s *Service) GetProfile(id string) (object.User, error) {
	u, err := s.r.FindOneByID(id)
	if err != nil {
		return object.User{}, err
	}
	return object.User{
		ID:          u.ID,
		Name:        u.Name,
		AccessToken: u.AccessToken,
		ExpiresAt:   u.ExpiresAt,
		CreatedAt:   u.CreatedAt,
	}, nil
}

// Delete delete a user record from repository
func (s *Service) Delete(id string) error {
	u, err := s.r.FindOneByID(id)
	if err != nil {
		return err
	}
	return s.r.DeleteOne(&u)
}
