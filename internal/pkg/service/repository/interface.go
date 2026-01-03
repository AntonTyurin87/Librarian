package repository

import (
	"Librarian/internal/pkg/domain/entity"
	"Librarian/internal/pkg/service/repository/dto"
	"Librarian/internal/pkg/service/repository/query/articles"
	"Librarian/internal/pkg/service/repository/query/books"
	"Librarian/internal/pkg/service/repository/query/fragments"
	"Librarian/internal/pkg/service/repository/query/photos"
	"Librarian/internal/pkg/service/repository/query/regions"
	"Librarian/internal/pkg/service/repository/query/sources"
	"context"
)

type Repository interface {
	SelectSources(ctx context.Context, q sources.Select) ([]*entity.Source, error)
	SelectRegions(ctx context.Context, q regions.Select) ([]*entity.Region, error)

	InsertSource(ctx context.Context, q sources.Insert) (*entity.Source, error)
	InsertBook(ctx context.Context, q books.Insert) (dto.Books, error)
	InsertArticle(ctx context.Context, q articles.Insert) (dto.Articles, error)
	InsertFragment(ctx context.Context, q fragments.Insert) (dto.Fragments, error)
	InsertPhotos(ctx context.Context, q photos.Insert) (dto.Photos, error)
}
