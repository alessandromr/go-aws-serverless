package permission

import (
	"github.com/alessandromr/go-aws-serverless/utils/auth"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/lambda"
)

//LambdaPermission rappresent an ApiGateway Resource on AWS
type LambdaPermission struct {
	Action       string
	FunctionName string
	Principal    string
	SourceArn    string
	StatementId  string
}

//Create the given resources
func (resource *LambdaPermission) Create() error {
	auth.MakeClient(auth.Sess)
	svc := auth.Client.LambdaConn
	permissionsInput := &lambda.AddPermissionInput{
		Action:       aws.String(resource.Action),
		FunctionName: aws.String(resource.FunctionName),
		Principal:    aws.String(resource.Principal),
		SourceArn:    aws.String(resource.SourceArn),
		StatementId:  aws.String(resource.StatementId),
	}
	_, err := svc.AddPermission(permissionsInput)
	return err
}

//Delete the given resources
func (resource *LambdaPermission) Delete() error {
	auth.MakeClient(auth.Sess)
	svc := auth.Client.LambdaConn
	permissionsInput := &lambda.RemovePermissionInput{
		FunctionName: &resource.FunctionName,
		StatementId:  &resource.StatementId,
	}
	_, err := svc.RemovePermission(permissionsInput)
	return err
}
