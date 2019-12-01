package function

import (
	"github.com/aws/aws-sdk-go/service/lambda"
)

//ReadFunctionInput is an interface to Read a serverless function and the relative trigger
type ReadFunctionInput interface {
	ReadDependencies(*lambda.FunctionConfiguration) map[string]interface{}
	GetFunctionConfiguration() *lambda.GetFunctionConfigurationInput
}

//HTTPReadFunctionInput is an implementation of ReadFunctionInput
//Read serveless function with http trigger
type HTTPReadFunctionInput struct {
	FunctionConfigurationInput *lambda.GetFunctionConfigurationInput
	HTTPReadEvent
}

//S3ReadFunctionInput is an implementation of ReadFunctionInput
//Read serveless function with s3 trigger
type S3ReadFunctionInput struct {
	FunctionConfigurationInput *lambda.GetFunctionConfigurationInput
	S3ReadEvent
}
