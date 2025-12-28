package dto

import "Librarian/internal/pkg/domain/entity"

const (
	PhotosTableName = "photos"

	PhotosColumnID              = "id"
	PhotosColumnGroupId         = "group_id"
	PhotosColumnNameRu          = "ru_name"
	PhotosColumnNameNative      = "native_name"
	PhotosColumnFilmingLocation = "filming_location"
	PhotosColumnRegions         = "regions"
	PhotosColumnTimePeriods     = "time_periods"
	PhotosColumnDescription     = "description"
	PhotosColumnFileFormat      = "file_format"
)

// PhotosColumns ...
var PhotosColumns = []string{
	PhotosColumnID,
	PhotosColumnGroupId,
	PhotosColumnNameRu,
	PhotosColumnNameNative,
	PhotosColumnFilmingLocation,
	PhotosColumnRegions,
	PhotosColumnTimePeriods,
	PhotosColumnDescription,
	PhotosColumnFileFormat,
}

// Photo ...
type Photo struct {
	ID              int32  `db:"id"`
	GroupID         int32  `db:"group_id"`
	NameRu          string `db:"name_ru"`
	NameNative      string `db:"name_native"`
	FilmingLocation string `db:"filming_location"`
	Regions         string `db:"regions"`
	TimePeriods     string `db:"time_periods"`
	Description     string `db:"description"`
	FileFormat      string `db:"file_format"`
}

// GetID возвращает ID фотографии
func (p *Photo) GetID() int32 {
	if p == nil {
		return 0
	}
	return p.ID
}

// GetGroupID возвращает ID группы фотографий
func (p *Photo) GetGroupID() int32 {
	if p == nil {
		return 0
	}
	return p.GroupID
}

// GetNameRu возвращает название на русском
func (p *Photo) GetNameRu() string {
	if p == nil {
		return ""
	}
	return p.NameRu
}

// GetNameNative возвращает оригинальное название
func (p *Photo) GetNameNative() string {
	if p == nil {
		return ""
	}
	return p.NameNative
}

// GetFilmingLocation возвращает место съемки
func (p *Photo) GetFilmingLocation() string {
	if p == nil {
		return ""
	}
	return p.FilmingLocation
}

// GetRegions возвращает регионы
func (p *Photo) GetRegions() string {
	if p == nil {
		return ""
	}
	return p.Regions
}

// GetTimePeriods возвращает временные периоды
func (p *Photo) GetTimePeriods() string {
	if p == nil {
		return ""
	}
	return p.TimePeriods
}

// GetDescription возвращает описание
func (p *Photo) GetDescription() string {
	if p == nil {
		return ""
	}
	return p.Description
}

// GetFileFormat возвращает формат файла
func (p *Photo) GetFileFormat() string {
	if p == nil {
		return ""
	}
	return p.FileFormat
}

// Entity ...
func (p *Photo) Entity() *entity.Photo {
	if p == nil {
		return nil
	}

	return &entity.Photo{
		ID:              p.GetID(),
		GroupID:         p.GetGroupID(),
		NameRu:          p.GetNameRu(),
		NameNative:      p.GetNameNative(),
		FilmingLocation: p.GetFilmingLocation(),
		Regions:         p.GetRegions(),
		TimePeriods:     p.GetTimePeriods(),
		Description:     p.GetDescription(),
		FileFormat:      p.GetFileFormat(),
	}
}

// Photos ...
type Photos []*Photo

// Entity ...
func (p Photos) Entity() []*entity.Photo { return ToEntitySlice[[]*entity.Photo](p) }
