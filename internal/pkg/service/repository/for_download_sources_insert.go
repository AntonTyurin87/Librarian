package repository

import (
	"Librarian/internal/pkg/domain/entity"
	"Librarian/internal/pkg/service/repository/dto"
	"Librarian/internal/pkg/service/repository/query/for_download_sources"
	"context"
	"fmt"
	"strings"

	sq "github.com/Masterminds/squirrel"
)

// InsertForDownloadSources ...
func (r *repository) InsertForDownloadSources(ctx context.Context, q for_download_sources.Insert) (*entity.ForDownloadSource, error) {
	var res dto.ForDownloadSources

	if err := Selectx(ctx, r.storage, &res, insertForDownloadSourcesQuery(q)); err != nil {
		return nil, fmt.Errorf("selectx(ctx, r.storage, &res, insertForDownloadSourcesQuery(q)): %w", err)
	}

	// сохраняем только один источник за раз
	source := res.Entity()[0]

	return source, nil
}

func insertForDownloadSourcesQuery(query for_download_sources.Insert) sq.InsertBuilder {
	insertQuery := sq.StatementBuilder.Insert(dto.ForDownloadSourcesTableName).
		Columns(
			dto.ForDownloadSourcesColumnUserID,
			dto.ForDownloadSourcesColumnType,
			dto.ForDownloadSourcesColumnNameRU,
			dto.ForDownloadSourcesColumnNameENG,
			dto.ForDownloadSourcesColumnAuthorRU,
			dto.ForDownloadSourcesColumnYear,
			dto.ForDownloadSourcesColumnDescription,
			dto.ForDownloadSourcesColumnDownloadURL,
			dto.ForDownloadSourcesColumnCreatedAt,
		).
		Prefix("--InsertForDownloadSources\n")

	for _, source := range query.ForDownloadSources {
		a := dto.ForDownloadSourceDtoFromEntity(source)
		insertQuery = insertQuery.Values(
			a.GetUserID(),
			a.GetType(),
			a.GetNameRU(),
			a.GetNameENG(),
			a.GetAuthorRU(),
			a.GetYear(),
			a.GetDescription(),
			a.GetDownloadURL(),
			a.GetCreatedAt(),
		)
	}

	insertQuery = insertQuery.Suffix(fmt.Sprintf("RETURNING %s", strings.Join(dto.ForDownloadSourcesColumns, ", ")))

	return insertQuery
}
