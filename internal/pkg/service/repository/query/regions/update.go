package regions

import "Librarian/internal/pkg/domain/entity"

// Update ...
type Update struct {
	Regions []*entity.Region
}

// GetRegions ...
func (s *Update) GetRegions() []*entity.Region {
	if s == nil {
		return nil
	}

	return s.Regions
}
