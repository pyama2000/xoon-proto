// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: server_status/v1/server_status_service.proto

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

// ServerStatusServiceClient is the client API for ServerStatusService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServerStatusServiceClient interface {
	GetStatus(ctx context.Context, in *GetStatusRequest, opts ...grpc.CallOption) (*GetStatusResponse, error)
}

type serverStatusServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewServerStatusServiceClient(cc grpc.ClientConnInterface) ServerStatusServiceClient {
	return &serverStatusServiceClient{cc}
}

func (c *serverStatusServiceClient) GetStatus(ctx context.Context, in *GetStatusRequest, opts ...grpc.CallOption) (*GetStatusResponse, error) {
	out := new(GetStatusResponse)
	err := c.cc.Invoke(ctx, "/server_status.v1.ServerStatusService/GetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServerStatusServiceServer is the server API for ServerStatusService service.
// All implementations should embed UnimplementedServerStatusServiceServer
// for forward compatibility
type ServerStatusServiceServer interface {
	GetStatus(context.Context, *GetStatusRequest) (*GetStatusResponse, error)
}

// UnimplementedServerStatusServiceServer should be embedded to have forward compatible implementations.
type UnimplementedServerStatusServiceServer struct {
}

func (UnimplementedServerStatusServiceServer) GetStatus(context.Context, *GetStatusRequest) (*GetStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}

// UnsafeServerStatusServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServerStatusServiceServer will
// result in compilation errors.
type UnsafeServerStatusServiceServer interface {
	mustEmbedUnimplementedServerStatusServiceServer()
}

func RegisterServerStatusServiceServer(s grpc.ServiceRegistrar, srv ServerStatusServiceServer) {
	s.RegisterService(&ServerStatusService_ServiceDesc, srv)
}

func _ServerStatusService_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerStatusServiceServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server_status.v1.ServerStatusService/GetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerStatusServiceServer).GetStatus(ctx, req.(*GetStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ServerStatusService_ServiceDesc is the grpc.ServiceDesc for ServerStatusService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServerStatusService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "server_status.v1.ServerStatusService",
	HandlerType: (*ServerStatusServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStatus",
			Handler:    _ServerStatusService_GetStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server_status/v1/server_status_service.proto",
}
