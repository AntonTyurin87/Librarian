package repository

import (
	"Librarian/internal/pkg/domain/entity"
	"Librarian/internal/pkg/service/repository/dto"
	"Librarian/internal/pkg/service/repository/query/by_regions"
	"context"
	"fmt"
	"strings"

	sq "github.com/Masterminds/squirrel"
)

// InsertByRegions ...
func (r *repository) InsertByRegions(ctx context.Context, q by_regions.Insert) (*entity.ByRegion, error) {
	var res dto.ByRegions

	if err := Selectx(ctx, r.storage, &res, insertByRegionsQuery(q)); err != nil {
		return nil, fmt.Errorf("selectx(ctx, r.storage, &res, insertByRegionsQuery(q)): %w", err)
	}

	// сохраняем только один источник за раз
	source := res.Entity()[0]

	return source, nil
}

func insertByRegionsQuery(query by_regions.Insert) sq.InsertBuilder {
	insertQuery := sq.StatementBuilder.Insert(dto.ByRegionsTableName).
		Columns(
			dto.ByRegionsColumnSourceID,
			dto.ByRegionsColumnRegionID,
			dto.ByRegionsColumnType,
			dto.ByRegionsColumnPages,
		).
		Prefix("--InsertByRegions\n")

	for _, source := range query.ByRegions {
		a := dto.ByRegionDtoFromEntity(source)
		insertQuery = insertQuery.Values(
			a.GetSourceID(),
			a.GetRegionID(),
			a.GetType(),
			a.GetPages(),
		)
	}

	insertQuery = insertQuery.Suffix(fmt.Sprintf("RETURNING %s", strings.Join(dto.ByRegionsColumns, ", ")))

	return insertQuery
}
