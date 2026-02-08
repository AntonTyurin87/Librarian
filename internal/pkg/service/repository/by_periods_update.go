package repository

import (
	"Librarian/internal/pkg/service/repository/dto"
	"Librarian/internal/pkg/service/repository/query/by_periods"
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

// UpdateByPeriods ...
func (r *repository) UpdateByPeriods(ctx context.Context, q by_periods.Update) (dto.ByPeriods, error) {
	var res dto.ByPeriods

	if err := Selectx(ctx, r.storage, &res, updateByPeriodsQuery(q)); err != nil {
		return nil, fmt.Errorf("selectx(ctx, r.storage, &res, updateByPeriodsQuery(q)): %w", err)
	}

	return res, nil
}

func updateByPeriodsQuery(query by_periods.Update) sq.UpdateBuilder {
	updateQuery := sq.StatementBuilder.Update(dto.ByPeriodsTableName).
		Prefix("--UpdateByPeriods\n")

	for _, source := range query.ByPeriods {

		// что дополнять
		if source.ID != 0 {
			updateQuery = updateQuery.Where(sq.Eq{dto.ByPeriodsColumnID: source.ID})
		}
		if source.SourceID != 0 {
			updateQuery = updateQuery.Where(sq.Eq{dto.ByPeriodsColumnSourceID: source.SourceID})
		}

		// чем дополнять
		if source.PeriodID != 0 {
			updateQuery = updateQuery.Set(dto.ByPeriodsColumnPeriodID, source.PeriodID)
		}
		if source.Type != "" {
			updateQuery = updateQuery.Set(dto.ByPeriodsColumnType, source.Type)
		}
		if source.Pages != "" {
			updateQuery = updateQuery.Set(dto.ByPeriodsColumnPages, source.Pages)
		}

		updateQuery = updateQuery.Suffix(fmt.Sprintf("RETURNING %s", dto.ByPeriodsColumnID))
	}

	return updateQuery
}
