package librarian

import (
	"context"
	"fmt"

	lib "github.com/AntonTyurin87/Recon_Com_protoc/gen/go/librarian"
)

type Server struct {
	lib.UnimplementedLibrarianServer
}

func (s *Server) SendFile(ctx context.Context, in *lib.SendFileRequest) (*lib.SendFileResponse, error) {
	if in.GetText() == "file" {
		fmt.Println("Отправляем файл")
		return &lib.SendFileResponse{File: []byte("some file")}, nil
	}

	fmt.Println("Не получено нужное слово")
	return &lib.SendFileResponse{}, nil
}
