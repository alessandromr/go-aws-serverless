package integration

import (
	"github.com/alessandromr/go-aws-serverless/utils/auth"
	"github.com/aws/aws-sdk-go/aws"
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

//Create the given resources
func (resource *ApiGatewayIntegration) Create() error {
	auth.MakeClient(auth.Sess)
	svc := auth.Client.ApigatewayConn

	integrationInput := &apigateway.PutIntegrationInput{
		Type:                  aws.String(resource.Type),
		Credentials:           aws.String(resource.Credentials),
		HttpMethod:            aws.String(resource.HttpMethod),
		RestApiId:             aws.String(resource.RestApiId),
		ResourceId:            aws.String(resource.ResourceId),
		IntegrationHttpMethod: aws.String(resource.IntegrationHTTPMethod),
		Uri:                   aws.String(resource.Uri),
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
