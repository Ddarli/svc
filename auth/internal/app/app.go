package app

import (
	"auth/conf"
	"auth/internal/client/blockchain"
	auth "auth/internal/grpc"
	"auth/internal/repo"
	"auth/internal/router"
	"auth/internal/service"
	pb "auth/pkg/proto"
	"context"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"log"
	"log/slog"
	"net"
)

func Run(ctx context.Context, configuration *conf.Configuration) {
	app := fiber.New()
	db := repo.ConnectDB(ctx, *configuration)

	repository := repo.NewRepo(db)
	blockchainClient := blockchain.New(ctx, configuration.BlockchainClient.Port)

	service := service.NewService(repository, configuration.ServiceConf, blockchainClient)

	router := router.NewRouter(app, *service)
	router.RegisterRoutes(configuration)

	listener, err := net.Listen("tcp", configuration.Server.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	authServer := auth.NewAuthServer(service)

	pb.RegisterAuthServiceServer(grpcServer, authServer)

	slog.Info("gRPC server is running", "port", configuration.Server.Port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
