package main

import (
	"flag"
	"fmt"

	"github.com/h-yamada/miniosample/download"
	"github.com/h-yamada/miniosample/upload"
)

func main() {
	commandType := flag.String("type", "", "command type d:Download u:Upload")
	bucket := flag.String("bucket", "", "S3 Bucket")
	localFile := flag.String("file", "", "local file")
	s3File := flag.String("key", "", "S3 file")
	flag.Parse()

	switch *commandType {
	case "u":
		u := upload.NewUpload()
		err := u.Upload(*bucket, *localFile, *s3File)
		if err == nil {
			fmt.Println("Upload: Success.")
		}
	case "d":
		d := download.NewDownload()
		err := d.Download(*bucket, *localFile, *s3File)
		if err == nil {
			fmt.Println("Download: Success.")
		}
	}

}
