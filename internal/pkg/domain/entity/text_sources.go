package entity

// TextSource ...
type TextSource struct {
	ID          int64      `json:"id"`
	UserID      int64      `json:"user_id"`
	Type        SourceType `json:"type"`
	NameRU      string     `json:"name_ru"`
	NameENG     string     `json:"name_eng"`
	AuthorRU    string     `json:"author_ru"`
	Year        int64      `json:"year"`
	Description string     `json:"description"`
	PlaceURL    string     `json:"place_url"`
	FromURL     string     `json:"from_url"`
	FileName    string     `json:"file_name"`
	FileFormat  string     `json:"file_format"`
	CreatedAt   string     `json:"created_at"`
	IsAvailable int64      `json:"is_available"`
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
func (ts *TextSource) GetType() SourceType {
	if ts == nil {
		return SourceTypeUnknown
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
