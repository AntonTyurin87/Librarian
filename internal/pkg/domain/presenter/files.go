package presenter

import (
	"Librarian/internal/pkg/domain/entity"
	"fmt"
	"os"

	lib "github.com/AntonTyurin87/Recon_Com_protoc/gen/go/librarian"
)

// FileFormatFromLibToEntity ...
func (p *presenter) FileFormatFromLibToEntity(format lib.FileFormat) entity.FileFormat {
	switch format {
	case lib.FileFormat_FILE_FORMAT_PDF:
		return entity.FileFormat_PDF
	case lib.FileFormat_FILE_FORMAT_DOC:
		return entity.FileFormat_DOC
	case lib.FileFormat_FILE_FORMAT_DOCX:
		return entity.FileFormat_DOCX
	case lib.FileFormat_FILE_FORMAT_DJVU:
		return entity.FileFileFormat_DJVU
	case lib.FileFormat_FILE_FORMAT_CBR:
		return entity.FileFileFormat_CBR
	case lib.FileFormat_FILE_FORMAT_CBZ:
		return entity.FileFileFormat_CBZ
	case lib.FileFormat_FILE_FORMAT_DJV:
		return entity.FileFileFormat_DJV
	case lib.FileFormat_FILE_FORMAT_PNG:
		return entity.FileFileFormat_PNG
	case lib.FileFormat_FILE_FORMAT_JPG:
		return entity.FileFileFormat_JPG
	case lib.FileFormat_FILE_FORMAT_JPEG:
		return entity.FileFileFormat_JPEG
	case lib.FileFormat_FILE_FORMAT_GIF:
		return entity.FileFileFormat_GIF
	case lib.FileFormat_FILE_FORMAT_BMP:
		return entity.FileFileFormat_BMP
	case lib.FileFormat_FILE_FORMAT_TIFF:
		return entity.FileFileFormat_TIFF
	case lib.FileFormat_FILE_FORMAT_TIF:
		return entity.FileFileFormat_TIF
	case lib.FileFormat_FILE_FORMAT_SVG:
		return entity.FileFileFormat_SVG

	default:
		return entity.FileFormat_UNKNOWN
	}
}

// DeleteFile ...
func (p *presenter) DeleteFile(address string) error {
	err := os.Remove(address)
	if err != nil {
		return fmt.Errorf("error while deleting file: %w", err)
	}
	return nil
}
