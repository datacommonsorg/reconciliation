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

// TODO(spaceenter): Cite this from a sharable repo as a lib.

package util

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io/ioutil"

	pb "github.com/datacommonsorg/reconciliation/internal/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	// BtReconIDMapPrefix for ID mapping for ID-based recon. The key excludes DCID.
	BtReconIDMapPrefix = "d/5/"

	// BtFamily is the key for the row.
	BtFamily = "csv"

	// BtBatchQuerySize is the size of BigTable batch query.
	BtBatchQuerySize = 1000
)

var (
	// RankedIDProps is a preferred list.
	// The props ranked higher are preferred over those ranked lower for resolving.
	RankedIDProps = []string{
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

// UnzipAndDecode decompresses the given contents using gzip and decodes it from base64.
func UnzipAndDecode(contents string) ([]byte, error) {
	// Decode from base64.
	decode, err := base64.StdEncoding.DecodeString(contents)
	if err != nil {
		return nil, err
	}

	// Unzip the string.
	gzReader, err := gzip.NewReader(bytes.NewReader(decode))
	if err != nil {
		return nil, err
	}
	defer gzReader.Close()
	gzResult, err := ioutil.ReadAll(gzReader)
	if err != nil {
		return nil, err
	}
	return gzResult, nil
}

// Get the value of a given property, assuming single value.
func GetPropVal(node *pb.McfGraph_PropertyValues, prop string) string {
	values, ok := (node.GetPvs())[prop]
	if !ok {
		return ""
	}
	typedValues := values.GetTypedValues()
	if len(typedValues) == 0 {
		return ""
	}
	return typedValues[0].GetValue()
}

// Get {ID prop, ID val} from EntitySubGraph.
func IDsFromEntitySubGraph(entity *pb.EntitySubGraph) (map[string]string, error) {
	sourceID := entity.GetSourceId()
	result := map[string]string{}

	switch t := entity.GraphRepresentation.(type) {
	case *pb.EntitySubGraph_SubGraph:
		node, ok := (entity.GetSubGraph().GetNodes())[sourceID]
		if !ok {
			return nil, status.Errorf(codes.Internal, "Node not found for %s", sourceID)
		}
		for _, idProp := range RankedIDProps {
			idVal := GetPropVal(node, idProp)
			if idVal == "" {
				continue
			}
			result[idProp] = idVal
		}
	case *pb.EntitySubGraph_EntityIds:
		idStore := map[string]string{} // Map: ID prop -> ID val.
		for _, id := range entity.GetEntityIds().GetIds() {
			idStore[id.GetProp()] = id.GetVal()
		}
		for _, idProp := range RankedIDProps {
			idVal, ok := idStore[idProp]
			if !ok {
				continue
			}
			result[idProp] = idVal
		}
	default:
		return nil, fmt.Errorf("Entity.GraphRepresentation has unexpected type %T", t)
	}

	return result, nil
}
