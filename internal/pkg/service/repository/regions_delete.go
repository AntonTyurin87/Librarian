package repository

import (
	"Librarian/internal/pkg/service/repository/dto"
	"Librarian/internal/pkg/service/repository/query/regions"
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

// DeleteRegions ...
func (r *repository) DeleteRegions(ctx context.Context, q regions.Delete) (dto.Regions, error) {
	var res dto.Regions

	if err := Selectx(ctx, r.storage, &res, deleteRegionsQuery(q)); err != nil {
		return nil, fmt.Errorf("selectx(ctx, r.storage, &res, deleteRegionsQuery(q)): %w", err)
	}

	return res, nil
}

func deleteRegionsQuery(query regions.Delete) sq.DeleteBuilder {
	deleteQuery := sq.StatementBuilder.Delete(dto.RegionsTableName).
		Prefix("--DeleteRegions\n")

	where := sq.Eq{}

	if len(query.IDs) > 0 {
		where[dto.RegionsColumnID] = query.IDs
	}
	if len(query.NamesRU) > 0 {
		where[dto.RegionsColumnNameRU] = query.NamesRU
	}

	deleteQuery = deleteQuery.Where(where)

	deleteQuery = deleteQuery.Suffix(fmt.Sprintf("RETURNING %s", dto.RegionsColumnID))

	return deleteQuery
}
