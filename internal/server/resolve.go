package server

import (
	"context"

	pb "github.com/datacommonsorg/reconciliation/internal/proto"
)

// ResolveEntities implements API for ReconServer.ResolveEntities.
func (s *Server) ResolveEntities(ctx context.Context, in *pb.ResolveEntitiesRequest) (
	*pb.ResolveEntitiesResponse, error) {
	// TODO(spaceenter): Implement.
	return &pb.ResolveEntitiesResponse{
		ResolvedEntities: []*pb.ResolveEntitiesResponse_ResolvedEntity{
			{
				SourceId: "aaa",
				ResolvedIds: []*pb.ResolveEntitiesResponse_ResolvedId{
					{
						Ids: []*pb.ResolveEntitiesResponse_ResolvedId_IdWithProperty{
							{
								Prop: "wikidataId",
								Val:  "Q1234",
							},
						},
						Probability: 0.58,
					},
				},
			},
		},
	}, nil
}
