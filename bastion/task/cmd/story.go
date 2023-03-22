package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/goccy/go-json"
	"github.com/project-nova/backend/bastion/config"
	"github.com/project-nova/backend/pkg/constant"
	"github.com/project-nova/backend/pkg/logger"
	"github.com/project-nova/backend/pkg/model"
	"github.com/project-nova/backend/pkg/notion"
	"github.com/project-nova/backend/pkg/s3"
	"github.com/spf13/cobra"
)

var (
	pageId string
	s3Path string

	taskCfg      *config.TaskConfig
	notionClient notion.NotionClient
	s3Client     s3.S3Client
	bucket       string
	folder       string
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
		bucket = pathNames[0]
		folder = pathNames[1]
		// 2. Set configs
		cfg, err := config.GetConfig()
		if err != nil {
			logger.Fatalf("Failed to get configs: %v", err)
		}
		taskCfg = cfg
		// 3. Connect api server and S3
		// (Adding back in next PR)
		//apiGateway := gateway.NewApiHttpGateway(cfg.ApiGatewayUrl)

		awsSession, err := session.NewSession(&aws.Config{
			Region: aws.String(cfg.Region),
		})
		if err != nil {
			logger.Fatalf("Failed to create aws session: %v", err)
		}
		s3Client = s3.NewS3Client(awsSession)
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
		notionClient = notion.NewNotionClient()
		pageResponse, err := notionClient.RetrievePage(pageId, cfg.NotionAuthToken)
		if err != nil {
			logger.Fatalf("Failed to retrieve page data from notion: %v", err)
		}

		blockReponse, err := notionClient.RetrieveBlockChildren(pageId, cfg.NotionAuthToken)
		if err != nil {
			logger.Fatalf("Failed to retrieve content from notion: %v", err)
		}
		logger.Infof("length of blocks: %d", len(blockReponse.Results))
		//5.1. Create content json and add processed contents.
		content, err := formatContentAndUploadImage(pageResponse, blockReponse)
		if err != nil {
			logger.Fatalf("Failed to format and upload content: %v", err)
		}
		//5.2. Upload the final content json to S3
		err = createAndUploadContentJson(content)
		if err != nil {
			logger.Fatal("Failed to create and upload content json: %v", err)
		}
		// 6. (TO DO) Call api server to update DB and set a release time
		// 7. (TO DO) Refresh content cache

		logger.Infof("\nContent upload for %s %s completed\n", bucket, folder)
	},
}

func formatContentAndUploadImage(pageResp *notion.PageResponse, blockResp *notion.BlockChildrenResponse) (*model.StoryContentModel, error) {
	content := &model.StoryContentModel{}
	// 1. Add chapter name, heading and cover image to content model
	// 1.1 Handle chapter Heading and Title
	if pageResp.Properties == nil || pageResp.Properties.Title == nil || len(pageResp.Properties.Title.Title) == 0 {
		return nil, fmt.Errorf("title is not set")
	}

	title := pageResp.Properties.Title.Title[0].PlainText
	headingTitle := strings.Split(title, ":")
	if len(headingTitle) != 2 {
		return nil, fmt.Errorf("failed to process title, invalid format")
	}

	content.Heading = strings.TrimSpace(headingTitle[0])
	content.Title = strings.TrimSpace(headingTitle[1])
	// 1.2 Handle Cover image
	if pageResp.Cover != nil && pageResp.Cover.File != nil && pageResp.Cover.File.Url != "" {
		imageLocation, err := transferImage(pageResp.Cover.File.Url)
		if err != nil {
			return nil, fmt.Errorf("failed to transfer cover image: %v", err)
		}
		content.CoverImage = imageLocation
	}

	// 2. Go through each block and convert them to story data model.
	var curSection *model.StorySectionModel
	for idx, block := range blockResp.Results {
		switch block.Type {
		case notion.BlockTypes.Paragraph:
			if curSection == nil {
				curSection = &model.StorySectionModel{
					Type: model.StorySectionType.Paragraph,
					Data: []*model.StoryMediaModel{
						{
							Type: model.StoryMediaType.Text,
						},
					},
				}
			}

			newContent, err := getContent(block.Paragraph)
			if err != nil {
				return nil, fmt.Errorf("failed to get content for block %d, id %s: %v", idx, block.Id, err)
			}
			curSection.Data[0].Content = curSection.Data[0].Content + newContent

		case notion.BlockTypes.ColumnList:
			if curSection != nil {
				content.Content = append(content.Content, curSection)
				curSection = nil
			}

			imageSection, err := processColumnList(block)
			if err != nil {
				return nil, fmt.Errorf("failed to process column list block %s: %v", block.Id, err)
			}
			content.Content = append(content.Content, imageSection)

		default:
			return nil, fmt.Errorf("unknown block type %s", string(block.Type))
		}
	}

	if curSection != nil {
		content.Content = append(content.Content, curSection)
		curSection = nil
	}

	return content, nil
}

func getImageName(url string) string {
	// Example url: https://s3.us-west-2.amazonaws.com/secure.notion-static.com/b3b54d2c-a596-4383-8cf9-9a3352631546/rayze_as_cheetah_looking_from_rooftops.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Content-Sha256=UNSIGNED-PAYLOAD&X-Amz-Credential=AKIAT73L2G45EIPT3X45%2F20230322%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Date=20230322T041608Z&X-Amz-Expires=3600&X-Amz-Signature=ed4de5d680af981f7ee5ce12a1366051d7d9314da0db93f249966cfe387212ce&X-Amz-SignedHeaders=host&x-id=GetObject
	splitStr1 := strings.Split(url, "?")
	splitStr2 := strings.Split(splitStr1[0], "/")
	return splitStr2[len(splitStr2)-1]
}

func transferImage(url string) (string, error) {
	imageName, err := downloadImage(url)
	if err != nil {
		return "", fmt.Errorf("failed to download image: %v", err)
	}

	output, err := s3Client.UploadObject(bucket, folder+"/"+constant.MediaFolder+"/"+imageName, imageName, true)
	if err != nil {
		return "", fmt.Errorf("failed to upload image to S3: %v", err)
	}

	return output.Location, nil
}

func downloadImage(url string) (string, error) {
	imageName := getImageName(url)
	imageFile, err := os.Create(imageName)
	if err != nil {
		return "", fmt.Errorf("failed to create image file %s: %v", imageName, err)
	}
	defer imageFile.Close()

	coverImage, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to download image: %v", err)
	}
	defer coverImage.Body.Close()

	_, err = io.Copy(imageFile, coverImage.Body)
	if err != nil {
		return "", fmt.Errorf("failed to copy image to the file %s: %v", imageName, err)
	}

	return imageName, nil
}

func createAndUploadContentJson(content interface{}) error {
	contentJson, err := json.MarshalIndent(content, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal blocks: %v", err)
	}

	fileHandle, err := os.Create(constant.ContentObject)
	if err != nil {
		return fmt.Errorf("failed to create %s", constant.ContentObject)
	}
	defer fileHandle.Close()

	_, err = fileHandle.Write(contentJson)
	if err != nil {
		return fmt.Errorf("failed to write to file: %v", err)
	}

	_, err = s3Client.UploadObject(bucket, folder+"/"+constant.ContentObject, constant.ContentObject, false)
	if err != nil {
		return fmt.Errorf("failed to download content from s3 for object %s: %v", jsonFilePath, err)
	}

	return nil
}

func getContent(paragraph *notion.Paragraph) (string, error) {
	if paragraph == nil || paragraph.RichTexts == nil || len(paragraph.RichTexts) > 1 {
		return "", fmt.Errorf("invalid paragraph")
	}

	if len(paragraph.RichTexts) == 0 {
		return "\n", nil
	}

	return paragraph.RichTexts[0].PlainText + "\n", nil
}

func processColumnList(block *notion.Block) (*model.StorySectionModel, error) {
	if !block.HasChildren {
		return nil, fmt.Errorf("invalid column list: No children")
	}

	blockReponse, err := notionClient.RetrieveBlockChildren(block.Id, taskCfg.NotionAuthToken)
	if err != nil {
		return nil, fmt.Errorf("failed to get children of column list block: %v", err)
	}
	if len(blockReponse.Results) != 2 {
		return nil, fmt.Errorf("invalid column list: The list length != 2")
	}

	var imageSection *model.StorySectionModel
	type1, contentOrUrl1, err := processColumn(blockReponse.Results[0])
	if err != nil {
		return nil, fmt.Errorf("failed to process first column: %v", err)
	}
	type2, contentOrUrl2, err := processColumn(blockReponse.Results[1])
	if err != nil {
		return nil, fmt.Errorf("failed to process second column: %v", err)
	}

	if type1 == model.StoryMediaType.Image && type2 == model.StoryMediaType.Text {
		imageSection = &model.StorySectionModel{
			Type: model.StorySectionType.ImageText,
			Data: []*model.StoryMediaModel{
				{
					Type: type1,
					Url:  contentOrUrl1,
				},
				{
					Type:    type2,
					Content: contentOrUrl2,
				},
			},
		}
	} else if type1 == model.StoryMediaType.Text && type2 == model.StoryMediaType.Image {
		imageSection = &model.StorySectionModel{
			Type: model.StorySectionType.TextImage,
			Data: []*model.StoryMediaModel{
				{
					Type:    type1,
					Content: contentOrUrl1,
				},
				{
					Type: type2,
					Url:  contentOrUrl2,
				},
			},
		}
	} else {
		return nil, fmt.Errorf("invalid column type combination: %s %s", type1, type2)
	}

	return imageSection, nil
}

func processColumn(block *notion.Block) (string, string, error) {
	if !block.HasChildren || block.Type != notion.BlockTypes.Column {
		return "", "", fmt.Errorf("invalid column block, id: %s", block.Id)
	}

	blockReponse, err := notionClient.RetrieveBlockChildren(block.Id, taskCfg.NotionAuthToken)
	if err != nil {
		return "", "", fmt.Errorf("failed to get children of column block: %v", err)
	}
	if len(blockReponse.Results) == 0 {
		return "", "", fmt.Errorf("invalid column block children response, children == 0")
	}

	var mediaType string
	var contentOrUrl string
	blockType := blockReponse.Results[0].Type
	switch blockType {

	case notion.BlockTypes.Image:
		mediaType = model.StoryMediaType.Image
		imageBlock := blockReponse.Results[0]
		if imageBlock.Image == nil || imageBlock.Image.File == nil || imageBlock.Image.File.Url == "" {
			return "", "", fmt.Errorf("invalid image block: id %s", imageBlock.Id)
		}

		contentOrUrl, err = transferImage(imageBlock.Image.File.Url)
		if err != nil {
			return "", "", fmt.Errorf("failed to transfer image block %s: %v", imageBlock.Id, err)
		}

	case notion.BlockTypes.Paragraph:
		mediaType = model.StoryMediaType.Text
		contentOrUrl, err = transformParagraphToText(blockReponse.Results)
		if err != nil {
			return "", "", fmt.Errorf("failed to transform paragraph to text %s: %v", block.Id, err)
		}

	default:
		return "", "", fmt.Errorf("invalid block type within the column: id %s", block.Id)
	}

	return mediaType, contentOrUrl, nil
}

func transformParagraphToText(blocks []*notion.Block) (string, error) {
	var text string
	for idx, block := range blocks {
		if block.Type != notion.BlockTypes.Paragraph {
			return "", fmt.Errorf("invalid paragraph block %s", block.Id)
		}
		newContent, err := getContent(block.Paragraph)
		if err != nil {
			return "", fmt.Errorf("failed to get content for block %d, id %s: %v", idx, block.Id, err)
		}
		text = text + newContent
	}
	return text, nil
}

func init() {
	rootCmd.AddCommand(uploadContentCmd)

	uploadContentCmd.Flags().StringVarP(&pageId, "page_id", "p", "", "the id of the notion page")
	uploadContentCmd.Flags().StringVarP(&s3Path, "s3_path", "s", "", "the s3 path to store the content")
}
