package by_regions

import "Librarian/internal/pkg/domain/entity"

// Update ...
type Update struct {
	ByRegions []*entity.ByRegion
}

// GetByRegions ...
func (s *Update) GetByRegions() []*entity.ByRegion {
	if s == nil {
		return nil
	}

	return s.ByRegions
}
