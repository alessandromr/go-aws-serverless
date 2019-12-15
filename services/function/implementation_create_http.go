package function

import (
	"log"
	"time"

	"github.com/alessandromr/go-aws-serverless/manager/rollback"
	"github.com/alessandromr/go-aws-serverless/resource/apigateway/integration"
	"github.com/alessandromr/go-aws-serverless/resource/apigateway/method"
	"github.com/alessandromr/go-aws-serverless/resource/apigateway/resource"
	"github.com/alessandromr/go-aws-serverless/resource/apigateway/rest"
	"github.com/alessandromr/go-aws-serverless/utils"
	"github.com/alessandromr/go-aws-serverless/utils/auth"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/apigateway"
	"github.com/aws/aws-sdk-go/service/lambda"
)

//CreateDependencies create all the dependencies for the HTTPEvent
func (input HTTPCreateFunctionInput) CreateDependencies(lambdaResult *lambda.FunctionConfiguration) (map[string]interface{}, error) {
	auth.MakeClient(auth.Sess)
	svc := auth.Client.ApigatewayConn
	var err error

	time.Sleep(utils.LongSleep * time.Millisecond)

	//apigateway.CreateRestApi
	if !input.HTTPCreateEvent.Existing {
		apiInput := &apigateway.CreateRestApiInput{
			Name: input.HTTPCreateEvent.ApiName,
		}
		response, err := svc.CreateRestApi(apiInput)
		if err != nil {
			return nil, err
		}
		rollback.ResourcesList = append(
			rollback.ResourcesList,
			&rest.ApiGatewayRestApi{
				RestApiId: *response.Id,
				ApiName:   *input.HTTPCreateEvent.ApiName,
			},
		)
		input.HTTPCreateEvent.ApiId = response.Id
	}

	time.Sleep(utils.ShortSleep * time.Millisecond)

	//Get Root Resource
	//apigateway.GetResources
	getResourceInput := &apigateway.GetResourcesInput{
		RestApiId: input.HTTPCreateEvent.ApiId,
	}
	getResourceOutput, err := svc.GetResources(getResourceInput)
	if err != nil {
		return nil, err
	}

	var rootParent string
	for _, v := range getResourceOutput.Items {
		if *v.Path == "/" {
			rootParent = *v.Id
		}
	}

	time.Sleep(utils.ShortSleep * time.Millisecond)

	//apigateway.CreateResource
	resourceInput := &apigateway.CreateResourceInput{
		PathPart:  input.HTTPCreateEvent.Path,
		RestApiId: input.HTTPCreateEvent.ApiId,
		ParentId:  aws.String(rootParent),
	}
	createResourceOutput, err := svc.CreateResource(resourceInput)
	if err != nil {
		log.Println(err)
		rollback.ExecuteRollback()
		return nil, err
	}
	rollback.ResourcesList = append(
		rollback.ResourcesList,
		&resource.ApiGatewayResource{
			ResourceId: *createResourceOutput.Id,
			RestApiId:  *input.HTTPCreateEvent.ApiId,
			Path:       *input.HTTPCreateEvent.Path,
			ParentId:   rootParent,
		},
	)

	time.Sleep(utils.ShortSleep * time.Millisecond)

	//apigateway.PutMethod
	methodInput := &apigateway.PutMethodInput{
		HttpMethod:        input.HTTPCreateEvent.Method,
		RestApiId:         input.HTTPCreateEvent.ApiId,
		ResourceId:        createResourceOutput.Id,
		AuthorizationType: aws.String("NONE"),
	}
	_, err = svc.PutMethod(methodInput)
	if err != nil {
		log.Println(err)
		rollback.ExecuteRollback()
		return nil, err
	}
	rollback.ResourcesList = append(
		rollback.ResourcesList,
		&method.ApiGatewayMethod{
			HttpMethod: *input.HTTPCreateEvent.Method,
			ResourceId: *createResourceOutput.Id,
			RestApiId:  *input.HTTPCreateEvent.ApiId,
		},
	)

	time.Sleep(utils.ShortSleep * time.Millisecond)

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
		log.Println(err)
		rollback.ExecuteRollback()
		return nil, err
	}
	rollback.ResourcesList = append(
		rollback.ResourcesList,
		&integration.ApiGatewayIntegration{
			HttpMethod:            *input.HTTPCreateEvent.Method,
			IntegrationHTTPMethod: "POST",
			ResourceId:            *createResourceOutput.Id,
			RestApiId:             *input.HTTPCreateEvent.ApiId,
			Uri:                   "arn:aws:apigateway:" + auth.Region + ":lambda:path/2015-03-31/functions/" + *lambdaResult.FunctionArn + "/invocations",
			Credentials:           *input.HTTPCreateEvent.ExecutionRole,
			Type:                  "AWS_PROXY",
		},
	)

	time.Sleep(utils.LongSleep * time.Millisecond)

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
