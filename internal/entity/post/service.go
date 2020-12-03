package post

import (
	"daanretard/internal/infra/errors"
	"daanretard/internal/infra/fbgraph"
	"daanretard/internal/infra/object"
	"daanretard/internal/infra/validator"
	"strconv"
	"strings"
)

// IUsecase usecase interface of Post, serve as the interface for other
// packages to work with post entity, other packages should depend on
// this interface instead of Service
type IUsecase interface {
	Submit(post object.Post) (uint32, error)
	Search(query object.PostQuery) ([]object.Post, error)
	Review(id uint32, review object.PostReview) error
	Publish(id uint32, attachedMedia []string, accessToken string) (string, error)
	Delete(id uint32, accessToken string) error
}

// toObject convert Post to object.Post
func toObject(p Post) object.Post {
	var attachments []uint32
	for _, a := range strings.Split(p.Attachments, ",") {
		mediaID, _ := strconv.ParseUint(a, 10, 32)
		attachments = append(attachments, uint32(mediaID))
	}
	return object.Post{
		ID:          p.ID,
		Status:      p.Status,
		UserID:      p.UserID,
		IPAddr:      p.IPAddr,
		UserAgent:   p.UserAgent,
		Message:     p.Message,
		Attachments: attachments,
		Review: object.PostReview{
			PostID:    p.Review.PostID,
			UserID:    p.Review.UserID,
			Result:    p.Review.Result,
			Message:   p.Review.Message,
			CreatedAt: p.Review.CreatedAt,
		},
		FacebookID: p.FacebookID,
		CreatedAt:  p.CreatedAt,
	}
}

// toPost convert object.Post to Post
func toPost(p object.Post) Post {
	var attachments []string
	for _, mediaID := range p.Attachments {
		attachments = append(attachments, strconv.Itoa(int(mediaID)))
	}
	return Post{
		ID:          p.ID,
		Status:      p.Status,
		UserID:      p.UserID,
		IPAddr:      p.IPAddr,
		UserAgent:   p.UserAgent,
		Message:     p.Message,
		Attachments: strings.Join(attachments, ","),
		Review: Review{
			PostID:    p.Review.PostID,
			UserID:    p.Review.UserID,
			Result:    p.Review.Result,
			Message:   p.Review.Message,
			CreatedAt: p.Review.CreatedAt,
		},
		FacebookID: p.FacebookID,
		CreatedAt:  p.CreatedAt,
	}
}

// toReview convert object.PostReview to Review
func toReview(r object.PostReview) Review {
	return Review{
		PostID:    r.PostID,
		UserID:    r.UserID,
		Result:    r.Result,
		Message:   r.Message,
		CreatedAt: r.CreatedAt,
	}
}

// NewService create a Service instance
func NewService(v *validator.Validator, r IRepository, fb fbgraph.IUsecase) *Service {
	s := new(Service)
	s.v = v
	s.r = r
	s.fb = fb
	return s
}

// Service implement IUsecase
type Service struct {
	v  *validator.Validator
	r  IRepository
	fb fbgraph.IUsecase
}

// Submit validate and insert a post record to repository
func (s *Service) Submit(post object.Post) (uint32, error) {
	if err := s.v.Validate(post); err != nil {
		return 0, err
	}
	if post.Message == "" || post.IPAddr == nil || post.UserAgent == "" {
		var errs errors.Errors
		if post.Message == "" {
			errs = append(errs, errors.New("required", "message"))
		}
		if post.IPAddr == nil {
			errs = append(errs, errors.New("required", "ip_addr"))
		}
		if post.UserAgent == "" {
			errs = append(errs, errors.New("required", "user_agent"))
		}
		return 0, errs
	}
	p := toPost(post)
	if err := s.r.InsertOne(&p); err != nil {
		return 0, err
	}
	return p.ID, nil
}

// Search return a list of posts filtered by object.PostQuery
func (s *Service) Search(query object.PostQuery) ([]object.Post, error) {
	records, err := s.r.FindMany(query)
	if err != nil {
		return nil, err
	}
	var result []object.Post
	for _, record := range records {
		result = append(result, toObject(record))
	}
	return result, nil
}

// Review set review field of a post record
func (s *Service) Review(id uint32, review object.PostReview) error {
	if err := s.v.Validate(review); err != nil {
		return err
	}
	p, err := s.r.FindOneByID(id)
	if err != nil {
		return err
	}
	p.Review = toReview(review)
	return s.r.UpdateOne(&p)
}

// Publish publish a post record to facebook, the record's status must be
// reviewed and approved before publishing
func (s *Service) Publish(id uint32, attachedMedia []string, accessToken string) (string, error) {
	p, err := s.r.FindOneByID(id)
	if err != nil {
		return "", err
	}
	fbID, err := s.fb.PublishPost(p.Message, attachedMedia, accessToken)
	if err != nil {
		return "", err
	}
	p.Status = StatusPublished
	p.FacebookID = fbID
	if err := s.r.UpdateOne(&p); err != nil {
		_ = s.fb.Delete(fbID, accessToken)
		return "", err
	}
	return fbID, nil
}

// Delete delete a post record from facebook and repository
func (s *Service) Delete(id uint32, accessToken string) error {
	p, err := s.r.FindOneByID(id)
	if err != nil {
		return err
	}
	if p.Status == StatusPublished {
		if err := s.fb.Delete(p.FacebookID, accessToken); err != nil {
			return err
		}
	}
	return s.r.DeleteOne(&p)
}
