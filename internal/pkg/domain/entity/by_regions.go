package entity

// ByRegion ...
type ByRegion struct {
	ID       int64      `json:"id"`
	SourceID int64      `json:"source_id"`
	RegionID int64      `json:"region_id"`
	Type     SourceType `json:"type"`
	Pages    string     `json:"pages"`
}

// GetID ...
func (br *ByRegion) GetID() int64 {
	if br == nil {
		return 0
	}
	return br.ID
}

// GetSourceID ...
func (br *ByRegion) GetSourceID() int64 {
	if br == nil {
		return 0
	}
	return br.SourceID
}

// GetRegionID ...
func (br *ByRegion) GetRegionID() int64 {
	if br == nil {
		return 0
	}
	return br.RegionID
}

// GetType ...
func (br *ByRegion) GetType() SourceType {
	if br == nil {
		return SourceTypeUnknown
	}
	return br.Type
}

// GetPages ...
func (br *ByRegion) GetPages() string {
	if br == nil {
		return ""
	}
	return br.Pages
}
