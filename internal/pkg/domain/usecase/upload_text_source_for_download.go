package usecase

import (
	"Librarian/internal/pkg/domain/entity"
	"Librarian/internal/pkg/service/repository/query/for_download_sources"
	"context"
	"fmt"
)

// UploadTextSourceForDownload ...
func (u *usecase) UploadTextSourceForDownload(ctx context.Context, req *entity.UploadTextSourceForDownloadRequest) (*entity.UploadTextSourceForDownloadResponse, error) {
	if req == nil {
		return nil, nil //TODO надо как-то такое логировать
	}

	source, err := u.repository.InsertForDownloadSources(ctx, for_download_sources.Insert{
		ForDownloadSources: []*entity.ForDownloadSource{
			{
				UserID:      req.GetForDownloadSource().GetUserID(),
				Type:        req.GetForDownloadSource().GetType(),
				NameRU:      req.GetForDownloadSource().GetNameRU(),
				NameENG:     req.GetForDownloadSource().GetNameENG(),
				AuthorRU:    req.GetForDownloadSource().GetAuthorRU(),
				Year:        req.GetForDownloadSource().GetYear(),
				Description: req.GetForDownloadSource().GetDescription(),
				DownloadURL: req.GetForDownloadSource().GetDownloadURL(),
				CreatedAt:   req.GetForDownloadSource().GetCreatedAt(),
			},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("u.repository.InsertForDownloadSources: %w", err)
	}

	return &entity.UploadTextSourceForDownloadResponse{
		ID: source.GetID(),
	}, nil
}
