package rest

import (
	"github.com/alessandromr/go-serverless-client/utils/auth"
	"github.com/aws/aws-sdk-go/service/apigateway"
)

//ApiGatewayRestApi rappresent an ApiGateway Resource on AWS
type ApiGatewayRestApi struct {
	RestApiId string
	ApiName   string
}

//Delete the given resources
func (resource ApiGatewayRestApi) Delete() error {
	svc := apigateway.New(auth.Sess)
	apiInput := &apigateway.DeleteRestApiInput{
		RestApiId: &resource.RestApiId,
	}
	_, err := svc.DeleteRestApi(apiInput)
	return err
}
