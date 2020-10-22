package user

import (
	"daanretard/internal/infrastructure/security"
	"daanretard/internal/infrastructure/validator"
	"daanretard/internal/object"
	"errors"
)

// NewService create a UserService
func NewService(r IRepository) *Service {
	u := new(Service)
	u.r = r
	return u
}

// Service implement IUsecase
type Service struct {
	r IRepository
}

// Register register user
func (s *Service) Register(props object.UserProps, profile object.UserProfileProps) (uint32, error) {
	user := New()
	if err := validator.User(props) ; err != nil {
		return 0, err
	}
	if err := validator.UserProfile(profile) ; err != nil {
		return 0, err
	}
	password, _ := security.GenerateFromPassword(props.Password)
	user.Name = props.Name
	user.Email = props.Email
	user.Password = password
	user.Profile.FirstName = profile.FirstName
	user.Profile.LastName = profile.LastName
	err := s.r.InsertOne(user)
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}

// Authenticate authenticate user
func (s *Service) Authenticate(props object.UserProps) error {
	user, err := s.r.FindOne(Query{
		Name: props.Name,
		Email: props.Email,
	})
	if err != nil {
		return err
	}
	if security.CompareHashAndPassword(user.Password, props.Password) != nil {
		return errors.New("wrong password")
	}
	return nil
}

// Delete delete user
func (s *Service) Delete(id uint32) error {
	user, err := s.r.FindOne(Query{ ID: id })
	if err != nil {
		return err
	}
	return s.r.DeleteOne(user)
}