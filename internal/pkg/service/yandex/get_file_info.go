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

// GetFileInfo ...
func (y *YandexDiskClient) GetFileInfo(_ context.Context, request *entity.GetFileInfoRequest) (*entity.GetFileInfoResponse, error) {
	// URL encode путь
	encodedPath := url.PathEscape(request.Address)

	url := fmt.Sprintf("%s/resources?path=%s", y.BaseURL, encodedPath)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest: %w", err)
	}

	req.Header.Set("Authorization", "OAuth "+y.OAuthToken)

	resp, err := y.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("y.HTTPClient.Do: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error: %s, response: %s", resp.Status, string(body))
	}

	var fileInfo entity.FileInfo
	if err := json.NewDecoder(resp.Body).Decode(&fileInfo); err != nil {
		return nil, fmt.Errorf("json.NewDecoder: %w", err)
	}

	return &entity.GetFileInfoResponse{
		FileInfo: &fileInfo,
	}, nil
}
