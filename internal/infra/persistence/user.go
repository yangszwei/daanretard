package persistence

import "daanretard/internal/entity/user"

// NewUserRepo create a UserRepo, db should be connected to
// database
func NewUserRepo(db *DB) *UserRepo {
	r := new(UserRepo)
	r.db = db
	return r
}

// UserRepo implement user.IRepository, it provide methods to
// persist user records
type UserRepo struct {
	db *DB
}

// AutoMigrate setup "users" table on database, this should be run
// if the schema in database is not up to date
func (u *UserRepo) AutoMigrate() error {
	return u.db.conn.AutoMigrate(&user.User{})
}

// InsertOne insert a user record to database
func (u *UserRepo) InsertOne(user *user.User) error {
	return parseError(u.db.conn.Create(user).Error)
}

// FindOneByID find a user record by ID
func (u *UserRepo) FindOneByID(id string) (result user.User, err error) {
	err = parseError(u.db.conn.Where("id = ?", id).First(&result).Error)
	return
}

// FindMany find a list of user records
func (u *UserRepo) FindMany(limit, offset int) (result []user.User, err error) {
	err = parseError(u.db.conn.Limit(limit).Offset(offset).Find(&result).Error)
	return
}

// UpdateOne save a user record to database
func (u *UserRepo) UpdateOne(user *user.User) error {
	return parseError(u.db.conn.Save(user).Error)
}

// DeleteOne delete a user record from database
func (u *UserRepo) DeleteOne(user *user.User) error {
	return parseError(u.db.conn.Delete(user).Error)
}
