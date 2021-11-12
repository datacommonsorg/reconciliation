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

package integration

import (
	"context"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net"
	"path"
	"runtime"
	"strings"

	"cloud.google.com/go/bigtable"
	pb "github.com/datacommonsorg/reconciliation/internal/proto"
	"github.com/datacommonsorg/reconciliation/internal/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
)

var generateGolden bool

func init() {
	flag.BoolVar(
		&generateGolden, "generate_golden", false, "generate golden files")
}

// This test runs against staging bt dataset.
// It needs Application Default Credentials to run locally or need to
// provide service account credential when running on GCP.
const (
	btInstance   = "prophet-cache"
	storeProject = "datcom-store"
)

func setup() (pb.ReconClient, error) {
	ctx := context.Background()
	_, filename, _, _ := runtime.Caller(0)
	btTableName, _ := ioutil.ReadFile(path.Join(path.Dir(filename),
		"../../deploy/storage/bigtable.version"))

	btTable, err := server.NewBtTable(
		ctx, storeProject, btInstance, strings.TrimSpace(string(btTableName)))
	if err != nil {
		return nil, err
	}

	return newClient(btTable)
}

func newClient(
	btTable *bigtable.Table) (pb.ReconClient, error) {
	s := server.NewServer(btTable)
	srv := grpc.NewServer()
	pb.RegisterReconServer(srv, s)
	reflection.Register(srv)
	lis, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return nil, err
	}
	// Start recon at localhost:0.
	go func() {
		if err := srv.Serve(lis); err != nil {
			log.Fatalf("failed to start recon in localhost:0")
		}
	}()

	// Create recon client.
	conn, err := grpc.Dial(
		lis.Addr().String(),
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(200000000 /* 100M */)))
	if err != nil {
		return nil, err
	}
	client := pb.NewReconClient(conn)
	return client, nil
}

func updateProtoGolden(
	resp protoreflect.ProtoMessage, root, fname string) {
	var err error
	marshaller := protojson.MarshalOptions{Indent: ""}
	// protojson doesn't and won't make stable output:
	// https://github.com/golang/protobuf/issues/1082
	// Use encoding/json to get stable output.
	data, err := marshaller.Marshal(resp)
	if err != nil {
		log.Printf("could not write golden files to %s", fname)
		return
	}
	var rm json.RawMessage = data
	jsonByte, err := json.MarshalIndent(rm, "", "  ")
	if err != nil {
		log.Printf("could not write golden files to %s", fname)
		return
	}
	err = ioutil.WriteFile(path.Join(root, fname), jsonByte, 0644)
	if err != nil {
		log.Printf("could not write golden files to %s", fname)
	}
}

func readJSON(dir, fname string, resp protoreflect.ProtoMessage) error {
	bytes, err := ioutil.ReadFile(path.Join(dir, fname))
	if err != nil {
		return err
	}
	err = protojson.Unmarshal(bytes, resp)
	if err != nil {
		return err
	}
	return nil
}
