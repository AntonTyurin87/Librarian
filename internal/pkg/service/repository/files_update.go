package repository

import (
	"Librarian/internal/pkg/service/repository/dto"
	"Librarian/internal/pkg/service/repository/query/files"
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

// UpdateFiles ...
func (r *repository) UpdateFiles(ctx context.Context, q files.Update) (dto.Files, error) {
	var res dto.Files

	if err := Selectx(ctx, r.storage, &res, updateFilesQuery(q)); err != nil {
		return nil, fmt.Errorf("selectx(ctx, r.storage, &res, updateFilesQuery(q)): %w", err)
	}

	return res, nil
}

func updateFilesQuery(query files.Update) sq.UpdateBuilder {
	updateQuery := sq.StatementBuilder.Update(dto.FilesTableName).
		Prefix("--UpdateFiles\n")

	for _, file := range query.Files {

		// что дополнять
		if file.ID != 0 {
			updateQuery = updateQuery.Where(sq.Eq{dto.FilesColumnID: file.ID})
		}
		if file.TextSourceID != 0 {
			updateQuery = updateQuery.Where(sq.Eq{dto.FilesColumnTextSourceID: file.TextSourceID})
		}

		// чем дополнять
		if file.FileData != nil {
			updateQuery = updateQuery.Set(dto.FilesColumnFileData, file.FileData)
		}

		updateQuery = updateQuery.Suffix(fmt.Sprintf("RETURNING %s", dto.FilesColumnID))
	}

	return updateQuery
}
