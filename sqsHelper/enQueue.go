package sqsHelper

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func Enqueue(message SqsModel) error {

	svc := getSqsSession()
	qURL := getQueueUrl()

	e, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Error in json marshall")
		return err
	}

	json_string := string(e)

	result, err := svc.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(10),
		MessageBody:  aws.String(json_string),
		QueueUrl:     &qURL,
	})

	if err != nil {
		fmt.Println("Error", err)
		return err
	}

	fmt.Println("Success", *result.MessageId)
	return err
}
