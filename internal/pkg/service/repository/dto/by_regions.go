package dto

import (
	"Librarian/internal/pkg/domain/entity"
)

const (
	ByRegionsTableName = "by_regions"

	ByRegionsColumnID       = "id"
	ByRegionsColumnSourceID = "source_id"
	ByRegionsColumnRegionID = "region_id"
	ByRegionsColumnType     = "type"
	ByRegionsColumnPages    = "pages"
)

var ByRegionsColumns = []string{
	ByRegionsColumnID,
	ByRegionsColumnSourceID,
	ByRegionsColumnRegionID,
	ByRegionsColumnType,
	ByRegionsColumnPages,
}

// ByRegion ...
type ByRegion struct {
	ID       int64  `json:"id"`
	SourceID int64  `json:"source_id"`
	RegionID int64  `json:"region_id"`
	Type     string `json:"type"`
	Pages    string `json:"pages"`
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
func (br *ByRegion) GetType() string {
	if br == nil {
		return ""
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

// Entity ...
func (br *ByRegion) Entity() *entity.ByRegion {
	if br == nil {
		return nil
	}

	return &entity.ByRegion{
		ID:       br.GetID(),
		SourceID: br.GetSourceID(),
		RegionID: br.GetRegionID(),
		Type:     entity.SourceType(br.GetType()),
		Pages:    br.GetPages(),
	}
}

// ByRegions ...
type ByRegions []*ByRegion

// Entity ...
func (br ByRegions) Entity() []*entity.ByRegion {
	return ToEntitySlice[[]*entity.ByRegion](br)
}

// ByRegionDtoFromEntity ...
func ByRegionDtoFromEntity(e *entity.ByRegion) *ByRegion {
	if e == nil {
		return nil
	}
	return &ByRegion{
		ID:       e.GetID(),
		SourceID: e.GetSourceID(),
		RegionID: e.GetRegionID(),
		Type:     string(e.GetType()),
		Pages:    e.GetPages(),
	}
}
