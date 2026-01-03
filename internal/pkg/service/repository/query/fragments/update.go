package fragments

import "Librarian/internal/pkg/domain/entity"

// Update ...
type Update struct {
	Fragments []*entity.Fragment
}

// GetFragments ...
func (s *Update) GetFragments() []*entity.Fragment {
	if s == nil {
		return nil
	}

	return s.Fragments
}
