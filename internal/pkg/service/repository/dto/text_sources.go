package dto

import (
	"Librarian/internal/pkg/domain/entity"
)

const (
	TextSourcesTableName = "text_sources"

	TextSourcesColumnID          = "id"
	TextSourcesColumnUserID      = "user_id"
	TextSourcesColumnType        = "type"
	TextSourcesColumnNameRU      = "name_ru"
	TextSourcesColumnNameENG     = "name_eng"
	TextSourcesColumnAuthorRU    = "author_ru"
	TextSourcesColumnYear        = "year"
	TextSourcesColumnDescription = "description"
	TextSourcesColumnPlaceURL    = "place_url"
	TextSourcesColumnFromURL     = "from_url"
	TextSourcesColumnFileName    = "file_name"
	TextSourcesColumnFileFormat  = "file_format"
	TextSourcesColumnCreatedAt   = "created_at"
	TextSourcesColumnIsAvailable = "is_available"
)

var TextSourcesColumns = []string{
	TextSourcesColumnID,
	TextSourcesColumnUserID,
	TextSourcesColumnType,
	TextSourcesColumnNameRU,
	TextSourcesColumnNameENG,
	TextSourcesColumnAuthorRU,
	TextSourcesColumnYear,
	TextSourcesColumnDescription,
	TextSourcesColumnPlaceURL,
	TextSourcesColumnFromURL,
	TextSourcesColumnFileName,
	TextSourcesColumnFileFormat,
	TextSourcesColumnCreatedAt,
	TextSourcesColumnIsAvailable,
}

// TextSource ...
type TextSource struct {
	ID          int64  `json:"id"`
	UserID      int64  `json:"user_id"`
	Type        string `json:"type"`
	NameRU      string `json:"name_ru"`
	NameENG     string `json:"name_eng"`
	AuthorRU    string `json:"author_ru"`
	Year        int64  `json:"year"`
	Description string `json:"description"`
	PlaceURL    string `json:"place_url"`
	FromURL     string `json:"from_url"`
	FileName    string `json:"file_name"`
	FileFormat  string `json:"file_format"`
	CreatedAt   string `json:"created_at"`
	IsAvailable int64  `json:"is_available"`
}

// GetID ...
func (ts *TextSource) GetID() int64 {
	if ts == nil {
		return 0
	}
	return ts.ID
}

// GetUserID ...
func (ts *TextSource) GetUserID() int64 {
	if ts == nil {
		return 0
	}
	return ts.UserID
}

// GetType ...
func (ts *TextSource) GetType() string {
	if ts == nil {
		return ""
	}
	return ts.Type
}

// GetNameRU ...
func (ts *TextSource) GetNameRU() string {
	if ts == nil {
		return ""
	}
	return ts.NameRU
}

// GetNameENG ...
func (ts *TextSource) GetNameENG() string {
	if ts == nil {
		return ""
	}
	return ts.NameENG
}

// GetAuthorRU ...
func (ts *TextSource) GetAuthorRU() string {
	if ts == nil {
		return ""
	}
	return ts.AuthorRU
}

// GetYear ...
func (ts *TextSource) GetYear() int64 {
	if ts == nil {
		return 0
	}
	return ts.Year
}

// GetDescription ...
func (ts *TextSource) GetDescription() string {
	if ts == nil {
		return ""
	}
	return ts.Description
}

// GetPlaceURL ...
func (ts *TextSource) GetPlaceURL() string {
	if ts == nil {
		return ""
	}
	return ts.PlaceURL
}

// GetFromURL ...
func (ts *TextSource) GetFromURL() string {
	if ts == nil {
		return ""
	}
	return ts.FromURL
}

// GetFileName ...
func (ts *TextSource) GetFileName() string {
	if ts == nil {
		return ""
	}
	return ts.FileName
}

// GetFileFormat ...
func (ts *TextSource) GetFileFormat() string {
	if ts == nil {
		return ""
	}
	return ts.FileFormat
}

// GetCreatedAt ...
func (ts *TextSource) GetCreatedAt() string {
	if ts == nil {
		return ""
	}
	return ts.CreatedAt
}

// GetIsAvailable ...
func (ts *TextSource) GetIsAvailable() int64 {
	if ts == nil {
		return 0
	}
	return ts.IsAvailable
}

// Entity ...
func (ts *TextSource) Entity() *entity.TextSource {
	if ts == nil {
		return nil
	}

	return &entity.TextSource{
		ID:          ts.GetID(),
		UserID:      ts.GetUserID(),
		Type:        entity.SourceType(ts.GetType()),
		NameRU:      ts.GetNameRU(),
		NameENG:     ts.GetNameENG(),
		AuthorRU:    ts.GetAuthorRU(),
		Year:        ts.GetYear(),
		Description: ts.GetDescription(),
		PlaceURL:    ts.GetPlaceURL(),
		FromURL:     ts.GetFromURL(),
		FileName:    ts.GetFileName(),
		FileFormat:  ts.GetFileFormat(),
		CreatedAt:   ts.GetCreatedAt(),
		IsAvailable: ts.GetIsAvailable(),
	}
}

// TextSources ...
type TextSources []*TextSource

// Entity ...
func (ts TextSources) Entity() []*entity.TextSource {
	return ToEntitySlice[[]*entity.TextSource](ts)
}

// TextSourceDtoFromEntity ...
func TextSourceDtoFromEntity(e *entity.TextSource) *TextSource {
	if e == nil {
		return nil
	}
	return &TextSource{
		ID:          e.GetID(),
		UserID:      e.GetUserID(),
		Type:        string(e.GetType()),
		NameRU:      e.GetNameRU(),
		NameENG:     e.GetNameENG(),
		AuthorRU:    e.GetAuthorRU(),
		Year:        e.GetYear(),
		Description: e.GetDescription(),
		PlaceURL:    e.GetPlaceURL(),
		FromURL:     e.GetFromURL(),
		FileName:    e.GetFileName(),
		FileFormat:  e.GetFileFormat(),
		CreatedAt:   e.GetCreatedAt(),
		IsAvailable: e.GetIsAvailable(),
	}
}
