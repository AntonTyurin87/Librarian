package main

import (
	"log"
	"net"
	"time"

	lib "github.com/AntonTyurin87/Recon_Com_protoc/gen/go/librarian"
	"google.golang.org/grpc"

	"Librarian/internal/grpc/librarian"
)

const (
	port = "0.0.0.0:50052"
)

func main() { //TODO камент...

	// gRPC сервис подключаем
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(grpc.ConnectionTimeout(10 * time.Second))
	lib.RegisterLibrarianServer(s, &librarian.Server{})

	log.Printf("Service A (TextProcessor) listening on %s", port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	///////////////////////////////////////////////////
}
