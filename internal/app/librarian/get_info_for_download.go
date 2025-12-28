package librarian

import (
	"Librarian/internal/pkg/domain/entity"
	"context"
	"fmt"

	lib "github.com/AntonTyurin87/Recon_Com_protoc/gen/go/librarian"
)

func (s *Server) GetInfoForDownload(ctx context.Context, req *lib.GetInfoForDownloadRequest) (*lib.GetInfoForDownloadResponse, error) {
	err := validateGetInfoForDownloadRequest(req)
	if err != nil {
		return nil, fmt.Errorf("validateGetInfoForDownloadRequest: %w", err)
	}

	info, err := s.usacase.GetInfoForDownload(ctx, &entity.GetInfoForDownloadRequest{
		Sources: []*entity.Source{
			{
				ID: req.GetSourceId(),
			},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("s.usacase.GetInfoForDownload: %w", err)
	}

	//TODO возможно тут отдавать через презентер
	return &lib.GetInfoForDownloadResponse{
		DownloadURL: info.GetURL(),
		FileInfo: &lib.FileInfo{
			Size:     info.GetFileInfo().GetSize(),
			MimeType: info.GetFileInfo().GetMimeType(),
		},
	}, nil
}

func validateGetInfoForDownloadRequest(req *lib.GetInfoForDownloadRequest) error {
	if req.GetSourceId() <= 0 {
		return fmt.Errorf("sourceId shuld not 0")
	}
	return nil
}
