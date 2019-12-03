package method

import (
	"github.com/alessandromr/go-serverless-client/utils/auth"
	"github.com/aws/aws-sdk-go/service/apigateway"
)

type ApiGatewayMethod struct {
	HttpMethod string
	ResourceId string
	RestApiId  string
}

//Delete the given resources
func (resource ApiGatewayMethod) Delete() error {
	svc := apigateway.New(auth.Sess)
	methodInput := &apigateway.DeleteMethodInput{
		HttpMethod: &resource.HttpMethod,
		ResourceId: &resource.ResourceId,
		RestApiId:  &resource.RestApiId,
	}
	_, err := svc.DeleteMethod(methodInput)
	return err
}
