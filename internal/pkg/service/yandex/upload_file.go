package yandex

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

// UploadFile ... TODO сделать свои структроуры запроса и ответа
func (y *YandexDiskClient) UploadFile(ctx context.Context, localPath, remotePath string) error {
	// 1. Получаем URL для загрузки
	uploadURL, err := y.getUploadURL(ctx, remotePath, true)
	if err != nil {
		return fmt.Errorf("failed to get upload URL: %w", err)
	}

	// 2. Открываем файл
	file, err := os.Open(localPath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// 3. Загружаем файл
	req, err := http.NewRequestWithContext(ctx, "PUT", uploadURL, file)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Получаем размер файла
	fileInfo, _ := file.Stat()
	req.ContentLength = fileInfo.Size()
	req.Header.Set("Content-Type", "application/octet-stream")

	resp, err := y.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("upload request failed: %w", err)
	}
	defer resp.Body.Close()

	// 4. Проверяем ответ
	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("upload failed with status %d: %s", resp.StatusCode, string(body))
	}

	return nil
}

// getUploadURL получает URL для загрузки файла
func (y *YandexDiskClient) getUploadURL(ctx context.Context, remotePath string, overwrite bool) (string, error) {
	queryParams := url.Values{}
	queryParams.Add("path", remotePath)
	if overwrite {
		queryParams.Add("overwrite", "true")
	}

	urlStr := fmt.Sprintf("%s/resources/upload?%s", y.BaseURL, queryParams.Encode())

	req, err := http.NewRequestWithContext(ctx, "GET", urlStr, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "OAuth "+y.OAuthToken)
	//req.Header.Set("User-Agent", y.UserAgent)

	resp, err := y.HTTPClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("failed to get upload URL (status %d): %s", resp.StatusCode, string(body))
	}

	var result struct {
		Href      string `json:"href"`
		Method    string `json:"method"`
		Templated bool   `json:"templated"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	return result.Href, nil
}
