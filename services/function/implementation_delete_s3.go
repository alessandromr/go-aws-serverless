package function

import (
	"github.com/alessandromr/go-aws-serverless/utils"
	"github.com/alessandromr/go-aws-serverless/utils/auth"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/s3"
	"time"
)

//DeleteDependencies implements the dependencies deletion for S3 Event
func (input S3DeleteFunctionInput) DeleteDependencies(lambdaResult *lambda.DeleteFunctionInput) {
	auth.MakeClient(auth.Sess)
	svc := auth.Client.S3Conn
	lambdaClient := auth.Client.LambdaConn
	var err error

	time.Sleep(utils.LongSleep * time.Millisecond)

	//Remove BucketNotificationConfiguration by setting to empty
	//s3.PutBucketNotificationConfiguration
	putNotConfig := &s3.PutBucketNotificationConfigurationInput{
		Bucket: input.S3DeleteEvent.Bucket,
		NotificationConfiguration: &s3.NotificationConfiguration{
			LambdaFunctionConfigurations: []*s3.LambdaFunctionConfiguration{},
		},
	}

	_, err = svc.PutBucketNotificationConfiguration(putNotConfig)
	utils.CheckAWSErrExpect404(err, "S3 Bucket Notification Configuration")
	time.Sleep(utils.ShortSleep * time.Millisecond)

	//lambda.RemovePermission
	permissionsInput := &lambda.RemovePermissionInput{
		FunctionName: lambdaResult.FunctionName,
		StatementId:  input.StatementId,
	}
	_, err = lambdaClient.RemovePermission(permissionsInput)
	utils.CheckAWSErrExpect404(err, "Lambda S3 Permission")
}

//GetFunctionInput return the DeleteFunctionInput from the custom input
func (input S3DeleteFunctionInput) GetFunctionInput() *lambda.DeleteFunctionInput {
	return input.FunctionInput
}
