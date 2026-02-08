package repository

import (
	"Librarian/internal/pkg/domain/entity"
	"Librarian/internal/pkg/service/repository/dto"
	"Librarian/internal/pkg/service/repository/query/regions"
	"context"
	"fmt"
	"strings"

	sq "github.com/Masterminds/squirrel"
)

// InsertRegions ...
func (r *repository) InsertRegions(ctx context.Context, q regions.Insert) (*entity.Region, error) {
	var res dto.Regions

	if err := Selectx(ctx, r.storage, &res, insertRegionsQuery(q)); err != nil {
		return nil, fmt.Errorf("selectx(ctx, r.storage, &res, insertRegionsQuery(q)): %w", err)
	}

	// сохраняем только один регион за раз
	region := res.Entity()[0]

	return region, nil
}

func insertRegionsQuery(query regions.Insert) sq.InsertBuilder {
	insertQuery := sq.StatementBuilder.Insert(dto.RegionsTableName).
		Columns(
			dto.RegionsColumnNameRU,
			dto.RegionsColumnDescription,
		).
		Prefix("--InsertRegions\n")

	for _, region := range query.Regions {
		a := dto.RegionDtoFromEntity(region)
		insertQuery = insertQuery.Values(
			a.GetNameRU(),
			a.GetDescription(),
		)
	}

	insertQuery = insertQuery.Suffix(fmt.Sprintf("RETURNING %s", strings.Join(dto.RegionsColumns, ", ")))

	return insertQuery
}
