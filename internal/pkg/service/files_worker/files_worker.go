package files_worker

import (
	"Librarian/internal/pkg/domain/presenter"
	"sync"
)

// fileWorker - специальная структура для сервиса обработки файлов
type fileWorker struct {
	presenter presenter.Interface

	uploadsDir    string
	chunkSize     int64 // Размер чанка в байтах
	activeUploads map[string]*UploadSession
	mu            sync.RWMutex
}

func NewFileWorker(
	presenter presenter.Interface,

	chunkSize int64, // Размер чанка в байтах
	activeUploads map[string]*UploadSession,
	mu sync.RWMutex,
) FileWorker {
	return &fileWorker{
		presenter: presenter,

		chunkSize:     chunkSize,
		activeUploads: activeUploads,
		mu:            mu,
	}
}
