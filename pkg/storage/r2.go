package storage

import (
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type R2Storage struct {
	svc    *s3.S3
	bucket string
	url    string
}

func NewR2Storage(accessKey, secretKey, accountID, bucket string) (*R2Storage, error) {
	endpoint := fmt.Sprintf("https://%s.r2.cloudflarestorage.com", accountID)

	sess, err := session.NewSession(&aws.Config{
		Region:           aws.String("auto"),
		Credentials:      credentials.NewStaticCredentials(accessKey, secretKey, ""),
		Endpoint:         aws.String(endpoint),
		S3ForcePathStyle: aws.Bool(true),
	})
	if err != nil {
		return nil, err
	}

	svc := s3.New(sess)

	return &R2Storage{
		svc:    svc,
		bucket: bucket,
		url:    endpoint,
	}, nil
}

func (r *R2Storage) UploadFile(key string, body io.ReadSeeker, contentType string) error {
	input := &s3.PutObjectInput{
		Body:        body,
		Bucket:      aws.String(r.bucket),
		Key:         aws.String(key),
		ContentType: aws.String(contentType),
	}

	_, err := r.svc.PutObject(input)
	if err != nil {
		return err
	}

	return nil
}
