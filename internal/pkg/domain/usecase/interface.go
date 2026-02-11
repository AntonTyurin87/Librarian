package usecase

import (
	"Librarian/internal/pkg/domain/entity"
	"context"
)

type Interface interface {
	GetAllRegions(ctx context.Context) (*entity.GetAllRegionsResponse, error)

	UploadTextSourceForDownload(ctx context.Context, req *entity.UploadTextSourceForDownloadRequest) (*entity.UploadTextSourceForDownloadResponse, error)
	GetURLForDownloadSource(ctx context.Context, req *entity.GetURLForDownloadSourceRequest) (*entity.GetURLForDownloadSourceResponse, error)
	//UploadFile(ctx context.Context, req *entity.UploadFileRequest) (*entity.UploadFileResponse, error)

	DownloadExternalSourceFile(ctx context.Context) error
}
