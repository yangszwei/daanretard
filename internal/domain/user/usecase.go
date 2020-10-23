package user

import "daanretard/internal/object"

// IUsecase interface
type IUsecase interface {
	Register(props object.UserProps, profile object.UserProfileProps) (uint32, error)
	Authenticate(props object.UserProps) (uint32, error)
	Delete(id uint32) error
}