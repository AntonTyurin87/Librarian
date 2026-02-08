package dto

import (
	"Librarian/internal/pkg/domain/entity"
)

const (
	ForDownloadSourcesTableName = "for_download_sources"

	ForDownloadSourcesColumnID             = "id"
	ForDownloadSourcesColumnUserID         = "user_id"
	ForDownloadSourcesColumnType           = "type"
	ForDownloadSourcesColumnNameRU         = "name_ru"
	ForDownloadSourcesColumnNameENG        = "name_eng"
	ForDownloadSourcesColumnAuthorRU       = "author_ru"
	ForDownloadSourcesColumnYear           = "year"
	ForDownloadSourcesColumnDescription    = "description"
	ForDownloadSourcesColumnDownloadURL    = "download_url"
	ForDownloadSourcesColumnCreatedAt      = "created_at"
	ForDownloadSourcesColumnIsFileDownload = "is_file_download"
	ForDownloadSourcesColumnIsDownload     = "is_download"
)

var ForDownloadSourcesColumns = []string{
	ForDownloadSourcesColumnID,
	ForDownloadSourcesColumnUserID,
	ForDownloadSourcesColumnType,
	ForDownloadSourcesColumnNameRU,
	ForDownloadSourcesColumnNameENG,
	ForDownloadSourcesColumnAuthorRU,
	ForDownloadSourcesColumnYear,
	ForDownloadSourcesColumnDescription,
	ForDownloadSourcesColumnDownloadURL,
	ForDownloadSourcesColumnCreatedAt,
	ForDownloadSourcesColumnIsFileDownload,
	ForDownloadSourcesColumnIsDownload,
}

// ForDownloadSource ...
type ForDownloadSource struct {
	ID             int64  `json:"id"`
	UserID         int64  `json:"user_id"`
	Type           string `json:"type"`
	NameRU         string `json:"name_ru"`
	NameENG        string `json:"name_eng"`
	AuthorRU       string `json:"author_ru"`
	Year           int64  `json:"year"`
	Description    string `json:"description"`
	DownloadURL    string `json:"download_url"`
	CreatedAt      string `json:"created_at"`
	IsFileDownload int64  `json:"isSent"`
	IsDownload     int64  `json:"is_download"`
}

// GetID ...
func (s *ForDownloadSource) GetID() int64 {
	if s == nil {
		return 0
	}
	return s.ID
}

// GetUserID ...
func (s *ForDownloadSource) GetUserID() int64 {
	if s == nil {
		return 0
	}
	return s.UserID
}

// GetType ...
func (s *ForDownloadSource) GetType() string {
	if s == nil {
		return ""
	}
	return s.Type
}

// GetNameRU ...
func (s *ForDownloadSource) GetNameRU() string {
	if s == nil {
		return ""
	}
	return s.NameRU
}

// GetNameENG ...
func (s *ForDownloadSource) GetNameENG() string {
	if s == nil {
		return ""
	}
	return s.NameENG
}

// GetAuthorRU ...
func (s *ForDownloadSource) GetAuthorRU() string {
	if s == nil {
		return ""
	}
	return s.AuthorRU
}

// GetYear ...
func (s *ForDownloadSource) GetYear() int64 {
	if s == nil {
		return 0
	}
	return s.Year
}

// GetDescription ...
func (s *ForDownloadSource) GetDescription() string {
	if s == nil {
		return ""
	}
	return s.Description
}

func (s *ForDownloadSource) GetDownloadURL() string {
	if s == nil {
		return ""
	}
	return s.DownloadURL
}

// GetCreatedAt ...
func (s *ForDownloadSource) GetCreatedAt() string {
	if s == nil {
		return ""
	}
	return s.CreatedAt
}

// GetIsFileDownload ...
func (s *ForDownloadSource) GetIsFileDownload() int64 {
	if s == nil {
		return 0
	}
	return s.IsFileDownload
}

// GetIsDownload ...
func (s *ForDownloadSource) GetIsDownload() int64 {
	if s == nil {
		return 0
	}
	return s.IsDownload
}

// Entity ...
func (s *ForDownloadSource) Entity() *entity.ForDownloadSource {
	if s == nil {
		return nil
	}

	return &entity.ForDownloadSource{
		ID:             s.GetID(),
		UserID:         s.GetUserID(),
		Type:           entity.SourceType(s.GetType()),
		NameRU:         s.GetNameRU(),
		NameENG:        s.GetNameENG(),
		AuthorRU:       s.GetAuthorRU(),
		Year:           s.GetYear(),
		Description:    s.GetDescription(),
		DownloadURL:    s.GetDownloadURL(),
		CreatedAt:      s.GetCreatedAt(),
		IsFileDownload: s.GetIsFileDownload(),
		IsDownload:     s.GetIsDownload(),
	}
}

// ForDownloadSources ...
type ForDownloadSources []*ForDownloadSource

// Entity ...
func (s ForDownloadSources) Entity() []*entity.ForDownloadSource {
	return ToEntitySlice[[]*entity.ForDownloadSource](s)
}

// ForDownloadSourceDtoFromEntity ...
func ForDownloadSourceDtoFromEntity(e *entity.ForDownloadSource) *ForDownloadSource {
	if e == nil {
		return nil
	}
	return &ForDownloadSource{
		ID:             e.GetID(),
		UserID:         e.GetUserID(),
		Type:           string(e.GetType()),
		NameRU:         e.GetNameRU(),
		NameENG:        e.GetNameENG(),
		AuthorRU:       e.GetAuthorRU(),
		Year:           e.GetYear(),
		Description:    e.GetDescription(),
		DownloadURL:    e.GetDownloadURL(),
		CreatedAt:      e.GetCreatedAt(),
		IsFileDownload: e.GetIsFileDownload(),
		IsDownload:     e.GetIsDownload(),
	}
}
