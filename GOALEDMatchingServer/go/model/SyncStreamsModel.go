package model

import (
	pb "GOALED/go/pb"
	"sync"
)

type SyncStreamsModel struct {
	GameStartStream map[string]pb.MatchingService_GetStartGameStreamServer
	DoneChannel     map[string]chan struct{}
	mux             sync.Mutex
}

func NewSyncStreamsModel() *SyncStreamsModel {
	return &SyncStreamsModel{
		GameStartStream: make(map[string]pb.MatchingService_GetStartGameStreamServer),
		DoneChannel:     make(map[string]chan struct{}),
	}
}

func (ssm *SyncStreamsModel) AddGameStartStream(done chan struct{}, playerId string, stream pb.MatchingService_GetStartGameStreamServer) {
	ssm.mux.Lock()
	defer ssm.mux.Unlock()
	ssm.GameStartStream[playerId] = stream
	ssm.DoneChannel[playerId] = done
}

func (ssm *SyncStreamsModel) RemoveStream(playerId string) {
	ssm.mux.Lock()
	defer ssm.mux.Unlock()
	delete(ssm.GameStartStream, playerId)
	close(ssm.DoneChannel[playerId])
	delete(ssm.DoneChannel, playerId)
}

func (ssm *SyncStreamsModel) SendGameStart(player_id string, f func(stream pb.MatchingService_GetStartGameStreamServer)) {
	ssm.mux.Lock()
	defer ssm.mux.Unlock()
	if stream, ok := ssm.GameStartStream[player_id]; ok {
		f(stream)
	}
}
