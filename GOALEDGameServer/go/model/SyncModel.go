package model

import (
	pb "GOALED/go/pb"
	"log"
	"sync"
	"time"

	"github.com/fatih/color"
)

type SyncStreamModel struct {
	streams    map[string]pb.GameSyncService_RecvGameObjectServer
	streamDone map[string]chan struct{}
	mux        sync.Mutex
}

func NewSyncStreamModel() *SyncStreamModel {
	return &SyncStreamModel{
		streams:    make(map[string]pb.GameSyncService_RecvGameObjectServer),
		streamDone: make(map[string]chan struct{}),
	}
}

func (r *SyncStreamModel) AddStream(done chan struct{}, playerId string, stream pb.GameSyncService_RecvGameObjectServer) {
	r.mux.Lock()
	defer r.mux.Unlock()
	r.streams[playerId] = stream
	r.streamDone[playerId] = done
}

func (r *SyncStreamModel) RemoveStream(playerId string) {
	r.mux.Lock()
	defer r.mux.Unlock()
	delete(r.streams, playerId)
	close(r.streamDone[playerId])
	delete(r.streamDone, playerId)
}

func (r *SyncStreamModel) RangeStreams(f func(stream pb.GameSyncService_RecvGameObjectServer)) {
	r.mux.Lock()
	defer r.mux.Unlock()
	for _, stream := range r.streams {
		f(stream)
	}
}

type SyncObjectModel struct {
	objects map[string]*pb.GameObject
	mux     sync.Mutex
}

func NewSyncObjectModel() *SyncObjectModel {
	return &SyncObjectModel{
		objects: make(map[string]*pb.GameObject),
	}
}

func (r *SyncObjectModel) AddObjects(objects []*pb.GameObject) {
	r.mux.Lock()
	defer r.mux.Unlock()
	for _, object := range objects {
		r.objects[object.GetObjectId()] = object
	}
}

func (r *SyncObjectModel) RemoveObject(objectId string) {
	r.mux.Lock()
	defer r.mux.Unlock()
	delete(r.objects, objectId)
}

func (r *SyncObjectModel) GetObjects() []*pb.GameObject {
	r.mux.Lock()
	defer r.mux.Unlock()
	objects := make([]*pb.GameObject, 0, len(r.objects))
	for _, object := range r.objects {
		objects = append(objects, object)
	}
	return objects
}

type SyncModel struct {
	roomId       string
	roomName     string
	hostPLayerId string
	players      map[string]*pb.Player
	syncStream   *SyncStreamModel
	syncObject   *SyncObjectModel
	done         chan struct{}
	count        int
	beforeTime   time.Time
}

func NewSyncModel(roomId string, roomName string, hostPlayerId string) *SyncModel {
	room := SyncModel{
		roomId:       roomId,
		roomName:     roomName,
		hostPLayerId: hostPlayerId,
		players:      make(map[string]*pb.Player),
		syncStream:   NewSyncStreamModel(),
		syncObject:   NewSyncObjectModel(),
		done:         make(chan struct{}),
		count:        0,
		beforeTime:   time.Now(),
	}
	go room.Run()
	return &room
}

func (r *SyncModel) Update() {
	r.count++
	if time.Now().Sub(r.beforeTime) > 1*time.Second {
		r.beforeTime = time.Now()
		// log.Println(color.GreenString("sync count: %d", r.count))
		log.Printf("[%v] room %v Sync %v object", color.GreenString("Update"), r.roomId, len(r.syncObject.GetObjects()))
		r.count = 0
	}
}

func (r *SyncModel) Run() {
	log.Printf("[%v] Sync %v is running", color.GreenString("Run"), r.roomId)
	ticker := time.NewTicker(1 * time.Second / 20)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			r.SendObject(r.syncObject.GetObjects())
			r.Update()
		case <-r.done:
			log.Printf("[%v] Room %v is closed", color.GreenString("Close"), r.roomId)
			return
		}
	}
}

func (r *SyncModel) Close() {
	for _, done := range r.syncStream.streamDone {
		close(done)
	}
	close(r.done)
}

func (r *SyncModel) SendObject(objects []*pb.GameObject) {
	r.syncStream.RangeStreams(func(stream pb.GameSyncService_RecvGameObjectServer) {
		stream.Send(&pb.RecvGameObjectResponse{
			GameObjects: objects,
		})
	})
}

func (r *SyncModel) AddStream(done chan struct{}, playerId string, stream pb.GameSyncService_RecvGameObjectServer) {
	r.syncStream.AddStream(done, playerId, stream)
}

func (r *SyncModel) RemoveStream(playerId string) {
	r.syncStream.RemoveStream(playerId)
}

func (r *SyncModel) UpdateObjects(objects []*pb.GameObject) {
	r.syncObject.AddObjects(objects)
}

func (r *SyncModel) GetPlayers() map[string]*pb.Player {
	return r.players
}
