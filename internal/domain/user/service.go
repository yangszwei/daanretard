package user

import (
	"daanretard/internal/infrastructure/utilities/security"
	"daanretard/internal/infrastructure/utilities/validator"
	"daanretard/internal/object"
	"errors"
)

// NewService create a Service
func NewService(r IRepository) *Service {
	s := new(Service)
	s.r = r
	return s
}

// containsRequired checks whether all required fields are present
func containsRequired(props object.UserProps) bool {
	var invalid bool
	invalid = invalid || props.Email == ""
	invalid = invalid || props.Password == ""
	invalid = invalid || props.Profile.DisplayName == ""
	invalid = invalid || props.Profile.FirstName == ""
	invalid = invalid || props.Profile.LastName == ""
	return !invalid
}

// toUser convert object.UserProps to User
func toUser(props object.UserProps) (*User, error) {
	u := New()
	u.ID = props.ID
	u.Email = props.Email
	password, err := security.GenerateFromPassword(props.Password)
	if err != nil {
		return nil, err
	}
	u.Password = password
	u.Profile.DisplayName = props.Profile.DisplayName
	u.Profile.FirstName = props.Profile.FirstName
	u.Profile.LastName = props.Profile.LastName
	u.IsVerified = props.IsVerified
	u.CreatedAt = props.CreatedAt
	return u, nil
}

// toObject convert User to object.UserProps
func toObject(user *User) object.UserProps {
	p := object.UserProps{}
	p.ID = user.ID
	p.Email = user.Email
	if len(user.Password) > 0 {
		p.Password = "hashed string"
	}
	p.Profile.DisplayName = user.Profile.DisplayName
	p.Profile.FirstName = user.Profile.FirstName
	p.Profile.LastName = user.Profile.LastName
	p.IsVerified = user.IsVerified
	p.CreatedAt = user.CreatedAt
	return p
}

// Service implement IUsecase
type Service struct {
	r IRepository
}

// Register validate & save a new user
func (s *Service) Register(props object.UserProps) (uint32, error) {
	if !containsRequired(props) {
		return 0, errors.New("invalid credentials")
	}
	if err := validator.User(props); err != nil {
		return 0, err
	}
	user, err := toUser(props)
	if err != nil {
		panic(err)
	}
	if _, err := s.r.FindOneByEmail(props.Email); err != nil {
		err = s.r.InsertOne(user)
		return user.ID, err
	}
	return 0, errors.New("email already taken")
}

// GetProps get user properties
func (s *Service) GetProps(id uint32) (object.UserProps, error) {
	u, err := s.r.FindOneByID(id)
	props := toObject(u)
	return props, err
}

// Authenticate authenticate user by email & password
func (s *Service) Authenticate(email, password string) (uint32, error) {
	u, err := s.r.FindOneByEmail(email)
	if err != nil {
		if err.Error() == "record not found" {
			return 0, errors.New("invalid credentials")
		}
		return 0, err
	}
	if security.CompareHashAndPassword(u.Password, password) != nil {
		return 0, errors.New("invalid credentials")
	}
	return u.ID, nil
}

// AuthenticateWithID authenticate user by email & password
func (s *Service) AuthenticateWithID(id uint32, password string) error {
	u, err := s.r.FindOneByID(id)
	if err != nil {
		if err.Error() == "record not found" {
			return errors.New("invalid credentials")
		}
		return err
	}
	if security.CompareHashAndPassword(u.Password, password) != nil {
		return errors.New("invalid credentials")
	}
	return nil
}

// UpdateEmail set user email and mark user as not verified
func (s *Service) UpdateEmail(id uint32, email string) error {
	u, err := s.r.FindOneByID(id)
	if err != nil {
		return err
	}
	changes := object.UserProps{Email: email}
	if err := validator.User(changes); err != nil {
		return err
	}
	u.Email = email
	u.IsVerified = false
	err = s.r.UpdateOne(u)
	return err
}

// Update password set user password
func (s *Service) UpdatePassword(id uint32, password string) error {
	u, err := s.r.FindOneByID(id)
	if err != nil {
		return err
	}
	changes := object.UserProps{Password: password}
	if err := validator.User(changes); err != nil {
		return err
	}
	u.Password, err = security.GenerateFromPassword(password)
	if err != nil {
		return err
	}
	err = s.r.UpdateOne(u)
	return err
}

// MarkAsVerified mark user as verified
func (s *Service) MarkAsVerified(id uint32) error {
	u, err := s.r.FindOneByID(id)
	if err != nil {
		return err
	}
	u.IsVerified = true
	err = s.r.UpdateOne(u)
	return err
}

// AddAdministrator mark user as an administrator
func (s *Service) AddAdministrator(id uint32) error {
	u, err := s.r.FindOneByID(id)
	if err != nil {
		return err
	}
	u.IsAdministrator = true
	err = s.r.UpdateOne(u)
	return err
}

// IsAdministrator return whether user is an administrator
func (s *Service) IsAdministrator(id uint32) (bool, error) {
	u, err := s.r.FindOneByID(id)
	if err != nil {
		return false, err
	}
	return u.IsAdministrator, nil
}

// RemoveAdministrator mark user as an administrator
func (s *Service) RemoveAdministrator(id uint32) error {
	u, err := s.r.FindOneByID(id)
	if err != nil {
		return err
	}
	u.IsAdministrator = false
	err = s.r.UpdateOne(u)
	return err
}

// Update profile validate and set user profile
func (s *Service) UpdateProfile(id uint32, profile object.UserProfileProps) error {
	u, err := s.r.FindOneByID(id)
	if err != nil {
		return err
	}
	if err := validator.UserProfile(profile); err != nil {
		return err
	}
	user, err := toUser(object.UserProps{Profile: profile})
	if err != nil {
		return err
	}
	if user.Profile.DisplayName != "" {
		u.Profile.DisplayName = user.Profile.DisplayName
	}
	if user.Profile.FirstName != "" {
		u.Profile.FirstName = user.Profile.FirstName
	}
	if user.Profile.LastName != "" {
		u.Profile.LastName = user.Profile.LastName
	}
	err = s.r.UpdateOne(u)
	return err
}

// Delete delete user
func (s *Service) Delete(id uint32) error {
	u, err := s.r.FindOneByID(id)
	if err != nil {
		return err
	}
	err = s.r.DeleteOne(u)
	return err
}
