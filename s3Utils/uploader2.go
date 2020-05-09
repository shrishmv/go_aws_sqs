package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func main() {

	var filename = flag.String("f", "invalid", "filename to upload")
	var bucketname = flag.String("b", "invalid", "bucket to upload to")
	flag.Parse()
	fmt.Println("file name - ", *filename)
	fmt.Println("bucket name - ", *bucketname)

	file, err := os.Open(*filename)
	if err != nil {
		// Do your error handling here
		fmt.Println("Error in file open !!")
		return
	}
	defer file.Close()

	var sess = connectAWS("us-west-1")
	uploader := s3manager.NewUploader(sess)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(*bucketname), // Bucket to be used
		Key:    aws.String(*filename),   // Name of the file to be saved
		Body:   file,                    // File
	})
	if err != nil {
		// Do your error handling here
		fmt.Println("Error in uploading !!")
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

func check(e error) {
	if e != nil {
		panic(e)
	}
}
