package dto

import (
	"Librarian/internal/pkg/domain/entity"
)

const (
	FilesTableName = "files"

	FilesColumnID           = "id"
	FilesColumnTextSourceID = "text_source_id"
	FilesColumnFileData     = "file_data"
)

var FilesColumns = []string{
	FilesColumnID,
	FilesColumnTextSourceID,
	FilesColumnFileData,
}

// File ...
type File struct {
	ID           int64  `json:"id"`
	TextSourceID int64  `json:"text_source_id"`
	FileData     []byte `json:"file_data"`
}

// GetID ...
func (f *File) GetID() int64 {
	if f == nil {
		return 0
	}
	return f.ID
}

// GetTextSourceID ...
func (f *File) GetTextSourceID() int64 {
	if f == nil {
		return 0
	}
	return f.TextSourceID
}

// GetFileData ...
func (f *File) GetFileData() []byte {
	if f == nil {
		return nil
	}
	return f.FileData
}

// Entity ...
func (f *File) Entity() *entity.File {
	if f == nil {
		return nil
	}

	return &entity.File{
		ID:           f.GetID(),
		TextSourceID: f.GetTextSourceID(),
		FileData:     f.GetFileData(),
	}
}

// Files ...
type Files []*File

// Entity ...
func (f Files) Entity() []*entity.File {
	return ToEntitySlice[[]*entity.File](f)
}

// FileDtoFromEntity ...
func FileDtoFromEntity(e *entity.File) *File {
	if e == nil {
		return nil
	}
	return &File{
		ID:           e.GetID(),
		TextSourceID: e.GetTextSourceID(),
		FileData:     e.GetFileData(),
	}
}
