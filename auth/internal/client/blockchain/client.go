package blockchain

import (
	"auth/pkg/transport/medicalpb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log/slog"
)

type BlockChainClient struct {
	client medicalpb.MedicalServiceClient
}

func New(ctx context.Context, port string) *BlockChainClient {
	client, err := createClient(ctx, port)
	if err != nil {
		slog.Error(err.Error())
		return nil
	}

	return &BlockChainClient{client: client}
}

func (c *BlockChainClient) GenerateNewAccount(ctx context.Context, req *medicalpb.Empty) (response *medicalpb.AccountResponse, err error) {
	response = &medicalpb.AccountResponse{}

	res, err := c.client.GenerateNewAccount(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to generate new account: %v", err)
	}

	response.PrivateKey = res.GetPrivateKey()
	response.Address = res.GetAddress()

	return response, err
}

func createClient(ctx context.Context, port string) (medicalpb.MedicalServiceClient, error) {
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := medicalpb.NewMedicalServiceClient(conn)

	slog.InfoContext(ctx, "run auth client on", "port", port)

	return client, nil
}
