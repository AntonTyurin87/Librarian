package repository

import (
	"Librarian/internal/pkg/service/repository/dto"
	"Librarian/internal/pkg/service/repository/query/text_sources"
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

// UpdateTextSources ...
func (r *repository) UpdateTextSources(ctx context.Context, q text_sources.Update) (dto.TextSources, error) {
	var res dto.TextSources

	if err := Selectx(ctx, r.storage, &res, updateTextSourcesQuery(q)); err != nil {
		return nil, fmt.Errorf("selectx(ctx, r.storage, &res, updateTextSourcesQuery(q)): %w", err)
	}

	return res, nil
}

func updateTextSourcesQuery(query text_sources.Update) sq.UpdateBuilder {
	updateQuery := sq.StatementBuilder.Update(dto.TextSourcesTableName).
		Prefix("--UpdateTextSource\n")

	for _, source := range query.TextSources {

		// что дополнять
		if source.ID != 0 {
			updateQuery = updateQuery.Where(sq.Eq{dto.TextSourcesColumnID: source.ID})
		}
		if source.UserID != 0 {
			updateQuery = updateQuery.Where(sq.Eq{dto.TextSourcesColumnUserID: source.UserID})
		}

		// чем дополнять
		if source.Type != "" {
			updateQuery = updateQuery.Set(dto.TextSourcesColumnType, source.Type)
		}
		if source.NameRU != "" {
			updateQuery = updateQuery.Set(dto.TextSourcesColumnNameRU, source.NameRU)
		}
		if source.NameENG != "" {
			updateQuery = updateQuery.Set(dto.TextSourcesColumnNameENG, source.NameENG)
		}
		if source.AuthorRU != "" {
			updateQuery = updateQuery.Set(dto.TextSourcesColumnAuthorRU, source.AuthorRU)
		}
		if source.Year != 0 {
			updateQuery = updateQuery.Set(dto.TextSourcesColumnYear, source.Year)
		}
		if source.Description != "" {
			updateQuery = updateQuery.Set(dto.TextSourcesColumnDescription, source.Description)
		}
		if source.PlaceURL != "" {
			updateQuery = updateQuery.Set(dto.TextSourcesColumnPlaceURL, source.PlaceURL)
		}
		if source.FromURL != "" {
			updateQuery = updateQuery.Set(dto.TextSourcesColumnFromURL, source.FromURL)
		}
		if source.FileName != "" {
			updateQuery = updateQuery.Set(dto.TextSourcesColumnFileName, source.FileName)
		}
		if source.FileFormat != "" {
			updateQuery = updateQuery.Set(dto.TextSourcesColumnFileFormat, source.FileFormat)
		}
		if source.CreatedAt != "" {
			updateQuery = updateQuery.Set(dto.ForDownloadSourcesColumnCreatedAt, source.CreatedAt)
		}
		if source.IsAvailable != 0 {
			updateQuery = updateQuery.Set(dto.TextSourcesColumnIsAvailable, source.IsAvailable)
		}

		updateQuery = updateQuery.Suffix(fmt.Sprintf("RETURNING %s", dto.TextSourcesColumnID))
	}

	return updateQuery
}
