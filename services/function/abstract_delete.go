package function

import (
	"github.com/aws/aws-sdk-go/service/lambda"
)

//DeleteFunctionInput is an interface to delete a serverless function and the relative triggger
type DeleteFunctionInput interface {
	DeleteDependencies(*lambda.DeleteFunctionInput)
	GetFunctionInput() *lambda.DeleteFunctionInput
}

//HTTPDeleteFunctionInput is an implementation of CreateFunctionInput
//Create serveless function with http trigger
type HTTPDeleteFunctionInput struct {
	FunctionInput *lambda.DeleteFunctionInput
	HTTPDeleteEvent
}

//S3DeleteFunctionInput is an implementation of CreateFunctionInput
//Create serveless function with s3 trigger
type S3DeleteFunctionInput struct {
	FunctionInput *lambda.DeleteFunctionInput
	S3DeleteEvent
}
