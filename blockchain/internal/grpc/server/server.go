package server

import (
	"blockchain/internal/service"
	"blockchain/internal/transport/medicalpb"
	"context"
	"fmt"
)

type BlockChainService struct {
	medicalpb.UnimplementedMedicalServiceServer
	svc *service.Service
}

func NewBlockChainService(svc *service.Service) *BlockChainService {
	return &BlockChainService{svc: svc}
}

func (s *BlockChainService) GenerateNewAccount(ctx context.Context, _ *medicalpb.Empty) (response *medicalpb.AccountResponse, err error) {
	response = &medicalpb.AccountResponse{}

	pKey, address, err := s.svc.GenerateNewAccount(ctx)
	if err != nil {
		return response, err
	}

	response.Address = address
	response.PrivateKey = pKey

	return response, err
}

func (s *BlockChainService) DepositAccount(ctx context.Context, req *medicalpb.DepositRequest) (response *medicalpb.Empty, err error) {
	err = s.svc.DepositAccount(ctx, req.ToAddress, req.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("can't deposit account: %v", err)
	}

	return response, err
}

func (s *BlockChainService) AddMedicalRecord(ctx context.Context, req *medicalpb.MedicalRecordRequest) (response *medicalpb.TransactionResponse, err error) {
	err = s.svc.AddMedicalRecord(ctx, req.PrivateKey, req.DataHash)
	if err != nil {
		return nil, fmt.Errorf("can't add medical record in blockchain: %v", err)
	}

	return response, err
}

func (s *BlockChainService) GetRecord(ctx context.Context, req *medicalpb.GetRecordRequest) (response *medicalpb.GetRecordResponse, err error) {
	record, err := s.svc.GetRecord(ctx, req.DataHash)
	if err != nil {
		return nil, fmt.Errorf("can't get medical record from blockchain: %v", err)
	}

	response.Owner = record.Owner.Hex()
	response.Timestamp = record.Timestamp.Int64()

	for _, address := range record.Addresses {
		response.AuthorizedAddresses = append(response.GetAuthorizedAddresses(), address.Hex())
	}

	return response, err
}

func (s *BlockChainService) GrantAccess(ctx context.Context, req *medicalpb.AccessRequest) (response *medicalpb.TransactionResponse, err error) {
	tHash, err := s.svc.GrantAccess(ctx, req.Address, req.DataHash, req.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("can't get medical record from blockchain: %v", err)
	}

	response.TransactionHash = tHash.Hex()

	return response, err
}

func (s *BlockChainService) RevokeAccess(ctx context.Context, req *medicalpb.AccessRequest) (response *medicalpb.TransactionResponse, err error) {
	tHash, err := s.svc.RevokeAccess(ctx, req.Address, req.DataHash, req.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("can't get medical record from blockchain: %v", err)
	}

	response.TransactionHash = tHash.Hex()

	return response, err
}
