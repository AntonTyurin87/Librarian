package dto

import "Librarian/internal/pkg/domain/entity"

const (
	BooksTableName = "books"

	BooksColumnID          = "id"
	BooksColumnNameRu      = "name_ru"
	BooksColumnNameNative  = "name_native"
	BooksColumnAuthorRu    = "author_ru"
	BooksColumnYear        = "year"
	BooksColumnRegions     = "regions"
	BooksColumnTimePeriods = "time_periods"
	BooksColumnDescription = "description"
	BooksColumnFileFormat  = "file_format"
)

// BooksColumns ...
var BooksColumns = []string{
	BooksColumnID,
	BooksColumnNameRu,
	BooksColumnNameNative,
	BooksColumnAuthorRu,
	BooksColumnYear,
	BooksColumnRegions,
	BooksColumnTimePeriods,
	BooksColumnDescription,
	BooksColumnFileFormat,
}

// Book ...
type Book struct {
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

// Entity ...
func (b *Book) Entity() *entity.Book {
	if b == nil {
		return nil
	}

	return &entity.Book{
		ID:          b.GetID(),
		NameRu:      b.GetNameRu(),
		NameNative:  b.GetNameNative(),
		AuthorRu:    b.GetAuthorRu(),
		Year:        b.GetYear(),
		Regions:     b.GetRegions(),
		TimePeriods: b.GetTimePeriods(),
		Description: b.GetDescription(),
		FileFormat:  b.GetFileFormat(),
	}
}

// Books ...
type Books []*Book

// Entity ...
func (b Books) Entity() []*entity.Book { return ToEntitySlice[[]*entity.Book](b) }
