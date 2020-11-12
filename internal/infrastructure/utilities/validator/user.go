package validator

import (
	"daanretard/internal/object"
	"errors"
	"github.com/go-playground/validator/v10"
)

var (
	userValidator = validator.New()
	userProfileValidator = validator.New()
)

// User validate user.Props
func User(props object.UserProps) error {
	errs := userValidator.Struct(props)
	if errs != nil {
		return errors.New("invalid credentials")
	}
	return nil
}

// User validate user.Props
func UserProfile(props object.UserProfileProps) error {
	errs := userProfileValidator.Struct(props)
	if errs != nil {
		return errors.New("invalid credentials")
	}
	return nil
}