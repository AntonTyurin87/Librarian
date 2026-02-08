package repository

import (
	"Librarian/internal/pkg/domain/entity"
	"Librarian/internal/pkg/service/repository/dto"
	"Librarian/internal/pkg/service/repository/query/for_download_sources"
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

// SelectForDownloadSources ...
func (r *repository) SelectForDownloadSources(ctx context.Context, q for_download_sources.Select) ([]*entity.ForDownloadSource, error) {
	var res dto.ForDownloadSources

	if err := Selectx(ctx, r.storage, &res, selectForDownloadSourcesQuery(q)); err != nil {
		return nil, fmt.Errorf("selectx(ctx, r.storage, &res, selectForDownloadSourcesQuery(q)): %w", err)
	}

	return res.Entity(), nil
}

func selectForDownloadSourcesQuery(query for_download_sources.Select) sq.SelectBuilder {
	selectQuery := sq.StatementBuilder.Select(dto.ForDownloadSourcesColumns...).
		From(dto.ForDownloadSourcesTableName).
		Prefix("--SelectForDownloadSources\n")

	where := sq.Eq{}

	if len(query.IDs) > 0 {
		where[dto.ForDownloadSourcesColumnID] = query.IDs
	}
	if len(query.UserIDs) > 0 {
		where[dto.ForDownloadSourcesColumnUserID] = query.UserIDs
	}
	if len(query.Types) > 0 {
		where[dto.ForDownloadSourcesColumnType] = query.Types
	}
	if len(query.NamesRU) > 0 {
		where[dto.ForDownloadSourcesColumnNameRU] = query.NamesRU
	}
	if len(query.NamesENG) > 0 {
		where[dto.ForDownloadSourcesColumnNameENG] = query.NamesENG
	}
	if len(query.AuthorRU) > 0 {
		where[dto.ForDownloadSourcesColumnAuthorRU] = query.AuthorRU
	}
	if len(query.Year) > 0 {
		where[dto.ForDownloadSourcesColumnYear] = query.Year
	}
	if len(query.IsFileDownloads) > 0 {
		where[dto.ForDownloadSourcesColumnIsFileDownload] = query.IsFileDownloads
	}
	if len(query.IsDownloads) > 0 {
		where[dto.ForDownloadSourcesColumnIsDownload] = query.IsDownloads
	}

	//по полю CreateAt не ищем
	selectQuery = selectQuery.Where(where)
	selectQuery = selectQuery.OrderBy(dto.ForDownloadSourcesColumnID)

	return selectQuery
}
