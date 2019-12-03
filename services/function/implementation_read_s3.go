package function

import (
	"github.com/alessandromr/go-serverless-client/utils/auth"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
)

//ReadDependencies implements the dependencies deletion for S3 Event
func (input S3ReadFunctionInput) ReadDependencies(lambdaResult *lambda.FunctionConfiguration) map[string]interface{} {
	svc := s3.New(auth.Sess)

	notInput := &s3.GetBucketNotificationConfigurationRequest{
		Bucket: input.S3ReadEvent.Bucket,
	}
	response, err := svc.GetBucketNotificationConfiguration(notInput)
	if err != nil {
		log.Fatal(err)
	}

	out := make(map[string]interface{})
	out["Bucket"] = *input.S3ReadEvent.Bucket
	out["S3LambdaFunctionConfigurations"] = response.LambdaFunctionConfigurations
	return out
}

//GetFunctionInput return the ReadFunctionInput from the custom input
func (input S3ReadFunctionInput) GetFunctionConfiguration() *lambda.GetFunctionConfigurationInput {
	return input.FunctionConfigurationInput
}
