package function

import (
	"time"

	"github.com/alessandromr/go-aws-serverless/manager/create"
	"github.com/alessandromr/go-aws-serverless/resource/apigateway/integration"
	"github.com/alessandromr/go-aws-serverless/resource/apigateway/method"
	"github.com/alessandromr/go-aws-serverless/resource/apigateway/resource"
	"github.com/alessandromr/go-aws-serverless/resource/apigateway/rest"
	"github.com/alessandromr/go-aws-serverless/utils"
	"github.com/alessandromr/go-aws-serverless/utils/auth"
	"github.com/aws/aws-sdk-go/service/apigateway"
	"github.com/aws/aws-sdk-go/service/lambda"
)

//CreateDependencies create all the dependencies for the HTTPEvent
func (input HTTPCreateFunctionInput) CreateDependencies(lambdaResult *lambda.FunctionConfiguration) (map[string]interface{}, error) {
	auth.MakeClient(auth.Sess)
	svc := auth.Client.ApigatewayConn
	var err error

	//apigateway.CreateRestApi
	if !input.HTTPCreateEvent.Existing {
		restAPI := rest.ApiGatewayRestApi{
			ApiName: *input.HTTPCreateEvent.ApiName,
		}
		create.ResourcesList = append(
			create.ResourcesList,
			&restAPI,
		)
		input.HTTPCreateEvent.ApiId = &restAPI.RestApiId
	}
	create.ExecutePartial()

	//Get Root Resource
	//apigateway.GetResources
	getResourceInput := &apigateway.GetResourcesInput{
		RestApiId: input.HTTPCreateEvent.ApiId,
	}
	time.Sleep(utils.ShortSleep * time.Millisecond)
	getResourceOutput, err := svc.GetResources(getResourceInput)
	time.Sleep(utils.ShortSleep * time.Millisecond)
	if err != nil {
		return nil, err
	}
	var rootParent string
	for _, v := range getResourceOutput.Items {
		if *v.Path == "/" {
			rootParent = *v.Id
		}
	}

	//apigateway.CreateResource
	apiResource := resource.ApiGatewayResource{
		RestApiId: *input.HTTPCreateEvent.ApiId,
		Path:      *input.HTTPCreateEvent.Path,
		ParentId:  rootParent,
	}
	create.ResourcesList = append(
		create.ResourcesList,
		&apiResource,
	)

	//apigateway.PutMethod
	apiMethod := method.ApiGatewayMethod{
		HttpMethod: *input.HTTPCreateEvent.Method,
		ResourceId: apiResource.ResourceId,
		RestApiId:  *input.HTTPCreateEvent.ApiId,
	}
	create.ResourcesList = append(
		create.ResourcesList,
		&apiMethod,
	)

	//Put integration between lambda and api gateway method
	//apigateway.PutIntegration
	apiIntegration := integration.ApiGatewayIntegration{
		HttpMethod:            *input.HTTPCreateEvent.Method,
		IntegrationHTTPMethod: "POST",
		ResourceId:            apiResource.ResourceId,
		RestApiId:             *input.HTTPCreateEvent.ApiId,
		Uri:                   "arn:aws:apigateway:" + auth.Region + ":lambda:path/2015-03-31/functions/" + *lambdaResult.FunctionArn + "/invocations",
		Credentials:           *input.HTTPCreateEvent.ExecutionRole,
		Type:                  "AWS_PROXY",
	}
	create.ResourcesList = append(
		create.ResourcesList,
		&apiIntegration,
	)

	out := make(map[string]interface{})
	out["RestApiId"] = *input.HTTPCreateEvent.ApiId
	out["Method"] = *input.HTTPCreateEvent.Method
	out["ResourceId"] = apiResource.ResourceId
	return out, nil
}

//GetFunctionInput return the CreateFunctionInput from the custom input
func (input HTTPCreateFunctionInput) GetFunctionInput() *lambda.CreateFunctionInput {
	return input.FunctionInput
}
