package repository

import (
	"Librarian/internal/pkg/service/repository/dto"
	"Librarian/internal/pkg/service/repository/query/by_regions"
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

// UpdateByRegions ...
func (r *repository) UpdateByRegions(ctx context.Context, q by_regions.Update) (dto.ByRegions, error) {
	var res dto.ByRegions

	if err := Selectx(ctx, r.storage, &res, updateByRegionsQuery(q)); err != nil {
		return nil, fmt.Errorf("selectx(ctx, r.storage, &res, updateByRegionsQuery(q)): %w", err)
	}

	return res, nil
}

func updateByRegionsQuery(query by_regions.Update) sq.UpdateBuilder {
	updateQuery := sq.StatementBuilder.Update(dto.ByRegionsTableName).
		Prefix("--UpdateByRegions\n")

	for _, source := range query.ByRegions {

		// что дополнять
		if source.ID != 0 {
			updateQuery = updateQuery.Where(sq.Eq{dto.ByRegionsColumnID: source.ID})
		}
		if source.SourceID != 0 {
			updateQuery = updateQuery.Where(sq.Eq{dto.ByRegionsColumnSourceID: source.SourceID})
		}

		// чем дополнять
		if source.RegionID != 0 {
			updateQuery = updateQuery.Set(dto.ByRegionsColumnRegionID, source.RegionID)
		}
		if source.Type != "" {
			updateQuery = updateQuery.Set(dto.ByRegionsColumnType, source.Type)
		}
		if source.Pages != "" {
			updateQuery = updateQuery.Set(dto.ByRegionsColumnPages, source.Pages)
		}

		updateQuery = updateQuery.Suffix(fmt.Sprintf("RETURNING %s", dto.ByRegionsColumnID))
	}

	return updateQuery
}
