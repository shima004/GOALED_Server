package model

import (
	pb "GOALED/go/pb"
	"log"
	"sync"

	"github.com/fatih/color"
)

type RoomSyncStreamModel struct {
	streams    map[string]pb.RoomService_SyncRoomUsersServer
	streamDone map[string]chan struct{}
	mux        sync.Mutex
}

func NewRoomSyncStreamModel() *RoomSyncStreamModel {
	return &RoomSyncStreamModel{
		streams:    make(map[string]pb.RoomService_SyncRoomUsersServer),
		streamDone: make(map[string]chan struct{}),
	}
}

func (r *RoomSyncStreamModel) AddStream(done chan struct{}, playerId string, stream pb.RoomService_SyncRoomUsersServer) {
	r.mux.Lock()
	defer r.mux.Unlock()
	r.streams[playerId] = stream
	r.streamDone[playerId] = done
}

func (r *RoomSyncStreamModel) RemoveStream(playerId string) {
	r.mux.Lock()
	defer r.mux.Unlock()
	delete(r.streams, playerId)
	close(r.streamDone[playerId])
	delete(r.streamDone, playerId)
}

func (r *RoomSyncStreamModel) RangeStreams(f func(stream pb.RoomService_SyncRoomUsersServer)) {
	r.mux.Lock()
	defer r.mux.Unlock()
	for _, stream := range r.streams {
		f(stream)
	}
}

type RoomModel struct {
	roomId         string
	roomName       string
	hostPLayerId   string
	maxPlayer      int32
	status         pb.RoomStatus
	isPrivate      bool
	players        map[string]*pb.Player
	roomSyncStream *RoomSyncStreamModel
	addPlayer      chan *pb.Player
	removePlayer   chan string
	done           chan struct{}
}

func NewRoomModel(roomId string, roomName string, hostPlayerId string, maxPlayer int32, isPrivate bool) *RoomModel {
	room := RoomModel{
		roomId:         roomId,
		roomName:       roomName,
		hostPLayerId:   hostPlayerId,
		maxPlayer:      maxPlayer,
		status:         pb.RoomStatus_WAITING,
		isPrivate:      isPrivate,
		players:        make(map[string]*pb.Player),
		roomSyncStream: NewRoomSyncStreamModel(),
		addPlayer:      make(chan *pb.Player),
		removePlayer:   make(chan string),
		done:           make(chan struct{}),
	}
	go room.Run()
	return &room
}

func (r *RoomModel) Run() {
	log.Printf("[%v] Room %v is running", color.GreenString("Run"), r.roomId)
	for {
		select {
		case player := <-r.addPlayer:
			log.Printf("[%v] Room %v add player %v", color.BlueString("Add"), r.roomId, player.PlayerId)
			r.SendSyncRoomUser(&pb.SyncRoomUsersResponse{
				Player: player,
				IsLeft: false,
			})
			r.players[player.PlayerId] = player
		case playerId := <-r.removePlayer:
			log.Printf("[%v] Room %v remove player %v", color.BlueString("Remove"), r.roomId, playerId)
			r.RemoveStream(playerId)
			r.SendSyncRoomUser(&pb.SyncRoomUsersResponse{
				Player: r.players[playerId],
				IsLeft: true,
			})
			delete(r.players, playerId)

		case <-r.done:
			log.Printf("[%v] Room %v is closed", color.GreenString("Close"), r.roomId)
			return
		}
	}
}

func (r *RoomModel) Close() {
	for _, done := range r.roomSyncStream.streamDone {
		close(done)
	}
	close(r.done)
}

func (r *RoomModel) ToRoom() *pb.Room {
	return &pb.Room{
		RoomId:        r.roomId,
		RoomName:      r.roomName,
		HostPlayerId:  r.hostPLayerId,
		MaxPlayer:     r.maxPlayer,
		CurrentPlayer: int32(len(r.players)),
		Status:        r.status,
	}
}

func (r *RoomModel) SendSyncRoomUser(syncRoomUsersResponse *pb.SyncRoomUsersResponse) {
	r.roomSyncStream.RangeStreams(func(stream pb.RoomService_SyncRoomUsersServer) {
		stream.Send(syncRoomUsersResponse)
	})
}

func (r *RoomModel) AddPlayer(player *pb.Player) {
	r.addPlayer <- player
}

func (r *RoomModel) RemovePlayer(playerId string) {
	r.removePlayer <- playerId
}

func (r *RoomModel) AddStream(done chan struct{}, playerId string, stream pb.RoomService_SyncRoomUsersServer) {
	for _, player := range r.players {
		stream.Send(&pb.SyncRoomUsersResponse{
			Player: player,
			IsLeft: false,
		})
	}
	r.roomSyncStream.AddStream(done, playerId, stream)
}

func (r *RoomModel) RemoveStream(playerId string) {
	r.roomSyncStream.RemoveStream(playerId)
}

func (r *RoomModel) GetPlayers() map[string]*pb.Player {
	return r.players
}
