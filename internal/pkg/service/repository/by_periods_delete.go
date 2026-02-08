package repository

import (
	"Librarian/internal/pkg/service/repository/dto"
	"Librarian/internal/pkg/service/repository/query/by_periods"
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

// DeleteByPeriods ...
func (r *repository) DeleteByPeriods(ctx context.Context, q by_periods.Delete) (dto.ByPeriods, error) {
	var res dto.ByPeriods

	if err := Selectx(ctx, r.storage, &res, deleteByPeriodsQuery(q)); err != nil {
		return nil, fmt.Errorf("selectx(ctx, r.storage, &res, deleteByPeriodsQuery(q)): %w", err)
	}

	return res, nil
}

func deleteByPeriodsQuery(query by_periods.Delete) sq.DeleteBuilder {
	deleteQuery := sq.StatementBuilder.Delete(dto.ByPeriodsTableName).
		Prefix("--DeleteByPeriods\n")

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

	deleteQuery = deleteQuery.Where(where)

	deleteQuery = deleteQuery.Suffix(fmt.Sprintf("RETURNING %s", dto.ByPeriodsColumnID))

	return deleteQuery
}
