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

	pb "github.com/datacommonsorg/reconciliation/internal/proto"
	"github.com/datacommonsorg/reconciliation/internal/util"
	"google.golang.org/protobuf/proto"
)

// CompareEntities implements API for ReconServer.CompareEntities.
func (s *Server) CompareEntities(ctx context.Context, in *pb.CompareEntitiesRequest) (
	*pb.CompareEntitiesResponse, error) {
	res := &pb.CompareEntitiesResponse{}

	for _, entityPair := range in.GetEntityPairs() {
		entity1 := entityPair.GetEntityOne()
		entity2 := entityPair.GetEntityTwo()

		ids1, err := util.IDsFromEntitySubGraph(entity1)
		if err != nil {
			return nil, err
		}
		ids2, err := util.IDsFromEntitySubGraph(entity2)
		if err != nil {
			return nil, err
		}

		sharedIDPropCnt := 0
		sharedIDPropAndValCnt := 0
		for idKey1, idVal1 := range ids1 {
			if idVal2, ok := ids2[idKey1]; ok {
				sharedIDPropCnt++
				if idVal1 == idVal2 {
					sharedIDPropAndValCnt++
				}
			}
		}

		var probability float64
		if sharedIDPropCnt == 0 {
			probability = 0
		} else {
			probability = float64(sharedIDPropAndValCnt) / float64(sharedIDPropCnt)
		}
		res.Comparisons = append(res.Comparisons, &pb.CompareEntitiesResponse_Comparison{
			SourceIdOne: entity1.GetSourceId(),
			SourceIdTwo: entity2.GetSourceId(),
			Probability: proto.Float64(probability),
		})
	}

	return res, nil
}
