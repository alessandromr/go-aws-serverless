package permission

import (
	"github.com/alessandromr/go-aws-serverless/utils/auth"
	"github.com/aws/aws-sdk-go/service/lambda"
)

//LambdaPermission rappresent an ApiGateway Resource on AWS
type LambdaPermission struct {
	StatementId  string
	FunctionName string
}

//Delete the given resources
func (resource LambdaPermission) Delete() error {
	auth.MakeClient(auth.Sess)
	svc := auth.Client.Lambdaconn
	permissionsInput := &lambda.RemovePermissionInput{
		FunctionName: &resource.FunctionName,
		StatementId:  &resource.StatementId,
	}
	_, err := svc.RemovePermission(permissionsInput)
	return err
}
