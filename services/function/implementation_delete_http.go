package function

import (
	"github.com/alessandromr/go-aws-serverless/utils"
	"github.com/alessandromr/go-aws-serverless/utils/auth"
	"github.com/aws/aws-sdk-go/service/apigateway"
	"github.com/aws/aws-sdk-go/service/lambda"
	"time"
)

//DeleteDependencies implements the dependencies deletion for HTTP Event
func (input HTTPDeleteFunctionInput) DeleteDependencies(lambdaResult *lambda.DeleteFunctionInput) {
	svc := apigateway.New(auth.Sess)
	var err error

	//delete existing integration
	integrationInput := &apigateway.DeleteIntegrationInput{
		HttpMethod: input.HTTPDeleteEvent.Method,
		ResourceId: input.HTTPDeleteEvent.ResourceId,
		RestApiId:  input.HTTPDeleteEvent.ApiId,
	}
	_, err = svc.DeleteIntegration(integrationInput)
	utils.CheckAWSErrExpect404(err, "API Gateway Integration")
	time.Sleep(utils.ShortSleep * time.Millisecond)

	//delete method
	methodInput := &apigateway.DeleteMethodInput{
		HttpMethod: input.HTTPDeleteEvent.Method,
		ResourceId: input.HTTPDeleteEvent.ResourceId,
		RestApiId:  input.HTTPDeleteEvent.ApiId,
	}
	_, err = svc.DeleteMethod(methodInput)
	utils.CheckAWSErrExpect404(err, "API Gateway Method")
	time.Sleep(utils.ShortSleep * time.Millisecond)

	//check if resource is empty
	getResourceInput := &apigateway.GetResourceInput{
		ResourceId: input.HTTPDeleteEvent.ResourceId,
		RestApiId:  input.HTTPDeleteEvent.ApiId,
	}
	resourceResponse, err := svc.GetResource(getResourceInput)
	utils.CheckAWSErrExpect404(err, "API Gateway Get Resources")
	time.Sleep(utils.ShortSleep * time.Millisecond)

	if len(resourceResponse.ResourceMethods) < 1 {
		//delete resource
		resourceInput := &apigateway.DeleteResourceInput{
			ResourceId: input.HTTPDeleteEvent.ResourceId,
			RestApiId:  input.HTTPDeleteEvent.ApiId,
		}
		_, err = svc.DeleteResource(resourceInput)
		utils.CheckAWSErrExpect404(err, "API Gateway Resource")
		time.Sleep(utils.ShortSleep * time.Millisecond)
	}

	//check if api is empty
	getResourcesInput := &apigateway.GetResourcesInput{
		RestApiId: input.HTTPDeleteEvent.ApiId,
	}
	getResourcesOutput, err := svc.GetResources(getResourcesInput)
	utils.CheckAWSErrExpect404(err, "API Gateway Get Resources")
	time.Sleep(utils.ShortSleep * time.Millisecond)

	if len(getResourcesOutput.Items) <= 1 {
		//delete api
		apiInput := &apigateway.DeleteRestApiInput{
			RestApiId: input.HTTPDeleteEvent.ApiId,
		}
		_, err = svc.DeleteRestApi(apiInput)
		utils.CheckAWSErrExpect404(err, "API Gateway Rest API")
	}

}

//GetFunctionInput return the DeleteFunctionInput from the custom input
func (input HTTPDeleteFunctionInput) GetFunctionInput() *lambda.DeleteFunctionInput {
	return input.FunctionInput
}
