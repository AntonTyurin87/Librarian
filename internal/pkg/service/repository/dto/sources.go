package dto

import (
	"Librarian/internal/pkg/domain/entity"
	"time"
)

// Название колонок и таблицы
const (
	SourcesTableName = "sources"

	SourcesColumnID        = "id"
	SourcesColumnType      = "type"
	SourcesColumnObjectID  = "object_id"
	SourcesColumnAddress   = "address"
	SourcesColumnCreatedAt = "created_at"
)

// SourcesColumns ...
var SourcesColumns = []string{
	SourcesColumnID,
	SourcesColumnType,
	SourcesColumnObjectID,
	SourcesColumnAddress,
	SourcesColumnCreatedAt,
}

type Source struct {
	ID        int32     `db:"id"`
	Type      string    `db:"type"`
	ObjectID  int32     `db:"object_id"`
	Address   string    `db:"address"`
	CreatedAt time.Time `db:"created_at"`
}

// GetID возвращает ID источника
func (s *Source) GetID() int32 {
	if s == nil {
		return 0
	}
	return s.ID
}

// GetType возвращает тип источника
func (s *Source) GetType() string {
	if s == nil {
		return ""
	}
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

// Entity ...
func (s *Source) Entity() *entity.Source {
	if s == nil {
		return nil
	}

	return &entity.Source{
		ID:        s.GetID(),
		Type:      s.GetType(),
		ObjectID:  s.GetObjectID(),
		Address:   s.GetAddress(),
		CreatedAt: s.GetCreatedAt(),
	}
}

// Sources ...
type Sources []*Source

// Entity ...
func (s Sources) Entity() []*entity.Source { return ToEntitySlice[[]*entity.Source](s) }
