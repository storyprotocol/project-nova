package cmd

import (
	"strings"

	"github.com/project-nova/backend/bastion/config"
	"github.com/project-nova/backend/bastion/task/workflow"
	"github.com/project-nova/backend/pkg/logger"
	"github.com/spf13/cobra"
)

var (
	pageId string
	s3Path string
)

// proofCmd represents the proof command
var uploadContentCmd = &cobra.Command{
	Use:   "content-upload",
	Short: "Upload chapter content from Notion to S3 and DB. Refresh content cache.",
	Long: `The task requires flags: --page_id, --s3_path. 
	--page_id=<notion page id> is required. It will be used to get content data from Notion.
	--s3_path=<s3 folder path> is required. (i.e. project-nova-content-staging/1:1:1) It will tell the task which bucket and folder to
	upload the content`,
	Run: func(cmd *cobra.Command, args []string) {
		// 1. Get flag values
		if pageId == "" {
			logger.Fatal("--page_id is not specified.")
		}
		if s3Path == "" {
			logger.Fatal("--s3_path is not specified")
		}
		pathNames := strings.Split(s3Path, "/")
		bucket := pathNames[0]
		folder := pathNames[1]
		// 2. Set configs
		cfg, err := config.GetConfig()
		if err != nil {
			logger.Fatalf("Failed to get configs: %v", err)
		}
		// 3. Connect api server and S3
		// (Adding back in next PR)
		//apiGateway := gateway.NewApiHttpGateway(cfg.ApiGatewayUrl)

		/* (Adding back in next PR)
		// 4. Create auth message
		kmsClient := keymanagement.NewKmsClient(cfg.Region)
		encryptedBytes, err := kmsClient.Encrypt([]byte(cfg.AdminAuthMessage), cfg.AuthKeyId)
		if err != nil {
			logger.Fatalf("Failed to encrypt with kms: %v", err)
		}

		encryptedBase64 := base64.StdEncoding.EncodeToString(encryptedBytes)
		*/

		// 5. Call Notion to fetch content and images.
		workflow, err := workflow.NewNotionToS3Workflow(cfg, pageId, bucket, folder)
		if err != nil {
			logger.Fatalf("Failed to create notion to s3 workflow: %v", err)
		}

		_, err = workflow.Run()
		if err != nil {
			logger.Fatalf("Failed to run notion to s3 workflow: %v", err)
		}
		// 6. (TO DO) Call api server to update DB and set a release time
		// 7. (TO DO) Refresh content cache

		logger.Infof("\nContent upload for %s %s completed\n", bucket, folder)
	},
}

func init() {
	rootCmd.AddCommand(uploadContentCmd)

	uploadContentCmd.Flags().StringVarP(&pageId, "page_id", "p", "", "the id of the notion page")
	uploadContentCmd.Flags().StringVarP(&s3Path, "s3_path", "s", "", "the s3 path to store the content")
}
