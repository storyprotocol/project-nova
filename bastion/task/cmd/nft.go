package cmd

import (
	"encoding/base64"

	"github.com/project-nova/backend/bastion/config"
	"github.com/project-nova/backend/pkg/gateway"
	"github.com/project-nova/backend/pkg/keymanagement"
	"github.com/project-nova/backend/pkg/logger"
	"github.com/spf13/cobra"
)

var (
	oldAddress string
	newAddress string
)

// proofCmd represents the proof command
var updateCollectionAddressCmd = &cobra.Command{
	Use:   "address-update",
	Short: "Update NFT collection's address",
	Long: `The task requires flags: --old, --new.
	--old=<the current collection address> is required. It will be used to find the collection record in DB that needs to be updated.
	--new=<the new collection address> is required. It's the address that we want to update to DB`,
	Run: func(cmd *cobra.Command, args []string) {
		// 1. Get allowlist and json file path
		if oldAddress == "" {
			logger.Fatal("--old is not specified.")
		}
		if newAddress == "" {
			logger.Fatal("--new is not specified")
		}
		// 2. Set configs
		cfg, err := config.GetConfig()
		if err != nil {
			logger.Fatalf("Failed to get configs: %v", err)
		}

		// 3. Connect api server
		apiGateway := gateway.NewApiHttpGateway(cfg.ApiGatewayUrl)

		// 4. Create auth message
		kmsClient, err := keymanagement.NewKmsClient(cfg.Region, "")
		if err != nil {
			logger.Fatalf("Failed to create new kms clint: %v", err)
		}
		encryptedBytes, err := kmsClient.Encrypt([]byte(cfg.AdminAuthMessage), cfg.AuthKeyId)
		if err != nil {
			logger.Fatalf("Failed to encrypt with kms: %v", err)
		}
		encryptedBase64 := base64.StdEncoding.EncodeToString(encryptedBytes)

		// 5. Update collection address
		err = apiGateway.UpdateCollectionAddress(oldAddress, newAddress, encryptedBase64)
		if err != nil {
			logger.Fatalf("Failed to update collection address: %v", err)
		}

		logger.Infof("\nSuccessfully update collection address from %s to %s\n", oldAddress, newAddress)
	},
}

func init() {
	rootCmd.AddCommand(updateCollectionAddressCmd)

	updateCollectionAddressCmd.Flags().StringVarP(&oldAddress, "old", "o", "", "the old collection address")
	updateCollectionAddressCmd.Flags().StringVarP(&newAddress, "new", "n", "", "the new collection address")
}
