package cmd

import (
	"encoding/base64"
	"strings"
	"time"

	"github.com/project-nova/backend/bastion/config"
	"github.com/project-nova/backend/bastion/task/workflow"
	"github.com/project-nova/backend/pkg/gateway"
	"github.com/project-nova/backend/pkg/keymanagement"
	"github.com/project-nova/backend/pkg/logger"
	"github.com/project-nova/backend/pkg/model"
	"github.com/spf13/cobra"
)

var (
	pageId    string
	s3Path    string
	releaseAt string
)

// proofCmd represents the proof command
var uploadContentCmd = &cobra.Command{
	Use:   "content-upload",
	Short: "Upload chapter content from Notion to S3 and DB. Refresh content cache.",
	Long: `The task requires flags: --page_id, --s3_path. Option flags: --release_at 
	--page_id=<notion page id> is required. It will be used to get content data from Notion.
	--s3_path=<s3 folder path> is required. (i.e. project-nova-content-staging/1:1:1) It will tell the task which bucket and folder to
	upload the content
	--release_at=<time to release the chapter> is in format <2006-01-02T15:04:05Z>, if not present, the default time is now.
	`,
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
		// Get story ids.
		franchiseStoryChapter := strings.Split(folder, ":") // Example folder name: 1:1:1
		if len(franchiseStoryChapter) != 3 {
			logger.Fatalf("Invalid folder structure: %s", folder)
		}
		franchiseId := franchiseStoryChapter[0]
		storyNum := franchiseStoryChapter[1]
		chapterNum := franchiseStoryChapter[2]

		// 2. Set configs
		cfg, err := config.GetConfig()
		if err != nil {
			logger.Fatalf("Failed to get configs: %v", err)
		}

		// 3. Connect api server
		apiGateway := gateway.NewApiHttpGateway(cfg.ApiGatewayUrl)

		// 4. Create auth message
		kmsClient := keymanagement.NewKmsClient(cfg.Region)
		encryptedBytes, err := kmsClient.Encrypt([]byte(cfg.AdminAuthMessage), cfg.AuthKeyId)
		if err != nil {
			logger.Fatalf("Failed to encrypt with kms: %v", err)
		}
		encryptedBase64 := base64.StdEncoding.EncodeToString(encryptedBytes)

		// 5. Call Notion to fetch content and images
		workflow, err := workflow.NewNotionToS3Workflow(cfg, pageId, bucket, folder)
		if err != nil {
			logger.Fatalf("Failed to create notion to s3 workflow: %v", err)
		}

		output, err := workflow.Run()
		if err != nil {
			logger.Fatalf("Failed to run notion to s3 workflow: %v", err)
		}

		content, ok := output.(*model.StoryContentModel)
		if !ok {
			logger.Fatalf("Failed to convert workflow output to content, output: %v", output)
		}

		// 6. Refresh content cache
		err = apiGateway.UpdateContentCache(franchiseId, storyNum, chapterNum, encryptedBase64)
		if err != nil {
			logger.Fatalf("Failed to update content cache: %v", err)
		}

		// 7. Call api server to update DB and set a release time
		requestBody := gateway.FromStoryChapterModel(content)
		// Convert release time
		if releaseAt != "" {
			layout := "2006-01-02T15:04:05Z"
			releaseTime, err := time.Parse(layout, releaseAt)
			if err != nil {
				logger.Fatalf("failed to parse release time %s: %v", releaseAt, err)
			}
			requestBody.ReleaseAt = &releaseTime
		}

		err = apiGateway.CreateStoryChapter(franchiseId, storyNum, chapterNum, requestBody, encryptedBase64)
		if err != nil {
			logger.Fatalf("Failed to create story chapter: %v", err)
		}

		logger.Infof("\nContent upload for %s %s completed\n", bucket, folder)
	},
}

func init() {
	rootCmd.AddCommand(uploadContentCmd)

	uploadContentCmd.Flags().StringVarP(&pageId, "page_id", "p", "", "the id of the notion page")
	uploadContentCmd.Flags().StringVarP(&s3Path, "s3_path", "s", "", "the s3 path to store the content")
	uploadContentCmd.Flags().StringVarP(&releaseAt, "release_at", "r", "", "the release time in utc for the content")
}
