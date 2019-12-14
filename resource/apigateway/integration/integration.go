package integration

import (
	"github.com/alessandromr/go-aws-serverless/utils/auth"
	"github.com/aws/aws-sdk-go/service/apigateway"
)

type ApiGatewayIntegration struct {
	HttpMethod            string
	ResourceId            string
	RestApiId             string
	Uri                   string
	Credentials           string
	Type                  string
	IntegrationHTTPMethod string
}

//Delete the given resources
func (resource ApiGatewayIntegration) Delete() error {
	auth.MakeClient(auth.Sess)
	svc := auth.Client.Apigatewayconn
	integrationInput := &apigateway.DeleteIntegrationInput{
		HttpMethod: &resource.HttpMethod,
		ResourceId: &resource.ResourceId,
		RestApiId:  &resource.RestApiId,
	}
	_, err := svc.DeleteIntegration(integrationInput)
	return err
}
