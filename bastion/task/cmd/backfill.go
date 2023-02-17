/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/project-nova/backend/bastion/config"
	"github.com/project-nova/backend/pkg/abi/erc721"
	"github.com/project-nova/backend/pkg/gateway"
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
			fmt.Printf("[ERROR] collection address required. pass it with the flag --collection")
			return
		}

		cfg, err := config.GetConfig()
		if err != nil {
			fmt.Printf("[ERROR] Failed to get configs: %v", err)
			return
		}

		apiGateway := gateway.NewApiHttpGateway(cfg.ApiGatewayUrl)

		ethClient, err := ethclient.Dial(cfg.ProviderURL)
		if err != nil {
			fmt.Printf("[ERROR] Failed to connect to the blockchain provider, error: %v", err)
			return
		}

		if startId == -1 {
			startId = 0
		}

		if endId == -1 {
			address := common.HexToAddress(collectionAddress)
			contract, err := erc721.NewErc721(address, ethClient)
			if err != nil {
				fmt.Printf("[ERROR] Failed to instantiate the contract: %v", err)
				return
			}

			totalSold, err := contract.TotalSold(nil)
			if err != nil {
				fmt.Printf("[ERROR] Failed to query uri: %v", err)
				return
			}

			endId = int(totalSold.Int64()) - 1
		}

		for i := startId; i <= endId; i++ {
			err = apiGateway.CreateNftRecord(i, collectionAddress)
			if err != nil {
				fmt.Printf("[ERROR] Failed to create nft record for id %d: %v", i, err)
			}
		}

		fmt.Printf("Backfill for collection %s completed", collectionAddress)
	},
}

func init() {
	rootCmd.AddCommand(backfillCmd)

	backfillCmd.Flags().IntVarP(&startId, "start", "s", -1, "start id of the nft collection")
	backfillCmd.Flags().IntVarP(&endId, "end", "e", -1, "end id of the nft collection")
	backfillCmd.Flags().StringVarP(&collectionAddress, "collection", "c", "", "the address of the nft collection to backfill")
}
