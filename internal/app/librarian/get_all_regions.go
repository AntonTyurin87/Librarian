package librarian

import (
	"context"
	"fmt"

	lib "github.com/AntonTyurin87/Recon_Com_protoc/gen/go/librarian"
)

// GetAllRegions ...
func (s *Server) GetAllRegions(ctx context.Context, _ *lib.GetAllRegionsRequest) (*lib.GetAllRegionsResponse, error) {
	allRegions, err := s.usacase.GetAllRegions(ctx)
	if err != nil {
		return nil, fmt.Errorf("s.usacase.GetAllRegions: %w", err)
	}

	return &lib.GetAllRegionsResponse{
		Region: s.presenter.RegionsFromEntityToLib(allRegions.GetRegions()),
	}, nil
}
