package repository

import (
	"Librarian/internal/pkg/domain/entity"
	"Librarian/internal/pkg/service/repository/dto"
	"Librarian/internal/pkg/service/repository/query/sources"
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

// SelectSources ...
func (r *repository) SelectSources(ctx context.Context, q sources.Select) ([]*entity.Source, error) {
	var res dto.Sources
	if err := Selectx(ctx, r.storage, &res, getSelectSourcesQuery(q)); err != nil {
		return nil, fmt.Errorf("selectx(ctx, r.storage, &res, getSelectSourcesQuery(q)): %w", err)
	}

	return res.Entity(), nil
}

func getSelectSourcesQuery(query sources.Select) sq.SelectBuilder {
	selectQuery := sq.StatementBuilder.Select(dto.SourcesColumns...).
		From(dto.SourcesTableName).
		Prefix("--SelectSources\n")

	where := sq.Eq{}

	if len(query.IDs) > 0 {
		where[dto.SourcesColumnID] = query.IDs
	}
	if len(query.ObjectIDs) > 0 {
		where[dto.SourcesColumnObjectID] = query.ObjectIDs
	}
	if len(query.Types) > 0 {
		where[dto.SourcesColumnType] = query.Types
	}
	if len(query.Address) > 0 {
		where[dto.SourcesColumnAddress] = query.Address
	}

	//по полю CreateAt не ищем

	selectQuery = selectQuery.Where(where)
	selectQuery = selectQuery.OrderBy(dto.SourcesColumnID)

	return selectQuery
}
