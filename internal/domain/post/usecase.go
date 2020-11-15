package post

import "daanretard/internal/object"

// IUsecase interface
type IUsecase interface {
	Submit(post object.Post) (uint32, error)
	GetOne(id uint32) (object.Post, error)
	GetManyByUserID(id uint32, offset, limit int) ([]object.Post, error)
	GetManyNotReviewed(offset, limit int) ([]object.Post, error)
	MarkAsPublished(id uint32, fbPostID string) error
	GetManyPublished(offset, limit int) ([]object.Post, error)
	Review(id uint32, review object.PostReview) error
	Delete(id uint32) error
}
