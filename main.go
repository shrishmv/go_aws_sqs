package main

import (
	"fmt"
	"go_aws_sqs/sqsHelper"
	"time"
)

func main() {
	fmt.Println("Go aws SQS example.....")

	message := sqsHelper.SqsModel{}
	message.CallUUID = "tom-and-jerry-kids"
	message.ChunkName = "wohoh-hoh-hohohooo"
	message.ChunkIndex = 0
	message.ChunkCount = 12

	err := sqsHelper.Enqueue(message)
	if err != nil {
		fmt.Println("Error in enquing !")
	}

	//loop receive
	for {
		dmsg, err := sqsHelper.DeQueue(true)
		if err != nil {
			fmt.Println("Error in dequeue, !!")
		} else if (dmsg == sqsHelper.SqsModel{}) {
			time.Sleep(1000 * time.Millisecond)
			fmt.Println("Empty queue sleeping !!")
		} else {
			fmt.Println(dmsg)
		}
	}
}
