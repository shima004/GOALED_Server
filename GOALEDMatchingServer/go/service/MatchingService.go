package service

import (
	pb "GOALED/go/pb"
	"context"

	"github.com/google/uuid"
)

type MatchingServer struct {
	pb.UnimplementedMatchingServiceServer
	Rooms map[string]*pb.Room
}

func (ms *MatchingServer) GetPlayerId(ctx context.Context, in *pb.GetPlayerIdRequest) (*pb.GetPlayerIdResponse, error) {
	return &pb.GetPlayerIdResponse{
		PlayerId: uuid.New().String(),
	}, nil
}

func (ms *MatchingServer) GetPublicRoom(ctx context.Context, in *pb.GetPublicRoomRequest) (*pb.GetPublicRoomResponse, error) {
	rooms := make([]*pb.Room, 0)
	for _, room := range ms.Rooms {
		if room.Password == "" {
			rooms = append(rooms, room)
		}
	}

	return &pb.GetPublicRoomResponse{
		Rooms: rooms,
	}, nil
}

func (ms *MatchingServer) CreatePublicRoom(ctx context.Context, in *pb.CreatePublicRoomRequest) (*pb.CreatePublicRoomResponse, error) {
	room := &pb.Room{
		Id:            uuid.New().String(),
		Name:          in.GetName(),
		Password:      "",
		Status:        pb.RoomStatus_WAITING,
		MaxPlayer:     in.GetMaxPlayer(),
		CurrentPlayer: 0,
		Players:       make([]*pb.Player, 0),
	}
	ms.Rooms[room.Id] = room
	return &pb.CreatePublicRoomResponse{
		Room: room,
	}, nil
}

func (ms *MatchingServer) CreatePrivateRoom(ctx context.Context, in *pb.CreatePrivateRoomRequest) (*pb.CreatePrivateRoomResponse, error) {
	room := &pb.Room{
		Id:            uuid.New().String(),
		Name:          in.GetName(),
		Password:      in.GetPassword(),
		Status:        pb.RoomStatus_WAITING,
		MaxPlayer:     in.GetMaxPlayer(),
		CurrentPlayer: 0,
		Players:       make([]*pb.Player, 0),
	}
	ms.Rooms[room.Id] = room
	return &pb.CreatePrivateRoomResponse{
		Room: room,
	}, nil
}

func (ms *MatchingServer) JoinPublicRoom(ctx context.Context, in *pb.JoinPublicRoomRequest) (*pb.JoinPublicRoomResponse, error) {
	room, ok := ms.Rooms[in.GetRoomId()]
	if !ok {
		return &pb.JoinPublicRoomResponse{
			Room: nil,
		}, nil
	}

	room.Players = append(room.Players, in.GetPlayer())
	room.CurrentPlayer += 1
	return &pb.JoinPublicRoomResponse{
		Room: room,
	}, nil
}

func (ms *MatchingServer) JoinPrivateRoom(ctx context.Context, in *pb.JoinPrivateRoomRequest) (*pb.JoinPrivateRoomResponse, error) {
	password := in.GetPassword()
	for _, room := range ms.Rooms {
		if room.Password == password {
			room.Players = append(room.Players, in.GetPlayer())
			room.CurrentPlayer += 1
			return &pb.JoinPrivateRoomResponse{
				Room: room,
			}, nil
		}
	}
	return &pb.JoinPrivateRoomResponse{
		Room: nil,
	}, nil
}

func (ms *MatchingServer) LeaveRoom(ctx context.Context, in *pb.LeaveRoomRequest) (*pb.LeaveRoomResponse, error) {
	room := ms.Rooms[in.GetRoomId()]
	for i, player := range room.Players {
		if player.Id == in.GetPlayerId() {
			room.Players = append(room.Players[:i], room.Players[i+1:]...)
			room.CurrentPlayer -= 1
			return &pb.LeaveRoomResponse{
				Success: true,
			}, nil
		}
	}
	return &pb.LeaveRoomResponse{
		Success: false,
	}, nil
}

func NewMatchingServer() *MatchingServer {
	return &MatchingServer{
		Rooms: make(map[string]*pb.Room),
	}
}
