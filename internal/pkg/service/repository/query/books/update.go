package books

import "Librarian/internal/pkg/domain/entity"

// Update ...
type Update struct {
	Books []*entity.Book
}

// GetBooks ...
func (s *Update) GetBooks() []*entity.Book {
	if s == nil {
		return nil
	}

	return s.Books
}
