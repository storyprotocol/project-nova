package repository

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type StoryContentRepository interface {
	GetContentByChapter(franchiseId int64, storyNum int, chapterNum int) ([]*StorySectionModel, error)
}

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

func NewStoryContentFsImpl(contentFilePath string) (StoryContentRepository, error) {
	contentMap := make(map[string][]*StorySectionModel)

	dirs, err := os.ReadDir(contentFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory %s, error: %v", contentFilePath, err)
	}
	// Load contents from the content directories
	for _, dir := range dirs {
		dirName := contentFilePath + "/" + dir.Name()
		files, err := os.ReadDir(dirName)
		if err != nil {
			return nil, fmt.Errorf("failed to read directory %s, error: %v", dirName, err)
		}
		for _, file := range files {
			rawData, err := os.ReadFile(dirName + "/" + file.Name())
			if err != nil {
				return nil, fmt.Errorf("failed to read file %s, error: %v", file.Name(), err)
			}

			var data []*StorySectionModel
			err = json.Unmarshal([]byte(rawData), &data)
			if err != nil {
				return nil, fmt.Errorf("failed to unmarshal story content file data, error: %v", err)
			}

			fileName := strings.Split(file.Name(), ".")[0]
			contentKey := dir.Name() + ":" + fileName
			contentMap[contentKey] = data
		}
	}

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
