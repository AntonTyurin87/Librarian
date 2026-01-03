package entity

// Photo ...
type Photo struct {
	ID              int32      `json:"id"`
	GroupID         int32      `json:"group_id"`
	NameRu          string     `json:"name_ru"`
	NameNative      string     `json:"name_native"`
	FilmingLocation string     `json:"filming_location"`
	Regions         []string   `json:"regions"`
	TimePeriods     []string   `json:"time_periods"`
	Description     string     `json:"description"`
	Format          FileFormat `json:"format"`
}

// GetID возвращает ID фотографии
func (p *Photo) GetID() int32 {
	if p == nil {
		return 0
	}
	return p.ID
}

// GetGroupID возвращает ID группы фотографий
func (p *Photo) GetGroupID() int32 {
	if p == nil {
		return 0
	}
	return p.GroupID
}

// GetNameRu возвращает название на русском
func (p *Photo) GetNameRu() string {
	if p == nil {
		return ""
	}
	return p.NameRu
}

// GetNameNative возвращает оригинальное название
func (p *Photo) GetNameNative() string {
	if p == nil {
		return ""
	}
	return p.NameNative
}

// GetFilmingLocation возвращает место съемки
func (p *Photo) GetFilmingLocation() string {
	if p == nil {
		return ""
	}
	return p.FilmingLocation
}

// GetRegions возвращает регионы
func (p *Photo) GetRegions() []string {
	if p == nil {
		return nil
	}
	return p.Regions
}

// GetTimePeriods возвращает временные периоды
func (p *Photo) GetTimePeriods() []string {
	if p == nil {
		return nil
	}
	return p.TimePeriods
}

// GetDescription возвращает описание
func (p *Photo) GetDescription() string {
	if p == nil {
		return ""
	}
	return p.Description
}

// GetFormat возвращает формат файла
func (p *Photo) GetFormat() FileFormat {
	if p == nil {
		return FileFormat_UNKNOWN
	}
	return p.Format
}
