package entity

// ByPeriod ...
type ByPeriod struct {
	ID       int64      `json:"id"`
	SourceID int64      `json:"source_id"`
	PeriodID int64      `json:"period_id"`
	Type     SourceType `json:"type"`
	Pages    string     `json:"pages"`
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
func (bp *ByPeriod) GetType() SourceType {
	if bp == nil {
		return SourceTypeUnknown
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
