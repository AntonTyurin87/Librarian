package by_periods

import "Librarian/internal/pkg/domain/entity"

// Update ...
type Update struct {
	ByPeriods []*entity.ByPeriod
}

// GetByPeriods ...
func (s *Update) GetByPeriods() []*entity.ByPeriod {
	if s == nil {
		return nil
	}

	return s.ByPeriods
}
