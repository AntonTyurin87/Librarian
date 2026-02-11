package helpers

import (
	"fmt"
	"io"
	"net/http"
)

// GetFileBytes получает файл в виде []byte
func GetFileBytes(fileURL string) ([]byte, error) {
	// Скачиваем файл
	resp, err := http.Get(fileURL)
	if err != nil {
		return nil, fmt.Errorf("ошибка скачивания: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ошибка сервера: %s", resp.Status)
	}

	// Читаем файл в []byte
	fileBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения файла: %w", err)
	}

	return fileBytes, nil
}
