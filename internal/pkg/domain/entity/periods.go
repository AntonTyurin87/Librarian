package entity

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
