package gateway

import (
	"fmt"
	"strconv"

	"github.com/project-nova/backend/pkg/http"
	"github.com/project-nova/backend/pkg/middleware"
)

type ApiGateway interface {
	GetCollectionAddresses(authMessage string) ([]string, error)
	UpdateNftOwner(tokenId int, collectionAddress string, authMessage string) error
	UpdateCollectionAddress(oldAddress string, newAddress string, authMessage string) error
	CreateOrUpdateNftRecord(tokenId int, collectionAddress string, authMessage string) error
	DeleteNftRecord(tokenId int, collectionAddress string, authMessage string) error
	CreateProof(address *string, proof *string, allowlistId *string, authMessage string) error
	CreateStoryChapter(franchiseId string, storyNum string, chapterNum string, requestBody *CreateStoryChapterRequestBody, authMessage string) error
	UpdateContentCache(franchiseId string, storyNum string, chapterNum string, authMessage string) error
	CreateCharacterWithBackstory(franchiseId int64, characterId int64, storyId int64, txHash string, authMessage string) error
}

func NewApiHttpGateway(url string) ApiGateway {
	return &apiHttpGateway{
		client: http.NewClient(&http.ClientConfig{
			BaseURL: url,
		}),
	}
}

type apiHttpGateway struct {
	client http.Client
}

func (s *apiHttpGateway) GetCollectionAddresses(authMessage string) ([]string, error) {
	requestURL := "/admin/v1/nft/collections"
	results := []struct {
		CollectionAddress string `json:"collectionAddress"`
	}{}
	headers := &map[string]string{
		middleware.AuthMessageHeaderKey: authMessage,
	}
	_, err := s.client.RequestAddHeaders("GET", requestURL, headers, nil, &results)
	if err != nil {
		return nil, fmt.Errorf("http request to get collections failed. error: %v ", err)
	}

	addresses := []string{}
	for _, result := range results {
		addresses = append(addresses, result.CollectionAddress)
	}

	return addresses, nil
}

func (s *apiHttpGateway) UpdateNftOwner(tokenId int, collectionAddress string, authMessage string) error {
	requestURL := fmt.Sprintf("/admin/v1/nft/%s/owner?collectionAddress=%s", strconv.Itoa(tokenId), collectionAddress)
	headers := &map[string]string{
		middleware.AuthMessageHeaderKey: authMessage,
	}
	_, err := s.client.RequestAddHeaders("POST", requestURL, headers, nil, nil)
	if err != nil {
		return fmt.Errorf("http request to update nft owner failed. error: %v ", err)
	}
	return nil
}

func (s *apiHttpGateway) UpdateCollectionAddress(oldAddress string, newAddress string, authMessage string) error {
	requestURL := fmt.Sprintf("/admin/v1/nft/collection/%s?newAddress=%s", oldAddress, newAddress)
	headers := &map[string]string{
		middleware.AuthMessageHeaderKey: authMessage,
	}
	_, err := s.client.RequestAddHeaders("POST", requestURL, headers, nil, nil)
	if err != nil {
		return fmt.Errorf("http request to update nft collection address failed. error: %v ", err)
	}
	return nil
}

func (s *apiHttpGateway) CreateOrUpdateNftRecord(tokenId int, collectionAddress string, authMessage string) error {
	requestURL := fmt.Sprintf("/admin/v1/nft/%s?collectionAddress=%s", strconv.Itoa(tokenId), collectionAddress)
	headers := &map[string]string{
		middleware.AuthMessageHeaderKey: authMessage,
	}
	_, err := s.client.RequestAddHeaders("POST", requestURL, headers, nil, nil)
	if err != nil {
		return fmt.Errorf("http request to create nft record failed. error: %v ", err)
	}
	return nil
}

func (s *apiHttpGateway) DeleteNftRecord(tokenId int, collectionAddress string, authMessage string) error {
	requestURL := fmt.Sprintf("/admin/v1/nft/%s?collectionAddress=%s", strconv.Itoa(tokenId), collectionAddress)
	headers := &map[string]string{
		middleware.AuthMessageHeaderKey: authMessage,
	}
	_, err := s.client.RequestAddHeaders("DELETE", requestURL, headers, nil, nil)
	if err != nil {
		return fmt.Errorf("http request to delete nft record failed. error: %v ", err)
	}
	return nil
}

func (s *apiHttpGateway) CreateProof(address *string, proof *string, allowlistId *string, authMessage string) error {
	requestURL := fmt.Sprintf("/admin/v1/wallet/%s/proof?allowlistId=%s&proof=%s", *address, *allowlistId, *proof)
	headers := &map[string]string{
		middleware.AuthMessageHeaderKey: authMessage,
	}

	_, err := s.client.RequestAddHeaders("POST", requestURL, headers, nil, nil)
	if err != nil {
		return fmt.Errorf("http request to create proof failed. error: %v ", err)
	}
	return nil
}

func (s *apiHttpGateway) CreateStoryChapter(franchiseId string, storyNum string, chapterNum string, requestBody *CreateStoryChapterRequestBody, authMessage string) error {
	requestURL := fmt.Sprintf("/admin/v1/story/%s/%s/%s", franchiseId, storyNum, chapterNum)
	headers := &map[string]string{
		middleware.AuthMessageHeaderKey: authMessage,
	}

	_, err := s.client.RequestAddHeaders("POST", requestURL, headers, requestBody, nil)
	if err != nil {
		return fmt.Errorf("http request to create story chapter failed. error: %v ", err)
	}
	return nil
}

func (s *apiHttpGateway) UpdateContentCache(franchiseId string, storyNum string, chapterNum string, authMessage string) error {
	requestURL := fmt.Sprintf("/admin/v1/story/%s/%s/%s/cache", franchiseId, storyNum, chapterNum)
	headers := &map[string]string{
		middleware.AuthMessageHeaderKey: authMessage,
	}

	_, err := s.client.RequestAddHeaders("POST", requestURL, headers, nil, nil)
	if err != nil {
		return fmt.Errorf("http request to update content cache failed. error: %v ", err)
	}
	return nil
}

func (s *apiHttpGateway) CreateCharacterWithBackstory(franchiseId int64, characterId int64, storyId int64, txHash string, authMessage string) error {
	requestURL := fmt.Sprintf("/admin/v2/character/%s/%s/%s", strconv.FormatInt(franchiseId, 10), strconv.FormatInt(characterId, 10), strconv.FormatInt(storyId, 10))

	headers := &map[string]string{
		middleware.AuthMessageHeaderKey: authMessage,
	}
	requestBody := &CreateCharacterWithBackstoryRequestBody{
		TxHash: txHash,
	}

	_, err := s.client.RequestAddHeaders("POST", requestURL, headers, requestBody, nil)
	if err != nil {
		return fmt.Errorf("http request to create character with backstory. error: %v ", err)
	}
	return nil
}
