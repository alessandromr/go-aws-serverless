package function

import (
	"github.com/aws/aws-sdk-go/service/lambda"
)

//UpdateFunctionInput is an interface to Update a serverless function and the relative trigger
type UpdateFunctionInput interface {
	UpdateDependencies(*lambda.FunctionConfiguration) (map[string]interface{}, error)
	GetUpdateFunctionConfiguration() *lambda.UpdateFunctionConfigurationInput
}

//HTTPUpdateFunctionInput is an implementation of UpdateFunctionInput
//Update serveless function with http trigger
type HTTPUpdateFunctionInput struct {
	UpdateFunctionConfigurationInput *lambda.UpdateFunctionConfigurationInput
	HTTPUpdateEvent
}

//S3UpdateFunctionInput is an implementation of UpdateFunctionInput
//Update serveless function with s3 trigger
type S3UpdateFunctionInput struct {
	UpdateFunctionConfigurationInput *lambda.UpdateFunctionConfigurationInput
	S3UpdateEvent
}
