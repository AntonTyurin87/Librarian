package entity

import "time"

// GetInfoForDownloadRequest ...
type GetInfoForDownloadRequest struct {
	Sources []*Source
}

// GetSources ...
func (g *GetInfoForDownloadRequest) GetSources() []*Source {
	if g == nil {
		return nil
	}
	return g.Sources
}

type Source struct {
	ID        int32     `json:"id"`
	Type      string    `json:"type"`
	ObjectID  int32     `json:"object_id"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
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

// GetInfoForDownloadResponse ...
type GetInfoForDownloadResponse struct {
	URL      string    `json:"url"`
	FileInfo *FileInfo `json:"file_info"`
}

// GetURL ...
func (r *GetInfoForDownloadResponse) GetURL() string {
	if r == nil {
		return ""
	}
	return r.URL
}

// GetFileInfo ...
func (r *GetInfoForDownloadResponse) GetFileInfo() *FileInfo {
	if r == nil {
		return nil
	}
	return r.FileInfo
}

// FileInfo ...
type FileInfo struct {
	Size     int64  `json:"size"`
	MimeType string `json:"mime_type"`
}

// GetSize ...
func (f *FileInfo) GetSize() int64 {
	if f == nil {
		return 0
	}
	return f.Size
}

// GetMimeType ...
func (f *FileInfo) GetMimeType() string {
	if f == nil {
		return ""
	}
	return f.MimeType
}
