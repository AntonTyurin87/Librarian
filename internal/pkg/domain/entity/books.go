package entity

type Book struct {
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
func (b *Book) GetID() int32 {
	if b == nil {
		return 0
	}
	return b.ID
}

// GetNameRu возвращает название книги на русском
func (b *Book) GetNameRu() string {
	if b == nil {
		return ""
	}
	return b.NameRu
}

// GetNameNative возвращает оригинальное название книги
func (b *Book) GetNameNative() string {
	if b == nil {
		return ""
	}
	return b.NameNative
}

// GetAuthorRu возвращает автора на русском
func (b *Book) GetAuthorRu() string {
	if b == nil {
		return ""
	}
	return b.AuthorRu
}

// GetYear возвращает год издания
func (b *Book) GetYear() int32 {
	if b == nil {
		return 0
	}
	return b.Year
}

// GetRegions возвращает регионы
func (b *Book) GetRegions() string {
	if b == nil {
		return ""
	}
	return b.Regions
}

// GetTimePeriods возвращает временные периоды
func (b *Book) GetTimePeriods() string {
	if b == nil {
		return ""
	}
	return b.TimePeriods
}

// GetDescription возвращает описание книги
func (b *Book) GetDescription() string {
	if b == nil {
		return ""
	}
	return b.Description
}

// GetFileFormat возвращает формат файла
func (b *Book) GetFileFormat() string {
	if b == nil {
		return ""
	}
	return b.FileFormat
}
