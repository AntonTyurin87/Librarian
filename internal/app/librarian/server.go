package librarian

import (
	"Librarian/internal/pkg/domain/presenter"
	"Librarian/internal/pkg/domain/usecase"
	"Librarian/internal/pkg/service/files_worker"
	"context"
	"fmt"

	lib "github.com/AntonTyurin87/Recon_Com_protoc/gen/go/librarian"
)

type Server struct {
	lib.UnimplementedLibrarianServer
	presenter presenter.Interface

	usacase    usecase.Interface
	fileWorker files_worker.FileWorker
}

func NewServer(
	presenter presenter.Interface,

	usacase usecase.Interface,
	fileWorker files_worker.FileWorker,
) *Server {
	return &Server{
		presenter: presenter,

		usacase:    usacase,
		fileWorker: fileWorker,
	}
}

func (s *Server) SendFile(ctx context.Context, in *lib.SendFileRequest) (*lib.SendFileResponse, error) {
	if in.GetText() == "file" {
		fmt.Println("Отправляем файл")
		//err := s.presenter.SomeFunc()
		return &lib.SendFileResponse{File: []byte("some file")}, nil
	}

	fmt.Println("Не получено нужное слово")
	return &lib.SendFileResponse{}, nil
}
