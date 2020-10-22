package mock_persistence

import (
	entity "daanretard/internal/domain/post"
	"errors"
	"time"
)

// NewPostRepository create a PostRepository
func NewPostRepository() *PostRepository {
	p := new(PostRepository)
	p.count = 0
	p.ids = make(map[uint32]*entity.Post)
	return p
}

// PostRepository implement post.IRepository
type PostRepository struct {
	count uint32
	ids map[uint32]*entity.Post
}

// InsertOne insert a post
func (p *PostRepository) InsertOne(post *entity.Post) error {
	if _, exist := p.ids[post.ID] ; exist {
		return errors.New("duplicate entry")
	}
	p.count += 1
	post.ID = p.count
	p.ids[post.ID] = post
	return nil
}

// FindOne find a post
func (p *PostRepository) FindOne(query entity.Query) (*entity.Post, error) {
	for _, post := range p.ids {
		if query.ID != 0 && query.ID == post.ID {
			return post, nil
		}
		if query.Status == 0 || query.Status != post.Status {
			continue
		}
		if query.SubmitterID == 0 || query.SubmitterID != post.Submission.SubmitterID {
			continue
		}
		if query.ReviewerID == 0 || query.ReviewerID != post.Review.ReviewerID {
			continue
		}
		return post, nil
	}
	return nil, errors.New("record not found")
}

// FindAll find posts
func (p *PostRepository) FindAll(query entity.Query) ([]*entity.Post, error) {
	var posts []*entity.Post
	for _, post := range p.ids {
		if query.ID != 0 && query.ID == post.ID {
			return []*entity.Post{ post }, nil
		}
		if query.Status != 0 && query.Status != post.Status {
			continue
		}
		if query.SubmitterID != 0 && query.SubmitterID != post.Submission.SubmitterID {
			continue
		}
		if query.ReviewerID != 0 && query.ReviewerID != post.Review.ReviewerID {
			continue
		}
		if query.Status == 0 && query.SubmitterID == 0 && query.ReviewerID == 0 {
			continue
		}
		posts = append(posts, post)
	}
	if len(posts) == 0 {
		return nil, errors.New("record not found")
	}
	return posts, nil
}

// SaveOne save a post
func (p *PostRepository) SaveOne(post *entity.Post) error {
	if _, exist := p.ids[post.ID] ; !exist {
		return p.InsertOne(post)
	}
	p.ids[post.ID] = post
	return nil
}

// DeleteOne delete a post
func (p *PostRepository) DeleteOne(post *entity.Post) error {
	if _, exist := p.ids[post.ID] ; !exist {
		return errors.New("record not found")
	}
	post.DeletedAt = time.Now()
	return p.SaveOne(post)
}
