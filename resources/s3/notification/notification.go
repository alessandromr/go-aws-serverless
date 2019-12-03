package notification

//S3NotificationConfiguration 
type S3NotificationConfiguration struct {
	Bucket string
	Events []string
	FunctionArn string
	FilterRules []s3.FilterRule
}


//Delete the given resources
func (resource S3NotificationConfiguration) Delete() error {
	putNotConfig := &s3.PutBucketNotificationConfigurationInput{
		Bucket: resource.Bucket,
		NotificationConfiguration: &s3.NotificationConfiguration{
			LambdaFunctionConfigurations: []*s3.LambdaFunctionConfiguration{
				LambdaFunctionArn: &FunctionArn,
			},
		},
	}
	_, err = svc.PutBucketNotificationConfiguration(putNotConfig)
	return err
}

