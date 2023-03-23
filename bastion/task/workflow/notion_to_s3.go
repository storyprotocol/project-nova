package workflow

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/project-nova/backend/bastion/config"
	"github.com/project-nova/backend/pkg/constant"
	"github.com/project-nova/backend/pkg/logger"
	"github.com/project-nova/backend/pkg/model"
	"github.com/project-nova/backend/pkg/notion"
	"github.com/project-nova/backend/pkg/s3"
)

func NewNotionToS3Workflow(cfg *config.TaskConfig, pageId string, bucket string, folder string) (Workflow, error) {
	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(cfg.Region),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create aws session: %v", err)
	}
	s3Client := s3.NewS3Client(awsSession)

	notionClient := notion.NewNotionClient()

	return &notionToS3Workflow{
		cfg:          cfg,
		pageId:       pageId,
		bucket:       bucket,
		folder:       folder,
		s3Client:     s3Client,
		notionClient: notionClient,
	}, nil
}

type notionToS3Workflow struct {
	cfg          *config.TaskConfig
	pageId       string
	bucket       string
	folder       string
	s3Client     s3.S3Client
	notionClient notion.NotionClient
}

func (n *notionToS3Workflow) Run() (interface{}, error) {

	pageResponse, err := n.notionClient.RetrievePage(n.pageId, n.cfg.NotionAuthToken)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve page data from notion: %v", err)
	}

	blockReponse, err := n.notionClient.RetrieveBlockChildren(n.pageId, n.cfg.NotionAuthToken)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve content from notion: %v", err)
	}
	logger.Infof("length of blocks: %d", len(blockReponse.Results))
	//5.1. Create content json and add processed contents.
	content, err := n.formatContentAndUploadImage(pageResponse, blockReponse)
	if err != nil {
		return nil, fmt.Errorf("failed to format and upload content: %v", err)
	}
	//5.2. Upload the final content json to S3
	err = n.createAndUploadContentJson(content)
	if err != nil {
		return nil, fmt.Errorf("failed to create and upload content json: %v", err)
	}

	return content, nil
}

func (n *notionToS3Workflow) formatContentAndUploadImage(pageResp *notion.PageResponse, blockResp *notion.BlockChildrenResponse) (*model.StoryContentModel, error) {
	// Example page and block:
	// https://developers.notion.com/reference/block
	// https://developers.notion.com/reference/page
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
		imageLocation, err := n.transferImage(pageResp.Cover.File.Url)
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

			newContent, err := n.getContent(block.Paragraph)
			if err != nil {
				return nil, fmt.Errorf("failed to get content for block %d, id %s: %v", idx, block.Id, err)
			}
			curSection.Data[0].Content = curSection.Data[0].Content + newContent

		case notion.BlockTypes.ColumnList:
			if curSection != nil {
				content.Content = append(content.Content, curSection)
				curSection = nil
			}

			imageSection, err := n.processColumnList(block)
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

func (n *notionToS3Workflow) getImageName(url string) string {
	// Example url: https://s3.us-west-2.amazonaws.com/secure.notion-static.com/b3b54d2c-a596-4383-8cf9-9a3352631546/rayze_as_cheetah_looking_from_rooftops.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Content-Sha256=UNSIGNED-PAYLOAD&X-Amz-Credential=AKIAT73L2G45EIPT3X45%2F20230322%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Date=20230322T041608Z&X-Amz-Expires=3600&X-Amz-Signature=ed4de5d680af981f7ee5ce12a1366051d7d9314da0db93f249966cfe387212ce&X-Amz-SignedHeaders=host&x-id=GetObject
	splitStr1 := strings.Split(url, "?")
	splitStr2 := strings.Split(splitStr1[0], "/")
	return splitStr2[len(splitStr2)-1]
}

func (n *notionToS3Workflow) transferImage(url string) (string, error) {
	imageName, err := n.downloadImage(url)
	if err != nil {
		return "", fmt.Errorf("failed to download image: %v", err)
	}

	output, err := n.s3Client.UploadObject(n.bucket, n.folder+"/"+constant.MediaFolder+"/"+imageName, imageName, true)
	if err != nil {
		return "", fmt.Errorf("failed to upload image to S3: %v", err)
	}

	err = os.Remove(imageName)
	if err != nil {
		logger.Errorf("Failed to remove image file %s: %v", imageName, err)
	}

	return output.Location, nil
}

func (n *notionToS3Workflow) downloadImage(url string) (string, error) {
	imageName := n.getImageName(url)
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

func (n *notionToS3Workflow) createAndUploadContentJson(content interface{}) error {
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

	_, err = n.s3Client.UploadObject(n.bucket, n.folder+"/"+constant.ContentObject, constant.ContentObject, false)
	if err != nil {
		return fmt.Errorf("failed to upload image to S3: %v", err)
	}

	err = os.Remove(constant.ContentObject)
	if err != nil {
		logger.Errorf("Failed to remove josn file: %v", err)
	}

	return nil
}

func (n *notionToS3Workflow) getContent(paragraph *notion.Paragraph) (string, error) {
	if paragraph == nil || paragraph.RichTexts == nil || len(paragraph.RichTexts) > 1 {
		return "", fmt.Errorf("invalid paragraph")
	}

	if len(paragraph.RichTexts) == 0 {
		return "\n", nil
	}

	return paragraph.RichTexts[0].PlainText + "\n", nil
}

func (n *notionToS3Workflow) processColumnList(block *notion.Block) (*model.StorySectionModel, error) {
	if !block.HasChildren {
		return nil, fmt.Errorf("invalid column list: No children")
	}

	blockReponse, err := n.notionClient.RetrieveBlockChildren(block.Id, n.cfg.NotionAuthToken)
	if err != nil {
		return nil, fmt.Errorf("failed to get children of column list block: %v", err)
	}
	if len(blockReponse.Results) != 2 {
		return nil, fmt.Errorf("invalid column list: The list length != 2")
	}

	var imageSection *model.StorySectionModel
	type1, contentOrUrl1, err := n.processColumn(blockReponse.Results[0])
	if err != nil {
		return nil, fmt.Errorf("failed to process first column: %v", err)
	}
	type2, contentOrUrl2, err := n.processColumn(blockReponse.Results[1])
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

func (n *notionToS3Workflow) processColumn(block *notion.Block) (string, string, error) {
	if !block.HasChildren || block.Type != notion.BlockTypes.Column {
		return "", "", fmt.Errorf("invalid column block, id: %s", block.Id)
	}

	blockReponse, err := n.notionClient.RetrieveBlockChildren(block.Id, n.cfg.NotionAuthToken)
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

		contentOrUrl, err = n.transferImage(imageBlock.Image.File.Url)
		if err != nil {
			return "", "", fmt.Errorf("failed to transfer image block %s: %v", imageBlock.Id, err)
		}

	case notion.BlockTypes.Paragraph:
		mediaType = model.StoryMediaType.Text
		contentOrUrl, err = n.transformParagraphToText(blockReponse.Results)
		if err != nil {
			return "", "", fmt.Errorf("failed to transform paragraph to text %s: %v", block.Id, err)
		}

	default:
		return "", "", fmt.Errorf("invalid block type within the column: id %s", block.Id)
	}

	return mediaType, contentOrUrl, nil
}

func (n *notionToS3Workflow) transformParagraphToText(blocks []*notion.Block) (string, error) {
	var text string
	for idx, block := range blocks {
		if block.Type != notion.BlockTypes.Paragraph {
			return "", fmt.Errorf("invalid paragraph block %s", block.Id)
		}
		newContent, err := n.getContent(block.Paragraph)
		if err != nil {
			return "", fmt.Errorf("failed to get content for block %d, id %s: %v", idx, block.Id, err)
		}
		text = text + newContent
	}
	return text, nil
}
