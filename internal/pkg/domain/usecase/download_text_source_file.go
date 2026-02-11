package usecase

import (
	"Librarian/internal/pkg/domain/entity"
	"Librarian/internal/pkg/domain/helpers"
	"Librarian/internal/pkg/service/repository/query/files"
	"Librarian/internal/pkg/service/repository/query/for_download_sources"
	"context"
	"fmt"
)

// DownloadExternalSourceFile ...
func (u *usecase) DownloadExternalSourceFile(ctx context.Context) error {
	// ищем данные об источниках с не скаченными файлами
	source, err := u.repository.SelectForDownloadSources(ctx, for_download_sources.Select{
		IsFileDownloads: []int64{0},
	})
	if err != nil {
		fmt.Printf("u.repository.SelectForDownloadSources: %v", err)
		return err //TODO вот тут надо логировать
	}

	if source == nil {
		return nil
	}

	// берём именно первый в списке источники и работаем с ним
	sourceForSave := source[0]

	// скачиваем первый из списка файл
	fileBytes, err := helpers.GetFileBytes(sourceForSave.DownloadURL)
	if err != nil {
		fmt.Printf("helpers.GetFileBytes: %v", err)
		return err //TODO вот тут надо логировать
	}

	// кладём файл в базу
	save, err := u.repository.InsertFiles(ctx, files.Insert{
		Files: []*entity.File{
			{
				TextSourceID: sourceForSave.ID,
				FileData:     fileBytes,
			},
		},
	})
	if err != nil || save == nil {
		fmt.Printf("u.repository.InsertFiles: %v", err)
		return err //TODO вот тут надо логировать
	}

	// если файл сохранили, то проставляем признак 1
	source[0].IsFileDownload = 1
	result, err := u.repository.UpdateForDownloadSources(ctx, for_download_sources.Update{
		ForDownloadSources: []*entity.ForDownloadSource{sourceForSave},
	})
	if err != nil || result == nil {
		fmt.Printf("u.repository.UpdateForDownloadSources: %v", err)
		return err //TODO вот тут надо логировать
	}

	return nil
}
