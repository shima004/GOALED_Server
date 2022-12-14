package service

import (
	"GOALED/go/model"
	pb "GOALED/go/pb"
	"context"
	"log"

	"github.com/fatih/color"
)

type GameServer struct {
	pb.UnimplementedGameServiceServer
	room map[string]*model.GameModel
}

func (gs *GameServer) CreateRoom(ctx context.Context, in *pb.Room) (*pb.Room, error) {
	room_id := in.Id
	log.Printf("{\"method\": %s, \"args\": {\"room_id\": %s}}", color.GreenString("[CreateRoom]"), room_id)
	if _, ok := gs.room[room_id]; !ok {
		gs.room[room_id] = model.NewGameModel(in)
		return in, nil
	} else {
		return gs.room[room_id].Room, nil
	}
}

func (gs *GameServer) SyncPlayerData(in *pb.SyncPlayerDataRequest, stream pb.GameService_SyncPlayerDataServer) error {
	done := make(chan struct{})
	player_id := in.PlayerId
	room_id := in.RoomId
	if room, ok := gs.room[room_id]; ok {
		room.GetStreams().AddPlayerDataStream(done, player_id, stream)
		log.Printf("{\"method\": %s, \"args\": {\"player_id\": %s, \"room_id\": %s}}", color.GreenString("[SyncPlayerData]"), player_id, room_id)
		<-done
		return nil
	} else {
		log.Printf("{\"method\": %s, \"args\": {\"player_id\": %s, \"room_id\": %s}}", color.RedString("[SyncPlayerData]"), player_id, room_id)
		return nil
	}
}

func (gs *GameServer) SendPlayerData(stream pb.GameService_SendPlayerDataServer) error {
	isFirst := true
	for {
		res, err := stream.Recv()
		if err != nil {
			return err
		}
		room_id := res.RoomId
		if room, ok := gs.room[room_id]; ok {
			if isFirst {
				log.Printf("{\"method\": %s, \"args\": {\"room_id\": %s}}", color.GreenString("[SendPlayerData]"), room_id)
				isFirst = false
			}
			room.AddPlayerData(res.PlayerData)
		} else {
			log.Printf("{\"method\": %s, \"args\": {\"room_id\": %s}}", color.RedString("[SendPlayerData]"), room_id)
		}
	}
}

func (gs *GameServer) SyncObject(in *pb.SyncObjectRequest, stream pb.GameService_SyncObjectServer) error {
	done := make(chan struct{})
	player_id := in.PlayerId
	room_id := in.RoomId
	if room, ok := gs.room[room_id]; ok {
		room.GetStreams().AddObjectStream(done, player_id, stream)
		log.Printf("{\"method\": %s, \"args\": {\"player_id\": %s, \"room_id\": %s}}", color.GreenString("[SyncObject]"), player_id, room_id)
		<-done
		return nil
	} else {
		log.Printf("{\"method\": %s, \"args\": {\"player_id\": %s, \"room_id\": %s}}", color.RedString("[SyncObject]"), player_id, room_id)
		return nil
	}
}

func (gs *GameServer) SendObject(stream pb.GameService_SendObjectServer) error {
	isFirst := true
	for {
		res, err := stream.Recv()
		if err != nil {
			return err
		}
		room_id := res.RoomId
		if room, ok := gs.room[room_id]; ok {
			if isFirst {
				log.Printf("{\"method\": %s, \"args\": {\"room_id\": %s}}", color.GreenString("[SendObject]"), room_id)
				isFirst = false
			}
			room.AddObject(res.Object)
		} else {
			log.Printf("{\"method\": %s, \"args\": {\"room_id\": %s}}", color.RedString("[SendObject]"), room_id)
		}
	}
}

func (gs *GameServer) CloseStream(ctx context.Context, in *pb.CloseStreamRequest) (*pb.CloseStreamResponse, error) {
	room_id := in.RoomId
	player_id := in.PlayerId
	object := in.Object
	if room, ok := gs.room[room_id]; ok {
		room.AddObject(object)
		isFinish := room.RemovePlayer(player_id)
		if isFinish {
			gs.room[room_id].Close()
			delete(gs.room, room_id)
			log.Printf("{\"method\": %s, \"args\": {\"room_id\": %s}}", color.GreenString("[CloseRoom]"), room_id)
		}
		log.Printf("{\"method\": %s, \"args\": {\"player_id\": %s, \"room_id\": %s}}", color.GreenString("[CloseStream]"), player_id, room_id)
		return &pb.CloseStreamResponse{
			Message: "success",
		}, nil
	} else {
		log.Printf("{\"method\": %s, \"args\": {\"player_id\": %s, \"room_id\": %s}}", color.RedString("[CloseStream]"), player_id, room_id)
		return &pb.CloseStreamResponse{
			Message: "fail",
		}, nil
	}
}

func NewGameServer() *GameServer {
	return &GameServer{
		room: make(map[string]*model.GameModel),
	}
}
