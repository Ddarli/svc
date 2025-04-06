package service

import (
	"blockchain/config"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"strings"
)

type Service struct {
	client *ethclient.Client
	cfg    config.Config
}

type RecordData struct {
	Owner     common.Address
	Timestamp *big.Int
	Addresses []common.Address
}

func New(cfg config.Config) (*Service, error) {
	client, err := ethclient.Dial(cfg.Host)
	if err != nil {
		log.Fatalf("Ошибка подключения к Ganache: %v", err)
		return nil, err
	}

	return &Service{client: client, cfg: cfg}, nil
}

func (s *Service) GenerateNewAccount(ctx context.Context) (privateKeyHex string, address string, err error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return "", "", err
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	addressBytes := crypto.PubkeyToAddress(privateKey.PublicKey)

	return hex.EncodeToString(privateKeyBytes), addressBytes.Hex(), nil
}

func (s *Service) DepositAccount(ctx context.Context, toAdd, privateKeyHex string) (err error) {
	toAddress := common.HexToAddress(toAdd)

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatalf("Ошибка преобразования приватного ключа: %v", err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatalf("Не удалось получить публичный ключ")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := s.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalf("Ошибка получения nonce: %v", err)
	}

	balance, err := s.client.BalanceAt(context.Background(), fromAddress, nil)
	if err != nil {
		log.Fatalf("Ошибка получения баланса: %v", err)
	}
	fmt.Printf("Баланс отправителя: %s wei\n", balance.String())

	gasPrice, err := s.client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Ошибка получения цены газа: %v", err)
	}

	tx := types.NewTransaction(nonce, toAddress, big.NewInt(1000000000000000000), uint64(21000), gasPrice, nil)

	// Подписываем транзакцию
	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(big.NewInt(1337)), privateKey) // Используйте правильный chain ID
	if err != nil {
		log.Fatalf("Ошибка подписи транзакции: %v", err)
	}

	err = s.client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatalf("Ошибка отправки транзакции: %v", err)
	}

	return err
}

func (s *Service) AddMedicalRecord(ctx context.Context, pKey, hash string) error {
	contractABI, err := abi.JSON(strings.NewReader(addMedicalRecord))
	if err != nil {
		log.Fatal("Ошибка разбора ABI:", err)
	}

	data, err := contractABI.Pack("addMedicalRecord", hash)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := sendTransaction(s.client, pKey, s.cfg.ContractAddress, data, big.NewInt(0))

	fmt.Printf("Транзакция отправлена: %s\n", tx.Hex())

	return nil
}

func (s *Service) GetRecord(ctx context.Context, hash string) (record RecordData, err error) {
	contractAddress := common.HexToAddress(s.cfg.ContractAddress)

	contractABI, err := abi.JSON(strings.NewReader(getRecord))
	if err != nil {
		log.Fatal("Ошибка разбора ABI:", err)
	}

	callData, err := contractABI.Pack("getRecord", hash)
	if err != nil {
		log.Fatal("Ошибка упаковки вызова:", err)
	}

	msg := ethereum.CallMsg{
		To:   &contractAddress,
		Data: callData,
	}

	result, err := s.client.CallContract(context.Background(), msg, nil)
	if err != nil {
		log.Fatal("Ошибка вызова контракта:", err)
	}

	var owner common.Address
	var timestamp *big.Int
	var addresses []common.Address

	err = contractABI.UnpackIntoInterface(&[]interface{}{&owner, &timestamp, &addresses}, "getRecord", result)
	if err != nil {
		log.Fatal("Ошибка распаковки данных:", err)
	}

	return RecordData{
		Owner:     owner,
		Timestamp: timestamp,
		Addresses: addresses,
	}, err
}

func (s *Service) GrantAccess(ctx context.Context, address, dataHash, pKey string) (txHash common.Hash, err error) {
	contractABI, err := abi.JSON(strings.NewReader(grantAccess))
	if err != nil {
		log.Fatal("Ошибка разбора ABI:", err)
	}

	addressToAuthorize := common.HexToAddress(address)

	data, err := contractABI.Pack("authorizeAddress", dataHash, addressToAuthorize)
	if err != nil {
		log.Fatalf("Ошибка упаковки данных: %v", err)
	}

	txHash, err = sendTransaction(s.client, pKey, s.cfg.ContractAddress, data, big.NewInt(0))
	if err != nil {
		log.Fatalf("Ошибка отправки транзакции: %v", err)
	}

	return txHash, err
}

func (s *Service) RevokeAccess(ctx context.Context, address, dataHash, pKey string) (txHash common.Hash, err error) {
	contractABI, err := abi.JSON(strings.NewReader(revokeAccess))
	if err != nil {
		log.Fatal("Ошибка разбора ABI:", err)
	}

	addressToRevoke := common.HexToAddress(address)

	data, err := contractABI.Pack("revokeAddress", dataHash, addressToRevoke)
	if err != nil {
		log.Fatalf("Ошибка упаковки данных: %v", err)
	}

	txHash, err = sendTransaction(s.client, pKey, s.cfg.ContractAddress, data, big.NewInt(0))
	if err != nil {
		log.Fatalf("Ошибка отправки транзакции: %v", err)
	}

	return txHash, err
}

func sendTransaction(client *ethclient.Client, pKey, contrAddress string, data []byte, value *big.Int) (common.Hash, error) {
	privateKey, err := crypto.HexToECDSA(pKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("не удалось получить публичный ключ")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasTipCap, _ := client.SuggestGasTipCap(context.Background())
	gasFeeCap, _ := client.SuggestGasPrice(context.Background())
	chainID, _ := client.ChainID(context.Background())

	contractAddress := common.HexToAddress(contrAddress)

	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
		Gas:       300000,
		To:        &contractAddress,
		Value:     value,
		Data:      data,
	})

	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chainID), privateKey)
	if err != nil {
		log.Fatal("Ошибка подписи:", err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal("Ошибка отправки транзакции:", err)
	}

	return signedTx.Hash(), nil
}
