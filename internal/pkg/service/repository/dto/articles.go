package dto

import "Librarian/internal/pkg/domain/entity"

const (
	ArticlesTableName = "articles"

	ArticlesColumnID          = "id"
	ArticlesColumnNameRu      = "name_ru"
	ArticlesColumnNameNative  = "name_native"
	ArticlesColumnAuthorRu    = "author_ru"
	ArticlesColumnYear        = "year"
	ArticlesColumnRegions     = "regions"
	ArticlesColumnTimePeriods = "time_periods"
	ArticlesColumnDescription = "description"
	ArticlesColumnFileFormat  = "file_format"
)

// ArticlesColumns ...
var ArticlesColumns = []string{
	ArticlesColumnID,
	ArticlesColumnNameRu,
	ArticlesColumnNameNative,
	ArticlesColumnAuthorRu,
	ArticlesColumnYear,
	ArticlesColumnRegions,
	ArticlesColumnTimePeriods,
	ArticlesColumnDescription,
	ArticlesColumnFileFormat,
}

// Article ...
type Article struct {
	ID          int32  `db:"id"`
	NameRu      string `db:"name_ru"`
	NameNative  string `db:"name_native"`
	AuthorRu    string `db:"author_ru"`
	Year        int32  `db:"year"`
	Regions     string `db:"regions"`
	TimePeriods string `db:"time_periods"`
	Description string `db:"description"`
	FileFormat  string `db:"file_format"`
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

// Entity ...
func (a *Article) Entity() *entity.Article {
	if a == nil {
		return nil
	}

	return &entity.Article{
		ID:          a.GetID(),
		NameRu:      a.GetNameRu(),
		NameNative:  a.GetNameNative(),
		AuthorRu:    a.GetAuthorRu(),
		Year:        a.GetYear(),
		Regions:     a.GetRegions(),
		TimePeriods: a.GetTimePeriods(),
		Description: a.GetDescription(),
		FileFormat:  a.GetFileFormat(),
	}
}

// Articles ...
type Articles []*Article

// Entity ...
func (a Articles) Entity() []*entity.Article { return ToEntitySlice[[]*entity.Article](a) }
