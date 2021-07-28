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

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/datacommonsorg/reconciliation/internal/healthcheck"
	pb "github.com/datacommonsorg/reconciliation/internal/proto"
	"github.com/datacommonsorg/reconciliation/internal/server"
	"golang.org/x/oauth2/google"

	"cloud.google.com/go/profiler"
	"google.golang.org/api/compute/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/alts"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

var (
	storeProject = flag.String("store_project", "", "GCP project stores Bigtable.")
	btTableName  = flag.String("bt_table", "", "Cache Bigtable table.")
	port         = flag.Int("port", 12345, "Port on which to run the server.")
	useALTS      = flag.Bool("use_alts", false, "Whether to use ALTS server authentication")
)

const (
	btInstance = "prophet-cache"
)

func main() {
	fmt.Println("Enter recon main() function")

	flag.Parse()
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	ctx := context.Background()

	// Profiler.
	credentials, error := google.FindDefaultCredentials(ctx, compute.ComputeScope)
	if error == nil && credentials.ProjectID != "" {
		if err := profiler.Start(profiler.Config{
			Service: "recon-service",
		}); err != nil {
			log.Printf("Failed to start profiler: %v", err)
		}
	}

	// Cache BT Table.
	btTable, err := server.NewBtTable(ctx, *storeProject, btInstance, *btTableName)
	if err != nil {
		log.Fatalf("Failed to create BigTable client: %v", err)
	}

	// Server opts.
	opts := []grpc.ServerOption{}
	if *useALTS {
		// Use ALTS server credential to bind to VM's private IPv6 interface.
		altsTC := alts.NewServerCreds(alts.DefaultServerOptions())
		opts = append(opts, grpc.Creds(altsTC))
	}

	// Start recon server.
	srv := grpc.NewServer(opts...)
	s := server.NewServer(btTable)
	pb.RegisterReconServer(srv, s)
	// Register reflection service on gRPC server.
	reflection.Register(srv)

	healthService := healthcheck.NewHealthChecker()
	grpc_health_v1.RegisterHealthServer(srv, healthService)

	// Listen on network
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen on network: %v", err)
	}
	fmt.Println("Recon ready to serve!!")
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
