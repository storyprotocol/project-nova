package main

import (
	"context"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/project-nova/backend/pkg/gateway"
	"github.com/project-nova/backend/pkg/logger"
)

func main() {
	Logger, err := logger.InitLogger(logger.Levels.Info)
	if err != nil {
		logger.Fatalf("Failed to init logger, error: %v", err)
	}
	defer func() {
		_ = Logger.Sync()
	}()

	apiGateway := gateway.NewApiHttpGateway("http://localhost:8090")

	client, err := ethclient.Dial("wss://eth-goerli.g.alchemy.com/v2/RWhchAQNylFZLnnPO7Rmj3b4T4uZIFuO")
	if err != nil {
		logger.Fatal(err)
	}

	contractAddress := common.HexToAddress("0x041c1fb84f8673F1Fc40bE20d45B3b012d37769b")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
		Topics: [][]common.Hash{
			{common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")},
		},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info("Streamer starting")
	for {
		select {
		case err := <-sub.Err():
			logger.Fatal(err)
		case vlog := <-logs:
			logger.Infof("vLog: %v", vlog)

			tokenId := vlog.Topics[3].Big().Uint64()
			collectionAddress := vlog.Address.String()

			err = apiGateway.UpdateNftOwner(int(tokenId), collectionAddress)
			if err != nil {
				logger.Errorf("Failed to update nft owner: %v", err)
			}
		}
	}
}
