package presenter

import (
	"Librarian/internal/pkg/domain/entity"
	"Librarian/internal/pkg/service/repository/query/text_sources"
	lib "github.com/AntonTyurin87/Recon_Com_protoc/gen/go/librarian"
)

// SourcesFromEntityToTextSources ...
func (p *presenter) SourcesFromEntityToTextSources(sources []*entity.TextSource) text_sources.Select {
	IDs := make([]int64, 0, len(sources))
	for _, source := range sources {
		IDs = append(IDs, source.ID)
	}

	return text_sources.Select{
		IDs: IDs,
	}
}

// SourceTypeFromLibToEntity ...
func (p *presenter) SourceTypeFromLibToEntity(source lib.SourceType) entity.SourceType {
	switch source {
	case lib.SourceType_SOURCE_TYPE_BOOK:
		return entity.SourceTypeBook
	case lib.SourceType_SOURCE_TYPE_ARTICLE:
		return entity.SourceTypeArticle
	case lib.SourceType_SOURCE_TYPE_FRAGMENT:
		return entity.SourceTypeFragment
	default:
		return entity.SourceTypeUnknown
	}
}
