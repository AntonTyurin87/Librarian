package repository

import (
	"Librarian/internal/pkg/domain/entity"
	"Librarian/internal/pkg/service/repository/dto"
	"Librarian/internal/pkg/service/repository/query/sources"
	"context"
	"fmt"
	"strings"

	sq "github.com/Masterminds/squirrel"
)

// InsertSource ...
func (r *repository) InsertSource(ctx context.Context, q sources.Insert) (*entity.Source, error) {
	var res dto.Sources

	if err := Selectx(ctx, r.storage, &res, insertSourceQuery(q)); err != nil {
		return nil, fmt.Errorf("selectx(ctx, r.storage, &res, insertBookQuery(q)): %w", err)
	}

	// сохраняем только один источник за раз
	source := res.Entity()[0]

	return source, nil
}

func insertSourceQuery(query sources.Insert) sq.InsertBuilder {
	insertQuery := sq.StatementBuilder.Insert(dto.SourcesTableName).
		Columns(
			dto.SourcesColumnType,
			dto.SourcesColumnObjectID,
			dto.SourcesColumnAddress,
			dto.SourcesColumnCreatedAt,
			dto.SourcesColumnAvailability,
			dto.SourcesColumnTimePeriods,
		).
		Prefix("--InsertSource\n")

	for _, source := range query.Sources {
		a := dto.SourceDtoFromEntity(source)
		insertQuery = insertQuery.Values(
			a.GetType(),
			a.GetObjectID(),
			a.GetAddress(),
			a.GetCreatedAt(),
			a.GetAvailability(),
			a.GetTimePeriods(),
		)
	}

	insertQuery = insertQuery.Suffix(fmt.Sprintf("RETURNING %s", strings.Join(dto.SourcesColumns, ", ")))

	return insertQuery
}
