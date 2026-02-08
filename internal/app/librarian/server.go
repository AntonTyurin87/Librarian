package librarian

import (
	"Librarian/internal/pkg/domain/presenter"
	"Librarian/internal/pkg/domain/usecase"

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
