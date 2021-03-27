// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// HealthServiceClient is the client API for HealthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HealthServiceClient interface {
	GetHealth(ctx context.Context, in *GetHealthRequest, opts ...grpc.CallOption) (*GetHealthResponse, error)
}

type healthServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHealthServiceClient(cc grpc.ClientConnInterface) HealthServiceClient {
	return &healthServiceClient{cc}
}

func (c *healthServiceClient) GetHealth(ctx context.Context, in *GetHealthRequest, opts ...grpc.CallOption) (*GetHealthResponse, error) {
	out := new(GetHealthResponse)
	err := c.cc.Invoke(ctx, "/inetmock.rpc.v1.HealthService/GetHealth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HealthServiceServer is the server API for HealthService service.
// All implementations must embed UnimplementedHealthServiceServer
// for forward compatibility
type HealthServiceServer interface {
	GetHealth(context.Context, *GetHealthRequest) (*GetHealthResponse, error)
	mustEmbedUnimplementedHealthServiceServer()
}

// UnimplementedHealthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHealthServiceServer struct {
}

func (UnimplementedHealthServiceServer) GetHealth(context.Context, *GetHealthRequest) (*GetHealthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHealth not implemented")
}
func (UnimplementedHealthServiceServer) mustEmbedUnimplementedHealthServiceServer() {}

// UnsafeHealthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HealthServiceServer will
// result in compilation errors.
type UnsafeHealthServiceServer interface {
	mustEmbedUnimplementedHealthServiceServer()
}

func RegisterHealthServiceServer(s grpc.ServiceRegistrar, srv HealthServiceServer) {
	s.RegisterService(&HealthService_ServiceDesc, srv)
}

func _HealthService_GetHealth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHealthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServiceServer).GetHealth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inetmock.rpc.v1.HealthService/GetHealth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServiceServer).GetHealth(ctx, req.(*GetHealthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HealthService_ServiceDesc is the grpc.ServiceDesc for HealthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HealthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inetmock.rpc.v1.HealthService",
	HandlerType: (*HealthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHealth",
			Handler:    _HealthService_GetHealth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc/v1/health.proto",
}
