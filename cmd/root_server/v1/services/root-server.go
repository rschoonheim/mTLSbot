package root_server_v1

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RootServerService struct{}

// GetRootServerConfiguration - returns the current configuration of the root server.
func (rs RootServerService) GetRootServerConfiguration(context.Context, *RootServerConfiguration) (*RootServerConfiguration, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRootServerConfiguration not implemented")
}

func (rs RootServerService) mustEmbedUnimplementedRootServerServer() {}
