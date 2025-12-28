package presenter

import (
	"Librarian/internal/pkg/domain/entity"
	"Librarian/internal/pkg/service/repository/query/sources"

	lib "github.com/AntonTyurin87/Recon_Com_protoc/gen/go/librarian"
)

type Interface interface {
	RegionsFromEntityToLib(regionsEntity []*entity.Region) []*lib.Region
	SourcesFromEntityToSourcies(sourcesEntity []*entity.Source) sources.Select
}
