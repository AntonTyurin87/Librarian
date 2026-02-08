package periods

import "Librarian/internal/pkg/domain/entity"

// Update ...
type Update struct {
	Periods []*entity.Period
}

// GetPeriods ...
func (s *Update) GetPeriods() []*entity.Period {
	if s == nil {
		return nil
	}

	return s.Periods
}
