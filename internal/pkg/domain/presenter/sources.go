package presenter

import (
	"Librarian/internal/pkg/domain/entity"
	"Librarian/internal/pkg/service/repository/query/sources"
)

// SourcesFromEntityToSourcies ...
func (p *presenter) SourcesFromEntityToSourcies(sourcesEntity []*entity.Source) sources.Select {
	sourcesSelect := sources.Select{}

	if sourcesEntity == nil {
		// если пустой запрос почему-то пришел, то отправим один заведомо несуществующий ID, чтобы не вернуть все источники пачкой
		sourcesSelect.IDs = append(sourcesSelect.IDs, 0)
		return sourcesSelect
	}

	for _, source := range sourcesEntity {
		if source.ID != 0 {
			sourcesSelect.IDs = append(sourcesSelect.IDs, source.ID)
		}
		if source.ObjectID != 0 {
			sourcesSelect.IDs = append(sourcesSelect.IDs, source.ObjectID)
		}
		if source.Type != "" {
			sourcesSelect.Types = append(sourcesSelect.Types, source.Type)
		}
		if source.Address != "" {
			sourcesSelect.Address = append(sourcesSelect.Address, source.Address)
		}
	}

	return sourcesSelect
}
