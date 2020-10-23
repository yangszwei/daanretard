package session

import "daanretard/internal/object"

// IUsecase interface
type IUsecase interface {
	// Open accept 0 ~ 1 argument
	Open(id ...uint32) (object.SessionProps, error)
	Extend(id uint64) error
	Close(id uint64) error
}