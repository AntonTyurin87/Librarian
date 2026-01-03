package photos

import "Librarian/internal/pkg/domain/entity"

// Update ...
type Update struct {
	Photos []*entity.Photo
}

// GetPhotos ...
func (s *Update) GetPhotos() []*entity.Photo {
	if s == nil {
		return nil
	}

	return s.Photos
}
