package entity

// GetAllRegionsResponse ...
type GetAllRegionsResponse struct {
	Regions []*Region `json:"regions"`
}

func (g *GetAllRegionsResponse) GetRegions() []*Region {
	if g.Regions == nil {
		return nil
	}
	return g.Regions
}

// Region ...
type Region struct {
	ID          int64  `json:"id"`
	NameRU      string `json:"name_ru"`
	Description string `json:"description"`
}

// GetID ...
func (r *Region) GetID() int64 {
	if r == nil {
		return 0
	}
	return r.ID
}

// GetNameRU ...
func (r *Region) GetNameRU() string {
	if r == nil {
		return ""
	}
	return r.NameRU
}

// GetDescription ...
func (r *Region) GetDescription() string {
	if r == nil {
		return ""
	}
	return r.Description
}
