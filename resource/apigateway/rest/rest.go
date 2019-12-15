package rest

import (
	"github.com/alessandromr/go-aws-serverless/utils/auth"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/apigateway"
)

//ApiGatewayRestApi rappresent an ApiGateway Resource on AWS
type ApiGatewayRestApi struct {
	RestApiId string
	ApiName   string
}

//Create the given resources
func (resource *ApiGatewayRestApi) Create() error {
	auth.MakeClient(auth.Sess)
	svc := auth.Client.ApigatewayConn

	apiInput := &apigateway.CreateRestApiInput{
		Name: aws.String(resource.ApiName),
	}
	_, err := svc.CreateRestApi(apiInput)
	return err
}

//Delete the given resources
func (resource *ApiGatewayRestApi) Delete() error {
	auth.MakeClient(auth.Sess)
	svc := auth.Client.ApigatewayConn
	apiInput := &apigateway.DeleteRestApiInput{
		RestApiId: &resource.RestApiId,
	}
	_, err := svc.DeleteRestApi(apiInput)
	return err
}
