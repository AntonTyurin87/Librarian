package presenter

import (
	"Librarian/internal/pkg/domain/entity"
	"Librarian/internal/pkg/service/repository/query/sources"

	lib "github.com/AntonTyurin87/Recon_Com_protoc/gen/go/librarian"
)

type Interface interface {
	RegionsFromEntityToLib(regionsEntity []*entity.Region) []*lib.Region

	SourcesFromEntityToSourcies(sourcesEntity []*entity.Source) sources.Select
	SourceTypeFromLibToEntity(source lib.SourceType) entity.SourceType

	BooksFromLibToEntity(book *lib.Book) *entity.Book
	ArticlesFromLibToEntity(article *lib.Article) *entity.Article
	FragmentsFromLibToEntity(fragment *lib.Fragment) *entity.Fragment
	PhotosFromLibToEntity(photo *lib.Photo) *entity.Photo

	FileFormatFromLibToEntity(format lib.FileFormat) entity.FileFormat

	DeleteFile(address string) error
}
