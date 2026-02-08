package repository

import (
	"Librarian/internal/pkg/domain/entity"
	"Librarian/internal/pkg/service/repository/dto"
	"Librarian/internal/pkg/service/repository/query/text_sources"
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

// SelectTextSources ...
func (r *repository) SelectTextSources(ctx context.Context, q text_sources.Select) ([]*entity.TextSource, error) {
	var res dto.TextSources

	if err := Selectx(ctx, r.storage, &res, selectTextSourcesQuery(q)); err != nil {
		return nil, fmt.Errorf("selectx(ctx, r.storage, &res, selectTextSourcesQuery(q)): %w", err)
	}

	return res.Entity(), nil
}

func selectTextSourcesQuery(query text_sources.Select) sq.SelectBuilder {
	selectQuery := sq.StatementBuilder.Select(dto.TextSourcesColumns...).
		From(dto.TextSourcesTableName).
		Prefix("--SelectTextSources\n")

	where := sq.Eq{}

	if len(query.IDs) > 0 {
		where[dto.TextSourcesColumnID] = query.IDs
	}
	if len(query.UserIDs) > 0 {
		where[dto.TextSourcesColumnUserID] = query.UserIDs
	}
	if len(query.Types) > 0 {
		where[dto.TextSourcesColumnType] = query.Types
	}
	if len(query.NamesRU) > 0 {
		where[dto.TextSourcesColumnNameRU] = query.NamesRU
	}
	if len(query.NamesENG) > 0 {
		where[dto.TextSourcesColumnNameENG] = query.NamesENG
	}
	if len(query.AuthorRU) > 0 {
		where[dto.TextSourcesColumnAuthorRU] = query.AuthorRU
	}
	if len(query.Year) > 0 {
		where[dto.TextSourcesColumnYear] = query.Year
	}
	if len(query.FileFormats) > 0 {
		where[dto.TextSourcesColumnFileFormat] = query.FileFormats
	}
	if len(query.IsAvailable) > 0 {
		where[dto.TextSourcesColumnIsAvailable] = query.IsAvailable
	}

	//по полю CreateAt не ищем
	selectQuery = selectQuery.Where(where)
	selectQuery = selectQuery.OrderBy(dto.TextSourcesColumnID)

	return selectQuery
}
