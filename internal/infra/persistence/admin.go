package persistence

import "daanretard/internal/entity/admin"

// NewAdminRepo create a AdminRepo, db should be connected to
// database
func NewAdminRepo(db *DB) *AdminRepo {
	r := new(AdminRepo)
	r.db = db
	return r
}

// AdminRepo implement admin.IRepository, it provide methods to
// persist admin records
type AdminRepo struct {
	db *DB
}

// AutoMigrate setup "admins" table on database, this should be run
// if the schema in database is not up to date
func (a *AdminRepo) AutoMigrate() error {
	return a.db.conn.AutoMigrate(&admin.Admin{})
}

// InsertOne insert an admin record to database
func (a *AdminRepo) InsertOne(admin *admin.Admin) error {
	return parseError(a.db.conn.Create(admin).Error)
}

// FindOneByUserID find an admin record by UserID
func (a *AdminRepo) FindOneByUserID(id string) (result admin.Admin, err error) {
	err = parseError(a.db.conn.Where("user_id = ?", id).First(&result).Error)
	return
}

// FindMany find a list of admin records
func (a *AdminRepo) FindMany(limit, offset int) (result []admin.Admin, err error) {
	err = parseError(a.db.conn.Limit(limit).Offset(offset).Find(&result).Error)
	return
}

// UpdateOne save an admin record to database
func (a *AdminRepo) UpdateOne(admin *admin.Admin) error {
	return parseError(a.db.conn.Save(admin).Error)
}

// DeleteOne delete an admin record from database
func (a *AdminRepo) DeleteOne(admin *admin.Admin) error {
	return parseError(a.db.conn.Delete(admin).Error)
}
