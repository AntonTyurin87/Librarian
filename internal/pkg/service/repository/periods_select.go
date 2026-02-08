package repository

import (
	"Librarian/internal/pkg/domain/entity"
	"Librarian/internal/pkg/service/repository/dto"
	"Librarian/internal/pkg/service/repository/query/periods"
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

// SelectPeriods ...
func (r *repository) SelectPeriods(ctx context.Context, q periods.Select) ([]*entity.Period, error) {
	var res dto.Periods

	if err := Selectx(ctx, r.storage, &res, selectPeriodsQuery(q)); err != nil {
		return nil, fmt.Errorf("selectx(ctx, r.storage, &res, selectPeriodsQuery(q)): %w", err)
	}

	return res.Entity(), nil
}

func selectPeriodsQuery(query periods.Select) sq.SelectBuilder {
	selectQuery := sq.StatementBuilder.Select(dto.PeriodsColumns...).
		From(dto.PeriodsTableName).
		Prefix("--SelectPeriods\n")

	where := sq.Eq{}

	if len(query.IDs) > 0 {
		where[dto.PeriodsColumnID] = query.IDs
	}
	if len(query.Centuries) > 0 {
		where[dto.PeriodsColumnCentury] = query.Centuries
	}
	if len(query.Eras) > 0 {
		where[dto.PeriodsColumnEra] = query.Eras
	}

	selectQuery = selectQuery.Where(where)
	selectQuery = selectQuery.OrderBy(dto.PeriodsColumnID)

	return selectQuery
}
