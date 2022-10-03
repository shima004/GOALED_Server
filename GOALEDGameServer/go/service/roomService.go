package service

import (
	model "GOALED/go/model"
	pb "GOALED/go/pb"
	"context"
	"log"

	"github.com/fatih/color"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RoomServer struct {
	pb.UnimplementedRoomServiceServer
	rooms map[string]*model.RoomModel
}

func (rs *RoomServer) CreateRoom(ctx context.Context, in *pb.CreateRoomRequest) (*pb.CreateRoomResponse, error) {
	roomId := uuid.New().String()
	room := model.NewRoomModel(roomId, in.GetRoomName(), in.GetPlayerId(), 4, false)
	rs.rooms[roomId] = room
	log.Printf("[%v] Room %v created", color.BlueString("Create"), roomId)
	return &pb.CreateRoomResponse{RoomId: roomId}, nil
}

func (rs *RoomServer) CreatePrivateRoom(ctx context.Context, in *pb.CreatePrivateRoomRequest) (*pb.CreatePrivateRoomResponse, error) {
	if rs.rooms[in.GetRoomId()] != nil {
		return &pb.CreatePrivateRoomResponse{RoomId: in.GetRoomId()}, nil
	}
	room := model.NewRoomModel(in.GetRoomId(), in.GetRoomName(), in.GetPlayerId(), 4, true)
	rs.rooms[in.GetRoomId()] = room
	log.Printf("[%v] Room(Private) %v created", color.BlueString("Create"), in.GetRoomId())
	return &pb.CreatePrivateRoomResponse{RoomId: in.GetRoomId()}, nil
}

func (rs *RoomServer) GetRooms(ctx context.Context, in *pb.GetRoomsRequest) (*pb.GetRoomsResponse, error) {
	rooms := make([]*pb.Room, 0)
	for _, room := range rs.rooms {
		rooms = append(rooms, room.ToRoom())
	}
	log.Printf("[%v] %v Rooms sent", color.BlueString("Get"), len(rooms))
	return &pb.GetRoomsResponse{Rooms: rooms}, nil
}

func (rs *RoomServer) JoinRoom(ctx context.Context, in *pb.JoinRoomRequest) (*pb.JoinRoomResponse, error) {
	room, ok := rs.rooms[in.GetRoomId()]
	if !ok {
		log.Printf("[%v] Room %v not found", color.RedString("Join"), in.GetRoomId())
		return nil, status.Errorf(codes.NotFound, "Room %v not found", in.GetRoomId())
	}
	room.AddPlayer(in.GetPlayer())
	log.Printf("[%v] Room %v found", color.BlueString("Join"), in.GetRoomId())
	return &pb.JoinRoomResponse{}, nil
}

func (rs *RoomServer) LeaveRoom(ctx context.Context, in *pb.LeaveRoomRequest) (*pb.LeaveRoomResponse, error) {
	room, ok := rs.rooms[in.GetRoomId()]
	if !ok {
		log.Printf("[%v] Room %v not found", color.RedString("Leave"), in.GetRoomId())
		return nil, status.Errorf(codes.NotFound, "Room %v not found", in.GetRoomId())
	}
	room.RemovePlayer(in.GetPlayerId())
	log.Printf("%v", room.GetPlayers())
	if len(room.GetPlayers()) == 0 {
		room.Close()
		delete(rs.rooms, in.GetRoomId())
	}
	log.Printf("[%v] Room %v found", color.BlueString("Leave"), in.GetRoomId())
	return &pb.LeaveRoomResponse{}, nil
}

func (rs *RoomServer) GetPlayerId(ctx context.Context, in *pb.GetPlayerIDRequest) (*pb.GetPlayerIDResponse, error) {
	playerId := uuid.New().String()
	log.Printf("[%v] Player %v created", color.BlueString("Create"), playerId)
	return &pb.GetPlayerIDResponse{PlayerId: playerId}, nil
}

func (rs *RoomServer) SyncRoomUsers(in *pb.SyncRoomUsersRequest, stream pb.RoomService_SyncRoomUsersServer) error {
	room, ok := rs.rooms[in.GetRoomId()]
	if !ok {
		log.Printf("[%v] Room %v not found", color.RedString("Sync"), in.GetRoomId())
		return status.Errorf(codes.NotFound, "Room %v not found", in.GetRoomId())
	}
	done := make(chan struct{})
	room.AddStream(done, in.GetPlayerId(), stream)
	log.Printf("[%v] Room %v found player %v add stream", color.BlueString("Sync"), in.GetRoomId(), in.GetPlayerId())
	<-done
	log.Printf("[%v] Room %v found player %v end stream", color.BlueString("Sync"), in.GetRoomId(), in.GetPlayerId())
	return nil
}

func NewRoomServer() *RoomServer {
	return &RoomServer{
		rooms: make(map[string]*model.RoomModel),
	}
}
