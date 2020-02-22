package integration

import (
	"github.com/alessandromr/go-aws-serverless/utils/auth"
	"github.com/aws/aws-sdk-go/service/apigateway"
)

//ApiGatewayIntegration
type ApiGatewayIntegration struct {
	HttpMethod            string
	ResourceId            string
	RestApiId             string
	Uri                   string
	Credentials           string
	Type                  string
	IntegrationHTTPMethod string
}

//Create the given resources
func (resource *ApiGatewayIntegration) Create() error {
	auth.MakeClient(auth.Sess)
	svc := auth.Client.ApigatewayConn

	integrationInput := &apigateway.PutIntegrationInput{
		Type:                  &resource.Type,
		Credentials:           &resource.Credentials,
		HttpMethod:            &resource.HttpMethod,
		RestApiId:             &resource.RestApiId,
		ResourceId:            &resource.ResourceId,
		IntegrationHttpMethod: &resource.IntegrationHTTPMethod,
		Uri:                   &resource.Uri,
	}
	_, err := svc.PutIntegration(integrationInput)
	return err
}

//Delete the given resources
func (resource *ApiGatewayIntegration) Delete() error {
	auth.MakeClient(auth.Sess)
	svc := auth.Client.ApigatewayConn
	integrationInput := &apigateway.DeleteIntegrationInput{
		HttpMethod: &resource.HttpMethod,
		ResourceId: &resource.ResourceId,
		RestApiId:  &resource.RestApiId,
	}
	_, err := svc.DeleteIntegration(integrationInput)
	return err
}
