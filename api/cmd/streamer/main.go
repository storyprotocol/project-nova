package main

import (
	"context"
	"encoding/base64"
	"flag"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/project-nova/backend/api/internal/config"
	"github.com/project-nova/backend/pkg/gateway"
	"github.com/project-nova/backend/pkg/keymanagement"
	"github.com/project-nova/backend/pkg/logger"
	"github.com/project-nova/backend/pkg/utils"
)

const (
	// Transfer Topic's Hash Identifier
	TransferTopic = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"

	// BackstoryCreated's Hash Identifier
	BackstoryCreatedTopic = "0x45633e1dbbfb9daee5c3fef4d7e2f7956f7ebe21ef019d7d73891807f373233f"
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

	kmsClient := keymanagement.NewKmsClient(cfg.Region)
	encryptedBytes, err := kmsClient.Encrypt([]byte(cfg.AdminAuthMessage), cfg.AuthKeyId)
	if err != nil {
		logger.Fatalf("Failed to encrypt with kms: %v", err)
	}

	encryptedBase64 := base64.StdEncoding.EncodeToString(encryptedBytes)

	client, err := ethclient.Dial(cfg.ProviderWebsocket)
	if err != nil {
		logger.Fatalf("Failed to connect to blockchain provider: %v", err)
	}

	monitorAddresses, err := apiGateway.GetCollectionAddresses(encryptedBase64)
	if err != nil {
		logger.Fatalf("Failed to get monitor addresses: %v", err)
	}
	// Monitor pfp collections
	contractAddresses := []common.Address{}
	for _, address := range monitorAddresses {
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

	// Monitor story block
	orchestratorContractAddresses := []common.Address{
		common.HexToAddress("0x318C47f07a25C7a3EeA5534948FBA2d7a2b4e63c"),
	}

	orchestratorQuery := ethereum.FilterQuery{
		Addresses: orchestratorContractAddresses,
		Topics: [][]common.Hash{
			{common.HexToHash(BackstoryCreatedTopic)},
		},
	}

	orchestratorLogs := make(chan types.Log)
	orchestratorSub, err := client.SubscribeFilterLogs(context.Background(), orchestratorQuery, orchestratorLogs)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Info("Streamer starting")

	for {
		select {
		case err := <-orchestratorSub.Err():
			logger.Errorf("subscription error: %v", err)
			orchestratorSub.Unsubscribe()
			orchestratorSub, err = client.SubscribeFilterLogs(context.Background(), orchestratorQuery, orchestratorLogs)
			if err != nil {
				logger.Fatal(err)
			}
			logger.Info("Resubscribed")
		case vlog := <-orchestratorLogs:
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
			logger.Infof("vlog: %+v", vlog)

			franchiseId := vlog.Topics[1].Big().Int64()
			characterId := vlog.Topics[2].Big().Int64()
			storyId := vlog.Topics[3].Big().Int64()

			err = apiGateway.CreateCharacterWithBackstory(franchiseId, characterId, storyId, encryptedBase64)
			if err != nil {
				logger.Errorf("Failed to create character with backstory: %v", err)
			}
		case err := <-sub.Err():
			logger.Errorf("subscription error: %v", err)
			sub.Unsubscribe()
			sub, err = client.SubscribeFilterLogs(context.Background(), query, logs)
			if err != nil {
				logger.Fatal(err)
			}
			logger.Info("Resubscribed")
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
			logger.Infof("vlog: %+v", vlog)
			collectionAddress := vlog.Address.String()
			fromAddress := vlog.Topics[1].String()
			toAddress := vlog.Topics[2].String()
			tokenId := vlog.Topics[3].Big().Uint64()

			if utils.IsZeroAddress(fromAddress) { // Mint
				logger.Infof("Mint event from zero address to %s", toAddress)
				err = apiGateway.CreateOrUpdateNftRecord(int(tokenId), collectionAddress, encryptedBase64)
				if err != nil {
					logger.Errorf("Failed to create nft record: %v", err)
				}
			} else if utils.IsZeroAddress(toAddress) { // Burn
				logger.Infof("Burn event from %s", fromAddress)
				err = apiGateway.DeleteNftRecord(int(tokenId), collectionAddress, encryptedBase64)
				if err != nil {
					logger.Errorf("Failed to delete nft record: %v", err)
				}
			} else { // Transfer
				logger.Infof("Transfer event from %s to %s", fromAddress, toAddress)
				err = apiGateway.CreateOrUpdateNftRecord(int(tokenId), collectionAddress, encryptedBase64)
				if err != nil {
					logger.Errorf("Failed to update nft record: %v", err)
				}
			}

		}
	}
}
