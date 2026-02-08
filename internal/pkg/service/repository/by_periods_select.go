package repository

import (
	"Librarian/internal/pkg/domain/entity"
	"Librarian/internal/pkg/service/repository/dto"
	"Librarian/internal/pkg/service/repository/query/by_periods"
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

// SelectByPeriods ...
func (r *repository) SelectByPeriods(ctx context.Context, q by_periods.Select) ([]*entity.ByPeriod, error) {
	var res dto.ByPeriods

	if err := Selectx(ctx, r.storage, &res, selectByPeriodsQuery(q)); err != nil {
		return nil, fmt.Errorf("selectx(ctx, r.storage, &res, selectByPeriodsQuery(q)): %w", err)
	}

	return res.Entity(), nil
}

func selectByPeriodsQuery(query by_periods.Select) sq.SelectBuilder {
	selectQuery := sq.StatementBuilder.Select(dto.ByPeriodsColumns...).
		From(dto.ByPeriodsTableName).
		Prefix("--SelectByPeriods\n")

	where := sq.Eq{}

	if len(query.IDs) > 0 {
		where[dto.ByPeriodsColumnID] = query.IDs
	}
	if len(query.SourceIDs) > 0 {
		where[dto.ByPeriodsColumnSourceID] = query.SourceIDs
	}
	if len(query.PeriodIDs) > 0 {
		where[dto.ByPeriodsColumnPeriodID] = query.PeriodIDs
	}
	if len(query.Types) > 0 {
		where[dto.ByPeriodsColumnType] = query.Types
	}

	//по полю CreateAt не ищем
	selectQuery = selectQuery.Where(where)
	selectQuery = selectQuery.OrderBy(dto.ByPeriodsColumnID)

	return selectQuery
}
