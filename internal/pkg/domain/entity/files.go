package entity

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
