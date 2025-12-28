package repository

import (
	"Librarian/internal/pkg/domain/entity"
	"Librarian/internal/pkg/service/repository/query/regions"
	"Librarian/internal/pkg/service/repository/query/sources"
	"context"
)

type Repository interface {
	SelectRegions(ctx context.Context, q regions.Select) ([]*entity.Region, error)
	SelectSources(ctx context.Context, q sources.Select) ([]*entity.Source, error)
}
