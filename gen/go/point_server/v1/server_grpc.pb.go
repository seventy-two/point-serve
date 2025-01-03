// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: point_server/v1/server.proto

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

// PointServiceClient is the client API for PointService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PointServiceClient interface {
	AddPoints(ctx context.Context, in *AddPointsRequest, opts ...grpc.CallOption) (*AddPointsResponse, error)
	RemovePoints(ctx context.Context, in *RemovePointsRequest, opts ...grpc.CallOption) (*RemovePointsResponse, error)
	GetPoints(ctx context.Context, in *GetPointsRequest, opts ...grpc.CallOption) (*GetPointsResponse, error)
	GetLeaderboard(ctx context.Context, in *GetLeaderboardRequest, opts ...grpc.CallOption) (*GetLeaderboardResponse, error)
}

type pointServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPointServiceClient(cc grpc.ClientConnInterface) PointServiceClient {
	return &pointServiceClient{cc}
}

func (c *pointServiceClient) AddPoints(ctx context.Context, in *AddPointsRequest, opts ...grpc.CallOption) (*AddPointsResponse, error) {
	out := new(AddPointsResponse)
	err := c.cc.Invoke(ctx, "/point_server.v1.PointService/AddPoints", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pointServiceClient) RemovePoints(ctx context.Context, in *RemovePointsRequest, opts ...grpc.CallOption) (*RemovePointsResponse, error) {
	out := new(RemovePointsResponse)
	err := c.cc.Invoke(ctx, "/point_server.v1.PointService/RemovePoints", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pointServiceClient) GetPoints(ctx context.Context, in *GetPointsRequest, opts ...grpc.CallOption) (*GetPointsResponse, error) {
	out := new(GetPointsResponse)
	err := c.cc.Invoke(ctx, "/point_server.v1.PointService/GetPoints", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pointServiceClient) GetLeaderboard(ctx context.Context, in *GetLeaderboardRequest, opts ...grpc.CallOption) (*GetLeaderboardResponse, error) {
	out := new(GetLeaderboardResponse)
	err := c.cc.Invoke(ctx, "/point_server.v1.PointService/GetLeaderboard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PointServiceServer is the server API for PointService service.
// All implementations should embed UnimplementedPointServiceServer
// for forward compatibility
type PointServiceServer interface {
	AddPoints(context.Context, *AddPointsRequest) (*AddPointsResponse, error)
	RemovePoints(context.Context, *RemovePointsRequest) (*RemovePointsResponse, error)
	GetPoints(context.Context, *GetPointsRequest) (*GetPointsResponse, error)
	GetLeaderboard(context.Context, *GetLeaderboardRequest) (*GetLeaderboardResponse, error)
}

// UnimplementedPointServiceServer should be embedded to have forward compatible implementations.
type UnimplementedPointServiceServer struct {
}

func (UnimplementedPointServiceServer) AddPoints(context.Context, *AddPointsRequest) (*AddPointsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPoints not implemented")
}
func (UnimplementedPointServiceServer) RemovePoints(context.Context, *RemovePointsRequest) (*RemovePointsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemovePoints not implemented")
}
func (UnimplementedPointServiceServer) GetPoints(context.Context, *GetPointsRequest) (*GetPointsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPoints not implemented")
}
func (UnimplementedPointServiceServer) GetLeaderboard(context.Context, *GetLeaderboardRequest) (*GetLeaderboardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLeaderboard not implemented")
}

// UnsafePointServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PointServiceServer will
// result in compilation errors.
type UnsafePointServiceServer interface {
	mustEmbedUnimplementedPointServiceServer()
}

func RegisterPointServiceServer(s grpc.ServiceRegistrar, srv PointServiceServer) {
	s.RegisterService(&PointService_ServiceDesc, srv)
}

func _PointService_AddPoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPointsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointServiceServer).AddPoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/point_server.v1.PointService/AddPoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointServiceServer).AddPoints(ctx, req.(*AddPointsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PointService_RemovePoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemovePointsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointServiceServer).RemovePoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/point_server.v1.PointService/RemovePoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointServiceServer).RemovePoints(ctx, req.(*RemovePointsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PointService_GetPoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPointsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointServiceServer).GetPoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/point_server.v1.PointService/GetPoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointServiceServer).GetPoints(ctx, req.(*GetPointsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PointService_GetLeaderboard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLeaderboardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointServiceServer).GetLeaderboard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/point_server.v1.PointService/GetLeaderboard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointServiceServer).GetLeaderboard(ctx, req.(*GetLeaderboardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PointService_ServiceDesc is the grpc.ServiceDesc for PointService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PointService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "point_server.v1.PointService",
	HandlerType: (*PointServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddPoints",
			Handler:    _PointService_AddPoints_Handler,
		},
		{
			MethodName: "RemovePoints",
			Handler:    _PointService_RemovePoints_Handler,
		},
		{
			MethodName: "GetPoints",
			Handler:    _PointService_GetPoints_Handler,
		},
		{
			MethodName: "GetLeaderboard",
			Handler:    _PointService_GetLeaderboard_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "point_server/v1/server.proto",
}
