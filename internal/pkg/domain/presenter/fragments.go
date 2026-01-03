package presenter

import (
	"Librarian/internal/pkg/domain/entity"

	lib "github.com/AntonTyurin87/Recon_Com_protoc/gen/go/librarian"
)

// FragmentsFromLibToEntity ...
func (p *presenter) FragmentsFromLibToEntity(fragment *lib.Fragment) *entity.Fragment {
	if fragment == nil {
		return nil
	}

	result := entity.Fragment{
		NameRu:      fragment.GetNameRu(),
		NameNative:  fragment.GetNameNative(),
		AuthorRu:    fragment.GetAuthorRu(),
		Year:        fragment.GetYear(),
		Regions:     fragment.GetRegions(),
		TimePeriods: fragment.GetTimePeriods(),
		Description: fragment.GetDescription(),
		Format:      p.FileFormatFromLibToEntity(fragment.GetFileFormat()),
	}
	return &result
}
