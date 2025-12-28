package presenter

import (
	"Librarian/internal/pkg/domain/entity"

	lib "github.com/AntonTyurin87/Recon_Com_protoc/gen/go/librarian"
)

// RegionsFromEntityToLib ...
func (p *presenter) RegionsFromEntityToLib(regionsEntity []*entity.Region) []*lib.Region {
	if regionsEntity == nil {
		return nil
	}

	// Создаем срез с нужной емкостью, но нулевой длиной
	regions := make([]*lib.Region, 0, len(regionsEntity))

	for _, region := range regionsEntity {
		regionLib := &lib.Region{
			Id:          region.GetID(),
			NameRu:      region.GetNameRu(),
			Description: region.GetDescription(),
		}
		regions = append(regions, regionLib)
	}

	return regions
}
