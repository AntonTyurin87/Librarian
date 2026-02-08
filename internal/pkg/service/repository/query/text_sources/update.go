package text_sources

import "Librarian/internal/pkg/domain/entity"

// Update ...
type Update struct {
	TextSources []*entity.TextSource
}

// GetTextSources ...
func (s *Update) GetTextSources() []*entity.TextSource {
	if s == nil {
		return nil
	}

	return s.TextSources
}
