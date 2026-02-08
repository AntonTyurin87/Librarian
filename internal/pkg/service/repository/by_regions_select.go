package repository

import (
	"Librarian/internal/pkg/domain/entity"
	"Librarian/internal/pkg/service/repository/dto"
	"Librarian/internal/pkg/service/repository/query/by_regions"
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

// SelectByRegions ...
func (r *repository) SelectByRegions(ctx context.Context, q by_regions.Select) ([]*entity.ByRegion, error) {
	var res dto.ByRegions

	if err := Selectx(ctx, r.storage, &res, selectByRegionsQuery(q)); err != nil {
		return nil, fmt.Errorf("selectx(ctx, r.storage, &res, selectByRegionsQuery(q)): %w", err)
	}

	return res.Entity(), nil
}

func selectByRegionsQuery(query by_regions.Select) sq.SelectBuilder {
	selectQuery := sq.StatementBuilder.Select(dto.ByRegionsColumns...).
		From(dto.ByRegionsTableName).
		Prefix("--SelectByRegions\n")

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
	if len(query.Types) > 0 {
		where[dto.ByRegionsColumnType] = query.Types
	}

	//по полю CreateAt не ищем
	selectQuery = selectQuery.Where(where)
	selectQuery = selectQuery.OrderBy(dto.ByRegionsColumnID)

	return selectQuery
}
