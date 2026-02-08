package repository

import (
	"Librarian/internal/pkg/service/repository/dto"
	"Librarian/internal/pkg/service/repository/query/periods"
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

// UpdatePeriods ...
func (r *repository) UpdatePeriods(ctx context.Context, q periods.Update) (dto.Periods, error) {
	var res dto.Periods

	if err := Selectx(ctx, r.storage, &res, updatePeriodsQuery(q)); err != nil {
		return nil, fmt.Errorf("selectx(ctx, r.storage, &res, updatePeriodsQuery(q)): %w", err)
	}

	return res, nil
}

func updatePeriodsQuery(query periods.Update) sq.UpdateBuilder {
	updateQuery := sq.StatementBuilder.Update(dto.PeriodsTableName).
		Prefix("--UpdatePeriods\n")

	for _, period := range query.Periods {

		// что дополнять
		if period.ID != 0 {
			updateQuery = updateQuery.Where(sq.Eq{dto.PeriodsColumnID: period.ID})
		}

		// чем дополнять
		if period.Century != 0 {
			updateQuery = updateQuery.Set(dto.PeriodsColumnCentury, period.Century)
		}
		if period.Era != "" {
			updateQuery = updateQuery.Set(dto.PeriodsColumnEra, period.Era)
		}
		if period.Description != "" {
			updateQuery = updateQuery.Set(dto.PeriodsColumnDescription, period.Description)
		}

		updateQuery = updateQuery.Suffix(fmt.Sprintf("RETURNING %s", dto.PeriodsColumnID))
	}

	return updateQuery
}
