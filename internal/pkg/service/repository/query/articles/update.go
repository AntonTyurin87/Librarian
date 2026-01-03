package articles

import "Librarian/internal/pkg/domain/entity"

// Update ...
type Update struct {
	Articles []*entity.Article
}

// GetArticles ...
func (s *Update) GetArticles() []*entity.Article {
	if s == nil {
		return nil
	}

	return s.Articles
}
