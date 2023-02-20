package s3

import (
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3Client interface {
	DownloadObject(output io.WriterAt, bucket string, key string) (int64, error)
}

type s3Client struct {
	downloader *s3manager.Downloader
}

func NewS3Client(sess *session.Session) S3Client {
	return &s3Client{
		downloader: s3manager.NewDownloader(sess),
	}
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
