package repository

import (
	"Librarian/internal/pkg/service/storage"
)

type repository struct {
	storage storage.Storage
}

func NewRepository(
	storage storage.Storage,
) Repository {
	return &repository{
		storage: storage,
	}
}
