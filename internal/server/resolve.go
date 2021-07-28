package server

import (
	"context"

	pb "github.com/datacommonsorg/reconciliation/internal/proto"
)

// ResolveEntities implements API for ReconServer.ResolveEntities.
func (s *Server) ResolveEntities(ctx context.Context, in *pb.ResolveEntitiesRequest) (
	*pb.ResolveEntitiesResponse, error) {
	// TODO(spaceenter): Implement.
	return nil, nil
}
