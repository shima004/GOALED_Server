package model

import (
	pb "GOALED/go/pb"
)

type SyncStreamsModel struct {
	ObjectStream     pb.GameService_SyncObjectServer
	PlayerDataStream pb.GameService_SyncPlayerDataServer
}
