package presenter

import (
	"Librarian/internal/pkg/domain/entity"

	lib "github.com/AntonTyurin87/Recon_Com_protoc/gen/go/librarian"
)

//// SourcesFromEntityToSources ...
//func (p *presenter) SourcesFromEntityToSources(sourcesEntity []*entity.Source) for_download_sources.Select {
//	sourcesSelect := for_download_sources.Select{}
//
//	if sourcesEntity == nil {
//		// если пустой запрос почему-то пришел, то отправим один заведомо несуществующий ID, чтобы не вернуть все источники пачкой
//		sourcesSelect.IDs = append(sourcesSelect.IDs, 0)
//		return sourcesSelect
//	}
//
//	for _, source := range sourcesEntity {
//		if source.ID != 0 {
//			sourcesSelect.IDs = append(sourcesSelect.IDs, source.ID)
//		}
//		if source.ObjectID != 0 {
//			sourcesSelect.IDs = append(sourcesSelect.IDs, source.ObjectID)
//		}
//		if source.Type != entity.SourceType_UNKNOWN {
//			sourcesSelect.Types = append(sourcesSelect.Types, int32(source.Type))
//		}
//		if source.Address != "" {
//			sourcesSelect.Address = append(sourcesSelect.Address, source.Address)
//		}
//		if source.TimePeriods != nil {
//			sourcesSelect.TimePeriods = append(sourcesSelect.TimePeriods, source.TimePeriods...)
//		}
//	}
//
//	return sourcesSelect
//}

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
