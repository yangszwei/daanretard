package registry

import (
	"daanretard/internal/entity/admin"
	"daanretard/internal/entity/media"
	"daanretard/internal/entity/post"
	"daanretard/internal/entity/user"
	"daanretard/internal/infra/fbgraph"
	"daanretard/internal/infra/validator"
)

// newService create an instance of service
func newService(repo repo, fb fbgraph.IUsecase) service {
	var s service
	v := validator.New()
	s.user = user.NewService(repo.user, fb)
	s.post = post.NewService(v, repo.post, fb)
	s.admin = admin.NewService(repo.admin, fb, v)
	s.media = media.NewService(repo.media, v)
	return s
}

// service contain all service
type service struct {
	user  user.IUsecase
	post  post.IUsecase
	admin admin.IUsecase
	media media.IUsecase
}
