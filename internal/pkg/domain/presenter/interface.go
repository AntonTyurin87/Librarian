package presenter

import (
	"Librarian/internal/pkg/domain/entity"
	"Librarian/internal/pkg/service/repository/query/text_sources"

	lib "github.com/AntonTyurin87/Recon_Com_protoc/gen/go/librarian"
)

type Interface interface {
	RegionsFromEntityToLib(regionsEntity []*entity.Region) []*lib.Region

	SourceTypeFromLibToEntity(source lib.SourceType) entity.SourceType

	SourcesFromEntityToTextSources(sources []*entity.TextSource) text_sources.Select
}
