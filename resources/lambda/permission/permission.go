package permission

//LambdaPermission rappresent an ApiGateway Resource on AWS
type LambdaPermission struct {
	StatementId string
	FunctionName string
}


//Delete the given resources
func (resource LambdaPermission) Delete() error {
	permissionsInput := &lambda.RemovePermissionInput{
		FunctionName: resource.FunctionName,
		StatementId:  resource.StatementId,
	}
	_, err = lambdaClient.RemovePermission(permissionsInput)	
	return err
}

