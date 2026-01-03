package repository

import (
	"Librarian/internal/pkg/service/repository/dto"
	"Librarian/internal/pkg/service/repository/query/articles"
	"context"
	"fmt"
	"strings"

	sq "github.com/Masterminds/squirrel"
)

// InsertArticle ...
func (r *repository) InsertArticle(ctx context.Context, q articles.Insert) (dto.Articles, error) {
	var res dto.Articles

	if err := Selectx(ctx, r.storage, &res, insertArticleQuery(q)); err != nil {
		return nil, fmt.Errorf("selectx(ctx, r.storage, &res, insertArticleQuery(q)): %w", err)
	}

	return res, nil
}

func insertArticleQuery(query articles.Insert) sq.InsertBuilder {
	insertQuery := sq.StatementBuilder.Insert(dto.ArticlesTableName).
		Columns(
			dto.ArticlesColumnNameRu,
			dto.ArticlesColumnNameNative,
			dto.ArticlesColumnAuthorRu,
			dto.ArticlesColumnYear,
			dto.ArticlesColumnRegions,
			dto.ArticlesColumnTimePeriods,
			dto.ArticlesColumnDescription,
			dto.ArticlesColumnFormat,
		).
		Prefix("--InsertArticle\n")

	for _, article := range query.Articles {
		a := dto.ArticleDtoFromEntity(article)
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

	insertQuery = insertQuery.Suffix(fmt.Sprintf("RETURNING %s", strings.Join(dto.ArticlesColumns, ", ")))

	return insertQuery
}
