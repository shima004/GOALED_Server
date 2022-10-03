package service

import (
	model "GOALED/go/model"
	pb "GOALED/go/pb"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type gameSyncServer struct {
	pb.UnimplementedGameSyncServiceServer
	rooms map[string]*model.SyncModel
}

func (gss *gameSyncServer) StartGame(ctx context.Context, in *pb.StartGameRequest) (*pb.StartGameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartGame not implemented")
}

func (gss *gameSyncServer) SendGameObject(stream pb.GameSyncService_SendGameObjectServer) error {
	for {
		res, err := stream.Recv()
		if err != nil {
			return err
		}
		roomId := res.GetRoomId()
		if _, ok := gss.rooms[roomId]; !ok {
			gss.rooms[roomId] = model.NewSyncModel(roomId, "test", res.GetPlayerId())
		}
		gss.rooms[roomId].UpdateObjects(res.GetGameObjects())
	}
}

func (gss *gameSyncServer) RecvGameObject(in *pb.RecvGameObjectRequest, stream pb.GameSyncService_RecvGameObjectServer) error {
	roomId := in.GetRoomId()
	playerId := in.GetPlayerId()
	if _, ok := gss.rooms[roomId]; !ok {
		gss.rooms[roomId] = model.NewSyncModel(roomId, "test", playerId)
	}
	done := make(chan struct{})
	gss.rooms[roomId].AddStream(done, playerId, stream)
	<-done
	return nil
}

func NewGameSyncServer() *gameSyncServer {
	return &gameSyncServer{
		rooms: make(map[string]*model.SyncModel),
	}
}
