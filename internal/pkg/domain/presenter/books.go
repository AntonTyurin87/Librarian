package presenter

import (
	"Librarian/internal/pkg/domain/entity"

	lib "github.com/AntonTyurin87/Recon_Com_protoc/gen/go/librarian"
)

// BooksFromLibToEntity ...
func (p *presenter) BooksFromLibToEntity(book *lib.Book) *entity.Book {
	if book == nil {
		return nil
	}

	result := entity.Book{
		NameRu:      book.GetNameRu(),
		NameNative:  book.GetNameNative(),
		AuthorRu:    book.GetAuthorRu(),
		Year:        book.GetYear(),
		Regions:     book.GetRegions(),
		TimePeriods: book.GetTimePeriods(),
		Description: book.GetDescription(),
		Format:      p.FileFormatFromLibToEntity(book.GetFileFormat()),
	}
	return &result
}
