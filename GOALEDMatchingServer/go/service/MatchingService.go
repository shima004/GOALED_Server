package service

import (
	"GOALED/go/model"
	pb "GOALED/go/pb"
	pbg "GOALED/go/pb/game"
	"context"
	"log"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	GAME_SERVER_ADDRESS = "goaledgameserver-game-1"
	GAME_SERVER_PORT    = "8080"
)

type MatchingServer struct {
	pb.UnimplementedMatchingServiceServer
	Rooms   map[string]*pb.Room
	Streams *model.SyncStreamsModel
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
		Owner:         in.GetOwner(),
		Password:      "",
		Status:        pb.RoomStatus_WAITING,
		MaxPlayer:     in.GetMaxPlayer(),
		CurrentPlayer: 0,
		Players:       make([]*pb.Player, 0),
	}
	log.Println("CreatePublicRoom", room.Status, room.CurrentPlayer, room.MaxPlayer, room.Owner)
	ms.Rooms[room.Id] = room
	return &pb.CreatePublicRoomResponse{
		Room: room,
	}, nil
}

func (ms *MatchingServer) CreatePrivateRoom(ctx context.Context, in *pb.CreatePrivateRoomRequest) (*pb.CreatePrivateRoomResponse, error) {
	room := &pb.Room{
		Id:            uuid.New().String(),
		Name:          in.GetName(),
		Owner:         in.GetOwner(),
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
	if room.Status != pb.RoomStatus_WAITING || room.CurrentPlayer >= room.MaxPlayer {
		log.Println(room.Status, room.CurrentPlayer, room.MaxPlayer)
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
			if room.Status != pb.RoomStatus_WAITING || room.CurrentPlayer >= room.MaxPlayer {
				return &pb.JoinPrivateRoomResponse{
					Room: nil,
				}, nil
			}
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
	if room, ok := ms.Rooms[in.GetRoomId()]; ok {
		for i, player := range room.Players {
			if player.Id == in.GetPlayerId() {
				room.Players = append(room.Players[:i], room.Players[i+1:]...)
				room.CurrentPlayer -= 1
				ms.Streams.RemoveStream(in.GetPlayerId())
				if room.CurrentPlayer == 0 {
					delete(ms.Rooms, room.Id)
				}
				return &pb.LeaveRoomResponse{
					Success: true,
				}, nil
			}
		}
	}
	return &pb.LeaveRoomResponse{
		Success: false,
	}, nil
}

func (ms *MatchingServer) StartGame(ctx context.Context, in *pb.StartGameRequest) (*pb.StartGameResponse, error) {
	conn, err := grpc.Dial(
		GAME_SERVER_ADDRESS+":"+GAME_SERVER_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	if room, ok := ms.Rooms[in.GetRoomId()]; ok {
		player_ids := make([]string, 0)
		for _, player := range room.Players {
			player_ids = append(player_ids, player.Id)
		}
		client := pbg.NewGameServiceClient(conn)
		_, e := client.CreateRoom(ctx, &pbg.Room{
			Id:        in.GetRoomId(),
			Name:      room.GetName(),
			Owner:     room.GetOwner(),
			PlayerIds: player_ids,
		})
		if e != nil {
			log.Fatalf("could not greet: %v", e)
		}

		room.Status = pb.RoomStatus_PLAYING
		for _, player := range room.Players {
			ms.Streams.SendGameStart(player.Id, func(stream pb.MatchingService_GetStartGameStreamServer) {
				stream.Send(&pb.GetStartGameStreamResponse{
					Success: true,
				})
			})
		}
		for _, player := range room.Players {
			ms.Streams.RemoveStream(player.Id)
		}
	}
	return &pb.StartGameResponse{
		Success: true,
	}, nil
}

func (ms *MatchingServer) GetStartGameStream(in *pb.GetStartGameStreamRequest, stream pb.MatchingService_GetStartGameStreamServer) error {
	done := make(chan struct{})
	player_id := in.GetPlayerId()
	ms.Streams.AddGameStartStream(done, player_id, stream)
	<-done
	log.Println("GetStartGameStream done")
	return nil
}

func NewMatchingServer() *MatchingServer {
	return &MatchingServer{
		Rooms:   make(map[string]*pb.Room),
		Streams: model.NewSyncStreamsModel(),
	}
}
