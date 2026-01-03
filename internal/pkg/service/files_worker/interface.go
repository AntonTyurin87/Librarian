package files_worker

import (
	lib "github.com/AntonTyurin87/Recon_Com_protoc/gen/go/librarian"
	"google.golang.org/grpc"
)

type FileWorker interface {
	UploadFile(stream grpc.ClientStreamingServer[lib.UploadFileRequest, lib.UploadFileResponse]) (*UploadFileResponse, error)
}
