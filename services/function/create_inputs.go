package function

import (
	"github.com/alessandromr/goserverlessclient/utils"
	"github.com/alessandromr/goserverlessclient/utils/auth"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/apigateway"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/s3"
)

//CreateFunctionInput is an interface to create a serverless function and the relative trigger
type CreateFunctionInput interface {
	CreateDependencies(*lambda.FunctionConfiguration) map[string]interface{}
	GetFunctionInput() *lambda.CreateFunctionInput
}

//HTTPCreateFunctionInput is an implementation of CreateFunctionInput
//Create serveless function with http trigger
type HTTPCreateFunctionInput struct {
	FunctionInput *lambda.CreateFunctionInput
	HTTPCreateEvent
}

//CreateDependencies create all the dependencies for the given trigger
func (input HTTPCreateFunctionInput) CreateDependencies(lambdaResult *lambda.FunctionConfiguration) map[string]interface{} {
	svc := apigateway.New(auth.Sess)
	var err error

	//apigateway.CreateRestApi
	if !input.HTTPCreateEvent.Existing {
		apiInput := &apigateway.CreateRestApiInput{
			Name: input.HTTPCreateEvent.ApiName,
		}
		response, err := svc.CreateRestApi(apiInput)
		utils.CheckErr(err)
		input.HTTPCreateEvent.ApiId = response.Id
	}

	//Get Root Resource
	//apigateway.GetResources
	getResourceInput := &apigateway.GetResourcesInput{
		RestApiId: input.HTTPCreateEvent.ApiId,
	}
	getResourceOutput, err := svc.GetResources(getResourceInput)
	utils.CheckErr(err)

	var rootParent string
	for _, v := range getResourceOutput.Items {
		if *v.Path == "/" {
			rootParent = *v.Id
		}
	}

	//apigateway.CreateResource
	resourceInput := &apigateway.CreateResourceInput{
		PathPart:  input.HTTPCreateEvent.Path,
		RestApiId: input.HTTPCreateEvent.ApiId,
		ParentId:  aws.String(rootParent),
	}
	createResourceOutput, err := svc.CreateResource(resourceInput)
	utils.CheckErr(err)

	//apigateway.PutMethod
	methodInput := &apigateway.PutMethodInput{
		HttpMethod:        input.HTTPCreateEvent.Method,
		RestApiId:         input.HTTPCreateEvent.ApiId,
		ResourceId:        createResourceOutput.Id,
		AuthorizationType: aws.String("NONE"),
	}
	_, err = svc.PutMethod(methodInput)
	utils.CheckErr(err)

	//Put integration between lambda and api gateway method
	//apigateway.PutIntegration

	integrationInput := &apigateway.PutIntegrationInput{
		Type:                  aws.String("AWS_PROXY"),
		Credentials:           input.HTTPCreateEvent.ExecutionRole,
		HttpMethod:            input.HTTPCreateEvent.Method,
		RestApiId:             input.HTTPCreateEvent.ApiId,
		ResourceId:            createResourceOutput.Id,
		IntegrationHttpMethod: aws.String("POST"),
		Uri:                   aws.String("arn:aws:apigateway:" + auth.Region + ":lambda:path/2015-03-31/functions/" + *lambdaResult.FunctionArn + "/invocations"),
	}
	_, err = svc.PutIntegration(integrationInput)
	utils.CheckErr(err)

	out := make(map[string]interface{})
	out["RestApiId"] = *input.HTTPCreateEvent.ApiId
	out["Method"] = *input.HTTPCreateEvent.Method
	out["ResourceId"] = *createResourceOutput.Id
	return out
}

//GetFunctionInput return the CreateFunctionInput from the custom input
func (input HTTPCreateFunctionInput) GetFunctionInput() *lambda.CreateFunctionInput {
	return input.FunctionInput
}

//S3CreateFunctionInput is an implementation of CreateFunctionInput
//Create serveless function with s3 trigger
type S3CreateFunctionInput struct {
	FunctionInput *lambda.CreateFunctionInput
	S3CreateEvent
}

//CreateDependencies create all the dependencies for the given trigger
func (input S3CreateFunctionInput) CreateDependencies(lambdaResult *lambda.FunctionConfiguration) {
	svc := s3.New(auth.Sess)
	var err error

	//s3.CreateBucket
	if !input.S3CreateEvent.Existing {
		createBucket := &s3.CreateBucketInput{
			Bucket: input.S3CreateEvent.Bucket,
		}
		_, err = svc.CreateBucket(createBucket)
		utils.CheckErr(err)
	}

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
}

//GetFunctionInput return the CreateFunctionInput from the custom input
func (input S3CreateFunctionInput) GetFunctionInput() *lambda.CreateFunctionInput {
	return input.FunctionInput
}