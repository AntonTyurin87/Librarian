package repository

import (
	"Librarian/internal/pkg/service/repository/dto"
	"Librarian/internal/pkg/service/repository/query/photos"
	"context"
	"fmt"
	"strings"

	sq "github.com/Masterminds/squirrel"
)

// InsertPhotos ...
func (r *repository) InsertPhotos(ctx context.Context, q photos.Insert) (dto.Photos, error) {
	var res dto.Photos

	if err := Selectx(ctx, r.storage, &res, insertPhotoQuery(q)); err != nil {
		return nil, fmt.Errorf("selectx(ctx, r.storage, &res, insertPhotoQuery(q)): %w", err)
	}

	return res, nil
}

func insertPhotoQuery(query photos.Insert) sq.InsertBuilder {
	insertQuery := sq.StatementBuilder.Insert(dto.PhotosTableName).
		Columns(
			dto.PhotosColumnGroupId,
			dto.PhotosColumnNameRu,
			dto.PhotosColumnNameNative,
			dto.PhotosColumnFilmingLocation,
			dto.PhotosColumnRegions,
			dto.PhotosColumnTimePeriods,
			dto.PhotosColumnDescription,
			dto.PhotosColumnFormat,
		).
		Prefix("--InsertPhoto\n")

	for _, photo := range query.Photos {
		a := dto.PhotoDtoFromEntity(photo)
		insertQuery = insertQuery.Values(
			a.GetGroupID(),
			a.GetNameRu(),
			a.GetNameNative(),
			a.GetFilmingLocation(),
			strings.Join(a.GetRegions(), ", "),
			strings.Join(a.GetTimePeriods(), ", "),
			a.GetDescription(),
			a.GetFormat(),
		)
	}

	insertQuery = insertQuery.Suffix(fmt.Sprintf("RETURNING %s", strings.Join(dto.PhotosColumns, ", ")))

	return insertQuery
}
