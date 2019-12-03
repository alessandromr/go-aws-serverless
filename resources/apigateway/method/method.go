package method

type ApiGatewayMethod struct {
	HttpMethod string
	ResourceId string
	RestApiId  string
}


//Delete the given resources
func (resource ApiGatewayMethod) Delete() error {
	methodInput := &apigateway.DeleteMethodInput{
		HttpMethod: &resource.HttpMethod,
		ResourceId: &resource.ResourceId,
		RestApiId:  &resource.RestApiId,
	}
	_, err = svc.DeleteMethod(methodInput)
	return err
}

