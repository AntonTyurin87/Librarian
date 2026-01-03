package presenter

import (
	"Librarian/internal/pkg/domain/entity"

	lib "github.com/AntonTyurin87/Recon_Com_protoc/gen/go/librarian"
)

// ArticlesFromLibToEntity ...
func (p *presenter) ArticlesFromLibToEntity(article *lib.Article) *entity.Article {
	if article == nil {
		return nil
	}

	result := entity.Article{
		NameRu:      article.GetNameRu(),
		NameNative:  article.GetNameNative(),
		AuthorRu:    article.GetAuthorRu(),
		Year:        article.GetYear(),
		Regions:     article.GetRegions(),
		TimePeriods: article.GetTimePeriods(),
		Description: article.GetDescription(),
		Format:      p.FileFormatFromLibToEntity(article.GetFileFormat()),
	}
	return &result
}
