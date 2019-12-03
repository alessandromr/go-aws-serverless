package resource

import (
	"github.com/alessandromr/go-serverless-client/utils/auth"
	"github.com/aws/aws-sdk-go/service/apigateway"
)

//ApiGatewayResource
type ApiGatewayResource struct {
	Path       string
	ResourceId string
	ParentId   string
	RestApiId  string
}

//Delete the given resources
func (resource ApiGatewayResource) Delete() error {
	svc := apigateway.New(auth.Sess)
	resourceInput := &apigateway.DeleteResourceInput{
		ResourceId: &resource.ResourceId,
		RestApiId:  &resource.RestApiId,
	}
	_, err := svc.DeleteResource(resourceInput)
	return err
}
