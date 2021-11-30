package server

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"cloud.google.com/go/bigtable"
	pb "github.com/datacommonsorg/reconciliation/internal/proto"
	"github.com/datacommonsorg/reconciliation/internal/util"
	"google.golang.org/protobuf/encoding/protojson"
)

func (s *Server) ResolveIds(ctx context.Context, in *pb.ResolveIdsRequest) (
	*pb.ResolveIdsResponse, error) {
	inProp := in.GetInProp()
	outProp := in.GetOutProp()
	ids := in.GetIds()
	if inProp == "" || outProp == "" || len(ids) == 0 {
		return nil, fmt.Errorf(
			"invalid input: in_prop: %s, out_prop: %s, ids: %v", inProp, outProp, ids)
	}

	// Read cache data.
	rowList := bigtable.RowList{}
	for _, id := range ids {
		rowList = append(rowList,
			fmt.Sprintf("%s%s^%s^%s", util.BtReconIDMapPrefix, inProp, id, outProp))
	}
	dataMap, err := bigTableReadRowsParallel(ctx, s.btTable, rowList,
		func(rowKey string) (string, error) {
			parts := strings.Split(rowKey, "^")
			if len(parts) != 3 {
				return "", fmt.Errorf("wrong rowKey: %s", rowKey)
			}
			return parts[1], nil
		},
		func(dcid string, jsonRaw []byte) (interface{}, error) {
			var reconEntities pb.ReconEntities
			if err := protojson.Unmarshal(jsonRaw, &reconEntities); err != nil {
				return nil, err
			}
			return &reconEntities, nil
		})
	if err != nil {
		return nil, err
	}

	// Assemble result.
	res := &pb.ResolveIdsResponse{}
	for inID, reconEntities := range dataMap {
		entity := &pb.ResolveIdsResponse_Entity{InId: inID}

		for _, reconEntity := range reconEntities.(*pb.ReconEntities).GetEntities() {
			if len(reconEntity.GetIds()) != 1 {
				return nil, fmt.Errorf("wrong cache result for %s: %v",
					inID, reconEntities)
			}
			entity.OutIds = append(entity.OutIds, reconEntity.GetIds()[0].GetVal())
		}

		// Sort to make the result deterministic.
		sort.Strings(entity.OutIds)

		res.Entities = append(res.Entities, entity)
	}

	// Sort to make the result deterministic.
	sort.Slice(res.Entities, func(i, j int) bool {
		return res.Entities[i].GetInId() > res.Entities[j].GetInId()
	})

	return res, nil
}
