package notification

import (
	"github.com/alessandromr/go-aws-serverless/utils/auth"
	"github.com/aws/aws-sdk-go/service/s3"
)

//S3NotificationConfiguration
type S3NotificationConfiguration struct {
	Bucket      string
	Events      []string
	FunctionArn string
	FilterRules []s3.FilterRule
}

//Delete the given resources
func (resource S3NotificationConfiguration) Delete() error {
	auth.MakeClient(auth.Sess)
	svc := auth.Client.S3conn
	putNotConfig := &s3.PutBucketNotificationConfigurationInput{
		Bucket: &resource.Bucket,
		NotificationConfiguration: &s3.NotificationConfiguration{
			LambdaFunctionConfigurations: []*s3.LambdaFunctionConfiguration{
				{
					LambdaFunctionArn: &resource.FunctionArn,
				},
			},
		},
	}
	_, err := svc.PutBucketNotificationConfiguration(putNotConfig)
	return err
}
