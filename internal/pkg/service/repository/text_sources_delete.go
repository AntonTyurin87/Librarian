package repository

import (
	"Librarian/internal/pkg/service/repository/dto"
	"Librarian/internal/pkg/service/repository/query/text_sources"
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

// DeleteTextSources ...
func (r *repository) DeleteTextSources(ctx context.Context, q text_sources.Delete) (dto.TextSources, error) {
	var res dto.TextSources

	if err := Selectx(ctx, r.storage, &res, deleteTextSourcesQuery(q)); err != nil {
		return nil, fmt.Errorf("selectx(ctx, r.storage, &res, deleteTextSourcesQuery(q)): %w", err)
	}

	return res, nil
}

func deleteTextSourcesQuery(query text_sources.Delete) sq.DeleteBuilder {
	deleteQuery := sq.StatementBuilder.Delete(dto.TextSourcesTableName).
		Prefix("--DeleteTextSources\n")

	where := sq.Eq{}

	if len(query.IDs) > 0 {
		where[dto.TextSourcesColumnID] = query.IDs
	}
	if len(query.UserIDs) > 0 {
		where[dto.TextSourcesColumnUserID] = query.UserIDs
	}

	deleteQuery = deleteQuery.Where(where)

	deleteQuery = deleteQuery.Suffix(fmt.Sprintf("RETURNING %s", dto.TextSourcesColumnID))

	return deleteQuery
}
