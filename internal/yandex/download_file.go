package yandex

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

// DownloadFile скачивает файл с Яндекс Диска в указанную локальную директорию
func (c *YandexDiskClient) DownloadFile(remotePath, localDir string) error {
	// Получаем ссылку для скачивания
	downloadLink, err := c.getDownloadLink(remotePath)
	if err != nil {
		return fmt.Errorf("ошибка получения ссылки для скачивания: %v", err)
	}

	fmt.Printf("🔗 Получена ссылка для скачивания\n")

	// Создаем HTTP запрос для скачивания
	req, err := http.NewRequest("GET", downloadLink, nil)
	if err != nil {
		return fmt.Errorf("ошибка создания запроса: %v", err)
	}

	// Выполняем запрос
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("ошибка скачивания файла: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("ошибка скачивания: статус %d", resp.StatusCode)
	}

	// Создаем локальную директорию если её нет
	if err := os.MkdirAll(localDir, 0755); err != nil {
		return fmt.Errorf("ошибка создания директории: %v", err)
	}

	// Извлекаем имя файла из пути
	fileName := filepath.Base(remotePath)
	localFilePath := filepath.Join(localDir, fileName)

	// Создаем файл для записи
	file, err := os.Create(localFilePath)
	if err != nil {
		return fmt.Errorf("ошибка создания файла: %v", err)
	}
	defer file.Close()

	// Копируем содержимое из ответа в файл
	written, err := io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("ошибка записи файла: %v", err)
	}

	fmt.Printf("✅ Файл успешно скачан: %s\n", localFilePath)
	fmt.Printf("📏 Размер: %d байт\n", written)

	return nil
}

// getDownloadLink получает временную ссылку для скачивания файла
func (c *YandexDiskClient) getDownloadLink(remotePath string) (string, error) {
	encodedPath := url.QueryEscape(remotePath)
	fullURL := fmt.Sprintf("%s/resources/download?path=%s", c.BaseURL, encodedPath)

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "OAuth "+c.OAuthToken)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API ошибка: %d", resp.StatusCode)
	}

	// Парсим ответ
	var result struct {
		Href string `json:"href"`
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}

	return result.Href, nil
}

// GetCurrentDir возвращает текущую директорию проекта
func GetCurrentDir() (string, error) {
	// Получаем текущую рабочую директорию
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return dir, nil
}
