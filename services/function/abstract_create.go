package function

import (
	"github.com/aws/aws-sdk-go/service/lambda"
)

//CreateFunctionInput is an interface to create a serverless function and the relative trigger
type CreateFunctionInput interface {
	CreateDependencies(*lambda.FunctionConfiguration) (map[string]interface{}, error)
	GetFunctionInput() *lambda.CreateFunctionInput
}

//HTTPCreateFunctionInput is an implementation of CreateFunctionInput
//Create serveless function with http trigger
type HTTPCreateFunctionInput struct {
	FunctionInput *lambda.CreateFunctionInput
	HTTPCreateEvent
}

//S3CreateFunctionInput is an implementation of CreateFunctionInput
//Create serveless function with s3 trigger
type S3CreateFunctionInput struct {
	FunctionInput *lambda.CreateFunctionInput
	S3CreateEvent
}

//SQSCreateFunctionInput is an implementation of CreateFunctionInput
//Create serveless function with SQS as trigger
type SQSCreateFunctionInput struct {
	FunctionInput *lambda.CreateFunctionInput
	SQSCreateEvent
}
