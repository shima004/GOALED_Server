package model

import (
	pb "GOALED/go/pb"
)

type GameModel struct {
	streams           SyncStreamsModel
	PlayerDataChannel chan []*pb.PlayerData
	ObjectsChannel    chan []*pb.Object
	DoneChannel       chan struct{}
	Room              *pb.Room
}

func (gm *GameModel) Run() {
	for {
		select {
		case playerData := <-gm.PlayerDataChannel:
			res := &pb.SyncPlayerDataResponse{
				PlayerData: playerData,
			}
			gm.streams.RangePlayerDataStream(func(stream pb.GameService_SyncPlayerDataServer) {
				stream.Send(res)
			})
		case object := <-gm.ObjectsChannel:
			res := pb.SyncObjectResponse{
				Object: object,
			}
			gm.streams.RangeObjectStream(func(stream pb.GameService_SyncObjectServer) {
				stream.Send(&res)
			})
		case <-gm.DoneChannel:
			return
		}
	}
}

func (gm *GameModel) Close() {
	for _, done := range gm.streams.DoneChannels {
		for _, d := range done {
			close(d)
		}
	}
	close(gm.DoneChannel)
}

func NewGameModel(room *pb.Room) *GameModel {
	gamemodel := &GameModel{
		streams:           *NewSyncStreamsModel(),
		PlayerDataChannel: make(chan []*pb.PlayerData),
		ObjectsChannel:    make(chan []*pb.Object),
		DoneChannel:       make(chan struct{}),
		Room:              room,
	}
	go gamemodel.Run()
	return gamemodel
}

func (gm *GameModel) GetStreams() *SyncStreamsModel {
	return &gm.streams
}

func (gm *GameModel) AddPlayerData(playerData []*pb.PlayerData) {
	gm.PlayerDataChannel <- playerData
}

func (gm *GameModel) AddObject(object []*pb.Object) {
	gm.ObjectsChannel <- object
}
