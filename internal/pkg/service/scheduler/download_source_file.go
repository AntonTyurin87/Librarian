package scheduler

import (
	"context"
	"fmt"
	"time"
)

// DownloadSourceFile ...
func (s *scheduler) DownloadSourceFile(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	fmt.Printf("Фоновая задача DownloadSourceFile запущена (каждые %v)\n", interval)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Фоновая задача DownloadSourceFile остановлена по контексту")
			return
		case <-ticker.C:
			go s.executeAsyncTask(ctx) // Асинхронное выполнение
		}
	}
}

func (s *scheduler) executeAsyncTask(ctx context.Context) {
	// Выполняем в отдельной горутине
	fmt.Printf("[Async Task] %v\n", time.Now())

	err := s.usecase.DownloadExternalSourceFile(ctx)
	if err != nil {
		fmt.Printf("s.usecase.DownloadExternalSourceFile: %v", err)
	}
	return
	// Ваша бизнес-логика
	// Например:
	// - синхронизация с внешним API
	// - обработка очереди сообщений
	// - генерация отчетов
	// - очистка временных файлов
}
