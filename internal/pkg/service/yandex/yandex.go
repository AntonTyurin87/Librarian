package yandex

import (
	"net/http"
	"time"
)

// YandexDiskClient представляет клиент для работы с Яндекс Диском
type YandexDiskClient struct {
	OAuthToken string
	BaseURL    string
	HTTPClient *http.Client
}

// NewYandexDiskClient создает новый клиент Яндекс Диска
func NewYandexDiskClient(token string) Interface {
	return &YandexDiskClient{
		OAuthToken: token,
		BaseURL:    "https://cloud-api.yandex.net/v1/disk",
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}
