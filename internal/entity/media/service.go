package media

import (
	"daanretard/internal/infra/object"
	"daanretard/internal/infra/validator"
)

// IUsecase usecase interface of Media, serve as the interface for other
// packages to work with user entity, other packages should depend on this
// interface instead of Service
type IUsecase interface {
	Add(media object.Media) (uint32, error)
	GetOne(id uint32) (object.Media, error)
	Delete(id uint32) error
}

// NewService create a Service
func NewService(r IRepository, v *validator.Validator) *Service {
	s := new(Service)
	s.r = r
	s.v = v
	return s
}

// Service implement IUsecase
type Service struct {
	r IRepository
	v *validator.Validator
}

// Add validate and add a media record to repository
func (s *Service) Add(media object.Media) (uint32, error) {
	if err := s.v.Validate(media); err != nil {
		return 0, err
	}
	m := Media{
		ID:         media.ID,
		UserID:     media.UserID,
		Name:       media.Name,
		FacebookID: media.FacebookID,
		CreatedAt:  media.CreatedAt,
	}
	if err := s.r.InsertOne(&m); err != nil {
		return 0, err
	}
	return m.ID, nil
}

// GetOne find a media record by ID from repository
func (s *Service) GetOne(id uint32) (object.Media, error) {
	media, err := s.r.FindOneByID(id)
	if err != nil {
		return object.Media{}, err
	}
	return object.Media{
		ID:         media.ID,
		UserID:     media.UserID,
		Name:       media.Name,
		FacebookID: media.FacebookID,
		CreatedAt:  media.CreatedAt,
	}, nil
}

// Delete delete a media record by ID from repository
func (s *Service) Delete(id uint32) error {
	m, err := s.r.FindOneByID(id)
	if err != nil {
		return err
	}
	return s.r.DeleteOne(&m)
}
