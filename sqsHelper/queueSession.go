package sqsHelper

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func getQueueUrl() string {
	// URL to our queue
	qURL := "https://your-aws-sqs-url"
	return qURL
}

func getSqsSession() *sqs.SQS {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sqs.New(sess)
	return svc
}
