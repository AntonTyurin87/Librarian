package entity

type Fragment struct {
	ID          int32      `json:"id"`
	NameRu      string     `json:"name_ru"`
	NameNative  string     `json:"name_native"`
	AuthorRu    string     `json:"author_ru"`
	Year        int32      `json:"year"`
	Regions     []string   `json:"regions"`
	TimePeriods []string   `json:"time_periods"`
	Description string     `json:"description"`
	Format      FileFormat `json:"format"`
}

// GetID возвращает ID книги
func (f *Fragment) GetID() int32 {
	if f == nil {
		return 0
	}
	return f.ID
}

// GetNameRu возвращает название книги на русском
func (f *Fragment) GetNameRu() string {
	if f == nil {
		return ""
	}
	return f.NameRu
}

// GetNameNative возвращает оригинальное название книги
func (f *Fragment) GetNameNative() string {
	if f == nil {
		return ""
	}
	return f.NameNative
}

// GetAuthorRu возвращает автора на русском
func (f *Fragment) GetAuthorRu() string {
	if f == nil {
		return ""
	}
	return f.AuthorRu
}

// GetYear возвращает год издания
func (f *Fragment) GetYear() int32 {
	if f == nil {
		return 0
	}
	return f.Year
}

// GetRegions возвращает регионы
func (f *Fragment) GetRegions() []string {
	if f == nil {
		return nil
	}
	return f.Regions
}

// GetTimePeriods возвращает временные периоды
func (f *Fragment) GetTimePeriods() []string {
	if f == nil {
		return nil
	}
	return f.TimePeriods
}

// GetDescription возвращает описание книги
func (f *Fragment) GetDescription() string {
	if f == nil {
		return ""
	}
	return f.Description
}

// GetFileFormat возвращает формат файла
func (f *Fragment) GetFileFormat() FileFormat {
	if f == nil {
		return FileFormat_UNKNOWN
	}
	return f.Format
}
