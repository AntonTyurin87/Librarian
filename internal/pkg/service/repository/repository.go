package repository

import (
	"Librarian/internal/pkg/service/storage"
)

type repository struct {
	//ex      executor
	storage storage.Storage
}

func NewRepository(
	//ex executor,
	storage storage.Storage,
) Repository {
	return &repository{
		//ex:      ex,
		storage: storage,
	}
}
