package main

import (
	"flag"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {

	var bucketname = flag.String("b", "invalid", "bucket to list files")
	var folder = flag.String("fd", "invalid", "folder of bucket")
	var file = flag.String("f", "invalid", "file to check")
	flag.Parse()
	fmt.Println("bucket name - ", *bucketname)
	fmt.Println("folder name - ", *folder)

	var sess = connectAWS("us-west-1")
	svc := s3.New(sess)

	filefull := *folder + "/" + *file
	output, err := svc.HeadObject(&s3.HeadObjectInput{
		Bucket: aws.String(*bucketname),
		Key:    aws.String(filefull),
	})

	if err != nil {
		fmt.Println("Object doesnt exist")
	} else {
		fmt.Println("Object exists !! - ", output)
	}
}

func connectAWS(region string) *session.Session {
	sess, err := session.NewSession(&aws.Config{Region: aws.String(region)})
	if err != nil {
		panic(err)
	}
	return sess
}
