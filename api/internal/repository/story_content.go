package repository

import (
	"encoding/json"
	"fmt"
	"os"
)

type StorySectionModel struct {
	Type string             `json:"type"`
	Data []*StoryMediaModel `json:"data"`
}

type StoryMediaModel struct {
	Type        string `json:"type"`
	Content     string `json:"content"`
	Url         string `json:"url"`
	Description string `json:"description,omitempty"`
}

type StoryContentRepository interface {
	GetContentByChapter(franchiseId int64, storyNum int, chapterNum int) ([]*StorySectionModel, error)
}

func NewStoryContentFsImpl(contentFilePath string) (StoryContentRepository, error) {
	file, err := os.ReadFile(contentFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read the story content file: %s, error: %v", contentFilePath, err)
	}

	var data []*StorySectionModel
	err = json.Unmarshal([]byte(file), &data)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal story content file data, error: %v", err)
	}

	contentMap := make(map[string][]*StorySectionModel)
	contentMap["1:1:1"] = data

	return &storyContentFsImpl{
		contentMap: contentMap,
	}, nil
}

type storyContentFsImpl struct {
	contentMap map[string][]*StorySectionModel
}

func (s *storyContentFsImpl) GetContentByChapter(franchiseId int64, storyNum int, chapterNum int) ([]*StorySectionModel, error) {
	key := fmt.Sprintf("%d:%d:%d", franchiseId, storyNum, chapterNum)
	val, ok := s.contentMap[key]
	if !ok {
		return nil, fmt.Errorf("content not found, key: %s", key)
	}

	return val, nil
}
