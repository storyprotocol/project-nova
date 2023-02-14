package main

import (
	"context"
	"flag"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/project-nova/backend/api/internal/config"
	"github.com/project-nova/backend/pkg/gateway"
	"github.com/project-nova/backend/pkg/logger"
	"github.com/project-nova/backend/pkg/utils"
)

const (
	TransferTopic = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
)

func main() {
	flag.Parse()

	Logger, err := logger.InitLogger(logger.Levels.Info)
	if err != nil {
		logger.Fatalf("Failed to init logger, error: %v", err)
	}
	defer func() {
		_ = Logger.Sync()
	}()

	cfg, err := config.GetStreamerConfig()
	if err != nil {
		logger.Fatalf("Failed to init config, error: %v", err)
	}

	apiGateway := gateway.NewApiHttpGateway(cfg.ApiGatewayUrl)

	client, err := ethclient.Dial(cfg.ProviderWebsocket)
	if err != nil {
		logger.Fatal(err)
	}

	contractAddresses := []common.Address{}
	for _, address := range cfg.MonitorAddresses {
		contractAddresses = append(contractAddresses, common.HexToAddress(address))
	}

	query := ethereum.FilterQuery{
		Addresses: contractAddresses,
		Topics: [][]common.Hash{
			{common.HexToHash(TransferTopic)},
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
			collectionAddress := vlog.Address.String()
			fromAddress := vlog.Topics[1].String()
			toAddress := vlog.Topics[2].String()
			tokenId := vlog.Topics[3].Big().Uint64()

			if utils.IsZeroAddress(fromAddress) { // Mint
				err = apiGateway.CreateNftRecord(int(tokenId), collectionAddress)
				if err != nil {
					logger.Errorf("Failed to create nft record: %v", err)
				}
			} else if utils.IsZeroAddress(toAddress) { // Burn
				err = apiGateway.DeleteNftRecord(int(tokenId), collectionAddress)
				if err != nil {
					logger.Errorf("Failed to delete nft record: %v", err)
				}
			} else { // Transfer
				err = apiGateway.UpdateNftOwner(int(tokenId), collectionAddress)
				if err != nil {
					logger.Errorf("Failed to update nft owner: %v", err)
				}
			}

		}
	}
}
