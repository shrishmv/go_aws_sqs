package main

import (
	"flag"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {

	var bucketname = flag.String("b", "invalid", "bucket to list files")
	var folder = flag.String("fd", "invalid", "folder of bucket")
	flag.Parse()
	fmt.Println("bucket name - ", *bucketname)
	fmt.Println("folder name - ", *folder)

	var sess = connectAWS("us-west-1")

	svc := s3.New(sess)
	input := &s3.ListObjectsInput{
		Bucket: aws.String(*bucketname),
		Prefix: aws.String((*folder)),
	}

	result, err := svc.ListObjects(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchBucket:
				fmt.Println(s3.ErrCodeNoSuchBucket, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		// Do your error handling here
		return
	}

	if len(result.Contents) > 0 {
		fmt.Println("There are ", len(result.Contents), " items in bucket")
	}

	for _, item := range result.Contents {
		fmt.Println("File - ", *item.Key)
	}
}

func connectAWS(region string) *session.Session {
	sess, err := session.NewSession(&aws.Config{Region: aws.String(region)})
	if err != nil {
		panic(err)
	}
	return sess
}
