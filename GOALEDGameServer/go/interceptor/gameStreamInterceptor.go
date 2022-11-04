package interceptor

import (
	"GOALED/go/pb"
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"log"
	"sync/atomic"
	"time"

	"google.golang.org/grpc"
)

var (
	requestObjectSize      uint64
	responseObjectSize     uint64
	requestPlayerDataSize  uint64
	responsePlayerDataSize uint64
	kb                     = 1024 * 8.0
	mb                     = 1024 * 1024 * 8.0
)

func PacketSizeAnalysis() {
	go func() {
		var lastRequestObjectSize uint64
		var lastResponseObjectSize uint64
		var lastRequestPlayerDataSize uint64
		var lastResponsePlayerDataSize uint64
		for {
			log.Printf("Object     In: %.4f kb/s total: %.4f mb Out: %.4f kb/s total: %.4f mb ", float64(requestObjectSize-lastRequestObjectSize)/kb, float64(requestObjectSize)/mb, float64(responseObjectSize-lastResponseObjectSize)/kb, float64(responseObjectSize)/mb)
			log.Printf("PlayerData In: %.4f kb/s total: %.4f mb Out: %.4f kb/s total: %.4f mb ", float64(requestPlayerDataSize-lastRequestPlayerDataSize)/kb, float64(requestPlayerDataSize)/mb, float64(responsePlayerDataSize-lastResponsePlayerDataSize)/kb, float64(responsePlayerDataSize)/mb)
			log.Println("============================================================================")
			lastRequestObjectSize = requestObjectSize
			lastResponseObjectSize = responseObjectSize
			lastRequestPlayerDataSize = requestPlayerDataSize
			lastResponsePlayerDataSize = responsePlayerDataSize
			time.Sleep(time.Second * 1)
		}
	}()
}

func binarySize(val interface{}) int {
	var buff bytes.Buffer
	gob.Register(pb.ParamValue_BoolValue{})
	enc := gob.NewEncoder(&buff)
	err := enc.Encode(val)
	if err != nil {
		log.Print(err)
		return 0
	}
	return binary.Size(buff.Bytes())
}

func GameStreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	if info.FullMethod == "/GameService.GameService/SendObject" || info.FullMethod == "/GameService.GameService/SyncObject" {
		wrapped := &GameStreamWrapper{ss}
		return handler(srv, wrapped)
	} else if info.FullMethod == "/GameService.GameService/SendPlayerData" || info.FullMethod == "/GameService.GameService/SyncPlayerData" {
		wrapped := &PlayerDataStreamWrapper{ss}
		return handler(srv, wrapped)
	} else {
		return handler(srv, ss)
	}
}

type GameStreamWrapper struct {
	grpc.ServerStream
}

func (i *GameStreamWrapper) RecvMsg(m interface{}) error {
	err := i.ServerStream.RecvMsg(m)
	atomic.AddUint64(&requestObjectSize, uint64(binarySize(m)))
	return err
}

func (i *GameStreamWrapper) SendMsg(m interface{}) error {
	err := i.ServerStream.SendMsg(m)
	atomic.AddUint64(&responseObjectSize, uint64(binarySize(m)))
	return err
}

type PlayerDataStreamWrapper struct {
	grpc.ServerStream
}

func (i *PlayerDataStreamWrapper) RecvMsg(m interface{}) error {
	err := i.ServerStream.RecvMsg(m)
	atomic.AddUint64(&requestPlayerDataSize, uint64(binarySize(m)))
	return err
}

func (i *PlayerDataStreamWrapper) SendMsg(m interface{}) error {
	err := i.ServerStream.SendMsg(m)
	atomic.AddUint64(&responsePlayerDataSize, uint64(binarySize(m)))
	return err
}
