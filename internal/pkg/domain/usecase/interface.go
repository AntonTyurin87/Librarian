package usecase

import (
	"Librarian/internal/pkg/domain/entity"
	"context"
)

type Interface interface {
	GetAllRegions(ctx context.Context) (*entity.GetAllRegionsResponse, error)

	UploadTextSourceForDownload(ctx context.Context, req *entity.UploadTextSourceForDownloadRequest) (*entity.UploadTextSourceForDownloadResponse, error)
	//GetInfoForDownload(ctx context.Context, req *entity.GetInfoForDownloadRequest) (*entity.GetInfoForDownloadResponse, error)
	//UploadFile(ctx context.Context, req *entity.UploadFileRequest) (*entity.UploadFileResponse, error)
}
