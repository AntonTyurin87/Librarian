package entity

// UploadTextSourceForDownloadRequest ...
type UploadTextSourceForDownloadRequest struct {
	ForDownloadSource *ForDownloadSource
}

// GetForDownloadSource ...
func (u *UploadTextSourceForDownloadRequest) GetForDownloadSource() *ForDownloadSource {
	if u == nil {
		return nil
	}
	return u.ForDownloadSource
}

// UploadTextSourceForDownloadResponse ...
type UploadTextSourceForDownloadResponse struct {
	ID int64
}

type SourceType string

const (
	SourceTypeUnknown  SourceType = "Unknown"
	SourceTypeBook     SourceType = "Книга"
	SourceTypeArticle  SourceType = "Статья"
	SourceTypeFragment SourceType = "Фрагмент"
)

// ForDownloadSource ...
type ForDownloadSource struct {
	ID             int64      `json:"id"`
	UserID         int64      `json:"user_id"`
	Type           SourceType `json:"type"`
	NameRU         string     `json:"name_ru"`
	NameENG        string     `json:"name_eng"`
	AuthorRU       string     `json:"author_ru"`
	Year           int64      `json:"year"`
	Description    string     `json:"description"`
	DownloadURL    string     `json:"download_url"`
	CreatedAt      string     `json:"created_at"`
	IsFileDownload int64      `json:"isSent"`
	IsDownload     int64      `json:"is_download"`
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
func (s *ForDownloadSource) GetType() SourceType {
	if s == nil {
		return SourceTypeUnknown
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

// GetDownloadURL ...
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
