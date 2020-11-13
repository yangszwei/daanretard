package administrator

// NewService create a Service
func NewService(r IRepository) *Service {
	s := new(Service)
	s.r = r
	return s
}

// Service implement IUsecase
type Service struct {
	r IRepository
}

// Add add a user to administrator repository
func (s *Service) Add(userID uint32) error {
	return s.r.InsertOne(userID)
}

// IsAdministrator return nil is user is administrator
func (s *Service) IsAdministrator(userID uint32) error {
	_, err := s.r.FindOneByID(userID)
	return err
}

// SetFbAccessToken set admin facebook access token
func (s *Service) SetFbAccessToken(userID uint32, accessToken string) error {
	admin, err := s.r.FindOneByID(userID)
	if err != nil {
		return err
	}
	admin.FbAccessToken = accessToken
	return s.r.UpdateOne(admin)
}

// GetFbAccessToken get admin facebook access token
func (s *Service) GetFbAccessToken(userID uint32) (string, error) {
	admin, err := s.r.FindOneByID(userID)
	if err != nil {
		return "", err
	}
	return admin.FbAccessToken, nil
}

// Delete remove a user from administrator repository
func (s *Service) Delete(userID uint32) error {
	admin, err := s.r.FindOneByID(userID)
	if err != nil {
		return err
	}
	return s.r.DeleteOne(admin)
}
