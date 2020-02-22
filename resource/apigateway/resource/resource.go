package resource

import (
	"github.com/alessandromr/go-aws-serverless/utils/auth"
	"github.com/aws/aws-sdk-go/service/apigateway"
)

//ApiGatewayResource
type ApiGatewayResource struct {
	Path       string
	ResourceId string
	ParentId   string
	RestApiId  string
}

//Create the given resources
func (resource *ApiGatewayResource) Create() error {
	auth.MakeClient(auth.Sess)
	svc := auth.Client.ApigatewayConn

	resourceInput := &apigateway.CreateResourceInput{
		PathPart:  &resource.Path,
		RestApiId: &resource.RestApiId,
		ParentId:  &resource.ParentId,
	}

	createResourceOutput, err := svc.CreateResource(resourceInput)
	resource.ResourceId = *createResourceOutput.Id
	return err
}

//Delete the given resources
func (resource *ApiGatewayResource) Delete() error {
	auth.MakeClient(auth.Sess)
	svc := auth.Client.ApigatewayConn
	resourceInput := &apigateway.DeleteResourceInput{
		ResourceId: &resource.ResourceId,
		RestApiId:  &resource.RestApiId,
	}
	_, err := svc.DeleteResource(resourceInput)
	return err
}
