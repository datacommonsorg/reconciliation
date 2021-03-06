package server

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"cloud.google.com/go/bigtable"
	pb "github.com/datacommonsorg/reconciliation/internal/proto"
	"github.com/datacommonsorg/reconciliation/internal/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
)

var (
	// This is a preferred list.
	// The props ranked higher are preferred over those ranked lower for resolving.
	rankedIDProps = []string{
		"dcid",
		"geoId",
		"isoCode",
		"nutsCode",
		"wikidataId",
		"geoNamesId",
		"istatId",
		"austrianMunicipalityKey",
		"indianCensusAreaCode2011",
	}
)

// ResolveEntities implements API for ReconServer.ResolveEntities.
func (s *Server) ResolveEntities(ctx context.Context, in *pb.ResolveEntitiesRequest) (
	*pb.ResolveEntitiesResponse, error) {
	rowList := bigtable.RowList{}
	idKeyToSourceIDs := map[string][]string{}
	sourceIDs := map[string]struct{}{}

	// Collect to-be-resolved IDs to rowList and idKeyToSourceID.
	for _, entity := range in.GetEntities() {
		sourceID := entity.GetSourceId()

		// Try to resolve all the supported IDs
		// For the resolved ones, only rely on the one ranked higher.
		switch t := entity.GraphRepresentation.(type) {
		case *pb.EntitySubGraph_SubGraph:
			node, ok := (entity.GetSubGraph().GetNodes())[sourceID]
			if !ok {
				continue
			}
			for _, idProp := range rankedIDProps {
				idVal := util.GetPropVal(node, idProp)
				if idVal == "" {
					continue
				}
				idKey := fmt.Sprintf("%s^%s", idProp, idVal)
				rowList = append(rowList, fmt.Sprintf("%s%s", util.BtReconIDMapPrefix, idKey))
				idKeyToSourceIDs[idKey] = append(idKeyToSourceIDs[idKey], sourceID)
			}
		case *pb.EntitySubGraph_EntityIds:
			idStore := map[string]string{} // Map: ID prop -> ID val.
			for _, id := range entity.GetEntityIds().GetIds() {
				idStore[id.GetProp()] = id.GetVal()
			}
			for _, idProp := range rankedIDProps {
				idVal, ok := idStore[idProp]
				if !ok {
					continue
				}
				idKey := fmt.Sprintf("%s^%s", idProp, idVal)
				rowList = append(rowList, fmt.Sprintf("%s%s", util.BtReconIDMapPrefix, idKey))
				idKeyToSourceIDs[idKey] = append(idKeyToSourceIDs[idKey], sourceID)
			}
		default:
			return nil, fmt.Errorf("Entity.GraphRepresentation has unexpected type %T", t)
		}

		sourceIDs[sourceID] = struct{}{}
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
		for _, idProp := range rankedIDProps {
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

	// Sort to make the result deterministic.
	sort.Slice(res.ResolvedEntities, func(i, j int) bool {
		return res.ResolvedEntities[i].GetSourceId() > res.ResolvedEntities[j].GetSourceId()
	})

	return res, nil
}
