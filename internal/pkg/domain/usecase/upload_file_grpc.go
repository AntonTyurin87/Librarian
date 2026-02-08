package usecase

import (
	"Librarian/internal/pkg/domain/entity"
	"Librarian/internal/pkg/service/repository/query/articles"
	"Librarian/internal/pkg/service/repository/query/books"
	"Librarian/internal/pkg/service/repository/query/for_download_sources"
	"Librarian/internal/pkg/service/repository/query/fragments"
	"Librarian/internal/pkg/service/repository/query/photos"
	"Librarian/internal/pkg/service/yandex"
	"context"
	"fmt"
	"strings"
	"time"
)

// UploadFile ...
func (u *usecase) UploadFile(ctx context.Context, req *entity.UploadFileRequest) (*entity.UploadFileResponse, error) {
	// Сохраняем описание источника в зависимости от типа
	object, err := u.saveSource(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("u.saveSource: %w", err)
	}

	fileYandexName := u.getYandexFileName(req.FileAddress, object.SourceType)

	// Сохраняем файл на Яндекс диск
	err = u.yandexClient.UploadFile(ctx, req.FileAddress, fileYandexName)
	if err != nil {
		return nil, fmt.Errorf("u.yandexClient.UploadFile: %w", err)
	}

	// Сохраняем данных в БД источника
	_, err = u.repository.InsertSource(ctx, for_download_sources.Insert{
		Sources: []*entity.Source{
			{
				Type:         req.GetFileSourceData().GetSourceType(),
				ObjectID:     object.ObjectID,
				Address:      fileYandexName,
				CreatedAt:    time.Now(),
				Availability: true,
				TimePeriods:  object.TimePeriods,
			},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("u.repository.InsertSource: %w", err)
	}

	return nil, nil
}

func (u *usecase) saveSource(ctx context.Context, req *entity.UploadFileRequest) (*entity.UploadFileResponse, error) {
	var objectID int32
	var timePeriods []string
	// Сохраняем описание источника в зависимости от типа
	switch req.GetFileSourceData().GetSourceType() {

	case entity.SourceType_BOOK:
		b, err := u.repository.InsertBook(ctx, books.Insert{
			Boooks: []*entity.Book{req.GetFileSourceData().GetBook()},
		})
		if err != nil {
			return nil, fmt.Errorf("u.repository.InsertBook: %w", err)
		}
		objectID = b[0].GetID()
		timePeriods = b[0].GetTimePeriods()

	case entity.SourceType_ARTICLE:
		a, err := u.repository.InsertArticle(ctx, articles.Insert{
			Articles: []*entity.Article{req.GetFileSourceData().GetArticles()},
		})
		if err != nil {
			return nil, fmt.Errorf("u.repository.InsertArticle: %w", err)
		}
		objectID = a[0].GetID()
		timePeriods = a[0].GetTimePeriods()

	case entity.SourceType_FRAGMENT:
		f, err := u.repository.InsertFragment(ctx, fragments.Insert{
			Fragments: []*entity.Fragment{req.GetFileSourceData().GetFragment()},
		})
		if err != nil {
			return nil, fmt.Errorf("u.repository.InsertFragment: %w", err)
		}
		objectID = f[0].GetID()
		timePeriods = f[0].GetTimePeriods()

	case entity.SourceType_PHOTO:
		p, err := u.repository.InsertPhotos(ctx, photos.Insert{
			Photos: []*entity.Photo{req.GetFileSourceData().GetPhoto()},
		})
		if err != nil {
			return nil, fmt.Errorf("u.repository.InsertPhotos: %w", err)
		}
		objectID = p[0].GetID()
		timePeriods = p[0].GetTimePeriods()

	default:
		return nil, fmt.Errorf("invalid file source type: %s", req.GetFileSourceData().GetSourceType())
	}

	return &entity.UploadFileResponse{
		ObjectID:    objectID,
		SourceType:  req.GetFileSourceData().GetSourceType(),
		TimePeriods: timePeriods,
	}, nil
}

func (u *usecase) getYandexFileName(fileAddress string, sourceType entity.SourceType) string {
	fileName := strings.TrimPrefix(fileAddress, "files_worker.FilesFolder") //TODO тут должна быть папка для файла

	var fileFolder string

	switch sourceType {
	case entity.SourceType_BOOK:
		fileFolder = yandex.BooksFolder
	case entity.SourceType_ARTICLE:
		fileFolder = yandex.ArticlesFolder
	case entity.SourceType_FRAGMENT:
		fileFolder = yandex.FragmentsFolder
	case entity.SourceType_PHOTO:
		fileFolder = yandex.PhotosFolder
	}
	fileLocation := fmt.Sprintf("%s/%s/%s", yandex.SharedFolder, fileFolder, fileName)

	return fileLocation
}
