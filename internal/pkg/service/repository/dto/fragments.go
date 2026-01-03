package dto

import "Librarian/internal/pkg/domain/entity"

const (
	FragmentsTableName = "fragments"

	FragmentsColumnID          = "id"
	FragmentsColumnNameRu      = "name_ru"
	FragmentsColumnNameNative  = "name_native"
	FragmentsColumnAuthorRu    = "author_ru"
	FragmentsColumnYear        = "year"
	FragmentsColumnRegions     = "regions"
	FragmentsColumnTimePeriods = "time_periods"
	FragmentsColumnDescription = "description"
	FragmentsColumnFormat      = "format"
)

// FragmentsColumns ...
var FragmentsColumns = []string{
	FragmentsColumnID,
	FragmentsColumnNameRu,
	FragmentsColumnNameNative,
	FragmentsColumnAuthorRu,
	FragmentsColumnYear,
	FragmentsColumnRegions,
	FragmentsColumnTimePeriods,
	FragmentsColumnDescription,
	FragmentsColumnFormat,
}

// Fragment ...
type Fragment struct {
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

// GetFormat возвращает формат файла
func (f *Fragment) GetFormat() int32 {
	return f.Format
}

// Entity ...
func (f *Fragment) Entity() *entity.Fragment {
	if f == nil {
		return nil
	}

	return &entity.Fragment{
		ID:          f.GetID(),
		NameRu:      f.GetNameRu(),
		NameNative:  f.GetNameNative(),
		AuthorRu:    f.GetAuthorRu(),
		Year:        f.GetYear(),
		Regions:     f.GetRegions(),
		TimePeriods: f.GetTimePeriods(),
		Description: f.GetDescription(),
		Format:      entity.FileFormat(f.GetFormat()),
	}
}

// Fragments ...
type Fragments []*Fragment

// Entity ...
func (f Fragments) Entity() []*entity.Fragment { return ToEntitySlice[[]*entity.Fragment](f) }

// FragmentDtoFromEntity ...
func FragmentDtoFromEntity(e *entity.Fragment) *Fragment {
	if e == nil {
		return nil
	}
	return &Fragment{
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
