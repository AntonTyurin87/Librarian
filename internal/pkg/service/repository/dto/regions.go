package dto

import "Librarian/internal/pkg/domain/entity"

// Название колонок и таблицы
const (
	RegionsTableName = "regions"

	RegionsColumnID          = "id"
	RegionsColumnNameRu      = "name_ru"
	RegionsColumnDescription = "description"
)

// RegionsColumns ...
var RegionsColumns = []string{
	RegionsColumnID,
	RegionsColumnNameRu,
	RegionsColumnDescription,
}

// Region ...
type Region struct {
	ID          int32  `db:"id"`
	NameRu      string `db:"name_ru"`
	Description string `db:"description"`
}

// GetID ...
func (r *Region) GetID() int32 {
	if r == nil {
		return 0
	}
	return r.ID
}

// GetNameRu ...
func (r *Region) GetNameRu() string {
	if r == nil {
		return ""
	}
	return r.NameRu
}

// GetDescription ...
func (r *Region) GetDescription() string {
	if r == nil {
		return ""
	}
	return r.Description
}

// Entity ...
func (r *Region) Entity() *entity.Region {
	if r == nil {
		return nil
	}

	return &entity.Region{
		ID:          r.GetID(),
		NameRu:      r.GetNameRu(),
		Description: r.GetDescription(),
	}
}

// Regions ...
type Regions []*Region

// Entity ...
func (r Regions) Entity() []*entity.Region { return ToEntitySlice[[]*entity.Region](r) }

/* TODO Возможно это не нужно
// RegionDtoFromEntity ...
func RegionDtoFromEntity(r *entity.Region) *Region {
	if r == nil {
		return nil
	}

	return &Region{
		ID:          r.GetID(),
		NameRu:      r.GetNameRu(),
		Description: r.GetDescription(),
	}
}
*/
