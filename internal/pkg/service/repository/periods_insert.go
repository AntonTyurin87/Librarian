package repository

import (
	"Librarian/internal/pkg/domain/entity"
	"Librarian/internal/pkg/service/repository/dto"
	"Librarian/internal/pkg/service/repository/query/periods"
	"context"
	"fmt"
	"strings"

	sq "github.com/Masterminds/squirrel"
)

// InsertPeriods ...
func (r *repository) InsertPeriods(ctx context.Context, q periods.Insert) (*entity.Period, error) {
	var res dto.Periods

	if err := Selectx(ctx, r.storage, &res, insertPeriodsQuery(q)); err != nil {
		return nil, fmt.Errorf("selectx(ctx, r.storage, &res, insertPeriodsQuery(q)): %w", err)
	}

	// сохраняем только один период за раз
	period := res.Entity()[0]

	return period, nil
}

func insertPeriodsQuery(query periods.Insert) sq.InsertBuilder {
	insertQuery := sq.StatementBuilder.Insert(dto.PeriodsTableName).
		Columns(
			dto.PeriodsColumnCentury,
			dto.PeriodsColumnEra,
			dto.PeriodsColumnDescription,
		).
		Prefix("--InsertPeriods\n")

	for _, period := range query.Periods {
		a := dto.PeriodDtoFromEntity(period)
		insertQuery = insertQuery.Values(
			a.GetCentury(),
			a.GetEra(),
			a.GetDescription(),
		)
	}

	insertQuery = insertQuery.Suffix(fmt.Sprintf("RETURNING %s", strings.Join(dto.PeriodsColumns, ", ")))

	return insertQuery
}
