package sqsHelper

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func DeQueue(delete bool) (SqsModel, error) {

	svc := getSqsSession()
	qURL := getQueueUrl()

	result, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		AttributeNames: []*string{
			aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
		},
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		QueueUrl:            &qURL,
		MaxNumberOfMessages: aws.Int64(1),
		VisibilityTimeout:   aws.Int64(20), // 20 seconds
		WaitTimeSeconds:     aws.Int64(2),
	})

	if err != nil {
		fmt.Println("Error", err)
		return SqsModel{}, err
	}

	if len(result.Messages) == 0 {
		fmt.Println("Received no messages")
		return SqsModel{}, err
	}

	message := SqsModel{}

	if len(result.Messages) != 1 {
		fmt.Println("Received more than one message")
		return message, errors.New("More than one msg")
	}

	m := result.Messages[0]
	body := *m.Body
	bytes := []byte(body)
	json.Unmarshal(bytes, &message)

	if delete {
		resultDelete, err := svc.DeleteMessage(&sqs.DeleteMessageInput{
			QueueUrl:      &qURL,
			ReceiptHandle: m.ReceiptHandle,
		})

		if err != nil {
			fmt.Println("Delete Error", err)
			return message, err
		}

		fmt.Println("Message Deleted")
		fmt.Println(*resultDelete)
	}

	fmt.Println("Successfully dequeued")
	//fmt.Println(message)
	return message, err

}
