package upload

import (
	"fmt"
	"os"

	"github.com/h-yamada/miniosample/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Upload struct {
	client *s3.S3
}

func (u *Upload) Upload(bucket string, localFile string, s3File string) error {
	fp, err := os.Open(localFile)
	defer fp.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = u.client.PutObject(&s3.PutObjectInput{
		Body:   fp,
		Bucket: &bucket,
		Key:    &s3File,
	})
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func NewUpload() *Upload {
	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(config.AccessKey, config.SecretKey, ""),
		Endpoint:         aws.String(config.Endpoint),
		Region:           aws.String(config.Region),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(true),
	}
	newSession := session.New(s3Config)

	return &Upload{client: s3.New(newSession)}
}
