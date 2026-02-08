package repository

import (
	"Librarian/internal/pkg/service/repository/dto"
	"Librarian/internal/pkg/service/repository/query/by_regions"
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

// DeleteByRegions ...
func (r *repository) DeleteByRegions(ctx context.Context, q by_regions.Delete) (dto.ByRegions, error) {
	var res dto.ByRegions

	if err := Selectx(ctx, r.storage, &res, deleteByRegionsQuery(q)); err != nil {
		return nil, fmt.Errorf("selectx(ctx, r.storage, &res, deleteByRegionsQuery(q)): %w", err)
	}

	return res, nil
}

func deleteByRegionsQuery(query by_regions.Delete) sq.DeleteBuilder {
	deleteQuery := sq.StatementBuilder.Delete(dto.ByRegionsTableName).
		Prefix("--DeleteByRegions\n")

	where := sq.Eq{}

	if len(query.IDs) > 0 {
		where[dto.ByRegionsColumnID] = query.IDs
	}
	if len(query.SourceIDs) > 0 {
		where[dto.ByRegionsColumnSourceID] = query.SourceIDs
	}
	if len(query.RegionIDs) > 0 {
		where[dto.ByRegionsColumnRegionID] = query.RegionIDs
	}

	deleteQuery = deleteQuery.Where(where)

	deleteQuery = deleteQuery.Suffix(fmt.Sprintf("RETURNING %s", dto.ByRegionsColumnID))

	return deleteQuery
}
