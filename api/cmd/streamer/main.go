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
	"github.com/project-nova/backend/pkg/abi/relationship"
)

const (
	// Transfer Topic's Hash Identifier
	TransferTopic = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"

	// BackstoryCreated's Hash Identifier
	BackstoryCreatedTopic = "0x45633e1dbbfb9daee5c3fef4d7e2f7956f7ebe21ef019d7d73891807f373233f"

	// RelationSet Event Id
	// `cast k "RelationSet(address,uint256,address,uint256,bytes32,uint256)"
	RelationSetTopic = "0xdac80e4156e67d10c07ce819561c6cd96452ad81db0c68e6a47a8687f3d59271"
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

	kmsClient, err := keymanagement.NewKmsClient(cfg.Region, cfg.SSOProfile)
	if err != nil {
		logger.Fatalf("Failed to create KMS client: %v", err)
	}

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

	relationshipContractAddress := common.HexToAddress(cfg.RelationshipContract)
	relationshipFilterer, err := relationship.NewRelationshipFilterer(relationshipContractAddress, client)
	if err != nil {
			logger.Fatalf("Failed to create relationship filterer: %v", err)
	}

	// Monitor relationship contract
	relationshipContractAddresses := []common.Address{
		relationshipContractAddress,
	}

	relationshipQuery := ethereum.FilterQuery{
		Addresses: relationshipContractAddresses,
		Topics: [][]common.Hash{
			{common.HexToHash(RelationSetTopic)},
		},
	}

	relationshipLogs := make(chan types.Log)
	relationshipSub, err := client.SubscribeFilterLogs(context.Background(), relationshipQuery, relationshipLogs)
	if err != nil {
		logger.Fatal(err)
	}

	// Monitor story block
	orchestratorContractAddresses := []common.Address{
		common.HexToAddress(cfg.OrchestratorContract),
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
		case err := <-relationshipSub.Err():
			logger.Errorf("subscription error: %v", err)
			relationshipSub.Unsubscribe()
			relationshipSub, err = client.SubscribeFilterLogs(context.Background(), relationshipQuery, relationshipLogs)
			if err != nil {
				logger.Fatal(err)
			}
			logger.Info("Resubscribed")
		case vlog := <-relationshipLogs:
			logger.Infof("vlog: %+v", vlog)
			event, err := relationshipFilterer.ParseRelationSet(vlog)
			if err != nil {
				logger.Fatal(err)
			}
			sourceContract := event.SourceContract.String()
			sourceId := event.SourceId.Uint64()
			destContract := event.DestContract.String()
			destId := event.DestId.Uint64()
			txHash := event.Raw.TxHash.String()

			err = apiGateway.CreateRelationship(sourceContract, sourceId, destContract, destId, txHash, encryptedBase64)
			if err != nil {
				logger.Errorf("Failed to create relation set: %v", err)
			}
		case err := <-orchestratorSub.Err():
			logger.Errorf("subscription error: %v", err)
			orchestratorSub.Unsubscribe()
			orchestratorSub, err = client.SubscribeFilterLogs(context.Background(), orchestratorQuery, orchestratorLogs)
			if err != nil {
				logger.Fatal(err)
			}
			logger.Info("Resubscribed")
		case vlog := <-orchestratorLogs:
			logger.Infof("vlog: %+v", vlog)
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

			franchiseId := vlog.Topics[1].Big().Int64()
			characterId := vlog.Topics[2].Big().Int64()
			storyId := vlog.Topics[3].Big().Int64()
			txHash := vlog.TxHash.String()

			err = apiGateway.CreateCharacterWithBackstory(franchiseId, characterId, storyId, txHash, encryptedBase64)
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
