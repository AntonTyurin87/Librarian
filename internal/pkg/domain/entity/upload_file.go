package entity

// UploadFileRequest ...
type UploadFileRequest struct {
	FileAddress    string
	FileSourceData *FileSourceData
}

// GetFileAddress ...
func (u *UploadFileRequest) GetFileAddress() string {
	if u == nil {
		return ""
	}
	return u.FileAddress
}

// GetFileSourceData ...
func (u *UploadFileRequest) GetFileSourceData() *FileSourceData {
	if u == nil {
		return nil
	}
	return u.FileSourceData
}

type UploadFileResponse struct {
	ObjectID    int32
	SourceType  SourceType
	TimePeriods []string
}

// SourceType - тип источников
type SourceType int32

// Константы типов источников
const (
	SourceType_UNKNOWN  SourceType = 0 //Неизвестный
	SourceType_BOOK     SourceType = 1 // Книга
	SourceType_ARTICLE  SourceType = 2 // Статья
	SourceType_FRAGMENT SourceType = 3 // Фрагмент
	SourceType_PHOTO    SourceType = 4 // Фотография или картинка
)

type FileFormat int32

const (
	FileFormat_UNKNOWN  FileFormat = 0  // Неизвестный формат
	FileFormat_PDF      FileFormat = 1  // PDF документ
	FileFormat_DOC      FileFormat = 2  // Microsoft Word документ
	FileFormat_DOCX     FileFormat = 3  // Microsoft Word Open XML
	FileFileFormat_DJVU FileFormat = 10 // Сканированный документ DJVU
	FileFileFormat_CBR  FileFormat = 11 // Комикс (RAR архив)
	FileFileFormat_CBZ  FileFormat = 12 // Комикс (ZIP архив)
	FileFileFormat_DJV  FileFormat = 13 // Сканированный документ DJVU (альтернативное расширение)
	FileFileFormat_PNG  FileFormat = 20 // Изображение PNG
	FileFileFormat_JPG  FileFormat = 21 // Изображение JPEG
	FileFileFormat_JPEG FileFormat = 22 // Изображение JPEG (альтернативное расширение)
	FileFileFormat_GIF  FileFormat = 23 // Изображение GIF
	FileFileFormat_BMP  FileFormat = 24 // Изображение BMP
	FileFileFormat_TIFF FileFormat = 25 // Изображение TIFF
	FileFileFormat_TIF  FileFormat = 26 // Изображение TIFF (альтернативное расширение)
	FileFileFormat_SVG  FileFormat = 27 // Векторное изображение SVG
)

// FileSourceData ...
type FileSourceData struct {
	SourceType SourceType `json:"source_type"`
	Book       *Book      `json:"book"`
	Articles   *Article   `json:"articles"`
	Fragment   *Fragment  `json:"fragment"`
	Photo      *Photo     `json:"photo"`
}

// GetSourceType возвращает тип источника
func (fsd *FileSourceData) GetSourceType() SourceType {
	if fsd == nil {
		return SourceType_UNKNOWN
	}
	return fsd.SourceType
}

// GetBook возвращает книгу
func (fsd *FileSourceData) GetBook() *Book {
	if fsd == nil {
		return nil
	}
	return fsd.Book
}

// GetArticles возвращает статью
func (fsd *FileSourceData) GetArticles() *Article {
	if fsd == nil {
		return nil
	}
	return fsd.Articles
}

// GetFragment возвращает фрагмент
func (fsd *FileSourceData) GetFragment() *Fragment {
	if fsd == nil {
		return nil
	}
	return fsd.Fragment
}

// GetPhoto возвращает фотографию
func (fsd *FileSourceData) GetPhoto() *Photo {
	if fsd == nil {
		return nil
	}
	return fsd.Photo
}
