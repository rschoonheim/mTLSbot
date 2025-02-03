package main

import (
	"google.golang.org/grpc"
	"log/slog"
	root_server_v1 "mTLS/cmd/root_server/v1/services"
	"net"
	"os"
)

var (
	// address - the address where the root server is running.
	address = ":50051"
)

func init() {
	if os.Getenv("SERVER_ADDRESS") != "" {
		address = os.Getenv("SERVER_ADDRESS")
	}
}

func main() {
	slog.Info("Starting root server", "address", address)

	// Create a listener on the address.
	//
	lis, err := net.Listen("tcp", address)
	if err != nil {
		slog.Error("Failed to listen", "error", err)
	}

	// Configure and start the gRPC server.
	//
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	root_server_v1.RegisterRootServerServer(grpcServer, root_server_v1.RootServerService{})
	grpcServer.Serve(lis)
}
