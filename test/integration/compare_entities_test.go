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

package integration

import (
	"context"
	"path"
	"runtime"
	"testing"

	pb "github.com/datacommonsorg/reconciliation/internal/proto"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestCompareEntities(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	client, err := setup()
	if err != nil {
		t.Fatalf("Failed to set up recon client")
	}
	_, filename, _, _ := runtime.Caller(0)
	goldenPath := path.Join(
		path.Dir(filename), "golden_response/compare_entities")

	for _, c := range []struct {
		req        *pb.CompareEntitiesRequest
		goldenFile string
	}{
		{
			&pb.CompareEntitiesRequest{
				EntityPairs: []*pb.EntityPair{
					{
						EntityOne: &pb.EntitySubGraph{
							SourceId: "aId/PlaceA",
							GraphRepresentation: &pb.EntitySubGraph_EntityIds{
								EntityIds: &pb.EntityIds{
									Ids: []*pb.IdWithProperty{
										{
											Prop: "aId",
											Val:  "PlaceA",
										},
										{
											Prop: "geoId",
											Val:  "0102",
										},
										{
											Prop: "wikidataId",
											Val:  "Q123",
										},
									},
								},
							},
						},
						EntityTwo: &pb.EntitySubGraph{
							SourceId: "bId/PlaceB",
							GraphRepresentation: &pb.EntitySubGraph_EntityIds{
								EntityIds: &pb.EntityIds{
									Ids: []*pb.IdWithProperty{
										{
											Prop: "bId",
											Val:  "PlaceB",
										},
										{
											Prop: "geoId",
											Val:  "0103",
										},
										{
											Prop: "wikidataId",
											Val:  "Q123",
										},
									},
								},
							},
						},
					},
					{
						EntityOne: &pb.EntitySubGraph{
							SourceId: "cId/PlaceC",
							GraphRepresentation: &pb.EntitySubGraph_SubGraph{
								SubGraph: &pb.McfGraph{
									Nodes: map[string]*pb.McfGraph_PropertyValues{
										"cId/PlaceC": {
											Pvs: map[string]*pb.McfGraph_Values{
												"geoId": {
													TypedValues: []*pb.McfGraph_TypedValue{
														{
															Type:  pb.ValueType_TEXT.Enum(),
															Value: proto.String("03933"),
														},
													},
												},
											},
										},
									},
								},
							},
						},
						EntityTwo: &pb.EntitySubGraph{
							SourceId: "dId/PlaceD",
							GraphRepresentation: &pb.EntitySubGraph_SubGraph{
								SubGraph: &pb.McfGraph{
									Nodes: map[string]*pb.McfGraph_PropertyValues{
										"dId/PlaceD": {
											Pvs: map[string]*pb.McfGraph_Values{
												"wikidataId": {
													TypedValues: []*pb.McfGraph_TypedValue{
														{
															Type:  pb.ValueType_TEXT.Enum(),
															Value: proto.String("Q345"),
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"result.json",
		},
	} {
		resp, err := client.CompareEntities(ctx, c.req)
		if err != nil {
			t.Errorf("could not CompareEntities: %s", err)
			continue
		}

		if generateGolden {
			updateProtoGolden(resp, goldenPath, c.goldenFile)
			continue
		}

		var expected pb.CompareEntitiesResponse
		if err = readJSON(goldenPath, c.goldenFile, &expected); err != nil {
			t.Errorf("Can not Unmarshal golden file")
			continue
		}

		if diff := cmp.Diff(resp, &expected, protocmp.Transform()); diff != "" {
			t.Errorf("payload got diff: %v", diff)
			continue
		}
	}
}
