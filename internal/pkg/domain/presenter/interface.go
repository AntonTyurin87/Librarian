package presenter

import (
	"Librarian/internal/pkg/domain/entity"

	lib "github.com/AntonTyurin87/Recon_Com_protoc/gen/go/librarian"
)

type Interface interface {
	RegionsFromEntityToLib(regionsEntity []*entity.Region) []*lib.Region

	SourceTypeFromLibToEntity(source lib.SourceType) entity.SourceType

	//SourcesFromEntityToSources(sourcesEntity []*entity.ForDownloadSource) for_download_sources.Select
}
