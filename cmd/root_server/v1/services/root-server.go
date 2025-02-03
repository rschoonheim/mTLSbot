package root_server_v1

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RootServerService struct{}

// Install - Installs the root server.
func (rs RootServerService) Install(context.Context, *RootServerInstallRequest) (*RootServerInstallResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Install not implemented")
}

func (rs RootServerService) mustEmbedUnimplementedRootServerServer() {}
