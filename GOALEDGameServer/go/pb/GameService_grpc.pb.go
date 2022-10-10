// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: GameService.proto

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

// GameServiceClient is the client API for GameService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GameServiceClient interface {
	CreateRoom(ctx context.Context, in *Room, opts ...grpc.CallOption) (*Room, error)
	SyncPlayerData(ctx context.Context, in *SyncPlayerDataRequest, opts ...grpc.CallOption) (GameService_SyncPlayerDataClient, error)
	SendPlayerData(ctx context.Context, opts ...grpc.CallOption) (GameService_SendPlayerDataClient, error)
	SyncObject(ctx context.Context, in *SyncObjectRequest, opts ...grpc.CallOption) (GameService_SyncObjectClient, error)
	SendObject(ctx context.Context, opts ...grpc.CallOption) (GameService_SendObjectClient, error)
}

type gameServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGameServiceClient(cc grpc.ClientConnInterface) GameServiceClient {
	return &gameServiceClient{cc}
}

func (c *gameServiceClient) CreateRoom(ctx context.Context, in *Room, opts ...grpc.CallOption) (*Room, error) {
	out := new(Room)
	err := c.cc.Invoke(ctx, "/GameService.GameService/CreateRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) SyncPlayerData(ctx context.Context, in *SyncPlayerDataRequest, opts ...grpc.CallOption) (GameService_SyncPlayerDataClient, error) {
	stream, err := c.cc.NewStream(ctx, &GameService_ServiceDesc.Streams[0], "/GameService.GameService/SyncPlayerData", opts...)
	if err != nil {
		return nil, err
	}
	x := &gameServiceSyncPlayerDataClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GameService_SyncPlayerDataClient interface {
	Recv() (*PlayerData, error)
	grpc.ClientStream
}

type gameServiceSyncPlayerDataClient struct {
	grpc.ClientStream
}

func (x *gameServiceSyncPlayerDataClient) Recv() (*PlayerData, error) {
	m := new(PlayerData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *gameServiceClient) SendPlayerData(ctx context.Context, opts ...grpc.CallOption) (GameService_SendPlayerDataClient, error) {
	stream, err := c.cc.NewStream(ctx, &GameService_ServiceDesc.Streams[1], "/GameService.GameService/SendPlayerData", opts...)
	if err != nil {
		return nil, err
	}
	x := &gameServiceSendPlayerDataClient{stream}
	return x, nil
}

type GameService_SendPlayerDataClient interface {
	Send(*PlayerData) error
	CloseAndRecv() (*SendPlayerDataResponse, error)
	grpc.ClientStream
}

type gameServiceSendPlayerDataClient struct {
	grpc.ClientStream
}

func (x *gameServiceSendPlayerDataClient) Send(m *PlayerData) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gameServiceSendPlayerDataClient) CloseAndRecv() (*SendPlayerDataResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SendPlayerDataResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *gameServiceClient) SyncObject(ctx context.Context, in *SyncObjectRequest, opts ...grpc.CallOption) (GameService_SyncObjectClient, error) {
	stream, err := c.cc.NewStream(ctx, &GameService_ServiceDesc.Streams[2], "/GameService.GameService/SyncObject", opts...)
	if err != nil {
		return nil, err
	}
	x := &gameServiceSyncObjectClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GameService_SyncObjectClient interface {
	Recv() (*SyncObjectResponse, error)
	grpc.ClientStream
}

type gameServiceSyncObjectClient struct {
	grpc.ClientStream
}

func (x *gameServiceSyncObjectClient) Recv() (*SyncObjectResponse, error) {
	m := new(SyncObjectResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *gameServiceClient) SendObject(ctx context.Context, opts ...grpc.CallOption) (GameService_SendObjectClient, error) {
	stream, err := c.cc.NewStream(ctx, &GameService_ServiceDesc.Streams[3], "/GameService.GameService/SendObject", opts...)
	if err != nil {
		return nil, err
	}
	x := &gameServiceSendObjectClient{stream}
	return x, nil
}

type GameService_SendObjectClient interface {
	Send(*SendObjectRequest) error
	CloseAndRecv() (*SendObjectResponse, error)
	grpc.ClientStream
}

type gameServiceSendObjectClient struct {
	grpc.ClientStream
}

func (x *gameServiceSendObjectClient) Send(m *SendObjectRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gameServiceSendObjectClient) CloseAndRecv() (*SendObjectResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SendObjectResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GameServiceServer is the server API for GameService service.
// All implementations must embed UnimplementedGameServiceServer
// for forward compatibility
type GameServiceServer interface {
	CreateRoom(context.Context, *Room) (*Room, error)
	SyncPlayerData(*SyncPlayerDataRequest, GameService_SyncPlayerDataServer) error
	SendPlayerData(GameService_SendPlayerDataServer) error
	SyncObject(*SyncObjectRequest, GameService_SyncObjectServer) error
	SendObject(GameService_SendObjectServer) error
	mustEmbedUnimplementedGameServiceServer()
}

// UnimplementedGameServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGameServiceServer struct {
}

func (UnimplementedGameServiceServer) CreateRoom(context.Context, *Room) (*Room, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRoom not implemented")
}
func (UnimplementedGameServiceServer) SyncPlayerData(*SyncPlayerDataRequest, GameService_SyncPlayerDataServer) error {
	return status.Errorf(codes.Unimplemented, "method SyncPlayerData not implemented")
}
func (UnimplementedGameServiceServer) SendPlayerData(GameService_SendPlayerDataServer) error {
	return status.Errorf(codes.Unimplemented, "method SendPlayerData not implemented")
}
func (UnimplementedGameServiceServer) SyncObject(*SyncObjectRequest, GameService_SyncObjectServer) error {
	return status.Errorf(codes.Unimplemented, "method SyncObject not implemented")
}
func (UnimplementedGameServiceServer) SendObject(GameService_SendObjectServer) error {
	return status.Errorf(codes.Unimplemented, "method SendObject not implemented")
}
func (UnimplementedGameServiceServer) mustEmbedUnimplementedGameServiceServer() {}

// UnsafeGameServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GameServiceServer will
// result in compilation errors.
type UnsafeGameServiceServer interface {
	mustEmbedUnimplementedGameServiceServer()
}

func RegisterGameServiceServer(s grpc.ServiceRegistrar, srv GameServiceServer) {
	s.RegisterService(&GameService_ServiceDesc, srv)
}

func _GameService_CreateRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Room)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).CreateRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GameService.GameService/CreateRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).CreateRoom(ctx, req.(*Room))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_SyncPlayerData_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SyncPlayerDataRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GameServiceServer).SyncPlayerData(m, &gameServiceSyncPlayerDataServer{stream})
}

type GameService_SyncPlayerDataServer interface {
	Send(*PlayerData) error
	grpc.ServerStream
}

type gameServiceSyncPlayerDataServer struct {
	grpc.ServerStream
}

func (x *gameServiceSyncPlayerDataServer) Send(m *PlayerData) error {
	return x.ServerStream.SendMsg(m)
}

func _GameService_SendPlayerData_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GameServiceServer).SendPlayerData(&gameServiceSendPlayerDataServer{stream})
}

type GameService_SendPlayerDataServer interface {
	SendAndClose(*SendPlayerDataResponse) error
	Recv() (*PlayerData, error)
	grpc.ServerStream
}

type gameServiceSendPlayerDataServer struct {
	grpc.ServerStream
}

func (x *gameServiceSendPlayerDataServer) SendAndClose(m *SendPlayerDataResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gameServiceSendPlayerDataServer) Recv() (*PlayerData, error) {
	m := new(PlayerData)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GameService_SyncObject_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SyncObjectRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GameServiceServer).SyncObject(m, &gameServiceSyncObjectServer{stream})
}

type GameService_SyncObjectServer interface {
	Send(*SyncObjectResponse) error
	grpc.ServerStream
}

type gameServiceSyncObjectServer struct {
	grpc.ServerStream
}

func (x *gameServiceSyncObjectServer) Send(m *SyncObjectResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _GameService_SendObject_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GameServiceServer).SendObject(&gameServiceSendObjectServer{stream})
}

type GameService_SendObjectServer interface {
	SendAndClose(*SendObjectResponse) error
	Recv() (*SendObjectRequest, error)
	grpc.ServerStream
}

type gameServiceSendObjectServer struct {
	grpc.ServerStream
}

func (x *gameServiceSendObjectServer) SendAndClose(m *SendObjectResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gameServiceSendObjectServer) Recv() (*SendObjectRequest, error) {
	m := new(SendObjectRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GameService_ServiceDesc is the grpc.ServiceDesc for GameService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GameService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "GameService.GameService",
	HandlerType: (*GameServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRoom",
			Handler:    _GameService_CreateRoom_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SyncPlayerData",
			Handler:       _GameService_SyncPlayerData_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SendPlayerData",
			Handler:       _GameService_SendPlayerData_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "SyncObject",
			Handler:       _GameService_SyncObject_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SendObject",
			Handler:       _GameService_SendObject_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "GameService.proto",
}
