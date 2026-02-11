package repository

import (
	"Librarian/internal/pkg/domain/entity"
	"Librarian/internal/pkg/service/repository/dto"
	"Librarian/internal/pkg/service/repository/query/files"
	"Librarian/internal/pkg/service/repository/query/for_download_sources"
	"Librarian/internal/pkg/service/repository/query/regions"
	"Librarian/internal/pkg/service/repository/query/text_sources"
	"context"
)

type Repository interface {
	InsertForDownloadSources(ctx context.Context, q for_download_sources.Insert) (*entity.ForDownloadSource, error)
	SelectForDownloadSources(ctx context.Context, q for_download_sources.Select) ([]*entity.ForDownloadSource, error)
	UpdateForDownloadSources(ctx context.Context, q for_download_sources.Update) (dto.ForDownloadSources, error)

	InsertTextSources(ctx context.Context, q text_sources.Insert) (*entity.TextSource, error)
	SelectTextSources(ctx context.Context, q text_sources.Select) ([]*entity.TextSource, error)

	SelectRegions(ctx context.Context, q regions.Select) ([]*entity.Region, error)

	InsertFiles(ctx context.Context, q files.Insert) (*entity.File, error)
}
