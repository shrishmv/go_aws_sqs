package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func main() {
	fmt.Println("vim-go")

	var filename = flag.String("f", "invalid", "filename to download")
	var bucketname = flag.String("b", "invalid", "bucket to download from")
	var folder = flag.String("fd", "invalid", "folder of bucket")
	flag.Parse()
	fmt.Println("download file name - ", *filename)
	fmt.Println("bucket name - ", *bucketname)
	fmt.Println("folder name - ", *folder)

	var sess = connectAWS("us-west-1")

	f, err := os.Create(*filename)
	if err != nil {
		fmt.Println("Error in file create")
		return
	}

	downloader := s3manager.NewDownloader(sess)
	_, err = downloader.Download(f, &s3.GetObjectInput{
		Bucket: aws.String(*bucketname),
		//Key:    aws.String((*folder) + "/" + (*filename)),
		Key: aws.String((*filename)),
	})
	if err != nil {
		fmt.Println("Error in file download")
		fmt.Println(err)
		return
	}
}

func connectAWS(region string) *session.Session {
	sess, err := session.NewSession(&aws.Config{Region: aws.String(region)})
	if err != nil {
		panic(err)
	}
	return sess
}
