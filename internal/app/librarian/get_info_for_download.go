package librarian

import (
	"Librarian/internal/pkg/domain/entity"
	"context"
	"fmt"

	lib "github.com/AntonTyurin87/Recon_Com_protoc/gen/go/librarian"
)

// GetURLForDownloadSource - Отдаёт по ID источника ссылку для скачивания этого источника с диска
func (s *Server) GetURLForDownloadSource(ctx context.Context, req *lib.GetURLForDownloadSourceRequest) (*lib.GetURLForDownloadSourceResponse, error) {
	err := validateGetInfoForDownloadRequest(req)
	if err != nil {
		return nil, fmt.Errorf("validateGetInfoForDownloadRequest: %w", err)
	}

	info, err := s.usacase.GetURLForDownloadSource(ctx, &entity.GetURLForDownloadSourceRequest{
		TextSources: []*entity.TextSource{
			{
				ID: req.GetSourceId(),
			},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("s.usacase.GetURLForDownloadSource: %w", err)
	}

	return &lib.GetURLForDownloadSourceResponse{
		URLForDownload: info.GetURL(),
		FileInfo: &lib.FileInfo{
			Size:     info.GetFileInfo().Size,
			FileName: info.GetFileInfo().FileName,
			FileType: info.GetFileInfo().FileType,
		},
	}, nil
}

func validateGetInfoForDownloadRequest(req *lib.GetURLForDownloadSourceRequest) error {
	if req.GetSourceId() <= 0 {
		return fmt.Errorf("sourceId shuld not be 0")
	}
	return nil
}
