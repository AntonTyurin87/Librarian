package usecase

import (
	"Librarian/internal/pkg/domain/entity"
	"Librarian/internal/pkg/service/repository/query/regions"
	"context"
	"fmt"
)

// GetAllRegions ...
func (u *usecase) GetAllRegions(ctx context.Context) (*entity.GetAllRegionsResponse, error) {
	selectRegions, err := u.repository.SelectRegions(ctx, regions.Select{})
	if err != nil {
		return nil, fmt.Errorf("u.repository.GetAllRegions: %w", err)
	}

	return &entity.GetAllRegionsResponse{
		Regions: selectRegions,
	}, nil
}
