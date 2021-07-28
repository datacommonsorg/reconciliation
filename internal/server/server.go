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
)

// Server holds resources for a recon server.
type Server struct {
	btTable *bigtable.Table
}

// NewServer creates a new server instance.
func NewServer(btTable *bigtable.Table) *Server {
	return &Server{btTable: btTable}
}

// NewBtTable creates a new bigtable.Table instance.
func NewBtTable(
	ctx context.Context, projectID, instanceID, tableID string) (
	*bigtable.Table, error) {
	btClient, err := bigtable.NewClient(ctx, projectID, instanceID)
	if err != nil {
		return nil, err
	}
	return btClient.Open(tableID), nil
}
