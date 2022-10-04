package service

import (
	"GOALED/go/model"
	pb "GOALED/go/pb"
	"context"
)

type GameServer struct {
	pb.UnimplementedGameServiceServer
	room    map[string]*model.GameModel
	players map[string]string
}

func (gs *GameServer) CreateRoom(ctx context.Context, in *pb.Room) (*pb.Room, error) {
	for _, user := range in.GetPlayerIds() {
		gs.players[user] = in.GetId()
	}
	gs.room[in.Id] = model.NewGameModel(in)
	return in, nil
}

func (gs *GameServer) SyncPlayerData(in *pb.SyncPlayerDataRequest, stream pb.GameService_SyncPlayerDataServer) error {
	done := make(chan struct{})
	player_id := in.PlayerId
	gs.room[in.RoomId].GetStreams().AddPlayerDataStream(done, player_id, stream)
	<-done
	return nil
}

func (gs *GameServer) SendPlayerData(stream pb.GameService_SendPlayerDataServer) error {
	for {
		playerData, err := stream.Recv()
		if err != nil {
			return err
		}
		gs.room[gs.players[playerData.Id]].AddPlayerData(playerData)
	}
}

func (gs *GameServer) SyncObject(in *pb.SyncObjectRequest, stream pb.GameService_SyncObjectServer) error {
	done := make(chan struct{})
	player_id := in.PlayerId
	gs.room[in.RoomId].GetStreams().AddObjectStream(done, player_id, stream)
	<-done
	return nil
}

func (gs *GameServer) SendObject(stream pb.GameService_SendObjectServer) error {
	for {
		object, err := stream.Recv()
		if err != nil {
			return err
		}
		gs.room[gs.players[object.Owner]].AddObject(object)
	}
}

func NewGameServer() *GameServer {
	return &GameServer{
		room:    make(map[string]*model.GameModel),
		players: make(map[string]string),
	}
}
