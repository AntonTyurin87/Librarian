package dto

import (
	"Librarian/internal/pkg/domain/entity"
	"time"
)

// Название колонок и таблицы
const (
	SourcesTableName = "sources"

	SourcesColumnID           = "id"
	SourcesColumnType         = "type"
	SourcesColumnObjectID     = "object_id"
	SourcesColumnAddress      = "address"
	SourcesColumnCreatedAt    = "created_at"
	SourcesColumnAvailability = "availability"
	SourcesColumnTimePeriods  = "time_periods"
)

// SourcesColumns ...
var SourcesColumns = []string{
	SourcesColumnID,
	SourcesColumnType,
	SourcesColumnObjectID,
	SourcesColumnAddress,
	SourcesColumnCreatedAt,
	SourcesColumnAvailability,
	SourcesColumnTimePeriods,
}

type Source struct {
	ID           int32     `db:"id"`
	Type         int32     `db:"type"`
	ObjectID     int32     `db:"object_id"`
	Address      string    `db:"address"`
	CreatedAt    time.Time `db:"created_at"`
	Availability bool      `db:"availability"`
	TimePeriods  []string  `db:"time_periods"`
}

// GetID возвращает ID источника
func (s *Source) GetID() int32 {
	if s == nil {
		return 0
	}
	return s.ID
}

// GetType возвращает тип источника
func (s *Source) GetType() int32 {
	return s.Type
}

// GetObjectID возвращает ID объекта
func (s *Source) GetObjectID() int32 {
	if s == nil {
		return 0
	}
	return s.ObjectID
}

// GetAddress возвращает адрес источника
func (s *Source) GetAddress() string {
	if s == nil {
		return ""
	}
	return s.Address
}

// GetCreatedAt возвращает время создания
func (s *Source) GetCreatedAt() time.Time {
	if s == nil {
		return time.Time{}
	}
	return s.CreatedAt
}

// GetAvailability признак доступности
func (s *Source) GetAvailability() bool {
	if s == nil {
		return false
	}
	return s.Availability
}

// GetTimePeriods ...
func (s *Source) GetTimePeriods() []string {
	if s == nil {
		return nil
	}
	return s.TimePeriods
}

// Entity ...
func (s *Source) Entity() *entity.Source {
	if s == nil {
		return nil
	}

	return &entity.Source{
		ID:           s.GetID(),
		Type:         entity.SourceType(s.GetType()),
		ObjectID:     s.GetObjectID(),
		Address:      s.GetAddress(),
		CreatedAt:    s.GetCreatedAt(),
		Availability: s.GetAvailability(),
		TimePeriods:  s.GetTimePeriods(),
	}
}

// Sources ...
type Sources []*Source

// Entity ...
func (s Sources) Entity() []*entity.Source { return ToEntitySlice[[]*entity.Source](s) }

func SourceDtoFromEntity(e *entity.Source) *Source {
	if e == nil {
		return nil
	}
	return &Source{
		ID:           e.ID,
		Type:         int32(e.Type),
		ObjectID:     e.ObjectID,
		Address:      e.Address,
		CreatedAt:    e.CreatedAt,
		Availability: e.Availability,
		TimePeriods:  e.TimePeriods,
	}
}
