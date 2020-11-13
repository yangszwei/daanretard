package post

import (
	"daanretard/internal/infrastructure/utilities/validator"
	"daanretard/internal/object"
	"errors"
	"strings"
)

// NewService create a Service
func NewService(r IRepository) *Service {
	s := new(Service)
	s.r = r
	return s
}

// containsRequired checks whether all required fields are present
func containsRequired(post object.Post) bool {
	var invalid bool
	invalid = invalid || post.UserID == 0
	invalid = invalid || len(post.IPAddr) == 0
	invalid = invalid || post.UserAgent == ""
	invalid = invalid || post.Message == ""
	return !invalid
}

// toPost convert object.Post to Post
func toPost(post object.Post) *Post {
	p := New()
	p.ID = post.ID
	p.Status = post.Status
	p.UserID = post.UserID
	p.IPAddr = post.IPAddr
	p.UserAgent = post.UserAgent
	p.Message = post.Message
	p.Attachments = strings.Join(post.Attachments, ",")
	p.Review.PostID = p.ID
	p.Review.UserID = post.Review.UserID
	p.Review.Result = post.Review.Result
	p.Review.Message = post.Review.Message
	p.Review.CreatedAt = post.Review.CreatedAt
	p.FacebookID = post.FacebookID
	p.CreatedAt = post.CreatedAt
	return p
}

// toObject convert Post to object.Post
func toObject(post *Post) object.Post {
	var p object.Post
	p.ID = post.ID
	p.Status = post.Status
	p.UserID = post.UserID
	p.IPAddr = post.IPAddr
	p.UserAgent = post.UserAgent
	p.Message = post.Message
	p.Attachments = strings.Split(post.Attachments, ",")
	p.Review.UserID = post.Review.UserID
	p.Review.Result = post.Review.Result
	p.Review.Message = post.Review.Message
	p.Review.CreatedAt = post.Review.CreatedAt
	p.FacebookID = post.FacebookID
	p.CreatedAt = post.CreatedAt
	return p
}

// Service implement IUsecase
type Service struct {
	r IRepository
}

// Submit submit a post
func (s *Service) Submit(post object.Post) (uint32, error) {
	if !containsRequired(post) {
		return 0, errors.New("invalid credentials")
	}
	if err := validator.Post(post); err != nil {
		return 0, err
	}
	post.Status = StatusSubmitted
	p := toPost(post)
	err := s.r.InsertOne(p)
	return p.ID, err
}

// GetOne get a post object by ID
func (s *Service) GetOne(id uint32) (object.Post, error) {
	post, err := s.r.FindOneByID(id)
	return toObject(post), err
}

// GetManyByUserID get many post objects by UserID
func (s *Service) GetManyByUserID(id uint32, offset, limit int) ([]object.Post, error) {
	records, err := s.r.FindMany(object.PostQuery{
		UserID: id,
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		return nil, err
	}
	var posts []object.Post
	for _, post := range records {
		posts = append(posts, toObject(post))
	}
	return posts, nil
}

// GetManyNotReviewed get many post objects that are not reviewed
func (s *Service) GetManyNotReviewed(offset, limit int) ([]object.Post, error) {
	records, err := s.r.FindMany(object.PostQuery{
		Status: StatusSubmitted,
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		return nil, err
	}
	var posts []object.Post
	for _, post := range records {
		posts = append(posts, toObject(post))
	}
	return posts, nil
}

// MarkAsPublished mark a post as published
func (s *Service) MarkAsPublished(id uint32) error {
	post, err := s.r.FindOneByID(id)
	if err != nil {
		return err
	}
	post.Status = StatusPublished
	return s.r.UpdateOne(post)
}

// GetManyPublished get many post objects that are published
func (s *Service) GetManyPublished(offset, limit int) ([]object.Post, error) {
	records, err := s.r.FindMany(object.PostQuery{
		Status: StatusPublished,
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		return nil, err
	}
	var posts []object.Post
	for _, post := range records {
		posts = append(posts, toObject(post))
	}
	return posts, nil
}

// Review review a post
func (s *Service) Review(id uint32, review object.PostReview) error {
	post, err := s.r.FindOneByID(id)
	if err != nil {
		return err
	}
	post.Review = toPost(object.Post{Review: review}).Review
	err = s.r.UpdateOne(post)
	return err
}

// Delete delete a post
func (s *Service) Delete(id uint32) error {
	post, err := s.r.FindOneByID(id)
	if err != nil {
		return err
	}
	return s.r.DeleteOne(post)
}
