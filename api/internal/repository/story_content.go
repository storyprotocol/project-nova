package repository

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/project-nova/backend/api/internal/constant"
	"github.com/project-nova/backend/pkg/s3"
)

type StoryContentRepository interface {
	GetContentByChapter(franchiseId int64, storyNum int, chapterNum int) (*StoryContentModel, error)
}

type StoryContentModel struct {
	Title      string               `json:"title"`
	Heading    string               `json:"heading"`
	CoverImage string               `json:"coverImage"`
	Content    []*StorySectionModel `json:"content"`
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

func NewStoryContentS3Impl(s3Client s3.S3Client, bucket string) (StoryContentRepository, error) {
	contentMap := make(map[string]*StoryContentModel)

	keys, err := s3Client.ListObjectsNonRecursive(bucket)
	if err != nil {
		return nil, fmt.Errorf("failed to list objects for %s: %v", bucket, err)
	}

	for _, key := range keys {
		buf := aws.NewWriteAtBuffer([]byte{})
		_, err := s3Client.DownloadObject(buf, bucket, *key+"/"+constant.S3ContentObject)
		if err != nil {
			return nil, fmt.Errorf("failed to download content from s3 for object %s: %v", *key, err)
		}

		var data StoryContentModel
		err = json.Unmarshal(buf.Bytes(), &data)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal story content file data, error: %v", err)
		}
		contentMap[*key] = &data
	}

	return &storyContentS3Impl{
		s3Client:   s3Client,
		contentMap: contentMap,
	}, nil
}

type storyContentS3Impl struct {
	s3Client   s3.S3Client
	contentMap map[string]*StoryContentModel
}

func (s *storyContentS3Impl) GetContentByChapter(franchiseId int64, storyNum int, chapterNum int) (*StoryContentModel, error) {
	key := fmt.Sprintf("%d:%d:%d", franchiseId, storyNum, chapterNum)
	val, ok := s.contentMap[key]
	if !ok {
		return nil, fmt.Errorf("content not found, key: %s", key)
	}

	return val, nil
}

func NewStoryContentFsImpl(contentFilePath string) (StoryContentRepository, error) {
	contentMap := make(map[string]*StoryContentModel)

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

			var data *StoryContentModel
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
	contentMap map[string]*StoryContentModel
}

func (s *storyContentFsImpl) GetContentByChapter(franchiseId int64, storyNum int, chapterNum int) (*StoryContentModel, error) {
	key := fmt.Sprintf("%d:%d:%d", franchiseId, storyNum, chapterNum)
	val, ok := s.contentMap[key]
	if !ok {
		return nil, fmt.Errorf("content not found, key: %s", key)
	}

	return val, nil
}
