/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/base64"
	"encoding/json"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/project-nova/backend/bastion/config"
	"github.com/project-nova/backend/pkg/gateway"
	"github.com/project-nova/backend/pkg/keymanagement"
	"github.com/project-nova/backend/pkg/logger"
	"github.com/project-nova/backend/pkg/s3"
	"github.com/spf13/cobra"
)

var (
	allowlistId  string
	jsonFilePath string
)

// proofCmd represents the proof command
var proofCmd = &cobra.Command{
	Use:   "proof",
	Short: "Upload Whitelist proof to DB",
	Long: `The task requires flags: --allowlist, --json.
	--allowlist=<allowlist id> is required. It will be used to associate the proofs uploaded to this allowlist id.
	--json=<s3 json file> is required. It will tell the task which json file to pull from s3 as the source file`,
	Run: func(cmd *cobra.Command, args []string) {
		// 1. Get allowlist and json file path
		if allowlistId == "" {
			logger.Fatal("--allowlist is not specified.")
		}
		if jsonFilePath == "" {
			logger.Fatal("--json is not specified")
		}
		// 2. Set configs
		cfg, err := config.GetConfig()
		if err != nil {
			logger.Fatalf("Failed to get configs: %v", err)
		}
		// 3. Connect api server and S3
		apiGateway := gateway.NewApiHttpGateway(cfg.ApiGatewayUrl)

		awsSession, err := session.NewSession(&aws.Config{
			Region: aws.String(cfg.Region),
		})
		if err != nil {
			logger.Fatalf("Failed to create aws session: %v", err)
		}
		s3Client := s3.NewS3Client(awsSession)
		// 4. Create auth message
		kmsClient := keymanagement.NewKmsClient(cfg.Region)
		encryptedBytes, err := kmsClient.Encrypt([]byte(cfg.AdminAuthMessage), cfg.AuthKeyId)
		if err != nil {
			logger.Fatalf("Failed to encrypt with kms: %v", err)
		}

		encryptedBase64 := base64.StdEncoding.EncodeToString(encryptedBytes)
		// 5. Download the json file from S3
		buf := aws.NewWriteAtBuffer([]byte{})
		_, err = s3Client.DownloadObject(buf, cfg.S3OperationBucketName, jsonFilePath)
		if err != nil {
			logger.Fatalf("Failed to download content from s3 for object %s: %v", jsonFilePath, err)
		}

		var data struct {
			Root   *string `json:"root"`
			Proofs []*struct {
				Address *string `json:"address"`
				Proof   *string `json:"proof"`
			} `json:"proofs"`
		}
		err = json.Unmarshal(buf.Bytes(), &data)
		if err != nil {
			logger.Fatalf("Failed to unmarshal the s3 proof json file, error: %v", err)
		}
		if data.Root == nil {
			logger.Fatalf("Failed to get root of the proofs")
		}
		// 6. Call api server to upload proofs
		for _, proof := range data.Proofs {
			if proof.Address == nil || proof.Proof == nil {
				logger.Fatalf("Failed to get completed proof for %v", *proof)
			}
			// process proof
			nodes := strings.Split(*proof.Proof, ",")
			marshaledProof, err := json.Marshal(nodes)
			if err != nil {
				logger.Fatalf("Failed to marshal proof for %s, proof %s: %v", *proof.Address, *proof.Proof, err)
			}
			encodedProof := base64.StdEncoding.EncodeToString(marshaledProof)
			// create proof
			err = apiGateway.CreateProof(proof.Address, &encodedProof, &allowlistId, encryptedBase64)
			if err != nil {
				logger.Fatalf("Failed to create proof for %s, proof %s: %v", *proof.Address, *proof.Proof, err)
			}
			logger.Infof("Created proof for address %s\n", *proof.Address)
		}

		logger.Infof("\nBackfill for allowlist %s with root %s completed\n", allowlistId, *data.Root)
	},
}

func init() {
	rootCmd.AddCommand(proofCmd)

	proofCmd.Flags().StringVarP(&allowlistId, "allowlist", "a", "", "the id of the allowlist")
	proofCmd.Flags().StringVarP(&jsonFilePath, "json", "j", "", "the s3 json file path that contains the proofs")
}
