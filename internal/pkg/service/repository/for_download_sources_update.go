package repository

import (
	"Librarian/internal/pkg/service/repository/dto"
	"Librarian/internal/pkg/service/repository/query/for_download_sources"
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

// UpdateForDownloadSources ...
func (r *repository) UpdateForDownloadSources(ctx context.Context, q for_download_sources.Update) (dto.ForDownloadSources, error) {
	var res dto.ForDownloadSources

	if err := Selectx(ctx, r.storage, &res, updateForDownloadSourcesQuery(q)); err != nil {
		return nil, fmt.Errorf("selectx(ctx, r.storage, &res, updateForDownloadSourcesQuery(q)): %w", err)
	}

	return res, nil
}

func updateForDownloadSourcesQuery(query for_download_sources.Update) sq.UpdateBuilder {
	updateQuery := sq.StatementBuilder.Update(dto.ForDownloadSourcesTableName).
		Prefix("--UpdateForDownloadSources\n")

	for _, source := range query.ForDownloadSources {

		// что дополнять
		if source.ID != 0 {
			updateQuery = updateQuery.Where(sq.Eq{dto.ForDownloadSourcesColumnID: source.ID})
		}
		if source.UserID != 0 {
			updateQuery = updateQuery.Where(sq.Eq{dto.ForDownloadSourcesColumnUserID: source.UserID})
		}

		// чем дополнять
		if source.Type != "" {
			updateQuery = updateQuery.Set(dto.ForDownloadSourcesColumnType, source.Type)
		}
		if source.NameRU != "" {
			updateQuery = updateQuery.Set(dto.ForDownloadSourcesColumnNameRU, source.NameRU)
		}
		if source.NameENG != "" {
			updateQuery = updateQuery.Set(dto.ForDownloadSourcesColumnNameENG, source.NameENG)
		}
		if source.AuthorRU != "" {
			updateQuery = updateQuery.Set(dto.ForDownloadSourcesColumnAuthorRU, source.AuthorRU)
		}
		if source.Year != 0 {
			updateQuery = updateQuery.Set(dto.ForDownloadSourcesColumnYear, source.Year)
		}
		if source.Description != "" {
			updateQuery = updateQuery.Set(dto.ForDownloadSourcesColumnDescription, source.Description)
		}
		if source.DownloadURL != "" {
			updateQuery = updateQuery.Set(dto.ForDownloadSourcesColumnDownloadURL, source.DownloadURL)
		}
		if source.CreatedAt != "" {
			updateQuery = updateQuery.Set(dto.ForDownloadSourcesColumnCreatedAt, source.CreatedAt)
		}
		if source.IsFileDownload != 0 {
			updateQuery = updateQuery.Set(dto.ForDownloadSourcesColumnIsFileDownload, source.IsFileDownload)
		}
		if source.IsDownload != 0 {
			updateQuery = updateQuery.Set(dto.ForDownloadSourcesColumnIsDownload, source.IsDownload)
		}

		updateQuery = updateQuery.Suffix(fmt.Sprintf("RETURNING %s", dto.ForDownloadSourcesColumnID))
	}

	return updateQuery
}
