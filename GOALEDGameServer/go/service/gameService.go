package service

import (
	pb "GOALED/go/pb"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GameServer struct {
	pb.UnimplementedGameServiceServer
}

func (gs *GameServer) CreateRoom(ctx context.Context, in *pb.Room) (*pb.Room, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRoom not implemented")
}

func (gs *GameServer) SyncPlayerData(in *pb.SyncPlayerDataRequest, stream pb.GameService_SyncPlayerDataServer) error {
	return status.Errorf(codes.Unimplemented, "method SyncPlayerData not implemented")
}

func (gs *GameServer) SendPlayerData(stream pb.GameService_SendPlayerDataServer) error {
	return status.Errorf(codes.Unimplemented, "method SendPlayerData not implemented")
}

func (gs *GameServer) SyncObject(in *pb.SyncObjectRequest, stream pb.GameService_SyncObjectServer) error {
	return status.Errorf(codes.Unimplemented, "method SyncObject not implemented")
}

func (gs *GameServer) SendObject(stream pb.GameService_SendObjectServer) error {
	return status.Errorf(codes.Unimplemented, "method SendObject not implemented")
}

func NewGameServer() *GameServer {
	return &GameServer{}
}
