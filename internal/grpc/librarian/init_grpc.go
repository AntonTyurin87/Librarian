package librarian

import (
	"log"
	"net"
	"os"
	"time"

	lib "github.com/AntonTyurin87/Recon_Com_protoc/gen/go/librarian"
	"google.golang.org/grpc"
)

func InitGRPC() {
	lis, err := net.Listen("tcp", os.Getenv("LIBRARIAN_PORT"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(grpc.ConnectionTimeout(10 * time.Second))
	lib.RegisterLibrarianServer(s, &Server{})

	log.Printf("Librarion listening on %s", os.Getenv("LIBRARIAN_PORT"))

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
