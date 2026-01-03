package usecase

import (
	"Librarian/internal/pkg/domain/presenter"
	"Librarian/internal/pkg/service/repository"
	"Librarian/internal/pkg/service/yandex"
)

type usecase struct {
	presenter presenter.Interface

	repository   repository.Repository
	yandexClient yandex.Interface
}

// NewUsacase ...
func NewUsacase(
	presenter presenter.Interface,

	repository repository.Repository,
	yandexClient yandex.Interface,
) *usecase {
	return &usecase{
		presenter:    presenter,
		repository:   repository,
		yandexClient: yandexClient,
	}
}
