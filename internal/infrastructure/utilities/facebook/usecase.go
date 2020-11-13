package facebook

// IUsecase interface
type IUsecase interface {
	PublishPost(message string, attachments []string, accessToken string) (string, error)
	PublishPhoto(path, accessToken string) (string, error)
	DeletePost(id, accessToken string) error
	ExchangeToken(accessToken string) (string, error)
}
