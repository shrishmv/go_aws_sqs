package main

import (
	"bytes"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {

	var filename = flag.String("f", "invalid", "filename to upload")
	var bucketname = flag.String("b", "invalid", "bucket to upload to")
	flag.Parse()
	fmt.Println("file name - ", *filename)
	fmt.Println("bucket name - ", *bucketname)

	aws_access_key_id := "youaccessid"
	aws_secret_access_key := "your access key"
	token := ""
	creds := credentials.NewStaticCredentials(aws_access_key_id, aws_secret_access_key, token)
	_, err := creds.Get()
	if err != nil {
		// handle error
	}
	cfg := aws.NewConfig().WithRegion("us-west-1").WithCredentials(creds)
	svc := s3.New(session.New(), cfg)

	file, err := os.Open(*filename)
	if err != nil {
		// handle error
	}
	defer file.Close()
	fileInfo, _ := file.Stat()
	size := fileInfo.Size()
	buffer := make([]byte, size) // read file content to buffer

	file.Read(buffer)
	fileBytes := bytes.NewReader(buffer)
	fileType := http.DetectContentType(buffer)
	path := "/blah/" + file.Name()
	params := &s3.PutObjectInput{
		Bucket:        aws.String(*bucketname),
		Key:           aws.String(path),
		Body:          fileBytes,
		ContentLength: aws.Int64(size),
		ContentType:   aws.String(fileType),
	}
	resp, err := svc.PutObject(params)
	if err != nil {
		// handle error
	}
	fmt.Printf("response %s", awsutil.StringValue(resp))
}
