package resource

//ApiGatewayResource
type ApiGatewayResource struct {
	Path       string
	ResourceId string
	ParentId   string
	RestApiId   string
}



//Delete the given resources
func (resource ApiGatewayResource) Delete() error {
	resourceInput := &apigateway.DeleteResourceInput{
		ResourceId: resource.ResourceId,
		RestApiId:  resource.RestApiId,
	}
	_, err = svc.DeleteResource(resourceInput)
	return err
}

