package attachment

// IUsecase interface
type IUsecase interface {
	Add(name string) (uint32, error)
	GetOne(id uint32) (string, error)
	Delete(id uint32) error
}