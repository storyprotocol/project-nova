package entity

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/project-nova/backend/pkg/gateway"
)

type StoryChapterResp struct {
	ChapterNum int    `json:"chapterNum"`
	Title      string `json:"title"`
	Heading    string `json:"heading"`
	CoverUrl   string `json:"coverUrl"`
}

// StoryChapterModel represents the story chapter's model in data storage
type StoryChapterModel struct {
	ID        string    `gorm:"primaryKey;column:id" json:"id"`
	StoryId   string    `json:"storyId"`
	SeqNum    int       `json:"seqNum"`
	Title     *string   `json:"title"`
	Heading   *string   `json:"heading"`
	CoverUrl  *string   `json:"coverUrl"`
	ReleaseAt time.Time `json:"releaseAt"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (StoryChapterModel) TableName() string {
	return "story_chapter"
}

func (s *StoryChapterModel) ToStoryChapterResp() *StoryChapterResp {
	resp := &StoryChapterResp{
		ChapterNum: s.SeqNum,
	}

	if s.Heading != nil {
		resp.Heading = *s.Heading
	}
	if s.Title != nil {
		resp.Title = *s.Title
	}
	if s.CoverUrl != nil {
		resp.CoverUrl = *s.CoverUrl
	}

	return resp
}

func (s *StoryChapterModel) FromCreateStoryChapterRequestBody(request *gateway.CreateStoryChapterRequestBody) {
	if request.Heading != nil {
		s.Heading = request.Heading
	}
	if request.Title != nil {
		s.Title = request.Title
	}
	if request.CoverUrl != nil {
		s.CoverUrl = request.CoverUrl
	}
	if request.ReleaseAt != nil {
		s.ReleaseAt = *request.ReleaseAt
	}
}

func ToStoryChapterModel(request *gateway.CreateStoryChapterRequestBody, storyId string, chapterNum int) *StoryChapterModel {
	storyChapterModel := &StoryChapterModel{
		ID:      uuid.New().String(),
		StoryId: storyId,
		SeqNum:  chapterNum,
	}

	if request.Heading != nil {
		storyChapterModel.Heading = request.Heading
	}
	if request.Title != nil {
		storyChapterModel.Title = request.Title
	}
	if request.CoverUrl != nil {
		storyChapterModel.CoverUrl = request.CoverUrl
	}
	if request.ReleaseAt == nil {
		curr := time.Now().UTC()
		request.ReleaseAt = &curr
	}
	storyChapterModel.ReleaseAt = *request.ReleaseAt

	return storyChapterModel
}

type StoryInfoV2Model struct {
	ID               string                `gorm:"primaryKey;column:id" json:"id,omitempty"`
	FranchiseId      int64                 `json:"franchiseId,omitempty"`
	StoryId          *int64                `json:"storyId,omitempty"`
	StoryName        string                `json:"storyName,omitempty"`
	StoryDescription *string               `json:"storyDescription,omitempty"`
	OwnerAddress     *string               `json:"ownerAddress,omitempty"`
	CoverUrl         *string               `json:"coverUrl,omitempty"`
	Content          *string               `json:"content,omitempty"`
	MediaUri         *string               `json:"mediaUri,omitempty"`
	Txhash           *string               `json:"txhash,omitempty"`
	Characters       []*CharacterInfoModel `json:"characters,omitempty"`
	CreatedAt        time.Time             `json:"createdAt,omitempty"`
	UpdatedAt        time.Time             `json:"updatedAt,omitempty"`
}

func (StoryInfoV2Model) TableName() string {
	return "story_info_v2"
}

func (s *StoryInfoV2Model) ToGetStoryResp() (*GetStoryResp, error) {
	if s.StoryId == nil {
		return nil, fmt.Errorf("story id is nil")
	}
	resp := &GetStoryResp{
		StoryId: *s.StoryId,
	}
	if len(s.StoryName) > 0 {
		resp.StoryName = &s.StoryName
	}
	if s.StoryDescription != nil {
		resp.StoryDescription = s.StoryDescription
	}
	if s.OwnerAddress != nil {
		resp.OwnerAddress = s.OwnerAddress
	}
	if s.MediaUri != nil {
		resp.MediaUri = s.MediaUri
	}

	return resp, nil
}

func (s *StoryInfoV2Model) ToGetStoryDetailsResp() (*GetStoryDetailsResp, error) {
	if s.StoryId == nil {
		return nil, fmt.Errorf("story id is nil")
	}
	resp := &GetStoryDetailsResp{
		StoryId: *s.StoryId,
	}
	if len(s.StoryName) > 0 {
		resp.StoryName = &s.StoryName
	}
	if s.StoryDescription != nil {
		resp.StoryDescription = s.StoryDescription
	}
	if s.OwnerAddress != nil {
		resp.OwnerAddress = s.OwnerAddress
	}
	if s.MediaUri != nil {
		resp.MediaUri = s.MediaUri
	}
	if s.Content != nil {
		resp.Content = s.Content
	}
	if s.Txhash != nil {
		resp.Txhash = s.Txhash
	}

	return resp, nil
}

type GetStoryResp struct {
	StoryId          int64   `json:"storyId"`
	StoryName        *string `json:"storyName"`
	StoryDescription *string `json:"storyDescription"`
	OwnerAddress     *string `json:"ownerAddress"`
	MediaUri         *string `json:"arweaveUrl"`
}

type GetStoryDetailsResp struct {
	StoryId          int64   `json:"storyId"`
	StoryName        *string `json:"storyName"`
	StoryDescription *string `json:"storyDescription"`
	OwnerAddress     *string `json:"ownerAddress"`
	Content          *string `json:"content"`
	MediaUri         *string `json:"arweaveUrl"`
	Txhash           *string `json:"txhash"`
}

type CreateStoryRequestBody struct {
	StoryName        *string               `json:"name"`
	StoryDescription *string               `json:"description"`
	OwnerAddress     *string               `json:"owner"`
	Content          *string               `json:"content"`
	Characters       []*CharacterInfoModel `json:"characters"`
}

func (c *CreateStoryRequestBody) Validate() error {
	if c.Content == nil || len(*c.Content) == 0 {
		return fmt.Errorf("Story content is empty")
	}
	return nil
}

func (c *CreateStoryRequestBody) ToStoryInfoV2Model() *StoryInfoV2Model {
	storyInfo := &StoryInfoV2Model{
		ID: uuid.New().String(),
	}
	// Temporary: Add story id for FE testing
	storyId := int64(rand.Uint32())
	storyInfo.StoryId = &storyId

	if c.StoryName != nil {
		storyInfo.StoryName = *c.StoryName
	}
	if c.StoryDescription != nil {
		storyInfo.StoryDescription = c.StoryDescription
	}
	if c.OwnerAddress != nil {
		storyInfo.OwnerAddress = c.OwnerAddress
	}
	if c.Content != nil {
		storyInfo.Content = c.Content
	}
	return storyInfo
}

type StoryMetadata struct {
	Content    string                `json:"content"`
	Characters []*CharacterInfoModel `json:"characters,omitempty"`
}

type CreateStoryResp struct {
	MediaUri string `json:"arweaveURI"`
}

func (f *IpAssetTheGraph) ToStoryInfoV2() (*StoryInfoV2Model, error) {
	franchiseId, err := strconv.ParseInt(f.FranchiseId, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to convert. franchise id: %s is not a valid int64. %v", f.FranchiseId, err)
	}

	ipAssetId, err := strconv.ParseInt(f.IpAssetId, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to convert. ipAssetId: %s is not a valid int64. %v", f.IpAssetId, err)
	}

	story := &StoryInfoV2Model{
		FranchiseId:  franchiseId,
		StoryId:      &ipAssetId,
		OwnerAddress: &f.Owner,
		StoryName:    f.Name,
		MediaUri:     &f.MediaUrl,
		Txhash:       &f.TxHash,
	}
	return story, err
}
