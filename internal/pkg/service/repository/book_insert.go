package repository

import (
	"Librarian/internal/pkg/service/repository/dto"
	"Librarian/internal/pkg/service/repository/query/books"
	"context"
	"fmt"
	"strings"

	sq "github.com/Masterminds/squirrel"
)

// InsertBook ...
func (r *repository) InsertBook(ctx context.Context, q books.Insert) (dto.Books, error) {
	var res dto.Books

	if err := Selectx(ctx, r.storage, &res, insertBookQuery(q)); err != nil {
		return nil, fmt.Errorf("selectx(ctx, r.storage, &res, insertBookQuery(q)): %w", err)
	}

	return res, nil
}

func insertBookQuery(query books.Insert) sq.InsertBuilder {
	insertQuery := sq.StatementBuilder.Insert(dto.BooksTableName).
		Columns(
			dto.BooksColumnNameRu,
			dto.BooksColumnNameNative,
			dto.BooksColumnAuthorRu,
			dto.BooksColumnYear,
			dto.BooksColumnRegions,
			dto.BooksColumnTimePeriods,
			dto.BooksColumnDescription,
			dto.BooksColumnFormat,
		).
		Prefix("--InsertBook\n")

	for _, book := range query.Boooks {
		a := dto.BookDtoFromEntity(book)
		insertQuery = insertQuery.Values(
			a.GetNameRu(),
			a.GetNameNative(),
			a.GetAuthorRu(),
			a.GetYear(),
			strings.Join(a.GetRegions(), ", "),
			strings.Join(a.GetTimePeriods(), ", "),
			a.GetDescription(),
			a.GetFormat(),
		)
	}

	insertQuery = insertQuery.Suffix(fmt.Sprintf("RETURNING %s", strings.Join(dto.BooksColumns, ", ")))

	return insertQuery
}
