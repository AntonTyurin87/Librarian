package repository

import (
	"Librarian/internal/pkg/domain/entity"
	"Librarian/internal/pkg/service/repository/dto"
	"Librarian/internal/pkg/service/repository/query/text_sources"
	"context"
	"fmt"
	"strings"

	sq "github.com/Masterminds/squirrel"
)

// InsertTextSources ...
func (r *repository) InsertTextSources(ctx context.Context, q text_sources.Insert) (*entity.TextSource, error) {
	var res dto.TextSources

	if err := Selectx(ctx, r.storage, &res, insertTextSourcesQuery(q)); err != nil {
		return nil, fmt.Errorf("selectx(ctx, r.storage, &res, insertTextSourcesQuery(q)): %w", err)
	}

	// сохраняем только один источник за раз
	source := res.Entity()[0]

	return source, nil
}

func insertTextSourcesQuery(query text_sources.Insert) sq.InsertBuilder {
	insertQuery := sq.StatementBuilder.Insert(dto.TextSourcesTableName).
		Columns(
			dto.TextSourcesColumnUserID,
			dto.TextSourcesColumnType,
			dto.TextSourcesColumnNameRU,
			dto.TextSourcesColumnNameENG,
			dto.TextSourcesColumnAuthorRU,
			dto.TextSourcesColumnYear,
			dto.TextSourcesColumnDescription,
			dto.TextSourcesColumnPlaceURL,
			dto.TextSourcesColumnFromURL,
			dto.TextSourcesColumnFileName,
			dto.TextSourcesColumnFileFormat,
			dto.TextSourcesColumnCreatedAt,
		).
		Prefix("--InsertForDownloadSources\n")

	for _, source := range query.TextSources {
		a := dto.TextSourceDtoFromEntity(source)
		insertQuery = insertQuery.Values(
			a.GetUserID(),
			a.GetType(),
			a.GetNameRU(),
			a.GetNameENG(),
			a.GetAuthorRU(),
			a.GetYear(),
			a.GetDescription(),
			a.GetPlaceURL(),
			a.GetFromURL(),
			a.GetFileName(),
			a.GetFileFormat(),
			a.GetCreatedAt(),
		)
	}

	insertQuery = insertQuery.Suffix(fmt.Sprintf("RETURNING %s", strings.Join(dto.TextSourcesColumns, ", ")))

	return insertQuery
}
