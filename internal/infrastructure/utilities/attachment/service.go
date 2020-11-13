package attachment

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

// Add add an attachment
func (s *Service) Add(name string) (uint32, error) {
	return s.r.InsertOne(name)
}

// Add add an attachment
func (s *Service) GetOne(id uint32) (string, error) {
	return s.r.FindOneByID(id)
}

// Add add an attachment
func (s *Service) Delete(id uint32) error {
	return s.r.DeleteOne(id)
}
