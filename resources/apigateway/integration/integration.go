package integration

import (
	"github.com/alessandromr/goserverlessclient/utils/auth"
	"github.com/aws/aws-sdk-go/service/apigateway"
)

type ApiGatewayIntegration struct {
	HttpMethod string
	ResourceId string
	RestApiId  string
}

//Delete the given resources
func (resource ApiGatewayIntegration) Delete() error {
	svc := apigateway.New(auth.Sess)
	integrationInput := &apigateway.DeleteIntegrationInput{
		HttpMethod: &resource.HttpMethod,
		ResourceId: &resource.ResourceId,
		RestApiId:  &resource.RestApiId,
	}
	_, err := svc.DeleteIntegration(integrationInput)
	return err
}
