package repository

import (
	"Librarian/internal/pkg/service/repository/dto"
	"Librarian/internal/pkg/service/repository/query/for_download_sources"
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

// DeleteForDownloadSources ...
func (r *repository) DeleteForDownloadSources(ctx context.Context, q for_download_sources.Delete) (dto.ForDownloadSources, error) {
	var res dto.ForDownloadSources

	if err := Selectx(ctx, r.storage, &res, deleteForDownloadSourcesQuery(q)); err != nil {
		return nil, fmt.Errorf("selectx(ctx, r.storage, &res, deleteForDownloadSourcesQuery(q)): %w", err)
	}

	return res, nil
}

func deleteForDownloadSourcesQuery(query for_download_sources.Delete) sq.DeleteBuilder {
	deleteQuery := sq.StatementBuilder.Delete(dto.ForDownloadSourcesTableName).
		Prefix("--DeleteForDownloadSources\n")

	where := sq.Eq{}

	if len(query.IDs) > 0 {
		where[dto.ForDownloadSourcesColumnID] = query.IDs
	}
	if len(query.UserIDs) > 0 {
		where[dto.ForDownloadSourcesColumnUserID] = query.UserIDs
	}

	deleteQuery = deleteQuery.Where(where)

	deleteQuery = deleteQuery.Suffix(fmt.Sprintf("RETURNING %s", dto.ForDownloadSourcesColumnID))

	return deleteQuery
}
