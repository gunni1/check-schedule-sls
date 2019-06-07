package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func Signal(scheduleChange ScheduleChange, sqsQueueURL string) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	sqsClient := sqs.New(sess)

	changeAsJSON, marshalErr := json.Marshal(scheduleChange)
	if marshalErr != nil {
		log.Println("unable to json-marshal schedule change")
		return
	}

	result, err := sqsClient.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(10),
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"DateConcerned": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String(scheduleChange.DateConcerned),
			},
			"IssueDate": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String(scheduleChange.IssueDate),
			},
			"TeacherCode": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String(scheduleChange.TeacherCode),
			},
		},
		MessageBody: aws.String(string(changeAsJSON)),
		QueueUrl:    &sqsQueueURL,
	})
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println("Schedule Change signaled", *result.MessageId)

}
