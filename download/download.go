package download

import (
	"fmt"
	"os"

	"github.com/h-yamada/miniosample/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type Download struct {
	downloader *s3manager.Downloader
}

func (d *Download) Download(bucket string, localFile string, s3File string) error {
	file, err := os.Create(localFile)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()

	_, err = d.downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: &bucket,
			Key:    &s3File,
		})

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func NewDownload() *Download {
	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(config.AccessKey, config.SecretKey, ""),
		Endpoint:         aws.String(config.Endpoint),
		Region:           aws.String(config.Region),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(true),
	}
	newSession := session.New(s3Config)

	return &Download{downloader: s3manager.NewDownloader(newSession)}
}
