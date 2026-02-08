package repository

import (
	"Librarian/internal/pkg/service/repository/dto"
	"Librarian/internal/pkg/service/repository/query/regions"
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

// UpdateRegions ...
func (r *repository) UpdateRegions(ctx context.Context, q regions.Update) (dto.Regions, error) {
	var res dto.Regions

	if err := Selectx(ctx, r.storage, &res, updateRegionsQuery(q)); err != nil {
		return nil, fmt.Errorf("selectx(ctx, r.storage, &res, updateRegionsQuery(q)): %w", err)
	}

	return res, nil
}

func updateRegionsQuery(query regions.Update) sq.UpdateBuilder {
	updateQuery := sq.StatementBuilder.Update(dto.RegionsTableName).
		Prefix("--UpdateRegions\n")

	for _, region := range query.Regions {

		// что дополнять
		if region.ID != 0 {
			updateQuery = updateQuery.Where(sq.Eq{dto.RegionsColumnID: region.ID})
		}

		// чем дополнять
		if region.NameRU != "" {
			updateQuery = updateQuery.Set(dto.RegionsColumnNameRU, region.NameRU)
		}
		if region.Description != "" {
			updateQuery = updateQuery.Set(dto.RegionsColumnDescription, region.Description)
		}

		updateQuery = updateQuery.Suffix(fmt.Sprintf("RETURNING %s", dto.RegionsColumnID))
	}

	return updateQuery
}
