package repository

import (
	"Librarian/internal/pkg/service/repository/dto"
	"Librarian/internal/pkg/service/repository/query/files"
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

// DeleteFiles ...
func (r *repository) DeleteFiles(ctx context.Context, q files.Delete) (dto.Files, error) {
	var res dto.Files

	if err := Selectx(ctx, r.storage, &res, deleteFilesQuery(q)); err != nil {
		return nil, fmt.Errorf("selectx(ctx, r.storage, &res, deleteFilesQuery(q)): %w", err)
	}

	return res, nil
}

func deleteFilesQuery(query files.Delete) sq.DeleteBuilder {
	deleteQuery := sq.StatementBuilder.Delete(dto.FilesTableName).
		Prefix("--DeleteFiles\n")

	where := sq.Eq{}

	if len(query.IDs) > 0 {
		where[dto.FilesColumnID] = query.IDs
	}
	if len(query.TextSourceIDs) > 0 {
		where[dto.FilesColumnTextSourceID] = query.TextSourceIDs
	}

	deleteQuery = deleteQuery.Where(where)

	deleteQuery = deleteQuery.Suffix(fmt.Sprintf("RETURNING %s", dto.FilesColumnID))

	return deleteQuery
}
