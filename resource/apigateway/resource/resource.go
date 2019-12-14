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

//Delete the given resources
func (resource ApiGatewayResource) Delete() error {
	auth.MakeClient(auth.Sess)
	svc := auth.Client.Apigatewayconn
	resourceInput := &apigateway.DeleteResourceInput{
		ResourceId: &resource.ResourceId,
		RestApiId:  &resource.RestApiId,
	}
	_, err := svc.DeleteResource(resourceInput)
	return err
}
