package entity

// Article ...
type Article struct {
	ID          int32  `json:"id"`
	NameRu      string `json:"name_ru"`
	NameNative  string `json:"name_native"`
	AuthorRu    string `json:"author_ru"`
	Year        int32  `json:"year"`
	Regions     string `json:"regions"`
	TimePeriods string `json:"time_periods"`
	Description string `json:"description"`
	FileFormat  string `json:"file_format"`
}

// GetID возвращает ID книги
func (a *Article) GetID() int32 {
	if a == nil {
		return 0
	}
	return a.ID
}

// GetNameRu возвращает название книги на русском
func (a *Article) GetNameRu() string {
	if a == nil {
		return ""
	}
	return a.NameRu
}

// GetNameNative возвращает оригинальное название книги
func (a *Article) GetNameNative() string {
	if a == nil {
		return ""
	}
	return a.NameNative
}

// GetAuthorRu возвращает автора на русском
func (a *Article) GetAuthorRu() string {
	if a == nil {
		return ""
	}
	return a.AuthorRu
}

// GetYear возвращает год издания
func (a *Article) GetYear() int32 {
	if a == nil {
		return 0
	}
	return a.Year
}

// GetRegions возвращает регионы
func (a *Article) GetRegions() string {
	if a == nil {
		return ""
	}
	return a.Regions
}

// GetTimePeriods возвращает временные периоды
func (a *Article) GetTimePeriods() string {
	if a == nil {
		return ""
	}
	return a.TimePeriods
}

// GetDescription возвращает описание книги
func (a *Article) GetDescription() string {
	if a == nil {
		return ""
	}
	return a.Description
}

// GetFileFormat возвращает формат файла
func (a *Article) GetFileFormat() string {
	if a == nil {
		return ""
	}
	return a.FileFormat
}
