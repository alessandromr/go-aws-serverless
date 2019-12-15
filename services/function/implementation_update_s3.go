package function

import (
	"github.com/alessandromr/go-aws-serverless/utils"
	"github.com/alessandromr/go-aws-serverless/utils/auth"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
	"time"
)

//UpdateDependencies implements the dependencies deletion for S3 Event
func (input S3UpdateFunctionInput) UpdateDependencies(lambdaResult *lambda.FunctionConfiguration) (map[string]interface{}, error) {
	auth.MakeClient(auth.Sess)
	svc := auth.Client.S3Conn
	lambdaClient := auth.Client.LambdaConn
	var err error

	time.Sleep(utils.LongSleep * time.Millisecond)

	//lambda.RemovePermission (remove old permission)
	removePermissionsInput := &lambda.RemovePermissionInput{
		FunctionName: lambdaResult.FunctionName,
		StatementId:  input.StatementId,
	}
	_, err = lambdaClient.RemovePermission(removePermissionsInput)
	utils.CheckAWSErrExpect404(err, "Lambda S3 Permission")

	time.Sleep(utils.ShortSleep * time.Millisecond)

	// lambda.AddPermission
	addPermissionsInput := &lambda.AddPermissionInput{
		Action:       aws.String("lambda:InvokeFunction"),
		FunctionName: lambdaResult.FunctionName,
		Principal:    aws.String("s3.amazonaws.com"),
		SourceArn:    aws.String("arn:aws:s3:::" + *input.S3UpdateEvent.Bucket),
		StatementId:  aws.String("S3Event_" + *input.S3UpdateEvent.Bucket + "_" + *lambdaResult.FunctionName),
	}
	permissionsOutput, err := lambdaClient.AddPermission(addPermissionsInput)
	if err != nil {
		log.Println("Error") //ToDo
	}

	time.Sleep(utils.ShortSleep * time.Millisecond)

	//s3.PutBucketNotificationConfiguration
	putNotConfig := &s3.PutBucketNotificationConfigurationInput{
		Bucket: input.S3UpdateEvent.Bucket,
		NotificationConfiguration: &s3.NotificationConfiguration{
			LambdaFunctionConfigurations: []*s3.LambdaFunctionConfiguration{
				{
					LambdaFunctionArn: lambdaResult.FunctionName,
					Events:            input.S3UpdateEvent.Types,
				},
			},
		},
	}
	_, err = svc.PutBucketNotificationConfiguration(putNotConfig)
	if err != nil {
		log.Println("Error") //ToDo
	}

	out := make(map[string]interface{})
	out["LambdaPermission"] = permissionsOutput.Statement
	return out, nil
}

//GetUpdateFunctionConfiguration return the UpdateFunctionConfigurationInput from the custom input
func (input S3UpdateFunctionInput) GetUpdateFunctionConfiguration() *lambda.UpdateFunctionConfigurationInput {
	return input.UpdateFunctionConfigurationInput
}
