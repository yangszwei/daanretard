package persistence

import "daanretard/internal/entity/media"

// NewMediaRepo create a MediaRepo, db should be connected to
// database
func NewMediaRepo(db *DB) *MediaRepo {
	r := new(MediaRepo)
	r.db = db
	return r
}

// MediaRepo implement media.IRepository, it provide methods to
// persist media records
type MediaRepo struct {
	db *DB
}

// AutoMigrate setup "media" table on database, this should be run
// if the schema in database is not up to date
func (m *MediaRepo) AutoMigrate() error {
	return m.db.conn.AutoMigrate(&media.Media{})
}

// InsertOne insert a media record to database
func (m *MediaRepo) InsertOne(media *media.Media) error {
	return parseError(m.db.conn.Create(media).Error)
}

// FindOneByID find a media record by ID
func (m *MediaRepo) FindOneByID(id uint32) (result media.Media, err error) {
	err = parseError(m.db.conn.Where("id = ?", id).First(&result).Error)
	return
}

// UpdateOne save a media record to database
func (m *MediaRepo) UpdateOne(media *media.Media) error {
	return parseError(m.db.conn.Save(media).Error)
}

// DeleteOne delete a media record from database
func (m *MediaRepo) DeleteOne(media *media.Media) error {
	return parseError(m.db.conn.Delete(media).Error)
}
