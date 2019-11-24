package function

import (
	"github.com/aws/aws-sdk-go/service/lambda"
)

//CreateFunctionInput is an interface rapresenting a serverless function and the relative trigger
type CreateFunctionInput interface {
	CreateDependencies(*lambda.FunctionConfiguration)
	GetFunctionInput() *lambda.CreateFunctionInput
}

//HTTPCreateFunctionInput is an implementation of CreateFunctionInput
//Create serveless function with http trigger
type HTTPCreateFunctionInput struct {
	FunctionInput *lambda.CreateFunctionInput
	HTTPEvent
}

//CreateDependencies create all the dependencies for the given trigger
func (input HTTPCreateFunctionInput) CreateDependencies(lambdaResult *lambda.FunctionConfiguration) {
	svc := apigateway.New(session.New())

	//apigateway.CreateRestApi
	if !input.HTTPEvent.Existing{
		apiInput := &apigateway.CreateRestApiInput{
			Name: aws.String(input.HTTPEvent.ApiName),
		}
		response, err = svc.CreateRestApi(apiInput)
		utils.CheckErr(err)
		input.HTTPEvent.RestApiId = response.Id
	}

	//apigateway.CreateResource
	resourceInput := &apigateway.CreateResourceInput{
		PathPart: aws.String(input.HTTPEvent.Path),
		RestApiId: aws.String(input.HTTPEvent.ApiId),
	}
	response, err = svc.CreateResource(resourceInput)
	utils.CheckErr(err)

	//apigateway.PutMethod
	methodInput := &apigateway.PutMethodInput{
		HttpMethod: input.HTTPEvent.Method,
		RestApiId: aws.String(input.HTTPEvent.ApiId),
		ResourceId: response.Id,
	}
	_, err = svc.PutMethod(methodInput)
	utils.CheckErr(err)
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
func (input S3CreateFunctionInput) CreateDependencies(lambdaResult *lambda.FunctionConfiguration) {
	svc := s3.New(session.New())

	//s3.CreateBucket
	if !input.S3Event.Existing{
		createBucket := &s3.CreateBucketInput{
			Bucket: aws.String(input.S3Event.Bucket),
		}
		_, err = svc.CreateBucket(createBucket)
		utils.CheckErr(err)
	}

	//s3.PutBucketNotificationConfiguration
	putNotConfig := &s3.PutBucketNotificationConfigurationInput{
		Bucket: aws.String(input.S3Event.Bucket),
		NotificationConfiguration: &s3.NotificationConfiguration{
			LambdaFunctionConfiguration: &s3.LambdaFunctionConfiguration{
				LambdaFunctionArn: aws.String(lambdaResult.FunctionArn),
				Events: aws.StringSlice(input.S3Event.Types),
			},
		},
	}
	_, err = svc.PutBucketNotificationConfiguration(putNotConfig)
	utils.CheckErr(err)
}

//GetFunctionInput return the CreateFunctionInput from the custom input
func (input S3CreateFunctionInput) GetFunctionInput() *lambda.CreateFunctionInput {
	return input.FunctionInput
}
