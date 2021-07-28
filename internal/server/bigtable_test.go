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
	"testing"

	"cloud.google.com/go/bigtable"
	"github.com/google/go-cmp/cmp"
)

func TestReadOneTable(t *testing.T) {
	ctx := context.Background()
	data := map[string]string{
		"key1": "data1",
		"key2": "data2",
	}
	btTable, err := SetupBigtable(ctx, data)
	if err != nil {
		t.Errorf("setupBigtable got error: %v", err)
	}
	rowList := bigtable.RowList{"key1", "key2"}
	baseDataMap, err := bigTableReadRowsParallel(
		ctx,
		btTable,
		rowList,
		func(key string) (string, error) { return key, nil },
		func(dcid string, jsonRaw []byte) (interface{}, error) {
			return string(jsonRaw), nil
		},
	)
	if err != nil {
		t.Errorf("btReadRowsParallel got error: %v", err)
	}
	for dcid, result := range baseDataMap {
		if diff := cmp.Diff(data[dcid], result.(string)); diff != "" {
			t.Errorf("read rows got diff from table data %+v", diff)
		}
	}
}
