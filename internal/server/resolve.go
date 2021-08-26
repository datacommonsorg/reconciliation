// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"context"
	"fmt"
	"strings"

	"cloud.google.com/go/bigtable"
	pb "github.com/datacommonsorg/reconciliation/internal/proto"
	"github.com/datacommonsorg/reconciliation/internal/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
)

// ResolveEntities implements API for ReconServer.ResolveEntities.
func (s *Server) ResolveEntities(ctx context.Context, in *pb.ResolveEntitiesRequest) (
	*pb.ResolveEntitiesResponse, error) {
	rowList := bigtable.RowList{}
	idKeyToSourceIDs := map[string][]string{}
	sourceIDs := map[string]struct{}{}

	// Collect to-be-resolved IDs to rowList and idKeyToSourceID.
	for _, entity := range in.GetEntities() {
		ids, err := util.IDsFromEntitySubGraph(entity)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "IDsFromEntitySubGraph() = %s", err)
		}

		sourceID := entity.GetSourceId()
		sourceIDs[sourceID] = struct{}{}
		for idProp, idVal := range ids {
			idKey := fmt.Sprintf("%s^%s", idProp, idVal)
			rowList = append(rowList, fmt.Sprintf("%s%s", util.BtReconIDMapPrefix, idKey))
			idKeyToSourceIDs[idKey] = append(idKeyToSourceIDs[idKey], sourceID)
		}
	}

	// Read ReconIdMap cache.
	dataMap, err := bigTableReadRowsParallel(ctx, s.btTable, rowList,
		func(rowKey string) (string, error) {
			return strings.TrimPrefix(rowKey, util.BtReconIDMapPrefix), nil
		},
		func(dcid string, jsonRaw []byte) (interface{}, error) {
			var reconEntities pb.ReconEntities
			err := protojson.Unmarshal(jsonRaw, &reconEntities)
			if err != nil {
				return nil, err
			}
			return &reconEntities, nil
		})
	if err != nil {
		return nil, err
	}

	// Source ID -> ID Prop -> ReconEntities.
	reconEntityStore := map[string]map[string]*pb.ReconEntities{}

	// Group resolving cache result by source ID.
	for idKey, reconEntities := range dataMap {
		if reconEntities == nil {
			continue
		}
		sourceIDs, ok := idKeyToSourceIDs[idKey]
		if !ok {
			continue
		}

		parts := strings.Split(idKey, "^")
		if len(parts) != 2 {
			return nil, status.Errorf(codes.Internal, "Invalid id key %s", idKey)
		}
		idProp := parts[0]

		for _, sourceID := range sourceIDs {
			if _, ok := reconEntityStore[sourceID]; !ok {
				reconEntityStore[sourceID] = map[string]*pb.ReconEntities{}
			}
			if re := reconEntities.(*pb.ReconEntities); len(re.GetEntities()) > 0 {
				reconEntityStore[sourceID][idProp] = re
			}
		}
	}

	// Assemble response.
	res := &pb.ResolveEntitiesResponse{}
	for sourceId, idProp2ReconEntities := range reconEntityStore {
		var reconEntities *pb.ReconEntities
		for _, idProp := range util.RankedIDProps {
			if val, ok := idProp2ReconEntities[idProp]; ok {
				reconEntities = val
				break
			}
		}
		if reconEntities == nil {
			continue
		}

		// If it is resolved to multiple DC entities, each resolved entity has an equal probability.
		probability := float64(1.0 / len(reconEntities.GetEntities()))

		resolvedEntity := &pb.ResolveEntitiesResponse_ResolvedEntity{
			SourceId: sourceId,
		}

		for _, entity := range reconEntities.GetEntities() {
			resolvedId := &pb.ResolveEntitiesResponse_ResolvedId{
				Probability: probability,
			}
			for _, id := range entity.GetIds() {
				resolvedId.Ids = append(resolvedId.Ids,
					&pb.IdWithProperty{
						Prop: id.GetProp(),
						Val:  id.GetVal(),
					})
			}
			resolvedEntity.ResolvedIds = append(resolvedEntity.ResolvedIds, resolvedId)
		}

		res.ResolvedEntities = append(res.ResolvedEntities, resolvedEntity)
	}

	// Add entities that are not resolved as empty result.
	for sourceID := range sourceIDs {
		if _, ok := reconEntityStore[sourceID]; ok { // Resolved.
			continue
		}
		res.ResolvedEntities = append(res.ResolvedEntities,
			&pb.ResolveEntitiesResponse_ResolvedEntity{
				SourceId: sourceID,
			})
	}

	return res, nil
}
