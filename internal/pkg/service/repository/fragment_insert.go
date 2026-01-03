package repository

import (
	"Librarian/internal/pkg/service/repository/dto"
	"Librarian/internal/pkg/service/repository/query/fragments"
	"context"
	"fmt"
	"strings"

	sq "github.com/Masterminds/squirrel"
)

// InsertFragment ...
func (r *repository) InsertFragment(ctx context.Context, q fragments.Insert) (dto.Fragments, error) {
	var res dto.Fragments

	if err := Selectx(ctx, r.storage, &res, insertFragmentQuery(q)); err != nil {
		return nil, fmt.Errorf("selectx(ctx, r.storage, &res, insertFragmentQuery(q)): %w", err)
	}

	return res, nil
}

func insertFragmentQuery(query fragments.Insert) sq.InsertBuilder {
	insertQuery := sq.StatementBuilder.Insert(dto.FragmentsTableName).
		Columns(
			dto.FragmentsColumnNameRu,
			dto.FragmentsColumnNameNative,
			dto.FragmentsColumnAuthorRu,
			dto.FragmentsColumnYear,
			dto.FragmentsColumnRegions,
			dto.FragmentsColumnTimePeriods,
			dto.FragmentsColumnDescription,
			dto.FragmentsColumnFormat,
		).
		Prefix("--InsertFragment\n")

	for _, fragment := range query.Fragments {
		a := dto.FragmentDtoFromEntity(fragment)
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

	insertQuery = insertQuery.Suffix(fmt.Sprintf("RETURNING %s", strings.Join(dto.FragmentsColumns, ", ")))

	return insertQuery
}
