package notification

import (
	"github.com/alessandromr/go-aws-serverless/utils/auth"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

//S3NotificationConfiguration
type S3NotificationConfiguration struct {
	Bucket      string
	Events      []string
	FunctionArn string
	FilterRules []s3.FilterRule
}

//Create the given resources
func (resource *S3NotificationConfiguration) Create() error {
	auth.MakeClient(auth.Sess)
	svc := auth.Client.S3Conn

	putNotConfig := &s3.PutBucketNotificationConfigurationInput{
		Bucket: aws.String(resource.Bucket),
		NotificationConfiguration: &s3.NotificationConfiguration{
			LambdaFunctionConfigurations: []*s3.LambdaFunctionConfiguration{
				{
					LambdaFunctionArn: aws.String(resource.FunctionArn),
					Events:            aws.StringSlice(resource.Events),
				},
			},
		},
	}
	_, err := svc.PutBucketNotificationConfiguration(putNotConfig)
	return err
}

//Delete the given resources
func (resource S3NotificationConfiguration) Delete() error {
	auth.MakeClient(auth.Sess)
	svc := auth.Client.S3Conn
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
