package function

import (
	"github.com/alessandromr/go-aws-serverless/utils/auth"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
)

//ReadDependencies implements the dependencies deletion for S3 Event
func (input S3ReadFunctionInput) ReadDependencies(lambdaResult *lambda.FunctionConfiguration) map[string]interface{} {
	auth.MakeClient(auth.Sess)
	svc := auth.Client.S3Conn

	notInput := &s3.GetBucketNotificationConfigurationRequest{
		Bucket: input.S3ReadEvent.Bucket,
	}
	response, err := svc.GetBucketNotificationConfiguration(notInput)
	if err != nil {
		log.Fatal(err)
	}

	out := make(map[string]interface{})
	out["Bucket"] = *input.S3ReadEvent.Bucket
	out["StatementId"] = "S3Event_" + *input.S3ReadEvent.Bucket + "_" + *lambdaResult.FunctionName //ToDo improve
	out["S3Events"] = response.LambdaFunctionConfigurations[0].Events                              //ToDo improve [0]
	out["S3LambdaFunctionArn"] = response.LambdaFunctionConfigurations[0].LambdaFunctionArn
	return out
}

//GetFunctionInput return the ReadFunctionInput from the custom input
func (input S3ReadFunctionInput) GetFunctionConfiguration() *lambda.GetFunctionConfigurationInput {
	return input.FunctionConfigurationInput
}
