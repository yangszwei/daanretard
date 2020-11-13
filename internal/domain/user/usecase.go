package user

import "daanretard/internal/object"

// IUsecase interface
type IUsecase interface {
	Register(props object.UserProps) (uint32, error)
	GetProps(id uint32) (object.UserProps, error)
	Authenticate(email, password string) (uint32, error)
	AuthenticateWithID(id uint32, password string) error
	UpdateEmail(id uint32, email string) error
	UpdatePassword(id uint32, password string) error
	MarkAsVerified(id uint32) error
	UpdateProfile(id uint32, profile object.UserProfileProps) error
	Delete(id uint32) error
}
