package function

import (
	"github.com/alessandromr/goserverlessclient/utils"
	"github.com/alessandromr/goserverlessclient/utils/auth"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/s3"
)

//CreateDependencies create all the dependencies for S3Event
func (input S3CreateFunctionInput) CreateDependencies(lambdaResult *lambda.FunctionConfiguration) map[string]interface{} {
	svc := s3.New(auth.Sess)
	lambdaClient := lambda.New(auth.Sess)
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
	utils.CheckErr(err)

	//s3.PutBucketNotificationConfiguration
	putNotConfig := &s3.PutBucketNotificationConfigurationInput{
		Bucket: input.S3CreateEvent.Bucket,
		NotificationConfiguration: &s3.NotificationConfiguration{
			LambdaFunctionConfigurations: []*s3.LambdaFunctionConfiguration{
				&s3.LambdaFunctionConfiguration{
					LambdaFunctionArn: lambdaResult.FunctionArn,
					Events:            input.S3CreateEvent.Types,
				},
			},
		},
	}
	_, err = svc.PutBucketNotificationConfiguration(putNotConfig)
	utils.CheckErr(err)

	out := make(map[string]interface{})
	out["Bucket"] = *input.S3CreateEvent.Bucket
	out["LambdaPermission"] = permissionsOutput.Statement
	return out

}

//GetFunctionInput return the CreateFunctionInput from the custom input
func (input S3CreateFunctionInput) GetFunctionInput() *lambda.CreateFunctionInput {
	return input.FunctionInput
}
