package function

import (
	"log"

	"github.com/alessandromr/go-aws-serverless/utils/auth"
	"github.com/aws/aws-sdk-go/service/apigateway"
	"github.com/aws/aws-sdk-go/service/lambda"
)

//UpdateDependencies create all the dependencies for the HTTPEvent
func (input HTTPUpdateFunctionInput) UpdateDependencies(lambdaResult *lambda.FunctionConfiguration) {
	auth.MakeClient(auth.Sess)
	svc := auth.Client.Apigatewayconn
	var err error

	//Update integration between lambda and api gateway method
	//apigateway.UpdateIntegration
	integrationInput := &apigateway.UpdateIntegrationInput{
		HttpMethod: input.HTTPUpdateEvent.Method,
		RestApiId:  input.HTTPUpdateEvent.ApiId,
		ResourceId: input.HTTPUpdateEvent.ResourceId,
	}
	_, err = svc.UpdateIntegration(integrationInput)
	if err != nil {
		log.Println("Error") //ToDo
	}

}

//GetUpdateFunctionConfiguration return the UpdateFunctionConfigurationInput from the custom input
func (input HTTPUpdateFunctionInput) GetUpdateFunctionConfiguration() *lambda.UpdateFunctionConfigurationInput {
	return input.UpdateFunctionConfigurationInput
}
