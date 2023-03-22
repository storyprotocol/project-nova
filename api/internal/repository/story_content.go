package repository

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/project-nova/backend/pkg/constant"
	"github.com/project-nova/backend/pkg/model"
	"github.com/project-nova/backend/pkg/s3"
)

type StoryContentRepository interface {
	GetContentByChapter(franchiseId int64, storyNum int, chapterNum int) (*model.StoryContentModel, error)
}

func NewStoryContentS3Impl(s3Client s3.S3Client, bucket string) (StoryContentRepository, error) {
	contentMap := make(map[string]*model.StoryContentModel)

	keys, err := s3Client.ListObjectsNonRecursive(bucket)
	if err != nil {
		return nil, fmt.Errorf("failed to list objects for %s: %v", bucket, err)
	}

	for _, key := range keys {
		buf := aws.NewWriteAtBuffer([]byte{})
		_, err := s3Client.DownloadObject(buf, bucket, *key+"/"+constant.ContentObject)
		if err != nil {
			return nil, fmt.Errorf("failed to download content from s3 for object %s: %v", *key, err)
		}

		var data model.StoryContentModel
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
	contentMap map[string]*model.StoryContentModel
}

func (s *storyContentS3Impl) GetContentByChapter(franchiseId int64, storyNum int, chapterNum int) (*model.StoryContentModel, error) {
	key := fmt.Sprintf("%d:%d:%d", franchiseId, storyNum, chapterNum)
	val, ok := s.contentMap[key]
	if !ok {
		return nil, fmt.Errorf("content not found, key: %s", key)
	}

	return val, nil
}

func NewStoryContentFsImpl(contentFilePath string) (StoryContentRepository, error) {
	contentMap := make(map[string]*model.StoryContentModel)

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

			var data *model.StoryContentModel
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
	contentMap map[string]*model.StoryContentModel
}

func (s *storyContentFsImpl) GetContentByChapter(franchiseId int64, storyNum int, chapterNum int) (*model.StoryContentModel, error) {
	key := fmt.Sprintf("%d:%d:%d", franchiseId, storyNum, chapterNum)
	val, ok := s.contentMap[key]
	if !ok {
		return nil, fmt.Errorf("content not found, key: %s", key)
	}

	return val, nil
}
