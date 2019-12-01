package function

import (
	"github.com/alessandromr/goserverlessclient/utils"
	"github.com/alessandromr/goserverlessclient/utils/auth"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/apigateway"
	"github.com/aws/aws-sdk-go/service/lambda"
)

//CreateDependencies create all the dependencies for the HTTPEvent
func (input HTTPCreateFunctionInput) CreateDependencies(lambdaResult *lambda.FunctionConfiguration) (map[string]interface{}, error) {
	svc := apigateway.New(auth.Sess)
	var err error

	//Prepare a rollback object in case of failure
	rollback := HTTPDeleteFunctionInput{
		FunctionInput: &lambda.DeleteFunctionInput{
			FunctionName: lambdaResult.FunctionArn,
		},
		HTTPDeleteEvent: HTTPDeleteEvent{},
	}

	//apigateway.CreateRestApi
	if !input.HTTPCreateEvent.Existing {
		apiInput := &apigateway.CreateRestApiInput{
			Name: input.HTTPCreateEvent.ApiName,
		}
		response, err := svc.CreateRestApi(apiInput)
		utils.CheckErr(err)
		input.HTTPCreateEvent.ApiId = response.Id
		rollback.HTTPDeleteEvent.ApiId = response.Id
	}

	//Get Root Resource
	//apigateway.GetResources
	getResourceInput := &apigateway.GetResourcesInput{
		RestApiId: input.HTTPCreateEvent.ApiId,
	}
	getResourceOutput, err := svc.GetResources(getResourceInput)
	if err != nil {
		Rollback(rollback, err)
		return nil, err
	}

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
	rollback.HTTPDeleteEvent.ResourceId = createResourceOutput.Id
	if err != nil {
		Rollback(rollback, err)
		return nil, err
	}

	//apigateway.PutMethod
	methodInput := &apigateway.PutMethodInput{
		HttpMethod:        input.HTTPCreateEvent.Method,
		RestApiId:         input.HTTPCreateEvent.ApiId,
		ResourceId:        createResourceOutput.Id,
		AuthorizationType: aws.String("NONE"),
	}
	_, err = svc.PutMethod(methodInput)
	rollback.HTTPDeleteEvent.Method = input.HTTPCreateEvent.Method
	if err != nil {
		Rollback(rollback, err)
		return nil, err
	}

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
	if err != nil {
		Rollback(rollback, err)
		return nil, err
	}

	out := make(map[string]interface{})
	out["RestApiId"] = *input.HTTPCreateEvent.ApiId
	out["Method"] = *input.HTTPCreateEvent.Method
	out["ResourceId"] = *createResourceOutput.Id
	return out, nil
}

//GetFunctionInput return the CreateFunctionInput from the custom input
func (input HTTPCreateFunctionInput) GetFunctionInput() *lambda.CreateFunctionInput {
	return input.FunctionInput
}
