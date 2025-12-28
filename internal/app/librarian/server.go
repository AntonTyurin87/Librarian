package librarian

import (
	"Librarian/internal/pkg/domain/presenter"
	"Librarian/internal/pkg/domain/usecase"
	"context"
	"fmt"

	lib "github.com/AntonTyurin87/Recon_Com_protoc/gen/go/librarian"
)

type Server struct {
	lib.UnimplementedLibrarianServer
	presenter presenter.Interface

	usacase usecase.Interface
}

func NewServer(
	presenter presenter.Interface,

	usacase usecase.Interface,
) *Server {
	return &Server{
		presenter: presenter,

		usacase: usacase,
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
