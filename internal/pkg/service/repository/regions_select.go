package repository

import (
	"Librarian/internal/pkg/domain/entity"
	"Librarian/internal/pkg/service/repository/dto"
	"Librarian/internal/pkg/service/repository/query/regions"
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

// SelectRegions ...
func (r *repository) SelectRegions(ctx context.Context, q regions.Select) ([]*entity.Region, error) {
	var res dto.Regions
	if err := Selectx(ctx, r.storage, &res, getSelectRegionsQuery(q)); err != nil {
		return nil, fmt.Errorf("selectx(ctx, r.storage, &res, getSelectRegionsQuery(q)): %w", err)
	}

	return res.Entity(), nil
}

func getSelectRegionsQuery(query regions.Select) sq.SelectBuilder {
	selectQuery := sq.StatementBuilder.Select(dto.RegionsColumns...).
		From(dto.RegionsTableName).
		Prefix("--SelectRegions\n")

	where := sq.Eq{}

	if len(query.IDs) > 0 {
		where[dto.RegionsColumnID] = query.IDs
	}
	if len(query.NamesRu) > 0 {
		where[dto.RegionsColumnNameRu] = query.NamesRu
	}
	if len(query.Descriptions) > 0 {
		where[dto.RegionsColumnDescription] = query.Descriptions
	}

	selectQuery = selectQuery.Where(where)
	selectQuery = selectQuery.OrderBy(dto.RegionsColumnNameRu)

	return selectQuery
}
