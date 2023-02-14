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
	// Transfer Topic's Hash Identifier
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
			/* Sample log
			{
			  "address": "0x001c1fb84f8673f1fc40be20d45b3b012d300000",
			  "topics": [
			    "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
			    "0x0000000000000000000000000000000000000000000000000000000000000000",
			    "0x000000000000000000000000000b179bfd31635a387fc205f341d1fac6797327",
			    "0x0000000000000000000000000000000000000000000000000000000000000022"
			  ],
			  "data": "0x",
			  "blockNumber": "0x818800",
			  "transactionHash": "0x09165dec5964fce09e35c4f7a83caf49b9b3018935f9158b60f6512daae00000",
			  "transactionIndex": "0xc",
			  "blockHash": "0x0f2bf435e82f9772e9611c1f7918130587c6fc0d5b21e2eb8d5e88d918f00000",
			  "logIndex": "0x1a",
			  "removed": false
			}
			*/
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
