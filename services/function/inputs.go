package function

import (
	"github.com/alessandromr/goserverlessclient/utils"
	"github.com/alessandromr/goserverlessclient/utils/auth"
	"github.com/aws/aws-sdk-go/service/apigateway"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/s3"
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
	svc := apigateway.New(auth.Sess)
	var err error

	//apigateway.CreateRestApi
	if !input.HTTPEvent.Existing {
		apiInput := &apigateway.CreateRestApiInput{
			Name: input.HTTPEvent.ApiName,
		}
		response, err := svc.CreateRestApi(apiInput)
		utils.CheckErr(err)
		input.HTTPEvent.ApiId = response.Id
	}

	//apigateway.CreateResource
	resourceInput := &apigateway.CreateResourceInput{
		PathPart:  input.HTTPEvent.Path,
		RestApiId: input.HTTPEvent.ApiId,
	}
	response, err := svc.CreateResource(resourceInput)
	utils.CheckErr(err)

	//apigateway.PutMethod
	methodInput := &apigateway.PutMethodInput{
		HttpMethod: input.HTTPEvent.Method,
		RestApiId:  input.HTTPEvent.ApiId,
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
	svc := s3.New(auth.Sess)
	var err error

	//s3.CreateBucket
	if !input.S3Event.Existing {
		createBucket := &s3.CreateBucketInput{
			Bucket: input.S3Event.Bucket,
		}
		_, err = svc.CreateBucket(createBucket)
		utils.CheckErr(err)
	}

	//s3.PutBucketNotificationConfiguration
	putNotConfig := &s3.PutBucketNotificationConfigurationInput{
		Bucket: input.S3Event.Bucket,
		NotificationConfiguration: &s3.NotificationConfiguration{
			LambdaFunctionConfigurations: []*s3.LambdaFunctionConfiguration{
				&s3.LambdaFunctionConfiguration{
					LambdaFunctionArn: lambdaResult.FunctionArn,
					Events:            input.S3Event.Types,
				},
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
