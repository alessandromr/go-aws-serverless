package function

import (
	"time"

	"github.com/alessandromr/go-aws-serverless/manager/create"
	"github.com/alessandromr/go-aws-serverless/resource/apigateway/deployment"
	"github.com/alessandromr/go-aws-serverless/resource/apigateway/integration"
	"github.com/alessandromr/go-aws-serverless/resource/apigateway/method"
	"github.com/alessandromr/go-aws-serverless/resource/apigateway/resource"
	"github.com/alessandromr/go-aws-serverless/resource/apigateway/rest"
	"github.com/alessandromr/go-aws-serverless/resource/apigateway/stage"
	"github.com/alessandromr/go-aws-serverless/resource/iam/role"
	"github.com/alessandromr/go-aws-serverless/resource/lambda/permission"
	"github.com/alessandromr/go-aws-serverless/utils"
	"github.com/alessandromr/go-aws-serverless/utils/auth"
	"github.com/aws/aws-sdk-go/service/apigateway"
	"github.com/aws/aws-sdk-go/service/lambda"
)

var executionRoleAssumeRoleString string = `{"Version":"2012-10-17","Statement":[{"Sid":"","Effect":"Allow","Principal":{"Service":"apigateway.amazonaws.com"},"Action":"sts:AssumeRole"}]}`
var executionRolePolicyString string = `{"Version":"2012-10-17","Statement":[{"Effect":"Allow","Action":"lambda:InvokeFunction","Resource":"*"}]}`

//CreateDependencies create all the dependencies for the HTTPEvent
func (input HTTPCreateFunctionInput) CreateDependencies(lambdaResult *lambda.FunctionConfiguration) (map[string]interface{}, error) {
	auth.MakeClient(auth.Sess)
	svc := auth.Client.ApigatewayConn
	var err error

	accountID := auth.GetAccountID()

	var executionRole role.IamRole
	if input.HTTPCreateEvent.ExecutionRoleArn == nil || len(*input.HTTPCreateEvent.ExecutionRoleArn) < 20 {
		//iam.CreateRole
		executionRole = role.IamRole{
			AssumeRolePolicyDocument: executionRoleAssumeRoleString,
			Description:              "Role to allow API Gateway to invoke Lambda functions on behalf of the API caller.",
			RoleName:                 "ApiExecutionRole-" + *lambdaResult.FunctionName,
		}
		create.ResourcesList = append(
			create.ResourcesList,
			&executionRole,
		)
		input.HTTPCreateEvent.ExecutionRoleArn = &executionRole.RoleArn
	}

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

	//Create Rest Api if necessary
	err = create.ExecutePartial()
	if err != nil {
		return nil, err
	}

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

	//Create Rest Resource
	err = create.ExecutePartial()
	if err != nil {
		return nil, err
	}

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
		// Credentials:           *input.HTTPCreateEvent.ExecutionRoleArn,
		Type: "AWS_PROXY",
	}
	create.ResourcesList = append(
		create.ResourcesList,
		&apiIntegration,
	)

	//Create Lambda Permission
	permission := permission.LambdaPermission{
		StatementId:  "HTTPEvent_" + *input.HTTPCreateEvent.ApiId + "_" + *lambdaResult.FunctionName,
		FunctionName: *lambdaResult.FunctionArn,
		SourceArn:    "arn:aws:execute-api:" + auth.Region + ":" + accountID + ":" + *input.HTTPCreateEvent.ApiId + "/*/*/" + *input.HTTPCreateEvent.Path,
		Principal:    "apigateway.amazonaws.com",
		Action:       "lambda:InvokeFunction",
	}
	create.ResourcesList = append(
		create.ResourcesList,
		&permission,
	)

	//API Deployment
	apiDeployment := deployment.ApiGatewayDeployment{
		RestApiId:        *input.HTTPCreateEvent.ApiId,
		StageName:        "default",
		StageDescription: "Default Stage",
	}
	create.ResourcesList = append(
		create.ResourcesList,
		&apiDeployment,
	)

	//API Stage
	apiStage := stage.ApiGatewayStage{
		RestApiId:    *input.HTTPCreateEvent.ApiId,
		StageName:    "default",
		DeploymentID: apiDeployment.DeploymentId,
	}
	create.ResourcesList = append(
		create.ResourcesList,
		&apiStage,
	)

	//Create Resources
	err = create.ExecuteCreate()
	if err != nil {
		return nil, err
	}

	out := make(map[string]interface{})
	out["RestApiId"] = *input.HTTPCreateEvent.ApiId
	out["Method"] = *input.HTTPCreateEvent.Method
	out["ResourceId"] = apiResource.ResourceId
	out["ExecutionRoleName"] = executionRole.RoleName
	out["ExecutionRoleArn"] = executionRole.RoleArn
	return out, nil
}

//GetFunctionInput return the CreateFunctionInput from the custom input
func (input HTTPCreateFunctionInput) GetFunctionInput() *lambda.CreateFunctionInput {
	return input.FunctionInput
}
