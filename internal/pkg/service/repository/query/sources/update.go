package sources

import "Librarian/internal/pkg/domain/entity"

// Update ...
type Update struct {
	Sources []*entity.Source
}

// GetSources ...
func (s *Update) GetSources() []*entity.Source {
	if s == nil {
		return nil
	}

	return s.Sources
}
