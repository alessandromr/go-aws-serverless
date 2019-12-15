package function

import (
	"time"

	"github.com/alessandromr/go-aws-serverless/manager/rollback"
	"github.com/alessandromr/go-aws-serverless/resource/lambda/permission"
	"github.com/alessandromr/go-aws-serverless/resource/s3/notification"
	"github.com/alessandromr/go-aws-serverless/utils"
	"github.com/alessandromr/go-aws-serverless/utils/auth"
	"github.com/alessandromr/go-aws-serverless/utils/convert"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/s3"
)

//CreateDependencies create all the dependencies for S3Event
func (input S3CreateFunctionInput) CreateDependencies(lambdaResult *lambda.FunctionConfiguration) (map[string]interface{}, error) {
	auth.MakeClient(auth.Sess)
	svc := auth.Client.S3Conn
	lambdaClient := auth.Client.LambdaConn
	var err error

	//lambda.AddPermission
	permissionsInput := &lambda.AddPermissionInput{
		Action:       aws.String("lambda:InvokeFunction"),
		FunctionName: lambdaResult.FunctionArn,
		Principal:    aws.String("s3.amazonaws.com"),
		SourceArn:    aws.String("arn:aws:s3:::" + *input.S3CreateEvent.Bucket),
		StatementId:  aws.String("S3Event_" + *input.S3CreateEvent.Bucket + "_" + *lambdaResult.FunctionName),
	}
	permissionsOutput, err := lambdaClient.AddPermission(permissionsInput)
	if err != nil {
		rollback.ExecuteRollback()
		return nil, err
	}
	rollback.ResourcesList = append(
		rollback.ResourcesList,
		permission.LambdaPermission{
			StatementId:  "S3Event_" + *input.S3CreateEvent.Bucket + "_" + *lambdaResult.FunctionName,
			FunctionName: *lambdaResult.FunctionArn,
		},
	)

	time.Sleep(utils.ShortSleep * time.Millisecond)

	//s3.PutBucketNotificationConfiguration
	putNotConfig := &s3.PutBucketNotificationConfigurationInput{
		Bucket: input.S3CreateEvent.Bucket,
		NotificationConfiguration: &s3.NotificationConfiguration{
			LambdaFunctionConfigurations: []*s3.LambdaFunctionConfiguration{
				{
					LambdaFunctionArn: lambdaResult.FunctionArn,
					Events:            input.S3CreateEvent.Types,
				},
			},
		},
	}
	_, err = svc.PutBucketNotificationConfiguration(putNotConfig)
	if err != nil {
		rollback.ExecuteRollback()
		return nil, err
	}
	rollback.ResourcesList = append(
		rollback.ResourcesList,
		notification.S3NotificationConfiguration{
			Bucket:      *input.S3CreateEvent.Bucket,
			Events:      convert.StringSlice(input.S3CreateEvent.Types),
			FunctionArn: *lambdaResult.FunctionArn,
		},
	)
	out := make(map[string]interface{})
	out["Bucket"] = *input.S3CreateEvent.Bucket
	out["LambdaPermission"] = permissionsOutput.Statement
	return out, nil

}

//GetFunctionInput return the CreateFunctionInput from the custom input
func (input S3CreateFunctionInput) GetFunctionInput() *lambda.CreateFunctionInput {
	return input.FunctionInput
}
