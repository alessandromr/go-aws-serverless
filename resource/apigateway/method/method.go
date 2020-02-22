package method

import (
	"github.com/alessandromr/go-aws-serverless/utils/auth"
	"github.com/aws/aws-sdk-go/service/apigateway"
)

//ApiGatewayMethod
type ApiGatewayMethod struct {
	HttpMethod        string
	ResourceId        string
	RestApiId         string
	AuthorizationType string
}

//Create the given resources
func (resource *ApiGatewayMethod) Create() error {
	auth.MakeClient(auth.Sess)
	svc := auth.Client.ApigatewayConn

	methodInput := &apigateway.PutMethodInput{
		HttpMethod:        &resource.HttpMethod,
		RestApiId:         &resource.RestApiId,
		ResourceId:        &resource.ResourceId,
		AuthorizationType: &resource.AuthorizationType,
	}

	_, err := svc.PutMethod(methodInput)
	return err
}

//Delete the given resources
func (resource *ApiGatewayMethod) Delete() error {
	auth.MakeClient(auth.Sess)
	svc := auth.Client.ApigatewayConn
	methodInput := &apigateway.DeleteMethodInput{
		HttpMethod: &resource.HttpMethod,
		ResourceId: &resource.ResourceId,
		RestApiId:  &resource.RestApiId,
	}
	_, err := svc.DeleteMethod(methodInput)
	return err
}
