package app

import (
	"blockchain/config"
	"blockchain/internal/grpc/server"
	"blockchain/internal/service"
	"blockchain/internal/transport/medicalpb"
	"context"
	"google.golang.org/grpc"
	"log"
	"log/slog"
	"net"
)

func Run(ctx context.Context, configuration config.Config) {
	svc, err := service.New(configuration)
	if err != nil {
		log.Fatalf("error creating service")
	}

	listener, err := net.Listen("tcp", configuration.Server.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	blockchainServer := server.NewBlockChainService(svc)

	medicalpb.RegisterMedicalServiceServer(grpcServer, blockchainServer)

	slog.Info("gRPC server is running", "port", configuration.Server.Port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
