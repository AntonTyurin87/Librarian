package repository

import (
	"Librarian/internal/pkg/service/repository/dto"
	"Librarian/internal/pkg/service/repository/query/periods"
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

// DeletePeriods ...
func (r *repository) DeletePeriods(ctx context.Context, q periods.Delete) (dto.Periods, error) {
	var res dto.Periods

	if err := Selectx(ctx, r.storage, &res, deletePeriodsQuery(q)); err != nil {
		return nil, fmt.Errorf("selectx(ctx, r.storage, &res, deletePeriodsQuery(q)): %w", err)
	}

	return res, nil
}

func deletePeriodsQuery(query periods.Delete) sq.DeleteBuilder {
	deleteQuery := sq.StatementBuilder.Delete(dto.PeriodsTableName).
		Prefix("--DeletePeriods\n")

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

	deleteQuery = deleteQuery.Where(where)

	deleteQuery = deleteQuery.Suffix(fmt.Sprintf("RETURNING %s", dto.PeriodsColumnID))

	return deleteQuery
}
