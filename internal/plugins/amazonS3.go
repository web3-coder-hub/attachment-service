package plugins

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"io"
)

type AwsS3 struct {
	BuketName   *string
	KeyName     *string
	ContentType *string
	Body        io.Reader
}

func UploaderS3PublicRead(s3 *AwsS3) (*s3manager.UploadOutput, error) {
	uploader := s3manager.NewUploader(sess)
	upParams := &s3manager.UploadInput{
		Bucket:      s3.BuketName,
		Key:         s3.KeyName,
		Body:        s3.Body,
		ContentType: s3.ContentType,
		ACL:         aws.String("public-read"),
	}
	// Perform an upload.
	d, err := uploader.Upload(upParams)

	return d, err
}
