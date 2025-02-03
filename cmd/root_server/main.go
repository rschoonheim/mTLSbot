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

	// dataPath - the path to store data of the root server.
	dataPath = "/etc/mTLSbot"
)

func init() {

	// Override the variables with the environment variables if they are set.
	//
	if os.Getenv("SERVER_ADDRESS") != "" {
		address = os.Getenv("SERVER_ADDRESS")
	}

	if os.Getenv("MTLS_BOT_DATA_PATH") != "" {
		dataPath = os.Getenv("MTLS_BOT_DATA_PATH")
	}

	// Make sure that the data path exists.
	//
	if _, err := os.Stat(dataPath); os.IsNotExist(err) {
		slog.Error("The data path does not exist", "path", dataPath)
		os.Exit(1)
	}

	// Create `state.json` in the data path if it does not exist.
	//
	if _, err := os.Stat(dataPath + "/state.json"); os.IsNotExist(err) {
		err := os.WriteFile(dataPath+"/state.json", []byte("{}"), 0644)
		if err != nil {
			slog.Error("Failed to create state file", "error", err)
			os.Exit(1)
		}
	}

	// Make sure that the default directories exist.
	//
	directories := []string{
		dataPath + "/servers",
		dataPath + "/servers/clients",
	}

	for _, dir := range directories {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			err := os.MkdirAll(dir, 0755)
			if err != nil {
				slog.Error("Failed to create directory", "error", err)
				os.Exit(1)
			}
		}

		// Each directory must have a `state.json` file.
		//
		if _, err := os.Stat(dir + "/state.json"); os.IsNotExist(err) {
			err := os.WriteFile(dir+"/state.json", []byte("{}"), 0644)
			if err != nil {
				slog.Error("Failed to create state file", "error", err)
				os.Exit(1)
			}
		}
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
