package usecase

import (
	"Librarian/internal/pkg/domain/entity"
	"context"
	"fmt"
)

// GetInfoForDownload ...
func (u *usecase) GetInfoForDownload(ctx context.Context, req *entity.GetInfoForDownloadRequest) (*entity.GetInfoForDownloadResponse, error) {
	// поход адресом файла
	sourcesSelect := u.presenter.SourcesFromEntityToSourcies(req.GetSources())

	sources, err := u.repository.SelectSources(ctx, sourcesSelect)
	if err != nil {
		return nil, fmt.Errorf("u.repository.SelectSources: %w", err)
	}

	if len(sources) == 0 {
		return nil, fmt.Errorf("файл не найден в базе данных")
	}

	// адрес расположения файла на Яндекс диске
	address := sources[0].GetAddress()

	// поход за ссылкой
	downloadUrl, err := u.yandexClient.GetURLForDownload(ctx, &entity.GetURLForDownloadRequest{
		Address: address,
	})
	if err != nil {
		return nil, fmt.Errorf("u.yandexClient.GetURLForDownload: %w", err)
	}

	// поход за информацией о файле
	info, err := u.yandexClient.GetFileInfo(ctx, &entity.GetFileInfoRequest{Address: address})
	if err != nil {
		return nil, fmt.Errorf("u.yandexClient.GetFileInfo: %w", err)
	}

	return &entity.GetInfoForDownloadResponse{
		URL:      downloadUrl.URL,
		FileInfo: info.GetFileInfo(),
	}, nil
}
