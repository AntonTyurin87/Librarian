package dto

import (
	"Librarian/internal/pkg/domain/entity"
)

const (
	PeriodsTableName = "periods"

	PeriodsColumnID          = "id"
	PeriodsColumnCentury     = "century"
	PeriodsColumnEra         = "era"
	PeriodsColumnDescription = "description"
)

var PeriodsColumns = []string{
	PeriodsColumnID,
	PeriodsColumnCentury,
	PeriodsColumnEra,
	PeriodsColumnDescription,
}

// Period ...
type Period struct {
	ID          int64  `json:"id"`
	Century     int64  `json:"century"`
	Era         string `json:"era"`
	Description string `json:"description"`
}

// GetID ...
func (p *Period) GetID() int64 {
	if p == nil {
		return 0
	}
	return p.ID
}

// GetCentury ...
func (p *Period) GetCentury() int64 {
	if p == nil {
		return 0
	}
	return p.Century
}

// GetEra ...
func (p *Period) GetEra() string {
	if p == nil {
		return ""
	}
	return p.Era
}

// GetDescription ...
func (p *Period) GetDescription() string {
	if p == nil {
		return ""
	}
	return p.Description
}

// Entity ...
func (p *Period) Entity() *entity.Period {
	if p == nil {
		return nil
	}

	return &entity.Period{
		ID:          p.GetID(),
		Century:     p.GetCentury(),
		Era:         p.GetEra(),
		Description: p.GetDescription(),
	}
}

// Periods ...
type Periods []*Period

// Entity ...
func (p Periods) Entity() []*entity.Period {
	return ToEntitySlice[[]*entity.Period](p)
}

// PeriodDtoFromEntity ...
func PeriodDtoFromEntity(e *entity.Period) *Period {
	if e == nil {
		return nil
	}
	return &Period{
		ID:          e.GetID(),
		Century:     e.GetCentury(),
		Era:         e.GetEra(),
		Description: e.GetDescription(),
	}
}
