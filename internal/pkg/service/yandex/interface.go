package yandex

import (
	"Librarian/internal/pkg/domain/entity"
	"context"
)

// Interface ...
type Interface interface {
	GetURLForDownload(_ context.Context, request *entity.GetURLForDownloadRequest) (*entity.GetURLForDownloadResponse, error)
	GetFileInfo(_ context.Context, request *entity.GetFileInfoRequest) (*entity.GetFileInfoResponse, error)

	UploadFile(ctx context.Context, localPath, remotePath string) error
}
