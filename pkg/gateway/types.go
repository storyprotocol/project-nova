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
