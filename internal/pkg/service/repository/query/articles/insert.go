package articles

import "Librarian/internal/pkg/domain/entity"

// Insert ...
type Insert struct {
	Articles []*entity.Article
}
