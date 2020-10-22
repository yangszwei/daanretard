package persistence

import (
	entity "daanretard/internal/domain/post"
	"gorm.io/gorm"
	"time"
)

// NewPostRepository create a PostRepository
func NewPostRepository(db *gorm.DB) *PostRepository {
	p := new(PostRepository)
	p.DB = db
	return p
}

// PostRepository implement post.IRepository
type PostRepository struct {
	DB *gorm.DB
}

// AutoMigrate setup table
func (p *PostRepository) AutoMigrate() error {
	err := p.DB.AutoMigrate(
		&entity.Post{},
		&entity.Submission{},
		&entity.Review{},
	)
	return err
}

// InsertOne insert a post
func (p *PostRepository) InsertOne(post *entity.Post) error {
	result := p.DB.Create(post)
	return result.Error
}

// FindOne find a post
func (p *PostRepository) FindOne(query entity.Query) (*entity.Post, error) {
	var post entity.Post
	p.DB.Where(&entity.Post{
		ID: query.ID,
		Status: query.Status,
		Submission: entity.Submission{
			SubmitterID: query.SubmitterID,
		},
		Review: entity.Review{
			ReviewerID: query.ReviewerID,
		},
	}).Preload("Submission").Preload("Review").First(&post)
	return &post, nil
}

// FindAll find posts
func (p *PostRepository) FindAll(query entity.Query) ([]*entity.Post, error) {
	var posts []*entity.Post
	p.DB.Debug().Where(&entity.Post{
		ID: query.ID,
		Status: query.Status,
		Submission: entity.Submission{
			SubmitterID: query.SubmitterID,
		},
		Review: entity.Review{
			ReviewerID: query.ReviewerID,
		},
	}).Preload("Submission").Preload("Review").Find(&posts)
	return posts, nil
}

// SaveOne save a post
func (p *PostRepository) SaveOne(post *entity.Post) error {
	result := p.DB.Save(post)
	return result.Error
}

// DeleteOne delete a post
func (p *PostRepository) DeleteOne(post *entity.Post) error {
	post.DeletedAt = time.Now()
	result := p.DB.Save(post)
	return result.Error
}