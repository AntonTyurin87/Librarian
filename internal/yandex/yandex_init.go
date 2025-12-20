package yandex

import (
	"net/http"
	"time"
)

var YandexClient *YandexDiskClient

const YandexToken = "y0__xD-5p6GAhjblgMg19PL4RSMPdmJcFPz1TX5UgtL0SDttD5EoA"

// YandexDiskClient представляет клиент для работы с Яндекс Диском
type YandexDiskClient struct {
	OAuthToken string
	BaseURL    string
	HTTPClient *http.Client
}

// NewYandexDiskClient создает новый клиент Яндекс Диска
func NewYandexDiskClient(token string) *YandexDiskClient {
	return &YandexDiskClient{
		OAuthToken: token,
		BaseURL:    "https://cloud-api.yandex.net/v1/disk",
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}
