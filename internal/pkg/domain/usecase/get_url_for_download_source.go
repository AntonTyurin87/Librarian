package usecase

import (
	"Librarian/internal/pkg/domain/entity"
	"context"
	"fmt"
)

// GetURLForDownloadSource ...
func (u *usecase) GetURLForDownloadSource(ctx context.Context, req *entity.GetURLForDownloadSourceRequest) (*entity.GetURLForDownloadSourceResponse, error) {
	sources, err := u.repository.SelectTextSources(ctx, u.presenter.SourcesFromEntityToTextSources(req.GetSources()))
	if err != nil {
		return nil, fmt.Errorf("u.repository.SelectTextSources: %w", err)
	}
	if len(sources) == 0 {
		return nil, fmt.Errorf("файл не найден в базе данных")
	}

	// адрес расположения файла на Яндекс диске
	address := sources[0].GetPlaceURL()

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

	info.GetFileInfo().FileName = sources[0].GetFileName()
	info.GetFileInfo().FileType = sources[0].GetFileFormat()

	return &entity.GetURLForDownloadSourceResponse{
		URL:      downloadUrl.URL,
		FileInfo: info.GetFileInfo(),
	}, nil
}
