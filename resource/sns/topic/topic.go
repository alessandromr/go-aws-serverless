package topic

import (
	"github.com/alessandromr/go-aws-serverless/utils/auth"
	"github.com/aws/aws-sdk-go/service/sns"
)

//SNSTopic
type SNSTopic struct {
	Existing       bool
	TopicArn       *string
	DisplayName    *string
	Owner          *string
	KmsMasterKeyId *string
	// Policy
}

//Delete the given resources
func (resource SNSTopic) Delete() error {
	auth.MakeClient(auth.Sess)
	svc := auth.Client.SNSconn

	deleteTopicInput := &sns.DeleteTopicInput{
		TopicArn: resource.TopicArn,
	}
	_, err := svc.DeleteTopic(deleteTopicInput)
	return err
}
