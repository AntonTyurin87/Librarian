package usecase

import (
	"Librarian/internal/pkg/domain/entity"
	"context"
)

type Interface interface {
	GetAllRegions(ctx context.Context) (*entity.GetAllRegionsResponse, error)
	GetInfoForDownload(ctx context.Context, req *entity.GetInfoForDownloadRequest) (*entity.GetInfoForDownloadResponse, error)
}
