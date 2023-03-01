package cmd

import (
	"encoding/base64"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/project-nova/backend/bastion/config"
	"github.com/project-nova/backend/pkg/abi/erc721"
	"github.com/project-nova/backend/pkg/gateway"
	"github.com/project-nova/backend/pkg/keymanagement"
	"github.com/project-nova/backend/pkg/logger"
	"github.com/spf13/cobra"
)

var (
	startId           int
	endId             int
	collectionAddress string
)

// backfillCmd represents the backfill command
var backfillCmd = &cobra.Command{
	Use:   "backfill",
	Short: "backfill task for backfilling nft onchain data to database",
	Long: `The task accepts flags: --start, ---end, --collection.
--collection is required and will be used to determine which nft collection to backfill.
--start and --end tell the task the range of nft token id to backfill. When they are specified. The default is to backfill all nft data from a colleciton.	
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if collectionAddress == "" {
			logger.Error("Collection address required. pass it with the flag --collection")
			return
		}

		cfg, err := config.GetConfig()
		if err != nil {
			logger.Errorf("Failed to get configs: %v", err)
			return
		}

		apiGateway := gateway.NewApiHttpGateway(cfg.ApiGatewayUrl)

		kmsClient := keymanagement.NewKmsClient(cfg.Region)
		encryptedBytes, err := kmsClient.Encrypt([]byte(cfg.AdminAuthMessage), cfg.AuthKeyId)
		if err != nil {
			logger.Fatalf("Failed to encrypt with kms: %v", err)
		}

		encryptedBase64 := base64.StdEncoding.EncodeToString(encryptedBytes)

		ethClient, err := ethclient.Dial(cfg.ProviderURL)
		if err != nil {
			logger.Errorf("Failed to connect to the blockchain provider, error: %v", err)
			return
		}

		if startId == -1 {
			startId = 0
		}

		if endId == -1 {
			address := common.HexToAddress(collectionAddress)
			contract, err := erc721.NewErc721(address, ethClient)
			if err != nil {
				logger.Errorf("Failed to instantiate the contract: %v", err)
				return
			}

			totalSold, err := contract.TotalSold(nil)
			if err != nil {
				logger.Errorf("Failed to query uri: %v", err)
				return
			}

			endId = int(totalSold.Int64()) - 1
		}

		logger.Infof("Starting backfills from id %d to id %d\n\n", startId, endId)
		for i := startId; i <= endId; i++ {
			err = apiGateway.CreateNftRecord(i, collectionAddress, encryptedBase64)
			logger.Infof("Created nft record for id %d\n", i)
			if err != nil {
				logger.Errorf("Failed to create nft record for id %d: %v\n", i, err)
			}
			time.Sleep(500 * time.Millisecond)
		}

		logger.Infof("\nBackfill for collection %s completed\n", collectionAddress)
	},
}

func init() {
	rootCmd.AddCommand(backfillCmd)

	backfillCmd.Flags().IntVarP(&startId, "start", "s", -1, "start id of the nft collection")
	backfillCmd.Flags().IntVarP(&endId, "end", "e", -1, "end id of the nft collection")
	backfillCmd.Flags().StringVarP(&collectionAddress, "collection", "c", "", "the address of the nft collection to backfill")
}
