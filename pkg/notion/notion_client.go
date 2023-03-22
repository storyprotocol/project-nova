package notion

import (
	"fmt"

	"github.com/project-nova/backend/pkg/http"
)

const notionBaseUrl = "https://api.notion.com"

type NotionClient interface {
	RetrievePage(pageId string, token string) (*PageResponse, error)
	RetrieveBlockChildren(blockId string, token string) (*BlockChildrenResponse, error)
}

func NewNotionClient() NotionClient {
	return &notionClient{
		client: http.NewClient(&http.ClientConfig{
			BaseURL: notionBaseUrl,
		}),
	}
}

type notionClient struct {
	client http.Client
}

func (n *notionClient) RetrievePage(pageId string, token string) (*PageResponse, error) {
	requestURL := fmt.Sprintf("/v1/pages/%s", pageId)
	headers := &map[string]string{
		"accept":         "application/json",
		"Notion-Version": "2022-06-28",
		"Authorization":  "Bearer " + token,
	}

	var results PageResponse
	_, err := n.client.RequestAddHeaders("GET", requestURL, headers, nil, &results)
	if err != nil {
		return nil, fmt.Errorf("http request to get block children failed. error %v ", err)
	}

	return &results, nil
}

func (n *notionClient) RetrieveBlockChildren(blockId string, token string) (*BlockChildrenResponse, error) {
	requestURL := fmt.Sprintf("/v1/blocks/%s/children?page_size=100", blockId)
	headers := &map[string]string{
		"accept":         "application/json",
		"Notion-Version": "2022-06-28",
		"Authorization":  "Bearer " + token,
	}

	var results BlockChildrenResponse
	_, err := n.client.RequestAddHeaders("GET", requestURL, headers, nil, &results)
	if err != nil {
		return nil, fmt.Errorf("http request to get block children failed. error %v ", err)
	}

	return &results, nil
}
