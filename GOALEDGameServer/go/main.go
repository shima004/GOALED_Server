package main

import (
	"fmt"
	"log"
	"net"

	pb "GOALED/go/pb"
	service "GOALED/go/service"

	// grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	// grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	// grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	// "go.uber.org/zap"
	// "go.uber.org/zap/zapcore"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// opts := []grpc_zap.Option{
	// 	grpc_zap.WithDurationField(func(duration time.Duration) zapcore.Field {
	// 		return zap.Int64("grpc.time_ns", duration.Nanoseconds())
	// 	}),
	// }
	// zapLogger, _ := zap.NewProduction()
	// grpc_zap.ReplaceGrpcLogger(zapLogger)

	grpcServer := grpc.NewServer(
	// grpc_middleware.WithUnaryServerChain(
	// 	grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
	// 	grpc_zap.UnaryServerInterceptor(zapLogger, opts...),
	// ),
	)

	pb.RegisterGameServiceServer(grpcServer, service.NewGameServer())
	reflection.Register(grpcServer)

	fmt.Println("Server is running on port 8080")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
