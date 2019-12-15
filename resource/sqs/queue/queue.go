package queue

import (
	"github.com/alessandromr/go-aws-serverless/utils/auth"
	"github.com/aws/aws-sdk-go/service/sqs"
)

//SQSQueue
type SQSQueue struct {
	Existing                      bool
	QueueUrl                      *string
	QueueName                     *string
	DelaySeconds                  *int
	MaximumMessageSize            *int
	MessageRetentionPeriod        *int
	ReceiveMessageWaitTimeSeconds *int
	VisibilityTimeout             *int
	FifoQueue                     bool
	KmsMasterKeyId                *string
	ContentBasedDeduplication     bool
	// Policy
	// RedrivePolicy
}

//Delete the given resources
func (resource SQSQueue) Delete() error {
	auth.MakeClient(auth.Sess)
	svc := auth.Client.SQSconn

	deleteQueueInput := &sqs.DeleteQueueInput{
		QueueUrl: resource.QueueUrl,
	}
	_, err := svc.DeleteQueue(deleteQueueInput)
	return err
}
