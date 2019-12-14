package method

import (
	"github.com/alessandromr/go-aws-serverless/utils/auth"
	"github.com/aws/aws-sdk-go/service/apigateway"
)

type ApiGatewayMethod struct {
	HttpMethod string
	ResourceId string
	RestApiId  string
}

//Delete the given resources
func (resource ApiGatewayMethod) Delete() error {
	auth.MakeClient(auth.Sess)
	svc := auth.Client.Apigatewayconn
	methodInput := &apigateway.DeleteMethodInput{
		HttpMethod: &resource.HttpMethod,
		ResourceId: &resource.ResourceId,
		RestApiId:  &resource.RestApiId,
	}
	_, err := svc.DeleteMethod(methodInput)
	return err
}
