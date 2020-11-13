package administrator

// IUsecase interface
type IUsecase interface {
	Add(userID uint32) error
	IsAdministrator(userID uint32) error
	SetFbAccessToken(userID uint32, accessToken string) error
	GetFbAccessToken(userID uint32) (string, error)
	Delete(userID uint32) error
}
