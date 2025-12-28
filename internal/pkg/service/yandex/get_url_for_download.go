package yandex

import (
	"Librarian/internal/pkg/domain/entity"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// GetURLForDownload ...
func (y *YandexDiskClient) GetURLForDownload(_ context.Context, request *entity.GetURLForDownloadRequest) (*entity.GetURLForDownloadResponse, error) {
	encodedPath := url.QueryEscape(request.Address)

	fullURL := fmt.Sprintf("%s/resources/download?path=%s", y.BaseURL, encodedPath)

	downloadLinc := &entity.GetURLForDownloadResponse{}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return downloadLinc, fmt.Errorf("http.NewRequest: %w", err)
	}

	req.Header.Set("Authorization", "OAuth "+y.OAuthToken)

	resp, err := y.HTTPClient.Do(req)
	if err != nil {
		return downloadLinc, fmt.Errorf("y.HTTPClient.Do: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return downloadLinc, fmt.Errorf("API ошибка: %d", resp.StatusCode)
	}

	// Парсим ответ
	var result struct {
		Href string `json:"href"`
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return downloadLinc, fmt.Errorf("io.ReadAll: %w", err)
	}

	if err = json.Unmarshal(body, &result); err != nil {
		return downloadLinc, fmt.Errorf("json.Unmarshal: %w", err)
	}

	downloadLinc.URL = result.Href

	return downloadLinc, nil
}
