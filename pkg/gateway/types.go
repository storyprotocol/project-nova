package gateway

import (
	"time"

	"github.com/project-nova/backend/pkg/model"
)

// CreateStoryChapterRequestBody contains the data for CreateStoryChapterRequest
type CreateStoryChapterRequestBody struct {
	Title     *string    `json:"title"`
	Heading   *string    `json:"heading"`
	CoverUrl  *string    `json:"coverUrl"`
	ReleaseAt *time.Time `json:"releaseAt"`
}

type CreateRelationshipRequestBody struct {
	SourceContract string `json:"sourceContract"`
	SrcId          uint64 `json:"sourceId"`
	DestContract   string `json:"destContract"`
	DstId          uint64 `json:"destId"`
	TxHash         string `json:"txHash"`
	Type           string `json:"type"`
}

func FromStoryChapterModel(model *model.StoryContentModel) *CreateStoryChapterRequestBody {
	requestBody := &CreateStoryChapterRequestBody{}
	if model.Heading != "" {
		requestBody.Heading = &model.Heading
	}
	if model.Title != "" {
		requestBody.Title = &model.Title
	}
	if model.CoverImage != "" {
		requestBody.CoverUrl = &model.CoverImage
	}

	return requestBody
}

// CreateCharacterWithBackstory contains the data for CreateCharacterWithBackstory request
type CreateCharacterWithBackstoryRequestBody struct {
	SourceContract string
	SourceId       uint64
	DestContract   string
	DestId         uint64
	TxHash         string
}
