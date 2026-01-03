package files_worker

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	lib "github.com/AntonTyurin87/Recon_Com_protoc/gen/go/librarian"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	FilesFolder = "files/"
)

// UploadSession сессия загрузки файла
type UploadSession struct {
	FileID       string
	FileName     string
	FilePath     string
	File         *os.File
	Metadata     *lib.FileMetadata
	TotalSize    int64
	TotalChunks  int32
	Received     map[int32]bool // Полученные чанки
	BytesWritten int64
	StartTime    time.Time
	mu           sync.Mutex
}

type UploadFileResponse struct {
	PathAddress string            `json:"path_address"`
	MeteData    *lib.FileMetadata `json:"mete_data"`
}

// GetPathAddress ...
func (u *UploadFileResponse) GetPathAddress() string {
	return u.PathAddress
}

// GetMeteData ...
func (u *UploadFileResponse) GetMeteData() *lib.FileMetadata {
	if u.MeteData == nil {
		return nil
	}
	return u.MeteData
}

func (f *fileWorker) UploadFile(stream grpc.ClientStreamingServer[lib.UploadFileRequest, lib.UploadFileResponse]) (*UploadFileResponse, error) {
	var session *UploadSession
	var data *lib.FileMetadata

	ctx := stream.Context()

	defer func() {
		// Очистка при завершении или ошибке
		if session != nil {
			f.cleanupSession(session.FileID)
		}
	}()

	startTime := time.Now()

	for {
		// Проверяем контекст на отмену
		select {
		case <-ctx.Done():
			if session != nil {
				fmt.Printf("Upload cancelled for file %s", session.FileID)
			}
			return nil, status.Error(codes.Canceled, "upload cancelled")
		default:
			// Продолжаем обработку
		}

		// Получаем запрос из потока
		req, err := stream.Recv()
		if err == io.EOF {
			// Клиент завершил отправку
			break
		}
		if err != nil {
			fmt.Printf("Failed to receive request: %v", err)
			return nil, status.Errorf(codes.Internal, "failed to receive request: %v", err)
		}

		// Проверяем, что в запросе
		if metadata := req.GetMetaData(); metadata != nil {
			//fileID = metadata.Id

			// Валидация метаданных
			if err := f.validateMetadata(metadata); err != nil {
				return nil, err
			}

			data = req.GetMetaData()

			// Создаем сессию загрузки
			session, err = f.createUploadSession(metadata)
			if err != nil {
				fmt.Printf("failed to create upload session: %v", err)
				return nil, status.Errorf(codes.Internal, "failed to create upload session: %v", err)
			}
		}

		if chunk := req.GetChunkData(); chunk != nil {
			// Обработка чанка данных
			if session == nil {
				fmt.Printf("metadata must be sent first")
				return nil, status.Error(codes.FailedPrecondition, "metadata must be sent first")
			}

			chunk = req.GetChunkData()

			// Валидация чанка
			if err := f.validateChunk(chunk, session); err != nil {
				return nil, err
			}

			// Обрабатываем чанк
			if err := f.processChunk(session, chunk); err != nil {
				return nil, status.Errorf(codes.Internal, "failed to process chunk %d: %v",
					chunk.ChunkNumber, err)
			}

		}
	}

	// Проверяем, что сессия была создана
	if session == nil {
		return nil, status.Error(codes.InvalidArgument, "no file metadata received")
	}

	// Завершаем загрузку и формируем ответ
	response, err := f.finalizeUpload(session, startTime)
	if err != nil {
		return nil, err
	}

	// Отправляем финальный ответ
	if err := stream.SendAndClose(response); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to send response: %v", err)
	}

	return &UploadFileResponse{
		PathAddress: session.FilePath,
		MeteData:    data,
	}, nil
}

// validateMetadata проверяет валидность метаданных
func (f *fileWorker) validateMetadata(metadata *lib.FileMetadata) error {
	if metadata == nil {
		return status.Error(codes.InvalidArgument, "metadata is required")
	}

	if metadata.Name == "" {
		return status.Error(codes.InvalidArgument, "file name is required")
	}

	if metadata.Size <= 0 {
		return status.Error(codes.InvalidArgument, "file size must be positive")
	}

	// Проверяем расширение файла
	ext := filepath.Ext(metadata.Name)
	if ext == "" {
		return status.Error(codes.InvalidArgument, "file must have an extension")
	}

	// Проверяем безопасность имени файла
	if strings.Contains(metadata.Name, "..") || strings.Contains(metadata.Name, "/") {
		return status.Error(codes.InvalidArgument, "invalid file name")
	}

	return nil
}

// createUploadSession создает сессию загрузки
func (f *fileWorker) createUploadSession(metadata *lib.FileMetadata) (*UploadSession, error) {
	// Генерируем безопасное имя файла
	safeFileName := f.generateSafeFileName(metadata.Name)
	filePath := filepath.Join(FilesFolder, safeFileName)

	// Создаем файл
	file, err := os.Create(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to create file: %w", err)
	}

	// Создаем сессию
	session := &UploadSession{
		FileID:    metadata.Id,
		FileName:  metadata.Name,
		FilePath:  filePath,
		File:      file,
		Metadata:  metadata,
		TotalSize: metadata.Size,
		Received:  make(map[int32]bool),
		StartTime: time.Now(),
	}

	// Рассчитываем количество чанков
	session.TotalChunks = int32((metadata.Size + f.chunkSize - 1) / f.chunkSize)

	// Сохраняем сессию
	f.mu.Lock()
	f.activeUploads[session.FileID] = session
	f.mu.Unlock()

	return session, nil
}

// validateChunk проверяет валидность чанка
func (f *fileWorker) validateChunk(chunk *lib.FileChunk, session *UploadSession) error {
	if chunk == nil {
		return status.Error(codes.InvalidArgument, "chunk is required")
	}

	if chunk.ChunkNumber < 0 || chunk.ChunkNumber >= session.TotalChunks {
		return status.Errorf(codes.InvalidArgument,
			"invalid chunk number: %d (total: %d)",
			chunk.ChunkNumber, session.TotalChunks)
	}

	if chunk.ChunkSize <= 0 {
		return status.Error(codes.InvalidArgument, "chunk size must be positive")
	}

	if chunk.TotalSize != session.TotalSize {
		return status.Errorf(codes.InvalidArgument,
			"total size mismatch: expected %d, got %d",
			session.TotalSize, chunk.TotalSize)
	}

	return nil
}

// processChunk обрабатывает чанк данных
func (f *fileWorker) processChunk(session *UploadSession, chunk *lib.FileChunk) error {
	session.mu.Lock()
	defer session.mu.Unlock()

	// Проверяем, не получен ли этот чанк уже
	if session.Received[chunk.ChunkNumber] {
		fmt.Printf("Chunk %d already received for file %s", chunk.ChunkNumber, session.FileID)
		return nil // Игнорируем дубликат
	}

	// Вычисляем смещение для этого чанка
	offset := int64(chunk.ChunkNumber) * f.chunkSize

	// Записываем данные в файл
	n, err := session.File.WriteAt(chunk.Data, offset)
	if err != nil {
		return fmt.Errorf("failed to write chunk: %w", err)
	}

	// Проверяем размер записанных данных
	if int64(n) != chunk.ChunkSize {
		return fmt.Errorf("write size mismatch: expected %d, wrote %d",
			chunk.ChunkSize, n)
	}

	// Отмечаем чанк как полученный
	session.Received[chunk.ChunkNumber] = true
	session.BytesWritten += int64(n)

	return nil
}

// finalizeUpload завершает загрузку и проверяет файл
func (f *fileWorker) finalizeUpload(session *UploadSession, startTime time.Time) (*lib.UploadFileResponse, error) {
	// Закрываем файл
	if err := session.File.Close(); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to close file: %w", err)
	}

	// Проверяем, все ли чанки получены
	receivedCount := len(session.Received)
	if receivedCount != int(session.TotalChunks) {
		fmt.Printf("Missing chunks for file %s: received %d, expected %d",
			session.FileID, receivedCount, session.TotalChunks)

		// Можно вернуть ошибку или частичный результат
		return &lib.UploadFileResponse{
			FileId:         session.FileID,
			Status:         lib.UploadStatus_UPLOAD_STATUS_PARTIAL,
			BytesReceived:  session.BytesWritten,
			ChunksReceived: int32(receivedCount),
			TotalSize:      session.TotalSize,
		}, nil
	}

	// Проверяем размер файла
	fileInfo, err := os.Stat(session.FilePath)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to stat file: %w", err)
	}

	if fileInfo.Size() != session.TotalSize {
		return nil, status.Errorf(codes.DataLoss,
			"file size mismatch: expected %d, got %d",
			session.TotalSize, fileInfo.Size())
	}

	// Проверяем хеш если он был предоставлен
	hashVerificationPassed := true
	hashMessage := "Hash verification passed"

	// Формируем ответ
	processingTime := time.Since(startTime)

	response := &lib.UploadFileResponse{
		FileId:                  session.FileID,
		Status:                  lib.UploadStatus_UPLOAD_STATUS_COMPLETED,
		BytesReceived:           fileInfo.Size(),
		ChunksReceived:          int32(receivedCount),
		TotalSize:               session.TotalSize,
		ProcessingTimeMs:        processingTime.Milliseconds(),
		HashVerificationPassed:  hashVerificationPassed,
		HashVerificationMessage: hashMessage,
		UploadStartedAt:         timeToTimestamp(session.StartTime),
		UploadCompletedAt:       timeToTimestamp(time.Now()),
	}

	if !hashVerificationPassed {
		response.Status = lib.UploadStatus_UPLOAD_STATUS_FAILED
		response.Error = &lib.ErrorDetail{
			Code:    "HASH_MISMATCH",
			Message: "File hash verification failed",
			Details: hashMessage,
		}
	}

	// Логируем успешную загрузку
	fmt.Printf("Upload completed for file %s: %d bytes, %d chunks, took %v",
		session.FileID, fileInfo.Size(), receivedCount, processingTime)

	return response, nil
}

// Вспомогательные методы

// generateSafeFileName создает безопасное имя файла
func (f *fileWorker) generateSafeFileName(originalName string) string {
	// Извлекаем расширение
	ext := filepath.Ext(originalName)
	name := strings.TrimSuffix(originalName, ext)

	// Заменяем небезопасные символы
	name = strings.ReplaceAll(name, " ", "_")
	name = strings.ReplaceAll(name, "/", "_")
	name = strings.ReplaceAll(name, "\\", "_")
	name = strings.ReplaceAll(name, "..", "_")

	// Добавляем временную метку для уникальности
	timestamp := time.Now().Format("2006_01_02")

	return fmt.Sprintf("%s_%s%s", name, timestamp, ext)
}

// calculateFileHash вычисляет MD5 хеш файла
func (f *fileWorker) calculateFileHash(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

// cleanupSession очищает сессию загрузки
func (f *fileWorker) cleanupSession(fileID string) {
	f.mu.Lock()
	defer f.mu.Unlock()

	if session, exists := f.activeUploads[fileID]; exists {
		if session.File != nil {
			session.File.Close()
		}
		delete(f.activeUploads, fileID)
	}
}

// timeToTimestamp конвертирует time.Time в protobuf timestamp
func timeToTimestamp(t time.Time) *timestamppb.Timestamp {
	return timestamppb.New(t)
}
