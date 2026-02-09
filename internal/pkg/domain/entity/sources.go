package entity

// GetURLForDownloadSourceRequest ...
type GetURLForDownloadSourceRequest struct {
	TextSources []*TextSource
}

// GetSources ...
func (g *GetURLForDownloadSourceRequest) GetSources() []*TextSource {
	if g == nil {
		return nil
	}

	return g.TextSources
}

// GetURLForDownloadSourceResponse ...
type GetURLForDownloadSourceResponse struct {
	URL      string
	FileInfo *FileInfo
}

// GetURL ...
func (g *GetURLForDownloadSourceResponse) GetURL() string {
	if g == nil {
		return ""
	}
	return g.URL
}

// GetFileInfo ...
func (g *GetURLForDownloadSourceResponse) GetFileInfo() *FileInfo {
	if g == nil {
		return nil
	}
	return g.FileInfo
}

// FileInfo ...
type FileInfo struct {
	Size     int64
	FileName string
	FileType string
}
