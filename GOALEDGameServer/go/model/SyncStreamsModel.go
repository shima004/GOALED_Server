package model

import (
	pb "GOALED/go/pb"
	"sync"
)

type SyncStreamsModel struct {
	ObjectStream     map[string]pb.GameService_SyncObjectServer
	PlayerDataStream map[string]pb.GameService_SyncPlayerDataServer
	DoneChannel      map[string]chan struct{}
	mux              sync.Mutex
}

func NewSyncStreamsModel() *SyncStreamsModel {
	return &SyncStreamsModel{
		ObjectStream:     make(map[string]pb.GameService_SyncObjectServer),
		PlayerDataStream: make(map[string]pb.GameService_SyncPlayerDataServer),
		DoneChannel:      make(map[string]chan struct{}),
	}
}

func (ssm *SyncStreamsModel) AddObjectStream(done chan struct{}, playerId string, stream pb.GameService_SyncObjectServer) {
	ssm.mux.Lock()
	defer ssm.mux.Unlock()
	ssm.ObjectStream[playerId] = stream
	ssm.DoneChannel[playerId] = done
}

func (ssm *SyncStreamsModel) AddPlayerDataStream(done chan struct{}, playerId string, stream pb.GameService_SyncPlayerDataServer) {
	ssm.mux.Lock()
	defer ssm.mux.Unlock()
	ssm.PlayerDataStream[playerId] = stream
	ssm.DoneChannel[playerId] = done
}

func (ssm *SyncStreamsModel) RemoveStream(playerId string) {
	ssm.mux.Lock()
	defer ssm.mux.Unlock()
	delete(ssm.ObjectStream, playerId)
	delete(ssm.PlayerDataStream, playerId)
	close(ssm.DoneChannel[playerId])
	delete(ssm.DoneChannel, playerId)
}

func (ssm *SyncStreamsModel) RangeObjectStream(f func(stream pb.GameService_SyncObjectServer)) {
	ssm.mux.Lock()
	defer ssm.mux.Unlock()
	for _, stream := range ssm.ObjectStream {
		f(stream)
	}
}

func (ssm *SyncStreamsModel) RangePlayerDataStream(f func(stream pb.GameService_SyncPlayerDataServer)) {
	ssm.mux.Lock()
	defer ssm.mux.Unlock()
	for _, stream := range ssm.PlayerDataStream {
		f(stream)
	}
}
