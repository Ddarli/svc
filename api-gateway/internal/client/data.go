package client

import (
	"context"
	"fmt"
	pb "github.com/Ddarli/svc/gateway/pkg/proto"
	"google.golang.org/grpc"
	"log/slog"
)

type DataClient struct {
	client pb.FileServiceClient
}

func NewDataClient(ctx context.Context, cfg *Config) *DataClient {
	client, err := createDataClient(ctx, cfg.DataClient.Port)
	if err != nil {
		slog.Error(err.Error())
		return nil
	}

	return &DataClient{client: client}
}

func (c *DataClient) UploadFile(ctx context.Context, in *pb.UploadUserFileRequest) (*pb.UploadUserFileResponse, error) {
	resp, err := c.client.UploadUserFile(ctx, in)
	if err != nil {
		return nil, fmt.Errorf("error upload file grpc client: %v", err)
	}

	return resp, nil
}

func (c *DataClient) DownloadFile(ctx context.Context, in *pb.DownloadUserFileRequest) (*pb.DownloadUserFileResponse, error) {
	resp, err := c.client.DownloadUserFile(ctx, in)
	if err != nil {
		return nil, fmt.Errorf("error download file grpc client: %v", err)
	}

	return resp, nil
}

func (c *DataClient) ListUserFiles(ctx context.Context, in *pb.ListUserFilesRequest) (*pb.ListUserFilesResponse, error) {
	resp, err := c.client.ListUserFiles(ctx, in)
	if err != nil {
		return nil, fmt.Errorf("error list user files grpc client: %v", err)
	}

	return resp, nil
}

func createDataClient(ctx context.Context, port string) (pb.FileServiceClient, error) {
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := pb.NewFileServiceClient(conn)

	slog.InfoContext(ctx, "listening data client on", "port", port)

	return client, nil
}
