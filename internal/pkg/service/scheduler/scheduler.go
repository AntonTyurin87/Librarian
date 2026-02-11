package scheduler

import (
	"Librarian/internal/pkg/domain/presenter"
	"Librarian/internal/pkg/domain/usecase"
)

type scheduler struct {
	presenter presenter.Interface

	usecase usecase.Interface
}

// NewUsacase ...
func NewScheduler(
	presenter presenter.Interface,

	usecase usecase.Interface,
) *scheduler {
	return &scheduler{
		presenter: presenter,

		usecase: usecase,
	}
}
