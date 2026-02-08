package for_download_sources

import "Librarian/internal/pkg/domain/entity"

// Update ...
type Update struct {
	ForDownloadSources []*entity.ForDownloadSource
}

// GetForDownloadSources ...
func (s *Update) GetForDownloadSources() []*entity.ForDownloadSource {
	if s == nil {
		return nil
	}

	return s.ForDownloadSources
}
