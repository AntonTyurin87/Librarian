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
	ArticlesColumnFormat      = "format"
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
	ArticlesColumnFormat,
}

// Article ...
type Article struct {
	ID          int32    `db:"id"`
	NameRu      string   `db:"name_ru"`
	NameNative  string   `db:"name_native"`
	AuthorRu    string   `db:"author_ru"`
	Year        int32    `db:"year"`
	Regions     []string `db:"regions"`
	TimePeriods []string `db:"time_periods"`
	Description string   `db:"description"`
	Format      int32    `db:"format"`
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
func (a *Article) GetRegions() []string {
	if a == nil {
		return nil
	}
	return a.Regions
}

// GetTimePeriods возвращает временные периоды
func (a *Article) GetTimePeriods() []string {
	if a == nil {
		return nil
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

// GetFormat возвращает формат файла
func (a *Article) GetFormat() int32 {
	return a.Format
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
		Format:      entity.FileFormat(a.GetFormat()), //TODO проверить правильность работы!
	}
}

// Articles ...
type Articles []*Article

// Entity ...
func (a Articles) Entity() []*entity.Article { return ToEntitySlice[[]*entity.Article](a) }

// ArticleDtoFromEntity ...
func ArticleDtoFromEntity(e *entity.Article) *Article {
	if e == nil {
		return nil
	}
	return &Article{
		ID:          e.ID,
		NameRu:      e.NameRu,
		NameNative:  e.NameNative,
		AuthorRu:    e.AuthorRu,
		Year:        e.Year,
		Regions:     e.Regions,
		TimePeriods: e.TimePeriods,
		Description: e.Description,
		Format:      int32(e.Format),
	}
}
