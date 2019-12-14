package function

import (
	"time"

	"github.com/alessandromr/go-aws-serverless/utils"
	"github.com/alessandromr/go-aws-serverless/utils/auth"
	"github.com/aws/aws-sdk-go/service/apigateway"
	"github.com/aws/aws-sdk-go/service/lambda"
)

//ReadDependencies implements the dependencies deletion for HTTP Event
func (input HTTPReadFunctionInput) ReadDependencies(lambdaResult *lambda.FunctionConfiguration) map[string]interface{} {
	auth.MakeClient(auth.Sess)
	svc := auth.Client.Apigatewayconn
	var err error

	//get integration
	//get method
	//get resource
	//get rest api

	getRestInput := &apigateway.GetRestApiInput{
		RestApiId: input.HTTPReadEvent.ApiId,
	}
	restResponse, err := svc.GetRestApi(getRestInput)
	utils.CheckErr(err)

	time.Sleep(utils.ShortSleep * time.Millisecond)

	getResourceInput := &apigateway.GetResourceInput{
		RestApiId:  input.HTTPReadEvent.ApiId,
		ResourceId: input.HTTPReadEvent.ResourceId,
	}
	resourceResponse, err := svc.GetResource(getResourceInput)
	utils.CheckErr(err)

	time.Sleep(utils.ShortSleep * time.Millisecond)

	getIntegrationInput := &apigateway.GetIntegrationInput{
		RestApiId:  input.HTTPReadEvent.ApiId,
		ResourceId: input.HTTPReadEvent.ResourceId,
		HttpMethod: input.HTTPReadEvent.Method,
	}
	integrationResponse, err := svc.GetIntegration(getIntegrationInput)
	utils.CheckErr(err)

	out := make(map[string]interface{})
	out["RestApi"] = restResponse
	out["ApiResource"] = resourceResponse
	out["ApiIntegration"] = integrationResponse
	return out
}

//GetFunctionInput return the ReadFunctionInput from the custom input
func (input HTTPReadFunctionInput) GetFunctionConfiguration() *lambda.GetFunctionConfigurationInput {
	return input.FunctionConfigurationInput
}
