package mock_persistence

import (
	"daanretard/internal/domain/post"
	"daanretard/internal/object"
	"errors"
)

// NewPostRepository create a PostRepository
func NewPostRepository() *PostRepository {
	p := new(PostRepository)
	p.count = 0
	p.ids = make(map[uint32]*post.Post)
	return p
}

// PostRepository implement post.IRepository
type PostRepository struct {
	count uint32
	ids   map[uint32]*post.Post
}

// InsertOne insert a post
func (p *PostRepository) InsertOne(post *post.Post) error {
	if _, exist := p.ids[post.ID]; exist {
		return errors.New("error 1062")
	}
	if post.ID == 0 {
		p.count++
		post.ID = p.count
	}
	p.ids[post.ID] = post
	return nil
}

// FindOneByID find a pos by ID
func (p *PostRepository) FindOneByID(id uint32) (*post.Post, error) {
	if _, exist := p.ids[id]; !exist {
		return nil, errors.New("record not found")
	}
	return p.ids[id], nil
}

// FindMany find posts by query
func (p *PostRepository) FindMany(query object.PostQuery) ([]*post.Post, error) {
	var (
		records     []*post.Post
		offsetCount = 0
	)
	for _, record := range p.ids {
		if query.Status != 0 && query.Status != record.Status ||
			query.UserID != 0 && query.UserID != record.UserID ||
			query.IPAddr != nil && string(query.IPAddr) != string(record.IPAddr) ||
			!query.CreatedAfter.IsZero() && query.CreatedAfter.After(record.CreatedAt) ||
			!query.CreatedBefore.IsZero() && query.CreatedBefore.Before(record.CreatedAt) ||
			query.ReviewerID != 0 && query.ReviewerID != record.Review.UserID ||
			query.ReviewResult != 0 && query.ReviewResult != record.Review.Result {
			continue
		}
		if query.Status == 0 && query.UserID == 0 && query.IPAddr == nil && query.CreatedAfter.IsZero() &&
			query.CreatedBefore.IsZero() && query.ReviewerID == 0 && query.ReviewResult == 0 {
			continue
		}
		if query.Offset > 0 && query.Offset > offsetCount {
			offsetCount++
			continue
		}
		records = append(records, record)
		if query.Limit > 0 && len(records) >= query.Limit {
			break
		}
	}
	if len(records) == 0 {
		return nil, errors.New("record not found")
	}
	return records, nil
}

// UpdateOne save a updated post
func (p *PostRepository) UpdateOne(post *post.Post) error {
	p.ids[post.ID] = post
	return nil
}

// DeleteOne delete a post
func (p *PostRepository) DeleteOne(post *post.Post) error {
	delete(p.ids, post.ID)
	return nil
}
