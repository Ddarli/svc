package app

import (
	"context"
	"data-processor/config"
	"data-processor/internal/repository"
	"data-processor/internal/server"
	"data-processor/internal/service"
	"data-processor/internal/storage"
	"data-processor/pkg/data-processor/proto/fileservice"
	"google.golang.org/grpc"
	"log"
	"log/slog"
	"net"
)

func Run(ctx context.Context, configuration *config.Config) {
	db := repository.ConnectDB(ctx, *configuration)

	repo := repository.NewRepo(db)
	st := storage.New(*configuration)

	srv := service.New(st, repo)

	listener, err := net.Listen("tcp", configuration.Server.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	authServer := server.New(srv)

	fileservice.RegisterFileServiceServer(grpcServer, authServer)

	slog.Info("gRPC server is running", "port", configuration.Server.Port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
