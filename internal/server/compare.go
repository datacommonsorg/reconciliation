package server

import (
	"context"

	pb "github.com/datacommonsorg/reconciliation/internal/proto"
)

// CompareEntities implements API for ReconServer.CompareEntities.
func (s *Server) CompareEntities(ctx context.Context, in *pb.CompareEntitiesRequest) (
	*pb.CompareEntitiesResponse, error) {
	// TODO(spaceenter): Implement.
	return &pb.CompareEntitiesResponse{
		Comparisons: []*pb.CompareEntitiesResponse_Comparison{
			{
				SourceIds:   []string{"aaa", "bbb"},
				Probability: 0.67,
			},
		},
	}, nil
}
