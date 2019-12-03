package rest

//ApiGatewayRestApi rappresent an ApiGateway Resource on AWS
type ApiGatewayRestApi struct {
	RestApiId   string
	ApiName string
}


//Delete the given resources
func (resource ApiGatewayRestApi) Delete() error {
	apiInput := &apigateway.DeleteRestApiInput{
		RestApiId: resource.RestApiId,
	}
	_, err = svc.DeleteRestApi(apiInput)
	return err
}

