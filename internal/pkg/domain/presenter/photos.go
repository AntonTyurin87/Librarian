package presenter

import (
	"Librarian/internal/pkg/domain/entity"

	lib "github.com/AntonTyurin87/Recon_Com_protoc/gen/go/librarian"
)

// PhotosFromLibToEntity ...
func (p *presenter) PhotosFromLibToEntity(photo *lib.Photo) *entity.Photo {
	if photo == nil {
		return nil
	}

	result := entity.Photo{
		GroupID:         photo.GetGroupId(),
		NameRu:          photo.GetNameRu(),
		NameNative:      photo.GetNameNative(),
		FilmingLocation: photo.GetFilmingLocation(),
		Regions:         photo.GetRegions(),
		TimePeriods:     photo.GetTimePeriods(),
		Description:     photo.GetDescription(),
		Format:          p.FileFormatFromLibToEntity(photo.GetFileFormat()),
	}
	return &result
}
