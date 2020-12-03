package registry

import (
	"daanretard/internal/entity/admin"
	"daanretard/internal/entity/media"
	"daanretard/internal/entity/post"
	"daanretard/internal/entity/user"
	"daanretard/internal/infra/persistence"
	"log"
)

// autoMigrate migrate schemas
func autoMigrate(db *persistence.DB) {
	if err := persistence.NewUserRepo(db).AutoMigrate(); err != nil {
		log.Fatalln("Failed to migrate db table")
	}
	if err := persistence.NewPostRepo(db).AutoMigrate(); err != nil {
		log.Fatalln("Failed to migrate db table")
	}
	if err := persistence.NewAdminRepo(db).AutoMigrate(); err != nil {
		log.Fatalln("Failed to migrate db table")
	}
	if err := persistence.NewMediaRepo(db).AutoMigrate(); err != nil {
		log.Fatalln("Failed to migrate db table")
	}
}

// newRepo create an instance of repo
func newRepo(db *persistence.DB) repo {
	return repo{
		user:  persistence.NewUserRepo(db),
		post:  persistence.NewPostRepo(db),
		admin: persistence.NewAdminRepo(db),
		media: persistence.NewMediaRepo(db),
	}
}

// repo contain all repositories
type repo struct {
	user  user.IRepository
	post  post.IRepository
	admin admin.IRepository
	media media.IRepository
}
