package main

import (
	"fmt"
	"log"
	"net"

	pb "GOALED/go/pb"
	service "GOALED/go/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterGameSyncServiceServer(grpcServer, service.NewGameSyncServer())
	pb.RegisterRoomServiceServer(grpcServer, service.NewRoomServer())
	reflection.Register(grpcServer)

	fmt.Println("Now listening...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
