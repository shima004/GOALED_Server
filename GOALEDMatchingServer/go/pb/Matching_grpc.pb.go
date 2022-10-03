// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: Matching.proto

package pb

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

// MatchingServiceClient is the client API for MatchingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MatchingServiceClient interface {
	GetPlayerId(ctx context.Context, in *GetPlayerIdRequest, opts ...grpc.CallOption) (*GetPlayerIdResponse, error)
	GetPublicRoom(ctx context.Context, in *GetPublicRoomRequest, opts ...grpc.CallOption) (*GetPublicRoomResponse, error)
	CreatePublicRoom(ctx context.Context, in *CreatePublicRoomRequest, opts ...grpc.CallOption) (*CreatePublicRoomResponse, error)
	CreatePrivateRoom(ctx context.Context, in *CreatePrivateRoomRequest, opts ...grpc.CallOption) (*CreatePrivateRoomResponse, error)
	JoinPublicRoom(ctx context.Context, in *JoinPublicRoomRequest, opts ...grpc.CallOption) (*JoinPublicRoomResponse, error)
	JoinPrivateRoom(ctx context.Context, in *JoinPrivateRoomRequest, opts ...grpc.CallOption) (*JoinPrivateRoomResponse, error)
	LeaveRoom(ctx context.Context, in *LeaveRoomRequest, opts ...grpc.CallOption) (*LeaveRoomResponse, error)
}

type matchingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMatchingServiceClient(cc grpc.ClientConnInterface) MatchingServiceClient {
	return &matchingServiceClient{cc}
}

func (c *matchingServiceClient) GetPlayerId(ctx context.Context, in *GetPlayerIdRequest, opts ...grpc.CallOption) (*GetPlayerIdResponse, error) {
	out := new(GetPlayerIdResponse)
	err := c.cc.Invoke(ctx, "/MatchingService.MatchingService/GetPlayerId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matchingServiceClient) GetPublicRoom(ctx context.Context, in *GetPublicRoomRequest, opts ...grpc.CallOption) (*GetPublicRoomResponse, error) {
	out := new(GetPublicRoomResponse)
	err := c.cc.Invoke(ctx, "/MatchingService.MatchingService/GetPublicRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matchingServiceClient) CreatePublicRoom(ctx context.Context, in *CreatePublicRoomRequest, opts ...grpc.CallOption) (*CreatePublicRoomResponse, error) {
	out := new(CreatePublicRoomResponse)
	err := c.cc.Invoke(ctx, "/MatchingService.MatchingService/CreatePublicRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matchingServiceClient) CreatePrivateRoom(ctx context.Context, in *CreatePrivateRoomRequest, opts ...grpc.CallOption) (*CreatePrivateRoomResponse, error) {
	out := new(CreatePrivateRoomResponse)
	err := c.cc.Invoke(ctx, "/MatchingService.MatchingService/CreatePrivateRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matchingServiceClient) JoinPublicRoom(ctx context.Context, in *JoinPublicRoomRequest, opts ...grpc.CallOption) (*JoinPublicRoomResponse, error) {
	out := new(JoinPublicRoomResponse)
	err := c.cc.Invoke(ctx, "/MatchingService.MatchingService/JoinPublicRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matchingServiceClient) JoinPrivateRoom(ctx context.Context, in *JoinPrivateRoomRequest, opts ...grpc.CallOption) (*JoinPrivateRoomResponse, error) {
	out := new(JoinPrivateRoomResponse)
	err := c.cc.Invoke(ctx, "/MatchingService.MatchingService/JoinPrivateRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matchingServiceClient) LeaveRoom(ctx context.Context, in *LeaveRoomRequest, opts ...grpc.CallOption) (*LeaveRoomResponse, error) {
	out := new(LeaveRoomResponse)
	err := c.cc.Invoke(ctx, "/MatchingService.MatchingService/LeaveRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MatchingServiceServer is the server API for MatchingService service.
// All implementations must embed UnimplementedMatchingServiceServer
// for forward compatibility
type MatchingServiceServer interface {
	GetPlayerId(context.Context, *GetPlayerIdRequest) (*GetPlayerIdResponse, error)
	GetPublicRoom(context.Context, *GetPublicRoomRequest) (*GetPublicRoomResponse, error)
	CreatePublicRoom(context.Context, *CreatePublicRoomRequest) (*CreatePublicRoomResponse, error)
	CreatePrivateRoom(context.Context, *CreatePrivateRoomRequest) (*CreatePrivateRoomResponse, error)
	JoinPublicRoom(context.Context, *JoinPublicRoomRequest) (*JoinPublicRoomResponse, error)
	JoinPrivateRoom(context.Context, *JoinPrivateRoomRequest) (*JoinPrivateRoomResponse, error)
	LeaveRoom(context.Context, *LeaveRoomRequest) (*LeaveRoomResponse, error)
	mustEmbedUnimplementedMatchingServiceServer()
}

// UnimplementedMatchingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMatchingServiceServer struct {
}

func (UnimplementedMatchingServiceServer) GetPlayerId(context.Context, *GetPlayerIdRequest) (*GetPlayerIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlayerId not implemented")
}
func (UnimplementedMatchingServiceServer) GetPublicRoom(context.Context, *GetPublicRoomRequest) (*GetPublicRoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPublicRoom not implemented")
}
func (UnimplementedMatchingServiceServer) CreatePublicRoom(context.Context, *CreatePublicRoomRequest) (*CreatePublicRoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePublicRoom not implemented")
}
func (UnimplementedMatchingServiceServer) CreatePrivateRoom(context.Context, *CreatePrivateRoomRequest) (*CreatePrivateRoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePrivateRoom not implemented")
}
func (UnimplementedMatchingServiceServer) JoinPublicRoom(context.Context, *JoinPublicRoomRequest) (*JoinPublicRoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinPublicRoom not implemented")
}
func (UnimplementedMatchingServiceServer) JoinPrivateRoom(context.Context, *JoinPrivateRoomRequest) (*JoinPrivateRoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinPrivateRoom not implemented")
}
func (UnimplementedMatchingServiceServer) LeaveRoom(context.Context, *LeaveRoomRequest) (*LeaveRoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaveRoom not implemented")
}
func (UnimplementedMatchingServiceServer) mustEmbedUnimplementedMatchingServiceServer() {}

// UnsafeMatchingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MatchingServiceServer will
// result in compilation errors.
type UnsafeMatchingServiceServer interface {
	mustEmbedUnimplementedMatchingServiceServer()
}

func RegisterMatchingServiceServer(s grpc.ServiceRegistrar, srv MatchingServiceServer) {
	s.RegisterService(&MatchingService_ServiceDesc, srv)
}

func _MatchingService_GetPlayerId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlayerIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchingServiceServer).GetPlayerId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MatchingService.MatchingService/GetPlayerId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchingServiceServer).GetPlayerId(ctx, req.(*GetPlayerIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MatchingService_GetPublicRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPublicRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchingServiceServer).GetPublicRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MatchingService.MatchingService/GetPublicRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchingServiceServer).GetPublicRoom(ctx, req.(*GetPublicRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MatchingService_CreatePublicRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePublicRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchingServiceServer).CreatePublicRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MatchingService.MatchingService/CreatePublicRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchingServiceServer).CreatePublicRoom(ctx, req.(*CreatePublicRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MatchingService_CreatePrivateRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePrivateRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchingServiceServer).CreatePrivateRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MatchingService.MatchingService/CreatePrivateRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchingServiceServer).CreatePrivateRoom(ctx, req.(*CreatePrivateRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MatchingService_JoinPublicRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinPublicRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchingServiceServer).JoinPublicRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MatchingService.MatchingService/JoinPublicRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchingServiceServer).JoinPublicRoom(ctx, req.(*JoinPublicRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MatchingService_JoinPrivateRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinPrivateRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchingServiceServer).JoinPrivateRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MatchingService.MatchingService/JoinPrivateRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchingServiceServer).JoinPrivateRoom(ctx, req.(*JoinPrivateRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MatchingService_LeaveRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaveRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchingServiceServer).LeaveRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MatchingService.MatchingService/LeaveRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchingServiceServer).LeaveRoom(ctx, req.(*LeaveRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MatchingService_ServiceDesc is the grpc.ServiceDesc for MatchingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MatchingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "MatchingService.MatchingService",
	HandlerType: (*MatchingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPlayerId",
			Handler:    _MatchingService_GetPlayerId_Handler,
		},
		{
			MethodName: "GetPublicRoom",
			Handler:    _MatchingService_GetPublicRoom_Handler,
		},
		{
			MethodName: "CreatePublicRoom",
			Handler:    _MatchingService_CreatePublicRoom_Handler,
		},
		{
			MethodName: "CreatePrivateRoom",
			Handler:    _MatchingService_CreatePrivateRoom_Handler,
		},
		{
			MethodName: "JoinPublicRoom",
			Handler:    _MatchingService_JoinPublicRoom_Handler,
		},
		{
			MethodName: "JoinPrivateRoom",
			Handler:    _MatchingService_JoinPrivateRoom_Handler,
		},
		{
			MethodName: "LeaveRoom",
			Handler:    _MatchingService_LeaveRoom_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Matching.proto",
}