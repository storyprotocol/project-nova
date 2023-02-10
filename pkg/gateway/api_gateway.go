package gateway

import (
	"fmt"
	"strconv"

	"github.com/project-nova/backend/pkg/http"
)

type ApiGateway interface {
	UpdateNftOwner(tokenId int, collectionAddress string) error
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

func (s *apiHttpGateway) UpdateNftOwner(tokenId int, collectionAddress string) error {
	requestURL := fmt.Sprintf("/admin/v1/nft/%s/owner?collectionAddress=%s", strconv.Itoa(tokenId), collectionAddress)
	_, err := s.client.Request("POST", requestURL, nil, nil)
	if err != nil {
		return fmt.Errorf("http request to update nft owner failed. error %v ", err)
	}
	return nil
}
