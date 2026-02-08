package repository

import (
	"Librarian/internal/pkg/domain/entity"
	"Librarian/internal/pkg/service/repository/dto"
	"Librarian/internal/pkg/service/repository/query/by_periods"
	"context"
	"fmt"
	"strings"

	sq "github.com/Masterminds/squirrel"
)

// InsertByPeriods ...
func (r *repository) InsertByPeriods(ctx context.Context, q by_periods.Insert) (*entity.ByPeriod, error) {
	var res dto.ByPeriods

	if err := Selectx(ctx, r.storage, &res, insertByPeriodsQuery(q)); err != nil {
		return nil, fmt.Errorf("selectx(ctx, r.storage, &res, insertByPeriodsQuery(q)): %w", err)
	}

	// сохраняем только один источник за раз
	source := res.Entity()[0]

	return source, nil
}

func insertByPeriodsQuery(query by_periods.Insert) sq.InsertBuilder {
	insertQuery := sq.StatementBuilder.Insert(dto.ByPeriodsTableName).
		Columns(
			dto.ByPeriodsColumnSourceID,
			dto.ByPeriodsColumnPeriodID,
			dto.ByPeriodsColumnType,
			dto.ByPeriodsColumnPages,
		).
		Prefix("--InsertByPeriods\n")

	for _, source := range query.ByPeriods {
		a := dto.ByPeriodDtoFromEntity(source)
		insertQuery = insertQuery.Values(
			a.GetSourceID(),
			a.GetPeriodID(),
			a.GetType(),
			a.GetPages(),
		)
	}

	insertQuery = insertQuery.Suffix(fmt.Sprintf("RETURNING %s", strings.Join(dto.ByPeriodsColumns, ", ")))

	return insertQuery
}
