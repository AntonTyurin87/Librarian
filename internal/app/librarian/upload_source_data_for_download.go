package librarian

import (
	"Librarian/internal/pkg/domain/entity"
	"context"
	"fmt"

	lib "github.com/AntonTyurin87/Recon_Com_protoc/gen/go/librarian"
)

// UploadSourceDataForDownload принимает данные для последующей загрузки файла источника на Яндекс диск
func (s *Server) UploadSourceDataForDownload(ctx context.Context, req *lib.UploadSourceDataForDownloadRequest) (*lib.UploadSourceDataForDownloadResponse, error) {
	err := validateUploadSourceDataForDownloadRequest(req)
	if err != nil {
		return nil, fmt.Errorf("validateUploadSourceDataForDownloadRequest: %w", err)
	}

	result, err := s.usacase.UploadTextSourceForDownload(ctx, &entity.UploadTextSourceForDownloadRequest{
		ForDownloadSource: &entity.ForDownloadSource{
			UserID:      req.GetUserID(),
			Type:        s.presenter.SourceTypeFromLibToEntity(req.GetSourceType()),
			NameRU:      req.GetNameRu(),
			NameENG:     req.GetNameEng(),
			AuthorRU:    req.GetAuthorRu(),
			Year:        req.GetYear(),
			Description: req.GetDescription(),
			DownloadURL: req.GetDownloadUrl(),
			//CreatedAt:   req.GetCreatedAt(), //TODO
		},
	})
	if err != nil {
		return nil, fmt.Errorf("s.usacase.UploadTextSourceForDownload: %w", err)
	}

	return &lib.UploadSourceDataForDownloadResponse{
		Id:      result.ID,
		IsSaved: true,
	}, nil
}

func validateUploadSourceDataForDownloadRequest(req *lib.UploadSourceDataForDownloadRequest) error {
	//TODO написать проверку
	return nil
}
