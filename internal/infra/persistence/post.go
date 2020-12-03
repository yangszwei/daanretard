package persistence

import (
	"daanretard/internal/entity/post"
	"daanretard/internal/infra/object"
)

// NewPostRepo create a PostRepo, db should be connected to
// database
func NewPostRepo(db *DB) *PostRepo {
	r := new(PostRepo)
	r.db = db
	return r
}

// PostRepo implement post.IRepository, it provide methods to
// persist post records
type PostRepo struct {
	db *DB
}

// AutoMigrate setup "posts" table on database, this should be run
// if the schema in database is not up to date
func (p *PostRepo) AutoMigrate() error {
	return p.db.conn.AutoMigrate(&post.Post{}, &post.Review{})
}

// InsertOne insert a post record to database
func (p *PostRepo) InsertOne(post *post.Post) error {
	return parseError(p.db.conn.Create(post).Error)
}

// FindOneByID find a post record by ID
func (p *PostRepo) FindOneByID(id uint32) (result post.Post, err error) {
	err = parseError(p.db.conn.Where("id = ?", id).First(&result).Error)
	return
}

// FindMany find a list of post records with object.PostQuery, only filters
// with non-zero values are applies
func (p *PostRepo) FindMany(query object.PostQuery) (result []post.Post, err error) {
	stmt := p.db.conn.Where(&post.Post{
		Status:     query.Status,
		UserID:     query.UserID,
		FacebookID: query.FacebookID,
	})
	if query.Message != "" {
		stmt = stmt.Where("posts.message LIKE ?", "%"+query.Message+"%")
	}
	if !query.CreatedAfter.IsZero() {
		stmt = stmt.Where("posts.created_at > ?", query.CreatedAfter)
	}
	if !query.CreatedBefore.IsZero() {
		stmt = stmt.Where("posts.created_at < ?", query.CreatedBefore)
	}
	if query.ReviewerID != "" || query.ReviewResult != 0 {
		stmt = stmt.Joins("Review")
	}
	if query.ReviewerID != "" {
		stmt = stmt.Where("Review.user_id = ?", query.ReviewerID)
	}
	if query.ReviewResult != 0 {
		stmt = stmt.Where("Review.result = ?", query.ReviewResult)
	}
	if query.Limit != 0 {
		stmt = stmt.Limit(query.Limit).Offset(query.Offset)
	}
	err = parseError(stmt.Find(&result).Error)
	return
}

// UpdateOne save a post record to database
func (p *PostRepo) UpdateOne(post *post.Post) error {
	return parseError(p.db.conn.Save(post).Error)
}

// DeleteOne delete a post record from database
func (p *PostRepo) DeleteOne(post *post.Post) error {
	return parseError(p.db.conn.Select("Review").Delete(post).Error)
}
