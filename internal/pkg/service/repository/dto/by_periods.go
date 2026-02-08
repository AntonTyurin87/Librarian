package dto

import (
	"Librarian/internal/pkg/domain/entity"
)

const (
	ByPeriodsTableName = "by_periods"

	ByPeriodsColumnID       = "id"
	ByPeriodsColumnSourceID = "source_id"
	ByPeriodsColumnPeriodID = "period_id"
	ByPeriodsColumnType     = "type"
	ByPeriodsColumnPages    = "pages"
)

var ByPeriodsColumns = []string{
	ByPeriodsColumnID,
	ByPeriodsColumnSourceID,
	ByPeriodsColumnPeriodID,
	ByPeriodsColumnType,
	ByPeriodsColumnPages,
}

// ByPeriod ...
type ByPeriod struct {
	ID       int64  `json:"id"`
	SourceID int64  `json:"source_id"`
	PeriodID int64  `json:"period_id"`
	Type     string `json:"type"`
	Pages    string `json:"pages"`
}

// GetID ...
func (bp *ByPeriod) GetID() int64 {
	if bp == nil {
		return 0
	}
	return bp.ID
}

// GetSourceID ...
func (bp *ByPeriod) GetSourceID() int64 {
	if bp == nil {
		return 0
	}
	return bp.SourceID
}

// GetPeriodID ...
func (bp *ByPeriod) GetPeriodID() int64 {
	if bp == nil {
		return 0
	}
	return bp.PeriodID
}

// GetType ...
func (bp *ByPeriod) GetType() string {
	if bp == nil {
		return ""
	}
	return bp.Type
}

// GetPages ...
func (bp *ByPeriod) GetPages() string {
	if bp == nil {
		return ""
	}
	return bp.Pages
}

// Entity ...
func (bp *ByPeriod) Entity() *entity.ByPeriod {
	if bp == nil {
		return nil
	}

	return &entity.ByPeriod{
		ID:       bp.GetID(),
		SourceID: bp.GetSourceID(),
		PeriodID: bp.GetPeriodID(),
		Type:     entity.SourceType(bp.GetType()),
		Pages:    bp.GetPages(),
	}
}

// ByPeriods ...
type ByPeriods []*ByPeriod

// Entity ...
func (bp ByPeriods) Entity() []*entity.ByPeriod {
	return ToEntitySlice[[]*entity.ByPeriod](bp)
}

// ByPeriodDtoFromEntity ...
func ByPeriodDtoFromEntity(e *entity.ByPeriod) *ByPeriod {
	if e == nil {
		return nil
	}
	return &ByPeriod{
		ID:       e.GetID(),
		SourceID: e.GetSourceID(),
		PeriodID: e.GetPeriodID(),
		Type:     string(e.GetType()),
		Pages:    e.GetPages(),
	}
}
