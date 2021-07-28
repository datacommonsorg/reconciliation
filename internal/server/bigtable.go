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

	"cloud.google.com/go/bigtable"
	"github.com/datacommonsorg/reconciliation/internal/util"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Generates a function to be used as the callback function in Bigtable Read.
// This utilizes the Golang closure so the arguments can be scoped in the
// generated function.
func readRowFn(
	errCtx context.Context,
	btTable *bigtable.Table,
	rowSetPart bigtable.RowSet,
	getToken func(string) (string, error),
	action func(string, []byte) (interface{}, error),
	elemChan chan chanData,
) func() error {
	return func() error {
		if err := btTable.ReadRows(errCtx, rowSetPart,
			func(btRow bigtable.Row) bool {
				if len(btRow[util.BtFamily]) == 0 {
					return true
				}
				raw := btRow[util.BtFamily][0].Value

				token, err := getToken(btRow.Key())
				if err != nil {
					return false
				}

				jsonRaw, err := util.UnzipAndDecode(string(raw))
				if err != nil {
					return false
				}
				elem, err := action(token, jsonRaw)
				if err != nil {
					return false
				}
				elemChan <- chanData{token, elem}
				return true
			}); err != nil {
			return err
		}
		return nil
	}
}

// bigTableReadRowsParallel reads BigTable rows from Bigtable in parallel.
//
// Reading multiple rows is chunked as the size limit for RowSet is 500KB.
//
// Args:
// btTable: The bigtable that holds the cache.
// rowSet: BigTable rowSet containing the row keys.
// getToken: A function to get back the indexed token (like place dcid) from
//		bigtable row key.
// action: A callback function that converts the raw bytes into appropriate
//		go struct based on the cache content.
func bigTableReadRowsParallel(
	ctx context.Context,
	btTable *bigtable.Table,
	rowSet bigtable.RowSet,
	getToken func(string) (string, error),
	action func(string, []byte) (interface{}, error),
) (
	map[string]interface{},
	error,
) {
	if btTable == nil || getToken == nil || action == nil {
		return nil, status.Errorf(
			codes.InvalidArgument, "Invalid argument: btTable, getToken, action")
	}

	// Function start
	var rowSetSize int
	var rowList bigtable.RowList
	var rowRangeList bigtable.RowRangeList
	switch v := rowSet.(type) {
	case bigtable.RowList:
		rowList = rowSet.(bigtable.RowList)
		rowSetSize = len(rowList)
	case bigtable.RowRangeList:
		rowRangeList = rowSet.(bigtable.RowRangeList)
		rowSetSize = len(rowRangeList)
	default:
		return nil, status.Errorf(
			codes.Internal, "Unsupported RowSet type: %v", v)
	}
	if rowSetSize == 0 {
		return nil, nil
	}

	ch := make(chan chanData, rowSetSize)

	errs, errCtx := errgroup.WithContext(ctx)
	for i := 0; i <= rowSetSize/util.BtBatchQuerySize; i++ {
		left := i * util.BtBatchQuerySize
		right := (i + 1) * util.BtBatchQuerySize
		if right > rowSetSize {
			right = rowSetSize
		}
		var rowSetPart bigtable.RowSet
		if len(rowList) > 0 {
			rowSetPart = rowList[left:right]
		} else {
			rowSetPart = rowRangeList[left:right]
		}
		errs.Go(readRowFn(errCtx, btTable, rowSetPart, getToken, action, ch))
	}
	err := errs.Wait()
	if err != nil {
		return nil, err
	}
	close(ch)

	result := map[string]interface{}{}
	if btTable != nil {
		for elem := range ch {
			result[elem.token] = elem.data
		}
	}
	return result, nil
}
