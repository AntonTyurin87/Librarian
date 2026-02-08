package repository

import (
	"Librarian/internal/pkg/domain/entity"
	"Librarian/internal/pkg/service/repository/dto"
	"Librarian/internal/pkg/service/repository/query/files"
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

// SelectFiles ...
func (r *repository) SelectFiles(ctx context.Context, q files.Select) ([]*entity.File, error) {
	var res dto.Files

	if err := Selectx(ctx, r.storage, &res, selectFilesQuery(q)); err != nil {
		return nil, fmt.Errorf("selectx(ctx, r.storage, &res, selectFilesQuery(q)): %w", err)
	}

	return res.Entity(), nil
}

func selectFilesQuery(query files.Select) sq.SelectBuilder {
	selectQuery := sq.StatementBuilder.Select(dto.FilesColumns...).
		From(dto.FilesTableName).
		Prefix("--SelectFiles\n")

	where := sq.Eq{}

	if len(query.IDs) > 0 {
		where[dto.FilesColumnID] = query.IDs
	}
	if len(query.TextSourceIDs) > 0 {
		where[dto.FilesColumnTextSourceID] = query.TextSourceIDs
	}

	selectQuery = selectQuery.Where(where)
	selectQuery = selectQuery.OrderBy(dto.FilesColumnID)

	return selectQuery
}
