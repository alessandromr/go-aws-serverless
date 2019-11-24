package function

import (
	"github.com/aws/aws-sdk-go/service/lambda"
)

//CreateFunctionInput is an interface rapresenting a serverless function and the relative trigger
type CreateFunctionInput interface {
	CreateDependencies()
	GetFunctionInput() *lambda.CreateFunctionInput
}

//HTTPCreateFunctionInput is an implementation of CreateFunctionInput
//Create serveless function with http trigger
type HTTPCreateFunctionInput struct {
	FunctionInput *lambda.CreateFunctionInput
	HTTPEvent
}

//CreateDependencies create all the dependencies for the given trigger
func (input HTTPCreateFunctionInput) CreateDependencies() {

}

//GetFunctionInput return the CreateFunctionInput from the custom input
func (input HTTPCreateFunctionInput) GetFunctionInput() *lambda.CreateFunctionInput {
	return input.FunctionInput
}

//S3CreateFunctionInput is an implementation of CreateFunctionInput
//Create serveless function with s3 trigger
type S3CreateFunctionInput struct {
	FunctionInput *lambda.CreateFunctionInput
	S3Event
}

//CreateDependencies create all the dependencies for the given trigger
func (input S3CreateFunctionInput) CreateDependencies() {

}

//GetFunctionInput return the CreateFunctionInput from the custom input
func (input S3CreateFunctionInput) GetFunctionInput() *lambda.CreateFunctionInput {
	return input.FunctionInput
}
