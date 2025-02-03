// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: api/root-server.proto

package root_server_v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	RootServer_Install_FullMethodName = "/rschoonheim.mtls.v1.root_server.RootServer/Install"
)

// RootServerClient is the client API for RootServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RootServerClient interface {
	// Install - installs the root server.
	Install(ctx context.Context, in *RootServerInstallRequest, opts ...grpc.CallOption) (*RootServerInstallResponse, error)
}

type rootServerClient struct {
	cc grpc.ClientConnInterface
}

func NewRootServerClient(cc grpc.ClientConnInterface) RootServerClient {
	return &rootServerClient{cc}
}

func (c *rootServerClient) Install(ctx context.Context, in *RootServerInstallRequest, opts ...grpc.CallOption) (*RootServerInstallResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RootServerInstallResponse)
	err := c.cc.Invoke(ctx, RootServer_Install_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RootServerServer is the server API for RootServer service.
// All implementations must embed UnimplementedRootServerServer
// for forward compatibility.
type RootServerServer interface {
	// Install - installs the root server.
	Install(context.Context, *RootServerInstallRequest) (*RootServerInstallResponse, error)
	mustEmbedUnimplementedRootServerServer()
}

// UnimplementedRootServerServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRootServerServer struct{}

func (UnimplementedRootServerServer) Install(context.Context, *RootServerInstallRequest) (*RootServerInstallResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Install not implemented")
}
func (UnimplementedRootServerServer) mustEmbedUnimplementedRootServerServer() {}
func (UnimplementedRootServerServer) testEmbeddedByValue()                    {}

// UnsafeRootServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RootServerServer will
// result in compilation errors.
type UnsafeRootServerServer interface {
	mustEmbedUnimplementedRootServerServer()
}

func RegisterRootServerServer(s grpc.ServiceRegistrar, srv RootServerServer) {
	// If the following call pancis, it indicates UnimplementedRootServerServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RootServer_ServiceDesc, srv)
}

func _RootServer_Install_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RootServerInstallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RootServerServer).Install(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RootServer_Install_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RootServerServer).Install(ctx, req.(*RootServerInstallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RootServer_ServiceDesc is the grpc.ServiceDesc for RootServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RootServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rschoonheim.mtls.v1.root_server.RootServer",
	HandlerType: (*RootServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Install",
			Handler:    _RootServer_Install_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/root-server.proto",
}
