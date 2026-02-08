package repository

import (
	"Librarian/internal/pkg/domain/entity"
	"Librarian/internal/pkg/service/repository/dto"
	"Librarian/internal/pkg/service/repository/query/files"
	"context"
	"fmt"
	"strings"

	sq "github.com/Masterminds/squirrel"
)

// InsertFiles ...
func (r *repository) InsertFiles(ctx context.Context, q files.Insert) (*entity.File, error) {
	var res dto.Files

	if err := Selectx(ctx, r.storage, &res, insertFilesQuery(q)); err != nil {
		return nil, fmt.Errorf("selectx(ctx, r.storage, &res, insertFilesQuery(q)): %w", err)
	}

	// сохраняем только один файл за раз
	file := res.Entity()[0]

	return file, nil
}

func insertFilesQuery(query files.Insert) sq.InsertBuilder {
	insertQuery := sq.StatementBuilder.Insert(dto.FilesTableName).
		Columns(
			dto.FilesColumnTextSourceID,
			dto.FilesColumnFileData,
		).
		Prefix("--InsertFiles\n")

	for _, file := range query.Files {
		a := dto.FileDtoFromEntity(file)
		insertQuery = insertQuery.Values(
			a.GetTextSourceID(),
			a.GetFileData(),
		)
	}

	insertQuery = insertQuery.Suffix(fmt.Sprintf("RETURNING %s", strings.Join(dto.FilesColumns, ", ")))

	return insertQuery
}
