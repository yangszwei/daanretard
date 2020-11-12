package mock_persistence

import "errors"

// NewAttachmentRepository create a AttachmentRepository
func NewAttachmentRepository() *AttachmentRepository {
	a := new(AttachmentRepository)
	a.count = 0
	a.list = make(map[uint32]string)
	return a
}

// AttachmentRepository implement attachment.IRepository
type AttachmentRepository struct {
	count uint32
	list  map[uint32]string
}

// InsertOne add an attachment
func (a *AttachmentRepository) InsertOne(name string) (uint32, error) {
	a.count++
	a.list[a.count] = name
	return a.count, nil
}

// FindOneByID find an attachment
func (a *AttachmentRepository) FindOneByID(id uint32) (string, error) {
	if name, exist := a.list[id] ; exist {
		return name, nil
	}
	return "", errors.New("record not found")
}

// DeleteOne delete an attachment
func (a *AttachmentRepository) DeleteOne(id uint32) error {
	delete(a.list, id)
	return nil
}
