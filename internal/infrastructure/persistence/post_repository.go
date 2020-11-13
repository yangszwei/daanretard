package persistence

import (
	"daanretard/internal/domain/post"
	"daanretard/internal/object"
	"gorm.io/gorm"
)

// NewPostRepository create a post repository
func NewPostRepository(db *DB) *PostRepository {
	p := new(PostRepository)
	p.db = db.Conn
	return p
}

// PostRepository implement post.IRepository
type PostRepository struct {
	db *gorm.DB
}

// AutoMigrate set table schema
func (p *PostRepository) AutoMigrate() error {
	return p.db.AutoMigrate(&post.Post{}, &post.Review{})
}

// InsertOne insert a post
func (p *PostRepository) InsertOne(post *post.Post) error {
	result := p.db.Create(post)
	return result.Error
}

// FindOneByID find a pos by ID
func (p *PostRepository) FindOneByID(id uint32) (*post.Post, error) {
	var record post.Post
	result := p.db.Preload("Review").First(&record, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &record, nil
}

// FindMany find posts by query
func (p *PostRepository) FindMany(query object.PostQuery) ([]*post.Post, error) {
	var records []*post.Post
	result := p.db.
		Where(&post.Post{
			Status: query.Status,
			UserID: query.UserID,
			IPAddr: query.IPAddr,
			Review: post.Review{
				UserID: query.ReviewerID,
				Result: query.ReviewResult,
			},
		}).
		Where("created_at >= ?", query.CreatedAfter).
		Where("created_at <= ?", query.CreatedBefore).
		Limit(query.Limit).
		Offset(query.Offset).
		Preload("Review").
		Find(&records)
	if result.Error != nil {
		return nil, result.Error
	}
	if len(records) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return records, nil
}

// UpdateOne save a updated post
func (p *PostRepository) UpdateOne(post *post.Post) error {
	result := p.db.Save(post)
	return result.Error
}

// DeleteOne delete a post
func (p *PostRepository) DeleteOne(post *post.Post) error {
	result := p.db.Select("Review").Delete(post)
	return result.Error
}
