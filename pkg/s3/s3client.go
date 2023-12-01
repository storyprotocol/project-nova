package s3

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/project-nova/backend/pkg/logger"
)

type S3Client interface {
	RequestPreSignedUrl(bucket string, key string) (string, error)
	DownloadObject(output io.WriterAt, bucket string, key string) (int64, error)
	UploadObject(bucket string, key string, filename string, public bool) (*s3manager.UploadOutput, error)
	ListObjectsNonRecursive(bucket string) ([]*string, error)
}

type s3Client struct {
	client     *s3.S3
	downloader *s3manager.Downloader
	uploader   *s3manager.Uploader
}

func NewS3Client(sess *session.Session) S3Client {
	return &s3Client{
		client:     s3.New(sess),
		downloader: s3manager.NewDownloader(sess),
		uploader:   s3manager.NewUploader(sess),
	}
}

func (s *s3Client) RequestPreSignedUrl(bucket string, key string) (string, error) {
	req, _ := s.client.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})

	url, err := req.Presign(150 * 60) // 15 minutes
	if err != nil {
		return "", fmt.Errorf("failed to sign request: %v", err)
	}

	return url, nil
}

// DownloadObject downloads a file from S3 based on the input: bucket and key, and write to the output.
// The output can be of os.File type or WriteAtBuffer type, based on where the downloaded file should be kept
func (s *s3Client) DownloadObject(output io.WriterAt, bucket string, key string) (int64, error) {
	numBytes, err := s.downloader.Download(output,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(key),
		})

	return numBytes, err
}

func (s *s3Client) UploadObject(bucket string, key string, filename string, public bool) (*s3manager.UploadOutput, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %v. error: %v", filename, err)
	}

	input := &s3manager.UploadInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(key),
		Body:        file,
		ContentType: aws.String(getContentType(filename)),
	}

	if public {
		input.ACL = aws.String("public-read")
	}

	output, err := s.uploader.Upload(input)
	if err != nil {
		return nil, fmt.Errorf("unable to upload %v to %v: %v", filename, bucket, err)
	}

	logger.Infof("upload object %v to %v succeeded", filename, bucket)
	return output, nil
}

func (s *s3Client) ListObjectsNonRecursive(bucket string) ([]*string, error) {
	input := &s3.ListObjectsInput{
		Bucket:    aws.String(bucket),
		Delimiter: aws.String("/"),
	}

	resp, err := s.client.ListObjects(input)
	if err != nil {
		return nil, fmt.Errorf("failed to list s3 objects: %v", err)
	}

	var keys []*string
	for _, prefix := range resp.CommonPrefixes {
		if prefix.Prefix != nil {
			prefixProcessed := strings.TrimSuffix(*prefix.Prefix, "/")
			keys = append(keys, &prefixProcessed)
		}
	}

	return keys, nil
}
