package librarian

import (
	"Librarian/internal/pkg/domain/entity"
	"fmt"
	//"context"
	lib "github.com/AntonTyurin87/Recon_Com_protoc/gen/go/librarian"
	"google.golang.org/grpc"
)

// UploadFile ...
func (s *Server) UploadFile(stream grpc.ClientStreamingServer[lib.UploadFileRequest, lib.UploadFileResponse]) error {
	// сохраняем файл в папку
	metadata, err := s.fileWorker.UploadFile(stream)
	if err != nil {
		return fmt.Errorf("s.fileWorker.UploadFile: %w", err)
	}

	// сохраняем данные файла в БД и отправляем на Яндекс диск
	_, err = s.usacase.UploadFile(stream.Context(), &entity.UploadFileRequest{
		FileAddress: metadata.GetPathAddress(),
		FileSourceData: &entity.FileSourceData{
			SourceType: s.presenter.SourceTypeFromLibToEntity(metadata.GetMeteData().GetFileSourceData().GetSourceType()),
			Book:       s.presenter.BooksFromLibToEntity(metadata.GetMeteData().GetFileSourceData().GetBook()),
			Articles:   s.presenter.ArticlesFromLibToEntity(metadata.GetMeteData().GetFileSourceData().GetArticle()),
			Fragment:   s.presenter.FragmentsFromLibToEntity(metadata.GetMeteData().GetFileSourceData().GetFragment()),
			Photo:      s.presenter.PhotosFromLibToEntity(metadata.GetMeteData().GetFileSourceData().GetPhoto()),
		},
	})
	if err != nil {
		return fmt.Errorf("s.usacase.UploadFile: %w", err)
	}

	// удаляем файл
	err = s.presenter.DeleteFile(metadata.GetPathAddress())
	if err != nil {
		return fmt.Errorf("s.presenter.DeleteFile: %w", err)
	}

	return nil
}
